package apiserver

import (
	"errors"
	"fmt"

	k8s_auth "k8s.io/api/authentication/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client is a client for querying k8s API server
type Client interface {
	// GetNode returns the node object for the given node name
	GetNode(nodeName string) (*v1.Node, error)

	// GetPod returns the pod object for the given pod name and namespace
	GetPod(namespace, podName string) (*v1.Pod, error)

	// ValidateToken queries k8s token review API and returns information about the given token
	ValidateToken(token string, audiences []string) (*k8s_auth.TokenReviewStatus, error)
}

type client struct {
	kubeConfigFilePath string
}

// New creates a new Client.
// There are two cases:
// - If a kubeConfigFilePath is provided, config is taken from that file -> use for clients running out of a k8s cluster
// - If not (empty kubeConfigFilePath), InClusterConfig is used          -> use for clients running in a k8s cluster
func New(kubeConfigFilePath string) Client {
	return &client{
		kubeConfigFilePath: kubeConfigFilePath,
	}
}

func (c *client) GetPod(namespace, podName string) (*v1.Pod, error) {
	// Validate inputs
	if namespace == "" {
		return nil, errors.New("empty namespace")
	}
	if podName == "" {
		return nil, errors.New("empty pod name")
	}

	// Reload config
	clientset, err := c.loadClient()
	if err != nil {
		return nil, err
	}

	// Get pod
	pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if pod == nil {
		return nil, fmt.Errorf("got nil pod for pod with name: %v", podName)
	}

	return pod, nil
}

func (c *client) GetNode(nodeName string) (*v1.Node, error) {
	// Validate inputs
	if nodeName == "" {
		return nil, errors.New("empty node name")
	}

	// Reload config
	clientset, err := c.loadClient()
	if err != nil {
		return nil, err
	}

	// Get node
	node, err := clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if node == nil {
		return nil, fmt.Errorf("got nil node for node with name: %v", nodeName)
	}

	return node, nil
}

func (c *client) ValidateToken(token string, audiences []string) (*k8s_auth.TokenReviewStatus, error) {
	// Reload config
	clientset, err := c.loadClient()
	if err != nil {
		return nil, err
	}

	// Create token review request
	req := &k8s_auth.TokenReview{
		Spec: k8s_auth.TokenReviewSpec{
			Token:     token,
			Audiences: audiences,
		},
	}

	// Do request
	resp, err := clientset.AuthenticationV1().TokenReviews().Create(req)
	if resp == nil {
		return nil, errors.New("token review API response is nil")
	}

	// Evaluate token review response (review server will populate TokenReview.Status field)
	if err != nil {
		return nil, err
	}

	if resp.Status.Error != "" {
		return nil, errors.New(resp.Status.Error)
	}

	return &resp.Status, nil
}

func (c *client) loadClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if c.kubeConfigFilePath == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", c.kubeConfigFilePath)
	}
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

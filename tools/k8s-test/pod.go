package main

import (
	"context"
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
)

func FindPodNameByPrefix(ctx context.Context, prefix string) (string, error) {
	pods, err := GetPods(ctx)
	if err != nil {
		return "", nil
	}

	var names []string
	for _, pod := range pods {
		if strings.HasPrefix(pod.Name, prefix) {
			names = append(names, pod.Name)
		}
	}

	switch len(names) {
	case 0:
		return "", fmt.Errorf("no pod found with prefix %q", prefix)
	case 1:
		return names[0], nil
	default:
		return "", fmt.Errorf("more than one pod found with prefix %q: %q", prefix, names)
	}
}

func GetPods(ctx context.Context) ([]v1.Pod, error) {
	var list v1.PodList
	if err := getList(ctx, "pods", &list, nil); err != nil {
		return nil, err
	}
	if len(list.Items) == 0 {
		return nil, NotFound.New("no pods")
	}
	return list.Items, nil
}

func GetPod(ctx context.Context, name string) (*v1.Pod, error) {
	obj := new(v1.Pod)
	ok, err := getObject(ctx, "pods", name, obj, nil)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, NotFound.New("no such pod %q", name)
	}
	return obj, nil
}

func GetPodsByOwner(ctx context.Context, owner Object) ([]v1.Pod, error) {
	var list v1.PodList
	if err := getList(ctx, "pods", &list, &owner); err != nil {
		return nil, err
	}
	if len(list.Items) == 0 {
		return nil, NotFound.New("no pods owned by %s", owner)
	}
	return list.Items, nil
}

func GetPodByOwner(ctx context.Context, name string, owner Object) (*v1.Pod, error) {
	obj := new(v1.Pod)
	ok, err := getObject(ctx, "pods", name, obj, &owner)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, NotFound.New("no pod %q owned by %s", name, owner)
	}
	return obj, nil
}

func CheckPodReady(pod *v1.Pod) error {
	for _, status := range pod.Status.ContainerStatuses {
		if w := status.State.Waiting; w != nil {
			return fmt.Errorf("container %q is not running (%s)", status.Name, w.Reason)
		}
		if t := status.State.Terminated; t != nil {
			return fmt.Errorf("container %q is has terminated (%s)", status.Name, t.Reason)
		}
	}
	switch pod.Status.Phase {
	case "Running", "Succeeded":
	default:
		if len(pod.Status.Conditions) > 0 {
			return fmt.Errorf("not running or succeeded (phase=%q, condition=%q)", pod.Status.Phase, pod.Status.Conditions[0].Message)
		}
		return fmt.Errorf("not running or succeeded (phase=%q)", pod.Status.Phase)
	}
	return nil
}

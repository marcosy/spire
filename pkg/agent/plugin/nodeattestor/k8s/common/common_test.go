package common

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spiffe/spire/proto/agent/nodeattestor"
	"github.com/spiffe/spire/proto/common/plugin"
	"github.com/stretchr/testify/suite"
)

func TestCommonAttestorPlugin(t *testing.T) {
	suite.Run(t, new(CommonAttestorSuite))
}

type CommonAttestorSuite struct {
	suite.Suite

	dir      string
	attestor *nodeattestor.BuiltIn
}

func (s *CommonAttestorSuite) SetupTest() {
	var err error
	s.dir, err = ioutil.TempDir("", "spire-k8s-common-test-")
	s.Require().NoError(err)

	s.newAttestor()
	s.configure(CommonAttestorConfig{})
}

func (s *CommonAttestorSuite) TestFetchAttestationDataNotConfigured() {
	s.newAttestor()
	s.requireFetchError("k8s-sat: not configured")
}

func (s *CommonAttestorSuite) TestFetchAttestationDataNoToken() {
	s.configure(CommonAttestorConfig{
		TokenPath: s.joinPath("token"),
	})
	s.requireFetchError("unable to load token from")
}

func (s *CommonAttestorSuite) TestFetchAttestationDataSuccess() {
	s.configure(CommonAttestorConfig{
		TokenPath: s.writeValue("token", "TOKEN"),
	})

	stream, err := s.attestor.FetchAttestationData(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(stream)

	resp, err := stream.Recv()
	s.Require().NoError(err)
	s.Require().NotNil(resp)

	// assert attestation data
	s.Require().Equal("spiffe://example.org/spire/agent/k8s_sat/production/UUID", resp.SpiffeId)
	s.Require().NotNil(resp.AttestationData)
	s.Require().Equal("k8s_sat", resp.AttestationData.Type)
	s.Require().JSONEq(`{
		"cluster": "production",
		"uuid": "UUID",
		"token": "TOKEN"
	}`, string(resp.AttestationData.Data))

	// node attestor should return EOF now
	_, err = stream.Recv()
	s.Require().Equal(io.EOF, err)
}

func (s *CommonAttestorSuite) TestConfigure() {
	// malformed configuration
	resp, err := s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		GlobalConfig:  &plugin.ConfigureRequest_GlobalConfig{},
		Configuration: "blah",
	})
	s.requireErrorContains(err, "k8s-sat: unable to decode configuration")
	s.Require().Nil(resp)

	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{})
	s.requireErrorContains(err, "k8s-sat: global configuration is required")
	s.Require().Nil(resp)

	// missing trust domain
	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{GlobalConfig: &plugin.ConfigureRequest_GlobalConfig{}})
	s.Require().EqualError(err, "k8s-sat: global configuration missing trust domain")
	s.Require().Nil(resp)

	// missing cluster
	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		GlobalConfig: &plugin.ConfigureRequest_GlobalConfig{TrustDomain: "example.org"},
	})
	s.Require().EqualError(err, "k8s-sat: configuration missing cluster")
	s.Require().Nil(resp)

	// success
	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		GlobalConfig:  &plugin.ConfigureRequest_GlobalConfig{TrustDomain: "example.org"},
		Configuration: `cluster = "production"`,
	})
	s.Require().NoError(err)
	s.Require().Equal(resp, &plugin.ConfigureResponse{})
}

func (s *CommonAttestorSuite) TestGetPluginInfo() {
	resp, err := s.attestor.GetPluginInfo(context.Background(), &plugin.GetPluginInfoRequest{})
	s.Require().NoError(err)
	s.Require().Equal(resp, &plugin.GetPluginInfoResponse{})
}

func (s *CommonAttestorSuite) newAttestor() {
	attestor := NewCommonAttestorPlugin("k8s_sat", "/run/secrets/kubernetes.io/serviceaccount/token", "k8s-sat")
	attestor.hooks.newUUID = func() (string, error) {
		return "UUID", nil
	}
	s.attestor = nodeattestor.NewBuiltIn(attestor)
}

func (s *CommonAttestorSuite) configure(config CommonAttestorConfig) {
	_, err := s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		GlobalConfig: &plugin.ConfigureRequest_GlobalConfig{
			TrustDomain: "example.org",
		},
		Configuration: fmt.Sprintf(`
			cluster = "production"
			token_path = %q`, config.TokenPath),
	})
	s.Require().NoError(err)

}
func (s *CommonAttestorSuite) joinPath(path string) string {
	return filepath.Join(s.dir, path)
}

func (s *CommonAttestorSuite) writeValue(path, data string) string {
	valuePath := s.joinPath(path)
	err := os.MkdirAll(filepath.Dir(valuePath), 0755)
	s.Require().NoError(err)
	err = ioutil.WriteFile(valuePath, []byte(data), 0644)
	s.Require().NoError(err)
	return valuePath
}

func (s *CommonAttestorSuite) requireFetchError(contains string) {
	stream, err := s.attestor.FetchAttestationData(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(stream)

	resp, err := stream.Recv()
	s.requireErrorContains(err, contains)
	s.Require().Nil(resp)
}

func (s *CommonAttestorSuite) requireErrorContains(err error, contains string) {
	s.Require().Error(err)
	s.Require().Contains(err.Error(), contains)
}

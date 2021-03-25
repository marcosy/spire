package jointoken

import (
	"context"
	"errors"

	"github.com/spiffe/spire/pkg/common/catalog"
	nodeattestorv0 "github.com/spiffe/spire/proto/spire/agent/nodeattestor/v0"
	spi "github.com/spiffe/spire/proto/spire/common/plugin"
)

const (
	pluginName = "join_token"
)

func BuiltIn() catalog.Plugin {
	return builtin(New())
}

func builtin(p *Plugin) catalog.Plugin {
	return catalog.MakePlugin(pluginName, nodeattestorv0.PluginServer(p))
}

type Plugin struct {
	nodeattestorv0.UnsafeNodeAttestorServer
}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) FetchAttestationData(stream nodeattestorv0.NodeAttestor_FetchAttestationDataServer) error {
	return errors.New("join token attestation is currently implemented within the agent")
}

func (p *Plugin) Configure(ctx context.Context, req *spi.ConfigureRequest) (*spi.ConfigureResponse, error) {
	return &spi.ConfigureResponse{}, nil
}

func (*Plugin) GetPluginInfo(context.Context, *spi.GetPluginInfoRequest) (*spi.GetPluginInfoResponse, error) {
	return &spi.GetPluginInfoResponse{}, nil
}

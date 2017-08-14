// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_attestor.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	node_attestor.proto

It has these top-level messages:
	AttestedData
	AttestRequest
	AttestResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "github.com/spiffe/control-plane/plugins/common/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type ConfigureRequest proto2.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*proto2.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*proto2.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*proto2.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type ConfigureResponse proto2.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*proto2.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*proto2.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*proto2.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type GetPluginInfoRequest proto2.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*proto2.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*proto2.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type GetPluginInfoResponse proto2.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()          { (*proto2.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string  { return (*proto2.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()     {}
func (m *GetPluginInfoResponse) GetName() string { return (*proto2.GetPluginInfoResponse)(m).GetName() }
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string { return (*proto2.GetPluginInfoResponse)(m).GetType() }
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*proto2.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*proto2.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*proto2.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCompany()
}

// *A type which contains attestation data for specific platform.
type AttestedData struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AttestedData) Reset()                    { *m = AttestedData{} }
func (m *AttestedData) String() string            { return proto1.CompactTextString(m) }
func (*AttestedData) ProtoMessage()               {}
func (*AttestedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AttestedData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestedData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// *Represents a request to attest a node.
type AttestRequest struct {
	AttestedData   *AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	AttestedBefore bool          `protobuf:"varint,2,opt,name=attestedBefore" json:"attestedBefore,omitempty"`
}

func (m *AttestRequest) Reset()                    { *m = AttestRequest{} }
func (m *AttestRequest) String() string            { return proto1.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()               {}
func (*AttestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttestRequest) GetAttestedData() *AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *AttestRequest) GetAttestedBefore() bool {
	if m != nil {
		return m.AttestedBefore
	}
	return false
}

// *Represents a response when attesting a node.
type AttestResponse struct {
	Valid        bool   `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	BaseSPIFFEID string `protobuf:"bytes,2,opt,name=baseSPIFFEID" json:"baseSPIFFEID,omitempty"`
}

func (m *AttestResponse) Reset()                    { *m = AttestResponse{} }
func (m *AttestResponse) String() string            { return proto1.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()               {}
func (*AttestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AttestResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *AttestResponse) GetBaseSPIFFEID() string {
	if m != nil {
		return m.BaseSPIFFEID
	}
	return ""
}

func init() {
	proto1.RegisterType((*AttestedData)(nil), "proto.AttestedData")
	proto1.RegisterType((*AttestRequest)(nil), "proto.AttestRequest")
	proto1.RegisterType((*AttestResponse)(nil), "proto.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	// *Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *proto2.GetPluginInfoRequest, opts ...grpc.CallOption) (*proto2.GetPluginInfoResponse, error)
	// *Attesta a node.
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error) {
	out := new(proto2.ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *proto2.GetPluginInfoRequest, opts ...grpc.CallOption) (*proto2.GetPluginInfoResponse, error) {
	out := new(proto2.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/proto.NodeAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := grpc.Invoke(ctx, "/proto.NodeAttestor/Attest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	// *Responsible for configuration of the plugin.
	Configure(context.Context, *proto2.ConfigureRequest) (*proto2.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *proto2.GetPluginInfoRequest) (*proto2.GetPluginInfoResponse, error)
	// *Attesta a node.
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*proto2.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*proto2.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
		{
			MethodName: "Attest",
			Handler:    _NodeAttestor_Attest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_attestor.proto",
}

func init() { proto1.RegisterFile("node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xad, 0x11, 0x02, 0x63, 0xe1, 0xb0, 0x60, 0x24, 0xe8, 0x81, 0xf4, 0x60, 0xb8, 0x48,
	0x13, 0x8c, 0x7a, 0x33, 0x41, 0x11, 0x03, 0x07, 0x43, 0xd6, 0x07, 0x30, 0x0b, 0x9d, 0x62, 0x93,
	0xb2, 0xb3, 0x76, 0xb7, 0x26, 0x3e, 0xa6, 0x6f, 0x64, 0xdc, 0x6e, 0x0d, 0x45, 0x4f, 0xdd, 0xf9,
	0x67, 0xfe, 0xef, 0x9f, 0x0e, 0x74, 0x24, 0x45, 0xf8, 0x2a, 0x8c, 0x41, 0x6d, 0x28, 0x1b, 0xa9,
	0x8c, 0x0c, 0xb1, 0x9a, 0xfd, 0xf4, 0x27, 0x9b, 0xc4, 0xbc, 0xe5, 0xab, 0xd1, 0x9a, 0xb6, 0xa1,
	0x56, 0x49, 0x1c, 0x63, 0xb8, 0x26, 0x69, 0x32, 0x4a, 0x2f, 0x55, 0x2a, 0x24, 0x86, 0x2a, 0xcd,
	0x37, 0x89, 0xd4, 0xe1, 0x9a, 0xb6, 0x5b, 0x92, 0xa1, 0x75, 0xb9, 0xa2, 0x20, 0x05, 0x37, 0xe0,
	0x4f, 0x2c, 0x1b, 0xa3, 0xa9, 0x30, 0x82, 0x31, 0x38, 0x32, 0x9f, 0x0a, 0x7b, 0xde, 0xc0, 0x1b,
	0x36, 0xb9, 0x7d, 0xff, 0x68, 0x91, 0x30, 0xa2, 0x77, 0x38, 0xf0, 0x86, 0x3e, 0xb7, 0xef, 0x40,
	0x41, 0xab, 0xf0, 0x71, 0x7c, 0xcf, 0x51, 0x1b, 0x76, 0x0b, 0xbe, 0xd8, 0x01, 0x59, 0xc0, 0xf1,
	0xb8, 0x53, 0xc4, 0x8c, 0x76, 0x33, 0x78, 0x65, 0x90, 0x5d, 0x40, 0xbb, 0xac, 0xef, 0x31, 0xa6,
	0x0c, 0x6d, 0x4e, 0x83, 0xef, 0xa9, 0xc1, 0x02, 0xda, 0x65, 0xa2, 0x56, 0x24, 0x35, 0xb2, 0x2e,
	0xd4, 0x3e, 0x44, 0x9a, 0x44, 0x36, 0xab, 0xc1, 0x8b, 0x82, 0x05, 0xe0, 0xaf, 0x84, 0xc6, 0x97,
	0xe5, 0x7c, 0x36, 0x7b, 0x9c, 0x4f, 0x2d, 0xad, 0xc9, 0x2b, 0xda, 0xf8, 0xcb, 0x03, 0xff, 0x99,
	0x22, 0x9c, 0xb8, 0xb3, 0xb2, 0x3b, 0x68, 0x3e, 0x90, 0x8c, 0x93, 0x4d, 0x9e, 0x21, 0x3b, 0x75,
	0x4b, 0xff, 0x2a, 0xee, 0x1f, 0xfb, 0xbd, 0xbf, 0x0d, 0xb7, 0xca, 0x02, 0x5a, 0x4f, 0x68, 0x96,
	0xf6, 0xdc, 0x73, 0x19, 0x13, 0x3b, 0x73, 0xa3, 0x15, 0xb5, 0xe4, 0x9c, 0xff, 0xdf, 0x74, 0xac,
	0x6b, 0xa8, 0x17, 0x7b, 0xb1, 0x6e, 0xe5, 0x7a, 0xa5, 0xfb, 0x64, 0x4f, 0x2d, 0x6c, 0xcb, 0x83,
	0x55, 0xdd, 0xea, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x2a, 0x41, 0xba, 0x33, 0x02,
	0x00, 0x00,
}
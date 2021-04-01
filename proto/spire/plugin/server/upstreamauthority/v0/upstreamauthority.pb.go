// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spire/plugin/server/upstreamauthority/v0/upstreamauthority.proto

package upstreamauthorityv0

import (
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MintX509CARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Certificate signing request (PKCS#10)
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Preferred TTL is the TTL preferred by SPIRE server for signed CA. If
	// zero, the plugin should determine its own TTL value. Plugins are free to
	// ignore this and use their own policies around TTLs.
	PreferredTtl int32 `protobuf:"varint,2,opt,name=preferred_ttl,json=preferredTtl,proto3" json:"preferred_ttl,omitempty"`
}

func (x *MintX509CARequest) Reset() {
	*x = MintX509CARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintX509CARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintX509CARequest) ProtoMessage() {}

func (x *MintX509CARequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintX509CARequest.ProtoReflect.Descriptor instead.
func (*MintX509CARequest) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescGZIP(), []int{0}
}

func (x *MintX509CARequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

func (x *MintX509CARequest) GetPreferredTtl() int32 {
	if x != nil {
		return x.PreferredTtl
	}
	return 0
}

type MintX509CAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains ASN.1 encoded certificates representing the X.509 CA along with
	// any intermediates necessary to chain back to a certificate present in
	// the upstream_x509_roots.
	X509CaChain [][]byte `protobuf:"bytes,1,rep,name=x509_ca_chain,json=x509CaChain,proto3" json:"x509_ca_chain,omitempty"`
	// The trusted X.509 root authorities for the upstream authority
	UpstreamX509Roots [][]byte `protobuf:"bytes,2,rep,name=upstream_x509_roots,json=upstreamX509Roots,proto3" json:"upstream_x509_roots,omitempty"`
}

func (x *MintX509CAResponse) Reset() {
	*x = MintX509CAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintX509CAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintX509CAResponse) ProtoMessage() {}

func (x *MintX509CAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintX509CAResponse.ProtoReflect.Descriptor instead.
func (*MintX509CAResponse) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescGZIP(), []int{1}
}

func (x *MintX509CAResponse) GetX509CaChain() [][]byte {
	if x != nil {
		return x.X509CaChain
	}
	return nil
}

func (x *MintX509CAResponse) GetUpstreamX509Roots() [][]byte {
	if x != nil {
		return x.UpstreamX509Roots
	}
	return nil
}

type PublishJWTKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The JWT signing key to publish upstream
	JwtKey *common.PublicKey `protobuf:"bytes,1,opt,name=jwt_key,json=jwtKey,proto3" json:"jwt_key,omitempty"`
}

func (x *PublishJWTKeyRequest) Reset() {
	*x = PublishJWTKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishJWTKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishJWTKeyRequest) ProtoMessage() {}

func (x *PublishJWTKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishJWTKeyRequest.ProtoReflect.Descriptor instead.
func (*PublishJWTKeyRequest) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescGZIP(), []int{2}
}

func (x *PublishJWTKeyRequest) GetJwtKey() *common.PublicKey {
	if x != nil {
		return x.JwtKey
	}
	return nil
}

type PublishJWTKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The upstream JWT signing keys
	UpstreamJwtKeys []*common.PublicKey `protobuf:"bytes,1,rep,name=upstream_jwt_keys,json=upstreamJwtKeys,proto3" json:"upstream_jwt_keys,omitempty"`
}

func (x *PublishJWTKeyResponse) Reset() {
	*x = PublishJWTKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishJWTKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishJWTKeyResponse) ProtoMessage() {}

func (x *PublishJWTKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishJWTKeyResponse.ProtoReflect.Descriptor instead.
func (*PublishJWTKeyResponse) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescGZIP(), []int{3}
}

func (x *PublishJWTKeyResponse) GetUpstreamJwtKeys() []*common.PublicKey {
	if x != nil {
		return x.UpstreamJwtKeys
	}
	return nil
}

var File_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto protoreflect.FileDescriptor

var file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDesc = []byte{
	0x0a, 0x40, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x30, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x1a, 0x19, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x74, 0x58, 0x35, 0x30, 0x39, 0x43, 0x41, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x54, 0x74, 0x6c, 0x22, 0x68, 0x0a, 0x12, 0x4d,
	0x69, 0x6e, 0x74, 0x58, 0x35, 0x30, 0x39, 0x43, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x78, 0x35, 0x30, 0x39, 0x43, 0x61,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x11, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x58, 0x35, 0x30, 0x39,
	0x52, 0x6f, 0x6f, 0x74, 0x73, 0x22, 0x48, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x07, 0x6a, 0x77, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x6a, 0x77, 0x74, 0x4b, 0x65, 0x79, 0x22,
	0x5c, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x11, 0x75, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x0f, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x77, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x32, 0xce, 0x03,
	0x0a, 0x11, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x75, 0x0a, 0x0a, 0x4d, 0x69, 0x6e, 0x74, 0x58, 0x35, 0x30, 0x39, 0x43,
	0x41, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x58, 0x35, 0x30, 0x39, 0x43, 0x41, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x58, 0x35, 0x30, 0x39, 0x43, 0x41,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x7e, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5c,
	0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69,
	0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x30, 0x3b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescOnce sync.Once
	file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescData = file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDesc
)

func file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescGZIP() []byte {
	file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescOnce.Do(func() {
		file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescData)
	})
	return file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDescData
}

var file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_goTypes = []interface{}{
	(*MintX509CARequest)(nil),            // 0: spire.server.upstreamauthority.MintX509CARequest
	(*MintX509CAResponse)(nil),           // 1: spire.server.upstreamauthority.MintX509CAResponse
	(*PublishJWTKeyRequest)(nil),         // 2: spire.server.upstreamauthority.PublishJWTKeyRequest
	(*PublishJWTKeyResponse)(nil),        // 3: spire.server.upstreamauthority.PublishJWTKeyResponse
	(*common.PublicKey)(nil),             // 4: spire.common.PublicKey
	(*plugin.ConfigureRequest)(nil),      // 5: spire.common.plugin.ConfigureRequest
	(*plugin.GetPluginInfoRequest)(nil),  // 6: spire.common.plugin.GetPluginInfoRequest
	(*plugin.ConfigureResponse)(nil),     // 7: spire.common.plugin.ConfigureResponse
	(*plugin.GetPluginInfoResponse)(nil), // 8: spire.common.plugin.GetPluginInfoResponse
}
var file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_depIdxs = []int32{
	4, // 0: spire.server.upstreamauthority.PublishJWTKeyRequest.jwt_key:type_name -> spire.common.PublicKey
	4, // 1: spire.server.upstreamauthority.PublishJWTKeyResponse.upstream_jwt_keys:type_name -> spire.common.PublicKey
	0, // 2: spire.server.upstreamauthority.UpstreamAuthority.MintX509CA:input_type -> spire.server.upstreamauthority.MintX509CARequest
	2, // 3: spire.server.upstreamauthority.UpstreamAuthority.PublishJWTKey:input_type -> spire.server.upstreamauthority.PublishJWTKeyRequest
	5, // 4: spire.server.upstreamauthority.UpstreamAuthority.Configure:input_type -> spire.common.plugin.ConfigureRequest
	6, // 5: spire.server.upstreamauthority.UpstreamAuthority.GetPluginInfo:input_type -> spire.common.plugin.GetPluginInfoRequest
	1, // 6: spire.server.upstreamauthority.UpstreamAuthority.MintX509CA:output_type -> spire.server.upstreamauthority.MintX509CAResponse
	3, // 7: spire.server.upstreamauthority.UpstreamAuthority.PublishJWTKey:output_type -> spire.server.upstreamauthority.PublishJWTKeyResponse
	7, // 8: spire.server.upstreamauthority.UpstreamAuthority.Configure:output_type -> spire.common.plugin.ConfigureResponse
	8, // 9: spire.server.upstreamauthority.UpstreamAuthority.GetPluginInfo:output_type -> spire.common.plugin.GetPluginInfoResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_init() }
func file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_init() {
	if File_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintX509CARequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintX509CAResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishJWTKeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishJWTKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_goTypes,
		DependencyIndexes: file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_depIdxs,
		MessageInfos:      file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_msgTypes,
	}.Build()
	File_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto = out.File
	file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_rawDesc = nil
	file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_goTypes = nil
	file_spire_plugin_server_upstreamauthority_v0_upstreamauthority_proto_depIdxs = nil
}

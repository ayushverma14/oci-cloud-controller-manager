// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tenantagent.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import nodepools "bitbucket.oci.oraclecorp.com/oke/oke-common/types/nodepools"
import _ "bitbucket.oci.oraclecorp.com/oke/oke-common/types/nodes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TKWClusterRequest struct {
	JobID               string         `protobuf:"bytes,1,opt,name=JobID" json:"JobID,omitempty"`
	APIServerAddr       string         `protobuf:"bytes,8,opt,name=APIServerAddr" json:"APIServerAddr,omitempty"`
	ETCDAddrs           []string       `protobuf:"bytes,9,rep,name=ETCDAddrs" json:"ETCDAddrs,omitempty"`
	ServiceClusterCIDR  string         `protobuf:"bytes,10,opt,name=ServiceClusterCIDR" json:"ServiceClusterCIDR,omitempty"`
	KubeDNSServiceIP    string         `protobuf:"bytes,11,opt,name=KubeDNSServiceIP" json:"KubeDNSServiceIP,omitempty"`
	FlannelConfig       *FlannelConfig `protobuf:"bytes,13,opt,name=FlannelConfig" json:"FlannelConfig,omitempty"`
	CompartmentID       string         `protobuf:"bytes,14,opt,name=CompartmentID" json:"CompartmentID,omitempty"`
	TenancyID           string         `protobuf:"bytes,15,opt,name=TenancyID" json:"TenancyID,omitempty"`
	RegionName          string         `protobuf:"bytes,16,opt,name=RegionName" json:"RegionName,omitempty"`
	NetworkConfig       *NetworkConfig `protobuf:"bytes,17,opt,name=NetworkConfig" json:"NetworkConfig,omitempty"`
	RegionKey           string         `protobuf:"bytes,19,opt,name=RegionKey" json:"RegionKey,omitempty"`
	ProxymuxServicePort string         `protobuf:"bytes,20,opt,name=ProxymuxServicePort" json:"ProxymuxServicePort,omitempty"`
}

func (m *TKWClusterRequest) Reset()                    { *m = TKWClusterRequest{} }
func (m *TKWClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*TKWClusterRequest) ProtoMessage()               {}
func (*TKWClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *TKWClusterRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *TKWClusterRequest) GetAPIServerAddr() string {
	if m != nil {
		return m.APIServerAddr
	}
	return ""
}

func (m *TKWClusterRequest) GetETCDAddrs() []string {
	if m != nil {
		return m.ETCDAddrs
	}
	return nil
}

func (m *TKWClusterRequest) GetServiceClusterCIDR() string {
	if m != nil {
		return m.ServiceClusterCIDR
	}
	return ""
}

func (m *TKWClusterRequest) GetKubeDNSServiceIP() string {
	if m != nil {
		return m.KubeDNSServiceIP
	}
	return ""
}

func (m *TKWClusterRequest) GetFlannelConfig() *FlannelConfig {
	if m != nil {
		return m.FlannelConfig
	}
	return nil
}

func (m *TKWClusterRequest) GetCompartmentID() string {
	if m != nil {
		return m.CompartmentID
	}
	return ""
}

func (m *TKWClusterRequest) GetTenancyID() string {
	if m != nil {
		return m.TenancyID
	}
	return ""
}

func (m *TKWClusterRequest) GetRegionName() string {
	if m != nil {
		return m.RegionName
	}
	return ""
}

func (m *TKWClusterRequest) GetNetworkConfig() *NetworkConfig {
	if m != nil {
		return m.NetworkConfig
	}
	return nil
}

func (m *TKWClusterRequest) GetRegionKey() string {
	if m != nil {
		return m.RegionKey
	}
	return ""
}

func (m *TKWClusterRequest) GetProxymuxServicePort() string {
	if m != nil {
		return m.ProxymuxServicePort
	}
	return ""
}

type FlannelConfig struct {
	ETCDAddrs  []string `protobuf:"bytes,1,rep,name=ETCDAddrs" json:"ETCDAddrs,omitempty"`
	ETCDCACert []byte   `protobuf:"bytes,2,opt,name=ETCDCACert,proto3" json:"ETCDCACert,omitempty"`
	ETCDCert   []byte   `protobuf:"bytes,3,opt,name=ETCDCert,proto3" json:"ETCDCert,omitempty"`
	ETCDKey    []byte   `protobuf:"bytes,4,opt,name=ETCDKey,proto3" json:"ETCDKey,omitempty"`
}

func (m *FlannelConfig) Reset()                    { *m = FlannelConfig{} }
func (m *FlannelConfig) String() string            { return proto.CompactTextString(m) }
func (*FlannelConfig) ProtoMessage()               {}
func (*FlannelConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *FlannelConfig) GetETCDAddrs() []string {
	if m != nil {
		return m.ETCDAddrs
	}
	return nil
}

func (m *FlannelConfig) GetETCDCACert() []byte {
	if m != nil {
		return m.ETCDCACert
	}
	return nil
}

func (m *FlannelConfig) GetETCDCert() []byte {
	if m != nil {
		return m.ETCDCert
	}
	return nil
}

func (m *FlannelConfig) GetETCDKey() []byte {
	if m != nil {
		return m.ETCDKey
	}
	return nil
}

type GetTKWClusterRequest struct {
	JobID string `protobuf:"bytes,1,opt,name=JobID" json:"JobID,omitempty"`
}

func (m *GetTKWClusterRequest) Reset()                    { *m = GetTKWClusterRequest{} }
func (m *GetTKWClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTKWClusterRequest) ProtoMessage()               {}
func (*GetTKWClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *GetTKWClusterRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

type TKWClusterResponse struct {
	TAGitCommit    string                              `protobuf:"bytes,5,opt,name=TAGitCommit" json:"TAGitCommit,omitempty"`
	TAVersion      string                              `protobuf:"bytes,6,opt,name=TAVersion" json:"TAVersion,omitempty"`
	NodePoolStates map[string]*nodepools.NodePoolState `protobuf:"bytes,8,rep,name=NodePoolStates" json:"NodePoolStates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TKWClusterResponse) Reset()                    { *m = TKWClusterResponse{} }
func (m *TKWClusterResponse) String() string            { return proto.CompactTextString(m) }
func (*TKWClusterResponse) ProtoMessage()               {}
func (*TKWClusterResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *TKWClusterResponse) GetTAGitCommit() string {
	if m != nil {
		return m.TAGitCommit
	}
	return ""
}

func (m *TKWClusterResponse) GetTAVersion() string {
	if m != nil {
		return m.TAVersion
	}
	return ""
}

func (m *TKWClusterResponse) GetNodePoolStates() map[string]*nodepools.NodePoolState {
	if m != nil {
		return m.NodePoolStates
	}
	return nil
}

type BuildInfo struct {
	GitCommit   string `protobuf:"bytes,1,opt,name=GitCommit" json:"GitCommit,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=Version" json:"Version,omitempty"`
	Release     string `protobuf:"bytes,3,opt,name=Release" json:"Release,omitempty"`
	BuildHost   string `protobuf:"bytes,4,opt,name=BuildHost" json:"BuildHost,omitempty"`
	BuildDate   string `protobuf:"bytes,5,opt,name=BuildDate" json:"BuildDate,omitempty"`
	BuildBranch string `protobuf:"bytes,6,opt,name=BuildBranch" json:"BuildBranch,omitempty"`
}

func (m *BuildInfo) Reset()                    { *m = BuildInfo{} }
func (m *BuildInfo) String() string            { return proto.CompactTextString(m) }
func (*BuildInfo) ProtoMessage()               {}
func (*BuildInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *BuildInfo) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *BuildInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *BuildInfo) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *BuildInfo) GetBuildHost() string {
	if m != nil {
		return m.BuildHost
	}
	return ""
}

func (m *BuildInfo) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *BuildInfo) GetBuildBranch() string {
	if m != nil {
		return m.BuildBranch
	}
	return ""
}

type UpdateNodePoolRequest struct {
	NodePool *nodepools.NodePool `protobuf:"bytes,1,opt,name=NodePool" json:"NodePool,omitempty"`
}

func (m *UpdateNodePoolRequest) Reset()                    { *m = UpdateNodePoolRequest{} }
func (m *UpdateNodePoolRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodePoolRequest) ProtoMessage()               {}
func (*UpdateNodePoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UpdateNodePoolRequest) GetNodePool() *nodepools.NodePool {
	if m != nil {
		return m.NodePool
	}
	return nil
}

type DeleteNodePoolRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *DeleteNodePoolRequest) Reset()                    { *m = DeleteNodePoolRequest{} }
func (m *DeleteNodePoolRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodePoolRequest) ProtoMessage()               {}
func (*DeleteNodePoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *DeleteNodePoolRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*TKWClusterRequest)(nil), "types.TKWClusterRequest")
	proto.RegisterType((*FlannelConfig)(nil), "types.FlannelConfig")
	proto.RegisterType((*GetTKWClusterRequest)(nil), "types.GetTKWClusterRequest")
	proto.RegisterType((*TKWClusterResponse)(nil), "types.TKWClusterResponse")
	proto.RegisterType((*BuildInfo)(nil), "types.BuildInfo")
	proto.RegisterType((*UpdateNodePoolRequest)(nil), "types.UpdateNodePoolRequest")
	proto.RegisterType((*DeleteNodePoolRequest)(nil), "types.DeleteNodePoolRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TenantAgent service

type TenantAgentClient interface {
	UpdateNodePool(ctx context.Context, in *UpdateNodePoolRequest, opts ...grpc.CallOption) (*TKWClusterResponse, error)
	DeleteNodePool(ctx context.Context, in *DeleteNodePoolRequest, opts ...grpc.CallOption) (*TKWClusterResponse, error)
	GetTKWCluster(ctx context.Context, in *GetTKWClusterRequest, opts ...grpc.CallOption) (*TKWClusterResponse, error)
	Version(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*BuildInfo, error)
}

type tenantAgentClient struct {
	cc *grpc.ClientConn
}

func NewTenantAgentClient(cc *grpc.ClientConn) TenantAgentClient {
	return &tenantAgentClient{cc}
}

func (c *tenantAgentClient) UpdateNodePool(ctx context.Context, in *UpdateNodePoolRequest, opts ...grpc.CallOption) (*TKWClusterResponse, error) {
	out := new(TKWClusterResponse)
	err := grpc.Invoke(ctx, "/types.TenantAgent/UpdateNodePool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantAgentClient) DeleteNodePool(ctx context.Context, in *DeleteNodePoolRequest, opts ...grpc.CallOption) (*TKWClusterResponse, error) {
	out := new(TKWClusterResponse)
	err := grpc.Invoke(ctx, "/types.TenantAgent/DeleteNodePool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantAgentClient) GetTKWCluster(ctx context.Context, in *GetTKWClusterRequest, opts ...grpc.CallOption) (*TKWClusterResponse, error) {
	out := new(TKWClusterResponse)
	err := grpc.Invoke(ctx, "/types.TenantAgent/GetTKWCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantAgentClient) Version(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*BuildInfo, error) {
	out := new(BuildInfo)
	err := grpc.Invoke(ctx, "/types.TenantAgent/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TenantAgent service

type TenantAgentServer interface {
	UpdateNodePool(context.Context, *UpdateNodePoolRequest) (*TKWClusterResponse, error)
	DeleteNodePool(context.Context, *DeleteNodePoolRequest) (*TKWClusterResponse, error)
	GetTKWCluster(context.Context, *GetTKWClusterRequest) (*TKWClusterResponse, error)
	Version(context.Context, *google_protobuf.Empty) (*BuildInfo, error)
}

func RegisterTenantAgentServer(s *grpc.Server, srv TenantAgentServer) {
	s.RegisterService(&_TenantAgent_serviceDesc, srv)
}

func _TenantAgent_UpdateNodePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantAgentServer).UpdateNodePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.TenantAgent/UpdateNodePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantAgentServer).UpdateNodePool(ctx, req.(*UpdateNodePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantAgent_DeleteNodePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantAgentServer).DeleteNodePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.TenantAgent/DeleteNodePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantAgentServer).DeleteNodePool(ctx, req.(*DeleteNodePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantAgent_GetTKWCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTKWClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantAgentServer).GetTKWCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.TenantAgent/GetTKWCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantAgentServer).GetTKWCluster(ctx, req.(*GetTKWClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantAgent_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantAgentServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.TenantAgent/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantAgentServer).Version(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.TenantAgent",
	HandlerType: (*TenantAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateNodePool",
			Handler:    _TenantAgent_UpdateNodePool_Handler,
		},
		{
			MethodName: "DeleteNodePool",
			Handler:    _TenantAgent_DeleteNodePool_Handler,
		},
		{
			MethodName: "GetTKWCluster",
			Handler:    _TenantAgent_GetTKWCluster_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _TenantAgent_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenantagent.proto",
}

func init() { proto.RegisterFile("tenantagent.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0x55, 0x12, 0xba, 0x34, 0xd7, 0xb4, 0xb4, 0xd3, 0x2c, 0x9a, 0x0d, 0xab, 0x55, 0x14, 0x21,
	0x11, 0x21, 0xd6, 0x81, 0xf4, 0x05, 0x2d, 0x4f, 0xa9, 0x5d, 0xba, 0xa6, 0x52, 0x15, 0xb9, 0x5d,
	0x90, 0xe0, 0xc9, 0x76, 0x6e, 0xbd, 0x56, 0xed, 0x19, 0x33, 0x1e, 0x2f, 0x9b, 0x67, 0x5e, 0xf8,
	0x1b, 0xfc, 0x0e, 0x7e, 0x17, 0xef, 0x68, 0xc6, 0x63, 0x3b, 0x6e, 0x43, 0xd9, 0x87, 0x56, 0x73,
	0xcf, 0x39, 0x73, 0xe6, 0x7e, 0xcc, 0xc4, 0x70, 0x2c, 0x91, 0x05, 0x4c, 0x06, 0x31, 0x32, 0x69,
	0xe7, 0x82, 0x4b, 0x4e, 0xf6, 0xe4, 0x26, 0xc7, 0x62, 0x7c, 0x1a, 0x27, 0xf2, 0x6d, 0x19, 0xda,
	0x11, 0xcf, 0xe6, 0x31, 0x4f, 0x03, 0x16, 0xcf, 0x35, 0x1f, 0x96, 0xb7, 0xf3, 0x5c, 0x4b, 0xe6,
	0x98, 0xe5, 0x72, 0x53, 0xfd, 0xaf, 0xf6, 0x8e, 0xbf, 0xff, 0xff, 0x4d, 0x32, 0xc9, 0xb0, 0x90,
	0x41, 0x96, 0xb7, 0x2b, 0xb3, 0x19, 0xc2, 0x2c, 0x2a, 0xcc, 0xda, 0x42, 0x21, 0xb8, 0x30, 0xc1,
	0x88, 0xa1, 0xfc, 0x9d, 0x8b, 0xbb, 0x88, 0xb3, 0xdb, 0x24, 0xae, 0x25, 0xcf, 0x18, 0x5f, 0x63,
	0xce, 0x79, 0x5a, 0xcc, 0x9b, 0x95, 0xa1, 0x8e, 0x15, 0x50, 0xc1, 0x06, 0x9a, 0xfe, 0x33, 0x80,
	0xe3, 0x9b, 0xcb, 0x9f, 0x9d, 0xb4, 0x2c, 0x24, 0x0a, 0x1f, 0x7f, 0x2b, 0xb1, 0x90, 0x64, 0x04,
	0x7b, 0x3f, 0xf2, 0xd0, 0x73, 0x69, 0x6f, 0xd2, 0x9b, 0x0d, 0xfd, 0x2a, 0x20, 0x5f, 0xc0, 0xc1,
	0x72, 0xe5, 0x5d, 0xa3, 0x78, 0x87, 0x62, 0xb9, 0x5e, 0x0b, 0xba, 0xaf, 0xd9, 0x2e, 0x48, 0x9e,
	0xc3, 0xf0, 0xfc, 0xc6, 0x71, 0xd5, 0xba, 0xa0, 0xc3, 0xc9, 0x60, 0x36, 0xf4, 0x5b, 0x80, 0xd8,
	0x40, 0x94, 0x36, 0x89, 0xd0, 0x1c, 0xe9, 0x78, 0xae, 0x4f, 0x41, 0x1b, 0xed, 0x60, 0xc8, 0x57,
	0x70, 0x74, 0x59, 0x86, 0xe8, 0x5e, 0x5d, 0x1b, 0xd2, 0x5b, 0x51, 0x4b, 0xab, 0x1f, 0xe0, 0xe4,
	0x15, 0x1c, 0xfc, 0x90, 0x06, 0x8c, 0x61, 0xea, 0xe8, 0x8e, 0xd0, 0x83, 0x49, 0x6f, 0x66, 0x2d,
	0x46, 0xb6, 0xee, 0xb0, 0xdd, 0xe1, 0xfc, 0xae, 0x54, 0xd5, 0xe6, 0xf0, 0x2c, 0x0f, 0x84, 0xcc,
	0x90, 0x49, 0xcf, 0xa5, 0x87, 0x55, 0x6d, 0x1d, 0x50, 0xd5, 0x76, 0xa3, 0x2e, 0x46, 0xb4, 0xf1,
	0x5c, 0xfa, 0xa9, 0x56, 0xb4, 0x00, 0x79, 0x01, 0xe0, 0x63, 0x9c, 0x70, 0x76, 0x15, 0x64, 0x48,
	0x8f, 0x34, 0xbd, 0x85, 0xa8, 0xfc, 0xae, 0xaa, 0x89, 0x99, 0xfc, 0x8e, 0x3b, 0xf9, 0x75, 0x38,
	0xbf, 0x2b, 0x55, 0x27, 0x57, 0x4e, 0x97, 0xb8, 0xa1, 0x27, 0xd5, 0xc9, 0x0d, 0x40, 0xbe, 0x81,
	0x93, 0x95, 0xe0, 0xef, 0x37, 0x59, 0xf9, 0xde, 0xb4, 0x63, 0xc5, 0x85, 0xa4, 0x23, 0xad, 0xdb,
	0x45, 0x4d, 0xff, 0xe8, 0xdd, 0x6b, 0x56, 0x77, 0x6e, 0xbd, 0xfb, 0x73, 0x7b, 0x01, 0xa0, 0x02,
	0x67, 0xe9, 0xa0, 0x90, 0xb4, 0x3f, 0xe9, 0xcd, 0x3e, 0xf1, 0xb7, 0x10, 0x32, 0x86, 0x7d, 0x1d,
	0x29, 0x76, 0xa0, 0xd9, 0x26, 0x26, 0x14, 0x3e, 0x56, 0x6b, 0x95, 0xf9, 0x47, 0x9a, 0xaa, 0xc3,
	0xe9, 0xd7, 0x30, 0xba, 0x40, 0xf9, 0x81, 0xf7, 0x6f, 0xfa, 0x67, 0x1f, 0xc8, 0xb6, 0xb6, 0xc8,
	0x39, 0x2b, 0x90, 0x4c, 0xc0, 0xba, 0x59, 0x5e, 0x24, 0xd2, 0xe1, 0x59, 0x96, 0x48, 0xba, 0xa7,
	0xb7, 0x6c, 0x43, 0x7a, 0x6c, 0xcb, 0x9f, 0x50, 0x14, 0x09, 0x67, 0xf4, 0x89, 0x19, 0x5b, 0x0d,
	0x90, 0x37, 0x70, 0x78, 0xc5, 0xd7, 0xb8, 0xe2, 0x3c, 0xbd, 0x96, 0x81, 0xc4, 0x82, 0xee, 0x4f,
	0x06, 0x33, 0x6b, 0xf1, 0xd2, 0xcc, 0xe5, 0xe1, 0x91, 0x76, 0x57, 0x7f, 0xce, 0xa4, 0xd8, 0xf8,
	0xf7, 0x4c, 0xc6, 0xbf, 0xc2, 0xc9, 0x0e, 0x19, 0x39, 0x82, 0xc1, 0x1d, 0x6e, 0x4c, 0x61, 0x6a,
	0x49, 0x6c, 0xd8, 0x7b, 0x17, 0xa4, 0x25, 0xea, 0xae, 0x5a, 0x0b, 0x6a, 0xb7, 0xcf, 0xb6, 0x63,
	0xe0, 0x57, 0xb2, 0x57, 0xfd, 0xef, 0x7a, 0xd3, 0xbf, 0x7b, 0x30, 0x3c, 0x2b, 0x93, 0x74, 0xed,
	0xb1, 0x5b, 0xae, 0xea, 0x6b, 0xeb, 0xaf, 0x9c, 0x5b, 0x40, 0xb5, 0xbf, 0xae, 0xbd, 0xaf, 0xb9,
	0x3a, 0x54, 0x8c, 0x8f, 0x29, 0x06, 0x05, 0xea, 0x99, 0x0d, 0xfd, 0x3a, 0x54, 0x8e, 0xda, 0xfe,
	0x35, 0x2f, 0xa4, 0x1e, 0xda, 0xd0, 0x6f, 0x81, 0x86, 0x75, 0x03, 0x89, 0xa6, 0xdf, 0x2d, 0xa0,
	0xe6, 0xa1, 0x83, 0x33, 0x11, 0xb0, 0xe8, 0xad, 0xe9, 0xf7, 0x36, 0x34, 0x7d, 0x0d, 0x4f, 0xdf,
	0xe4, 0xeb, 0x40, 0x62, 0x5d, 0x5f, 0x3d, 0xf7, 0x39, 0xec, 0xd7, 0x90, 0xae, 0xc3, 0x5a, 0x9c,
	0xec, 0xe8, 0x86, 0xdf, 0x88, 0xa6, 0x5f, 0xc2, 0x53, 0x17, 0x53, 0x7c, 0xe8, 0x74, 0x08, 0xfd,
	0xe6, 0xfa, 0xf4, 0x3d, 0x77, 0xf1, 0x57, 0x1f, 0x2c, 0xfd, 0x52, 0xe5, 0x52, 0xfd, 0xa6, 0x13,
	0x0f, 0x0e, 0xbb, 0x29, 0x90, 0xe7, 0x66, 0xdc, 0x3b, 0x33, 0x1b, 0x3f, 0xfb, 0xcf, 0xcb, 0xa0,
	0xac, 0xba, 0x39, 0x34, 0x56, 0x3b, 0x53, 0x7b, 0xcc, 0xea, 0x02, 0x0e, 0x3a, 0xef, 0x81, 0x7c,
	0x6e, 0xb4, 0xbb, 0x5e, 0xc9, 0x63, 0x46, 0xa7, 0xcd, 0xcc, 0xc9, 0x67, 0x76, 0xcc, 0x79, 0x9c,
	0xa2, 0x5d, 0x7f, 0x71, 0xec, 0x73, 0xf5, 0x65, 0x1a, 0x1f, 0x99, 0xdd, 0xcd, 0x35, 0x3a, 0x3b,
	0xfd, 0xe5, 0xdb, 0x30, 0x91, 0x61, 0x19, 0xdd, 0xa1, 0xb4, 0x79, 0x94, 0xd8, 0x5c, 0x04, 0x51,
	0x8a, 0x11, 0x17, 0xb9, 0xfe, 0x72, 0xf1, 0x3b, 0x54, 0x7f, 0x2f, 0x23, 0x9e, 0x65, 0x9c, 0xcd,
	0xf5, 0xf6, 0xf0, 0x89, 0xb6, 0x3d, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xda, 0x35, 0x38,
	0x32, 0x07, 0x00, 0x00,
}
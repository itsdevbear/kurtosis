// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: kurtosis_backend_server_api.proto

package kurtosis_backend_server_rpc_api_bindings

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KurtosisCloudBackendServer_IsAvailable_FullMethodName                 = "/kurtosis_cloud.KurtosisCloudBackendServer/IsAvailable"
	KurtosisCloudBackendServer_GetCloudInstanceConfig_FullMethodName      = "/kurtosis_cloud.KurtosisCloudBackendServer/GetCloudInstanceConfig"
	KurtosisCloudBackendServer_GetOrCreateApiKey_FullMethodName           = "/kurtosis_cloud.KurtosisCloudBackendServer/GetOrCreateApiKey"
	KurtosisCloudBackendServer_GetOrCreateInstance_FullMethodName         = "/kurtosis_cloud.KurtosisCloudBackendServer/GetOrCreateInstance"
	KurtosisCloudBackendServer_GetOrCreatePaymentConfig_FullMethodName    = "/kurtosis_cloud.KurtosisCloudBackendServer/GetOrCreatePaymentConfig"
	KurtosisCloudBackendServer_RefreshDefaultPaymentMethod_FullMethodName = "/kurtosis_cloud.KurtosisCloudBackendServer/RefreshDefaultPaymentMethod"
	KurtosisCloudBackendServer_CancelPaymentSubscription_FullMethodName   = "/kurtosis_cloud.KurtosisCloudBackendServer/CancelPaymentSubscription"
)

// KurtosisCloudBackendServerClient is the client API for KurtosisCloudBackendServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KurtosisCloudBackendServerClient interface {
	IsAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCloudInstanceConfig(ctx context.Context, in *GetCloudInstanceConfigArgs, opts ...grpc.CallOption) (*GetCloudInstanceConfigResponse, error)
	GetOrCreateApiKey(ctx context.Context, in *GetOrCreateApiKeyRequest, opts ...grpc.CallOption) (*GetOrCreateApiKeyResponse, error)
	GetOrCreateInstance(ctx context.Context, in *GetOrCreateInstanceRequest, opts ...grpc.CallOption) (*GetOrCreateInstanceResponse, error)
	GetOrCreatePaymentConfig(ctx context.Context, in *GetOrCreatePaymentConfigArgs, opts ...grpc.CallOption) (*GetOrCreatePaymentConfigResponse, error)
	RefreshDefaultPaymentMethod(ctx context.Context, in *RefreshDefaultPaymentMethodArgs, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelPaymentSubscription(ctx context.Context, in *CancelPaymentSubscriptionArgs, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kurtosisCloudBackendServerClient struct {
	cc grpc.ClientConnInterface
}

func NewKurtosisCloudBackendServerClient(cc grpc.ClientConnInterface) KurtosisCloudBackendServerClient {
	return &kurtosisCloudBackendServerClient{cc}
}

func (c *kurtosisCloudBackendServerClient) IsAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_IsAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisCloudBackendServerClient) GetCloudInstanceConfig(ctx context.Context, in *GetCloudInstanceConfigArgs, opts ...grpc.CallOption) (*GetCloudInstanceConfigResponse, error) {
	out := new(GetCloudInstanceConfigResponse)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_GetCloudInstanceConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisCloudBackendServerClient) GetOrCreateApiKey(ctx context.Context, in *GetOrCreateApiKeyRequest, opts ...grpc.CallOption) (*GetOrCreateApiKeyResponse, error) {
	out := new(GetOrCreateApiKeyResponse)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_GetOrCreateApiKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisCloudBackendServerClient) GetOrCreateInstance(ctx context.Context, in *GetOrCreateInstanceRequest, opts ...grpc.CallOption) (*GetOrCreateInstanceResponse, error) {
	out := new(GetOrCreateInstanceResponse)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_GetOrCreateInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisCloudBackendServerClient) GetOrCreatePaymentConfig(ctx context.Context, in *GetOrCreatePaymentConfigArgs, opts ...grpc.CallOption) (*GetOrCreatePaymentConfigResponse, error) {
	out := new(GetOrCreatePaymentConfigResponse)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_GetOrCreatePaymentConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisCloudBackendServerClient) RefreshDefaultPaymentMethod(ctx context.Context, in *RefreshDefaultPaymentMethodArgs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_RefreshDefaultPaymentMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisCloudBackendServerClient) CancelPaymentSubscription(ctx context.Context, in *CancelPaymentSubscriptionArgs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KurtosisCloudBackendServer_CancelPaymentSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KurtosisCloudBackendServerServer is the server API for KurtosisCloudBackendServer service.
// All implementations should embed UnimplementedKurtosisCloudBackendServerServer
// for forward compatibility
type KurtosisCloudBackendServerServer interface {
	IsAvailable(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetCloudInstanceConfig(context.Context, *GetCloudInstanceConfigArgs) (*GetCloudInstanceConfigResponse, error)
	GetOrCreateApiKey(context.Context, *GetOrCreateApiKeyRequest) (*GetOrCreateApiKeyResponse, error)
	GetOrCreateInstance(context.Context, *GetOrCreateInstanceRequest) (*GetOrCreateInstanceResponse, error)
	GetOrCreatePaymentConfig(context.Context, *GetOrCreatePaymentConfigArgs) (*GetOrCreatePaymentConfigResponse, error)
	RefreshDefaultPaymentMethod(context.Context, *RefreshDefaultPaymentMethodArgs) (*emptypb.Empty, error)
	CancelPaymentSubscription(context.Context, *CancelPaymentSubscriptionArgs) (*emptypb.Empty, error)
}

// UnimplementedKurtosisCloudBackendServerServer should be embedded to have forward compatible implementations.
type UnimplementedKurtosisCloudBackendServerServer struct {
}

func (UnimplementedKurtosisCloudBackendServerServer) IsAvailable(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvailable not implemented")
}
func (UnimplementedKurtosisCloudBackendServerServer) GetCloudInstanceConfig(context.Context, *GetCloudInstanceConfigArgs) (*GetCloudInstanceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudInstanceConfig not implemented")
}
func (UnimplementedKurtosisCloudBackendServerServer) GetOrCreateApiKey(context.Context, *GetOrCreateApiKeyRequest) (*GetOrCreateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateApiKey not implemented")
}
func (UnimplementedKurtosisCloudBackendServerServer) GetOrCreateInstance(context.Context, *GetOrCreateInstanceRequest) (*GetOrCreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateInstance not implemented")
}
func (UnimplementedKurtosisCloudBackendServerServer) GetOrCreatePaymentConfig(context.Context, *GetOrCreatePaymentConfigArgs) (*GetOrCreatePaymentConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreatePaymentConfig not implemented")
}
func (UnimplementedKurtosisCloudBackendServerServer) RefreshDefaultPaymentMethod(context.Context, *RefreshDefaultPaymentMethodArgs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshDefaultPaymentMethod not implemented")
}
func (UnimplementedKurtosisCloudBackendServerServer) CancelPaymentSubscription(context.Context, *CancelPaymentSubscriptionArgs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPaymentSubscription not implemented")
}

// UnsafeKurtosisCloudBackendServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KurtosisCloudBackendServerServer will
// result in compilation errors.
type UnsafeKurtosisCloudBackendServerServer interface {
	mustEmbedUnimplementedKurtosisCloudBackendServerServer()
}

func RegisterKurtosisCloudBackendServerServer(s grpc.ServiceRegistrar, srv KurtosisCloudBackendServerServer) {
	s.RegisterService(&KurtosisCloudBackendServer_ServiceDesc, srv)
}

func _KurtosisCloudBackendServer_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_IsAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).IsAvailable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisCloudBackendServer_GetCloudInstanceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudInstanceConfigArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).GetCloudInstanceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_GetCloudInstanceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).GetCloudInstanceConfig(ctx, req.(*GetCloudInstanceConfigArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisCloudBackendServer_GetOrCreateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).GetOrCreateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_GetOrCreateApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).GetOrCreateApiKey(ctx, req.(*GetOrCreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisCloudBackendServer_GetOrCreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).GetOrCreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_GetOrCreateInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).GetOrCreateInstance(ctx, req.(*GetOrCreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisCloudBackendServer_GetOrCreatePaymentConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreatePaymentConfigArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).GetOrCreatePaymentConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_GetOrCreatePaymentConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).GetOrCreatePaymentConfig(ctx, req.(*GetOrCreatePaymentConfigArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisCloudBackendServer_RefreshDefaultPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshDefaultPaymentMethodArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).RefreshDefaultPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_RefreshDefaultPaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).RefreshDefaultPaymentMethod(ctx, req.(*RefreshDefaultPaymentMethodArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisCloudBackendServer_CancelPaymentSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPaymentSubscriptionArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisCloudBackendServerServer).CancelPaymentSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KurtosisCloudBackendServer_CancelPaymentSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisCloudBackendServerServer).CancelPaymentSubscription(ctx, req.(*CancelPaymentSubscriptionArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// KurtosisCloudBackendServer_ServiceDesc is the grpc.ServiceDesc for KurtosisCloudBackendServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KurtosisCloudBackendServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kurtosis_cloud.KurtosisCloudBackendServer",
	HandlerType: (*KurtosisCloudBackendServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvailable",
			Handler:    _KurtosisCloudBackendServer_IsAvailable_Handler,
		},
		{
			MethodName: "GetCloudInstanceConfig",
			Handler:    _KurtosisCloudBackendServer_GetCloudInstanceConfig_Handler,
		},
		{
			MethodName: "GetOrCreateApiKey",
			Handler:    _KurtosisCloudBackendServer_GetOrCreateApiKey_Handler,
		},
		{
			MethodName: "GetOrCreateInstance",
			Handler:    _KurtosisCloudBackendServer_GetOrCreateInstance_Handler,
		},
		{
			MethodName: "GetOrCreatePaymentConfig",
			Handler:    _KurtosisCloudBackendServer_GetOrCreatePaymentConfig_Handler,
		},
		{
			MethodName: "RefreshDefaultPaymentMethod",
			Handler:    _KurtosisCloudBackendServer_RefreshDefaultPaymentMethod_Handler,
		},
		{
			MethodName: "CancelPaymentSubscription",
			Handler:    _KurtosisCloudBackendServer_CancelPaymentSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kurtosis_backend_server_api.proto",
}

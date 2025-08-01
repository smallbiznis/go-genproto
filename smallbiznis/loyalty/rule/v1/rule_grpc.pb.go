// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: smallbiznis/loyalty/rule/v1/rule.proto

package rulev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Service_CreateEarnRules_FullMethodName   = "/smallbiznis.loyalty.rule.v1.Service/CreateEarnRules"
	Service_ListEarnRules_FullMethodName     = "/smallbiznis.loyalty.rule.v1.Service/ListEarnRules"
	Service_UpdateEarnRules_FullMethodName   = "/smallbiznis.loyalty.rule.v1.Service/UpdateEarnRules"
	Service_EvaluateEarnRuleS_FullMethodName = "/smallbiznis.loyalty.rule.v1.Service/EvaluateEarnRuleS"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateEarnRules(ctx context.Context, in *CreateEarnRuleRequest, opts ...grpc.CallOption) (*EarnRule, error)
	ListEarnRules(ctx context.Context, in *ListEarnRulesRequest, opts ...grpc.CallOption) (*ListEarnRulesResponse, error)
	UpdateEarnRules(ctx context.Context, in *UpdateEarnRuleRequest, opts ...grpc.CallOption) (*EarnRule, error)
	EvaluateEarnRuleS(ctx context.Context, in *EvaluateEarnRulesRequest, opts ...grpc.CallOption) (*EvaluateEarnRulesResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateEarnRules(ctx context.Context, in *CreateEarnRuleRequest, opts ...grpc.CallOption) (*EarnRule, error) {
	out := new(EarnRule)
	err := c.cc.Invoke(ctx, Service_CreateEarnRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListEarnRules(ctx context.Context, in *ListEarnRulesRequest, opts ...grpc.CallOption) (*ListEarnRulesResponse, error) {
	out := new(ListEarnRulesResponse)
	err := c.cc.Invoke(ctx, Service_ListEarnRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateEarnRules(ctx context.Context, in *UpdateEarnRuleRequest, opts ...grpc.CallOption) (*EarnRule, error) {
	out := new(EarnRule)
	err := c.cc.Invoke(ctx, Service_UpdateEarnRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) EvaluateEarnRuleS(ctx context.Context, in *EvaluateEarnRulesRequest, opts ...grpc.CallOption) (*EvaluateEarnRulesResponse, error) {
	out := new(EvaluateEarnRulesResponse)
	err := c.cc.Invoke(ctx, Service_EvaluateEarnRuleS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateEarnRules(context.Context, *CreateEarnRuleRequest) (*EarnRule, error)
	ListEarnRules(context.Context, *ListEarnRulesRequest) (*ListEarnRulesResponse, error)
	UpdateEarnRules(context.Context, *UpdateEarnRuleRequest) (*EarnRule, error)
	EvaluateEarnRuleS(context.Context, *EvaluateEarnRulesRequest) (*EvaluateEarnRulesResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateEarnRules(context.Context, *CreateEarnRuleRequest) (*EarnRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEarnRules not implemented")
}
func (UnimplementedServiceServer) ListEarnRules(context.Context, *ListEarnRulesRequest) (*ListEarnRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEarnRules not implemented")
}
func (UnimplementedServiceServer) UpdateEarnRules(context.Context, *UpdateEarnRuleRequest) (*EarnRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEarnRules not implemented")
}
func (UnimplementedServiceServer) EvaluateEarnRuleS(context.Context, *EvaluateEarnRulesRequest) (*EvaluateEarnRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluateEarnRuleS not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateEarnRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEarnRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateEarnRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateEarnRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateEarnRules(ctx, req.(*CreateEarnRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListEarnRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEarnRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListEarnRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListEarnRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListEarnRules(ctx, req.(*ListEarnRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateEarnRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEarnRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateEarnRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateEarnRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateEarnRules(ctx, req.(*UpdateEarnRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_EvaluateEarnRuleS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateEarnRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).EvaluateEarnRuleS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_EvaluateEarnRuleS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).EvaluateEarnRuleS(ctx, req.(*EvaluateEarnRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.loyalty.rule.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEarnRules",
			Handler:    _Service_CreateEarnRules_Handler,
		},
		{
			MethodName: "ListEarnRules",
			Handler:    _Service_ListEarnRules_Handler,
		},
		{
			MethodName: "UpdateEarnRules",
			Handler:    _Service_UpdateEarnRules_Handler,
		},
		{
			MethodName: "EvaluateEarnRuleS",
			Handler:    _Service_EvaluateEarnRuleS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/loyalty/rule/v1/rule.proto",
}

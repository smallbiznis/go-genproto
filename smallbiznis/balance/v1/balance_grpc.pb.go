// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: smallbiznis/balance/v1/balance.proto

package balance

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BalanceService_GetBalance_FullMethodName        = "/smallbiznis.balance.v1.BalanceService/GetBalance"
	BalanceService_CreateTransaction_FullMethodName = "/smallbiznis.balance.v1.BalanceService/CreateTransaction"
	BalanceService_ListTransaction_FullMethodName   = "/smallbiznis.balance.v1.BalanceService/ListTransaction"
	BalanceService_GetTransaction_FullMethodName    = "/smallbiznis.balance.v1.BalanceService/GetTransaction"
)

// BalanceServiceClient is the client API for BalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceServiceClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	ListTransaction(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*ListTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
}

type balanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceServiceClient(cc grpc.ClientConnInterface) BalanceServiceClient {
	return &balanceServiceClient{cc}
}

func (c *balanceServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, BalanceService_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transaction)
	err := c.cc.Invoke(ctx, BalanceService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceServiceClient) ListTransaction(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*ListTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTransactionResponse)
	err := c.cc.Invoke(ctx, BalanceService_ListTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transaction)
	err := c.cc.Invoke(ctx, BalanceService_GetTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServiceServer is the server API for BalanceService service.
// All implementations must embed UnimplementedBalanceServiceServer
// for forward compatibility.
type BalanceServiceServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error)
	ListTransaction(context.Context, *ListTransactionRequest) (*ListTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	mustEmbedUnimplementedBalanceServiceServer()
}

// UnimplementedBalanceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBalanceServiceServer struct{}

func (UnimplementedBalanceServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBalanceServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedBalanceServiceServer) ListTransaction(context.Context, *ListTransactionRequest) (*ListTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransaction not implemented")
}
func (UnimplementedBalanceServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedBalanceServiceServer) mustEmbedUnimplementedBalanceServiceServer() {}
func (UnimplementedBalanceServiceServer) testEmbeddedByValue()                        {}

// UnsafeBalanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceServiceServer will
// result in compilation errors.
type UnsafeBalanceServiceServer interface {
	mustEmbedUnimplementedBalanceServiceServer()
}

func RegisterBalanceServiceServer(s grpc.ServiceRegistrar, srv BalanceServiceServer) {
	// If the following call pancis, it indicates UnimplementedBalanceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BalanceService_ServiceDesc, srv)
}

func _BalanceService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalanceService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalanceService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceService_ListTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).ListTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalanceService_ListTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).ListTransaction(ctx, req.(*ListTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalanceService_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BalanceService_ServiceDesc is the grpc.ServiceDesc for BalanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.balance.v1.BalanceService",
	HandlerType: (*BalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _BalanceService_GetBalance_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _BalanceService_CreateTransaction_Handler,
		},
		{
			MethodName: "ListTransaction",
			Handler:    _BalanceService_ListTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _BalanceService_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/balance/v1/balance.proto",
}

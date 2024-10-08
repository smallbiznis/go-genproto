// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: smallbiznis/hosting/v1/hosting.proto

package hosting

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
	HostingService_ListDomain_FullMethodName   = "/smallbiznis.hosting.v1.HostingService/ListDomain"
	HostingService_GetDomain_FullMethodName    = "/smallbiznis.hosting.v1.HostingService/GetDomain"
	HostingService_AddDomain_FullMethodName    = "/smallbiznis.hosting.v1.HostingService/AddDomain"
	HostingService_UpdateDomain_FullMethodName = "/smallbiznis.hosting.v1.HostingService/UpdateDomain"
)

// HostingServiceClient is the client API for HostingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostingServiceClient interface {
	ListDomain(ctx context.Context, in *ListDomainRequest, opts ...grpc.CallOption) (*ListDomainResponse, error)
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	AddDomain(ctx context.Context, in *Domain, opts ...grpc.CallOption) (*Domain, error)
	UpdateDomain(ctx context.Context, in *Domain, opts ...grpc.CallOption) (*Domain, error)
}

type hostingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostingServiceClient(cc grpc.ClientConnInterface) HostingServiceClient {
	return &hostingServiceClient{cc}
}

func (c *hostingServiceClient) ListDomain(ctx context.Context, in *ListDomainRequest, opts ...grpc.CallOption) (*ListDomainResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDomainResponse)
	err := c.cc.Invoke(ctx, HostingService_ListDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostingServiceClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Domain)
	err := c.cc.Invoke(ctx, HostingService_GetDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostingServiceClient) AddDomain(ctx context.Context, in *Domain, opts ...grpc.CallOption) (*Domain, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Domain)
	err := c.cc.Invoke(ctx, HostingService_AddDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostingServiceClient) UpdateDomain(ctx context.Context, in *Domain, opts ...grpc.CallOption) (*Domain, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Domain)
	err := c.cc.Invoke(ctx, HostingService_UpdateDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostingServiceServer is the server API for HostingService service.
// All implementations must embed UnimplementedHostingServiceServer
// for forward compatibility.
type HostingServiceServer interface {
	ListDomain(context.Context, *ListDomainRequest) (*ListDomainResponse, error)
	GetDomain(context.Context, *GetDomainRequest) (*Domain, error)
	AddDomain(context.Context, *Domain) (*Domain, error)
	UpdateDomain(context.Context, *Domain) (*Domain, error)
	mustEmbedUnimplementedHostingServiceServer()
}

// UnimplementedHostingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHostingServiceServer struct{}

func (UnimplementedHostingServiceServer) ListDomain(context.Context, *ListDomainRequest) (*ListDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomain not implemented")
}
func (UnimplementedHostingServiceServer) GetDomain(context.Context, *GetDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomain not implemented")
}
func (UnimplementedHostingServiceServer) AddDomain(context.Context, *Domain) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDomain not implemented")
}
func (UnimplementedHostingServiceServer) UpdateDomain(context.Context, *Domain) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}
func (UnimplementedHostingServiceServer) mustEmbedUnimplementedHostingServiceServer() {}
func (UnimplementedHostingServiceServer) testEmbeddedByValue()                        {}

// UnsafeHostingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostingServiceServer will
// result in compilation errors.
type UnsafeHostingServiceServer interface {
	mustEmbedUnimplementedHostingServiceServer()
}

func RegisterHostingServiceServer(s grpc.ServiceRegistrar, srv HostingServiceServer) {
	// If the following call pancis, it indicates UnimplementedHostingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HostingService_ServiceDesc, srv)
}

func _HostingService_ListDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostingServiceServer).ListDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostingService_ListDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostingServiceServer).ListDomain(ctx, req.(*ListDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostingService_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostingServiceServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostingService_GetDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostingServiceServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostingService_AddDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Domain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostingServiceServer).AddDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostingService_AddDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostingServiceServer).AddDomain(ctx, req.(*Domain))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostingService_UpdateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Domain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostingServiceServer).UpdateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostingService_UpdateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostingServiceServer).UpdateDomain(ctx, req.(*Domain))
	}
	return interceptor(ctx, in, info, handler)
}

// HostingService_ServiceDesc is the grpc.ServiceDesc for HostingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.hosting.v1.HostingService",
	HandlerType: (*HostingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDomain",
			Handler:    _HostingService_ListDomain_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _HostingService_GetDomain_Handler,
		},
		{
			MethodName: "AddDomain",
			Handler:    _HostingService_AddDomain_Handler,
		},
		{
			MethodName: "UpdateDomain",
			Handler:    _HostingService_UpdateDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/hosting/v1/hosting.proto",
}

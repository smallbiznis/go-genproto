// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: smallbiznis/notification/v1/notification.proto

package notification

import (
	context "context"
	protobuf "github.com/smallbiznis/go-genproto/smallbiznis/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Service_ListNotification_FullMethodName        = "/smallbiznis.notification.v1.Service/ListNotification"
	Service_GetNotification_FullMethodName         = "/smallbiznis.notification.v1.Service/GetNotification"
	Service_CreateNotification_FullMethodName      = "/smallbiznis.notification.v1.Service/CreateNotification"
	Service_UpdateNotification_FullMethodName      = "/smallbiznis.notification.v1.Service/UpdateNotification"
	Service_DeleteNotification_FullMethodName      = "/smallbiznis.notification.v1.Service/DeleteNotification"
	Service_BatchUpdateNotification_FullMethodName = "/smallbiznis.notification.v1.Service/BatchUpdateNotification"
	Service_BatchDeleteNotification_FullMethodName = "/smallbiznis.notification.v1.Service/BatchDeleteNotification"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	ListNotification(ctx context.Context, in *ListNotificationRequest, opts ...grpc.CallOption) (*ListNotificationResponse, error)
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*Notification, error)
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*Notification, error)
	UpdateNotification(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*Notification, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
	BatchUpdateNotification(ctx context.Context, in *BatchNotificationRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
	BatchDeleteNotification(ctx context.Context, in *BatchNotificationRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ListNotification(ctx context.Context, in *ListNotificationRequest, opts ...grpc.CallOption) (*ListNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotificationResponse)
	err := c.cc.Invoke(ctx, Service_ListNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*Notification, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Notification)
	err := c.cc.Invoke(ctx, Service_GetNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*Notification, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Notification)
	err := c.cc.Invoke(ctx, Service_CreateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateNotification(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*Notification, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Notification)
	err := c.cc.Invoke(ctx, Service_UpdateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BatchUpdateNotification(ctx context.Context, in *BatchNotificationRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, Service_BatchUpdateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BatchDeleteNotification(ctx context.Context, in *BatchNotificationRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, Service_BatchDeleteNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	ListNotification(context.Context, *ListNotificationRequest) (*ListNotificationResponse, error)
	GetNotification(context.Context, *GetNotificationRequest) (*Notification, error)
	CreateNotification(context.Context, *CreateNotificationRequest) (*Notification, error)
	UpdateNotification(context.Context, *Notification) (*Notification, error)
	DeleteNotification(context.Context, *DeleteNotificationRequest) (*protobuf.Empty, error)
	BatchUpdateNotification(context.Context, *BatchNotificationRequest) (*protobuf.Empty, error)
	BatchDeleteNotification(context.Context, *BatchNotificationRequest) (*protobuf.Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) ListNotification(context.Context, *ListNotificationRequest) (*ListNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotification not implemented")
}
func (UnimplementedServiceServer) GetNotification(context.Context, *GetNotificationRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedServiceServer) CreateNotification(context.Context, *CreateNotificationRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedServiceServer) UpdateNotification(context.Context, *Notification) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotification not implemented")
}
func (UnimplementedServiceServer) DeleteNotification(context.Context, *DeleteNotificationRequest) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedServiceServer) BatchUpdateNotification(context.Context, *BatchNotificationRequest) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateNotification not implemented")
}
func (UnimplementedServiceServer) BatchDeleteNotification(context.Context, *BatchNotificationRequest) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteNotification not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}
func (UnimplementedServiceServer) testEmbeddedByValue()                 {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_ListNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListNotification(ctx, req.(*ListNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateNotification(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteNotification(ctx, req.(*DeleteNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BatchUpdateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BatchUpdateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_BatchUpdateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BatchUpdateNotification(ctx, req.(*BatchNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BatchDeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BatchDeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_BatchDeleteNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BatchDeleteNotification(ctx, req.(*BatchNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.notification.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNotification",
			Handler:    _Service_ListNotification_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _Service_GetNotification_Handler,
		},
		{
			MethodName: "CreateNotification",
			Handler:    _Service_CreateNotification_Handler,
		},
		{
			MethodName: "UpdateNotification",
			Handler:    _Service_UpdateNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _Service_DeleteNotification_Handler,
		},
		{
			MethodName: "BatchUpdateNotification",
			Handler:    _Service_BatchUpdateNotification_Handler,
		},
		{
			MethodName: "BatchDeleteNotification",
			Handler:    _Service_BatchDeleteNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/notification/v1/notification.proto",
}

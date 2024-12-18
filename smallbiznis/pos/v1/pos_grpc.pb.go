// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: smallbiznis/pos/v1/pos.proto

package pos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Service_ListStaff_FullMethodName        = "/smallbiznis.pos.v1.Service/ListStaff"
	Service_GetStaff_FullMethodName         = "/smallbiznis.pos.v1.Service/GetStaff"
	Service_AddStaff_FullMethodName         = "/smallbiznis.pos.v1.Service/AddStaff"
	Service_UpdateStaff_FullMethodName      = "/smallbiznis.pos.v1.Service/UpdateStaff"
	Service_DeleteMember_FullMethodName     = "/smallbiznis.pos.v1.Service/DeleteMember"
	Service_LookupPasscode_FullMethodName   = "/smallbiznis.pos.v1.Service/LookupPasscode"
	Service_VerifyPasscode_FullMethodName   = "/smallbiznis.pos.v1.Service/VerifyPasscode"
	Service_RegisterPasscode_FullMethodName = "/smallbiznis.pos.v1.Service/RegisterPasscode"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	ListStaff(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffResponse, error)
	GetStaff(ctx context.Context, in *GetStaffRequest, opts ...grpc.CallOption) (*Staff, error)
	AddStaff(ctx context.Context, in *AddStaffRequest, opts ...grpc.CallOption) (*Staff, error)
	UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*Staff, error)
	DeleteMember(ctx context.Context, in *DeleteStaffRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LookupPasscode(ctx context.Context, in *LookupPasscodeRequest, opts ...grpc.CallOption) (*LookupPasscodeResponse, error)
	VerifyPasscode(ctx context.Context, in *VerifyPasscodeRequest, opts ...grpc.CallOption) (*VerifyPasscodeResponse, error)
	RegisterPasscode(ctx context.Context, in *RegisterPasscodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ListStaff(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStaffResponse)
	err := c.cc.Invoke(ctx, Service_ListStaff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetStaff(ctx context.Context, in *GetStaffRequest, opts ...grpc.CallOption) (*Staff, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Staff)
	err := c.cc.Invoke(ctx, Service_GetStaff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddStaff(ctx context.Context, in *AddStaffRequest, opts ...grpc.CallOption) (*Staff, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Staff)
	err := c.cc.Invoke(ctx, Service_AddStaff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*Staff, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Staff)
	err := c.cc.Invoke(ctx, Service_UpdateStaff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteMember(ctx context.Context, in *DeleteStaffRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LookupPasscode(ctx context.Context, in *LookupPasscodeRequest, opts ...grpc.CallOption) (*LookupPasscodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookupPasscodeResponse)
	err := c.cc.Invoke(ctx, Service_LookupPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VerifyPasscode(ctx context.Context, in *VerifyPasscodeRequest, opts ...grpc.CallOption) (*VerifyPasscodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPasscodeResponse)
	err := c.cc.Invoke(ctx, Service_VerifyPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RegisterPasscode(ctx context.Context, in *RegisterPasscodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_RegisterPasscode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	ListStaff(context.Context, *ListStaffRequest) (*ListStaffResponse, error)
	GetStaff(context.Context, *GetStaffRequest) (*Staff, error)
	AddStaff(context.Context, *AddStaffRequest) (*Staff, error)
	UpdateStaff(context.Context, *UpdateStaffRequest) (*Staff, error)
	DeleteMember(context.Context, *DeleteStaffRequest) (*emptypb.Empty, error)
	LookupPasscode(context.Context, *LookupPasscodeRequest) (*LookupPasscodeResponse, error)
	VerifyPasscode(context.Context, *VerifyPasscodeRequest) (*VerifyPasscodeResponse, error)
	RegisterPasscode(context.Context, *RegisterPasscodeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) ListStaff(context.Context, *ListStaffRequest) (*ListStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStaff not implemented")
}
func (UnimplementedServiceServer) GetStaff(context.Context, *GetStaffRequest) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaff not implemented")
}
func (UnimplementedServiceServer) AddStaff(context.Context, *AddStaffRequest) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStaff not implemented")
}
func (UnimplementedServiceServer) UpdateStaff(context.Context, *UpdateStaffRequest) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaff not implemented")
}
func (UnimplementedServiceServer) DeleteMember(context.Context, *DeleteStaffRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedServiceServer) LookupPasscode(context.Context, *LookupPasscodeRequest) (*LookupPasscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupPasscode not implemented")
}
func (UnimplementedServiceServer) VerifyPasscode(context.Context, *VerifyPasscodeRequest) (*VerifyPasscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPasscode not implemented")
}
func (UnimplementedServiceServer) RegisterPasscode(context.Context, *RegisterPasscodeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPasscode not implemented")
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

func _Service_ListStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListStaff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListStaff(ctx, req.(*ListStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetStaff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetStaff(ctx, req.(*GetStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddStaff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddStaff(ctx, req.(*AddStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateStaff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateStaff(ctx, req.(*UpdateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteMember(ctx, req.(*DeleteStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_LookupPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupPasscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).LookupPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_LookupPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).LookupPasscode(ctx, req.(*LookupPasscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_VerifyPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPasscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).VerifyPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_VerifyPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).VerifyPasscode(ctx, req.(*VerifyPasscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RegisterPasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPasscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RegisterPasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RegisterPasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RegisterPasscode(ctx, req.(*RegisterPasscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.pos.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStaff",
			Handler:    _Service_ListStaff_Handler,
		},
		{
			MethodName: "GetStaff",
			Handler:    _Service_GetStaff_Handler,
		},
		{
			MethodName: "AddStaff",
			Handler:    _Service_AddStaff_Handler,
		},
		{
			MethodName: "UpdateStaff",
			Handler:    _Service_UpdateStaff_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _Service_DeleteMember_Handler,
		},
		{
			MethodName: "LookupPasscode",
			Handler:    _Service_LookupPasscode_Handler,
		},
		{
			MethodName: "VerifyPasscode",
			Handler:    _Service_VerifyPasscode_Handler,
		},
		{
			MethodName: "RegisterPasscode",
			Handler:    _Service_RegisterPasscode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/pos/v1/pos.proto",
}

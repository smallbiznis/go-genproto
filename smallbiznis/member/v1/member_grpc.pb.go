// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: smallbiznis/member/v1/member.proto

package member

import (
	context "context"
	protobuf "github.com/smallbiznis/go-genproto/smallbiznis/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MemberService_ListMember_FullMethodName   = "/smallbiznis.member.v1.MemberService/ListMember"
	MemberService_GetMember_FullMethodName    = "/smallbiznis.member.v1.MemberService/GetMember"
	MemberService_AddMember_FullMethodName    = "/smallbiznis.member.v1.MemberService/AddMember"
	MemberService_UpdateMember_FullMethodName = "/smallbiznis.member.v1.MemberService/UpdateMember"
	MemberService_DeleteMember_FullMethodName = "/smallbiznis.member.v1.MemberService/DeleteMember"
)

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberServiceClient interface {
	ListMember(ctx context.Context, in *ListMemberRequest, opts ...grpc.CallOption) (*ListMemberResponse, error)
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*Member, error)
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*Member, error)
	UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*Member, error)
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) ListMember(ctx context.Context, in *ListMemberRequest, opts ...grpc.CallOption) (*ListMemberResponse, error) {
	out := new(ListMemberResponse)
	err := c.cc.Invoke(ctx, MemberService_ListMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, MemberService_GetMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, MemberService_AddMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, MemberService_UpdateMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, MemberService_DeleteMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations must embed UnimplementedMemberServiceServer
// for forward compatibility
type MemberServiceServer interface {
	ListMember(context.Context, *ListMemberRequest) (*ListMemberResponse, error)
	GetMember(context.Context, *GetMemberRequest) (*Member, error)
	AddMember(context.Context, *AddMemberRequest) (*Member, error)
	UpdateMember(context.Context, *UpdateMemberRequest) (*Member, error)
	DeleteMember(context.Context, *DeleteMemberRequest) (*protobuf.Empty, error)
	mustEmbedUnimplementedMemberServiceServer()
}

// UnimplementedMemberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (UnimplementedMemberServiceServer) ListMember(context.Context, *ListMemberRequest) (*ListMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMember not implemented")
}
func (UnimplementedMemberServiceServer) GetMember(context.Context, *GetMemberRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedMemberServiceServer) AddMember(context.Context, *AddMemberRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedMemberServiceServer) UpdateMember(context.Context, *UpdateMemberRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMember not implemented")
}
func (UnimplementedMemberServiceServer) DeleteMember(context.Context, *DeleteMemberRequest) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedMemberServiceServer) mustEmbedUnimplementedMemberServiceServer() {}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_ListMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).ListMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_ListMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).ListMember(ctx, req.(*ListMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_GetMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_AddMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_UpdateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).UpdateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_UpdateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).UpdateMember(ctx, req.(*UpdateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_DeleteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).DeleteMember(ctx, req.(*DeleteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.member.v1.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMember",
			Handler:    _MemberService_ListMember_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _MemberService_GetMember_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _MemberService_AddMember_Handler,
		},
		{
			MethodName: "UpdateMember",
			Handler:    _MemberService_UpdateMember_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _MemberService_DeleteMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/member/v1/member.proto",
}

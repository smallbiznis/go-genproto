// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: smallbiznis/organization/v1/organization.proto

package organization

import (
	context "context"
	v1 "github.com/smallbiznis/go-genproto/smallbiznis/subscription/v1"
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
	Service_ListShippingRate_FullMethodName    = "/smallbiznis.organization.v1.Service/ListShippingRate"
	Service_GetShippingRate_FullMethodName     = "/smallbiznis.organization.v1.Service/GetShippingRate"
	Service_CreateShippingRate_FullMethodName  = "/smallbiznis.organization.v1.Service/CreateShippingRate"
	Service_ListTaxRule_FullMethodName         = "/smallbiznis.organization.v1.Service/ListTaxRule"
	Service_GetTaxRule_FullMethodName          = "/smallbiznis.organization.v1.Service/GetTaxRule"
	Service_CreateTaxRule_FullMethodName       = "/smallbiznis.organization.v1.Service/CreateTaxRule"
	Service_UpdateTaxRule_FullMethodName       = "/smallbiznis.organization.v1.Service/UpdateTaxRule"
	Service_ListOrg_FullMethodName             = "/smallbiznis.organization.v1.Service/ListOrg"
	Service_GetOrg_FullMethodName              = "/smallbiznis.organization.v1.Service/GetOrg"
	Service_CreateOrg_FullMethodName           = "/smallbiznis.organization.v1.Service/CreateOrg"
	Service_UpdateOrg_FullMethodName           = "/smallbiznis.organization.v1.Service/UpdateOrg"
	Service_DeleteOrg_FullMethodName           = "/smallbiznis.organization.v1.Service/DeleteOrg"
	Service_ListLocation_FullMethodName        = "/smallbiznis.organization.v1.Service/ListLocation"
	Service_GetLocation_FullMethodName         = "/smallbiznis.organization.v1.Service/GetLocation"
	Service_CrateLocation_FullMethodName       = "/smallbiznis.organization.v1.Service/CrateLocation"
	Service_UpdateLocation_FullMethodName      = "/smallbiznis.organization.v1.Service/UpdateLocation"
	Service_GetSubscription_FullMethodName     = "/smallbiznis.organization.v1.Service/GetSubscription"
	Service_CreateBillingPortal_FullMethodName = "/smallbiznis.organization.v1.Service/CreateBillingPortal"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	ListShippingRate(ctx context.Context, in *ListShippingRateRequest, opts ...grpc.CallOption) (*ListShippingRateResponse, error)
	GetShippingRate(ctx context.Context, in *ShippingRate, opts ...grpc.CallOption) (*ShippingRate, error)
	CreateShippingRate(ctx context.Context, in *ShippingRate, opts ...grpc.CallOption) (*ShippingRate, error)
	ListTaxRule(ctx context.Context, in *LisTaxRequest, opts ...grpc.CallOption) (*ListTaxResponse, error)
	GetTaxRule(ctx context.Context, in *TaxRule, opts ...grpc.CallOption) (*TaxRule, error)
	CreateTaxRule(ctx context.Context, in *TaxRule, opts ...grpc.CallOption) (*TaxRule, error)
	UpdateTaxRule(ctx context.Context, in *TaxRule, opts ...grpc.CallOption) (*TaxRule, error)
	ListOrg(ctx context.Context, in *ListOrganizationRequest, opts ...grpc.CallOption) (*ListOrganizationResponse, error)
	GetOrg(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	CreateOrg(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	UpdateOrg(ctx context.Context, in *Organization, opts ...grpc.CallOption) (*Organization, error)
	DeleteOrg(ctx context.Context, in *Organization, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListLocation(ctx context.Context, in *ListLocationRequest, opts ...grpc.CallOption) (*ListLocationResponse, error)
	GetLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	CrateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	GetSubscription(ctx context.Context, in *v1.Subscription, opts ...grpc.CallOption) (*v1.Subscription, error)
	CreateBillingPortal(ctx context.Context, in *CreateBillingRequest, opts ...grpc.CallOption) (*v1.BillingPortalSession, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ListShippingRate(ctx context.Context, in *ListShippingRateRequest, opts ...grpc.CallOption) (*ListShippingRateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListShippingRateResponse)
	err := c.cc.Invoke(ctx, Service_ListShippingRate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetShippingRate(ctx context.Context, in *ShippingRate, opts ...grpc.CallOption) (*ShippingRate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShippingRate)
	err := c.cc.Invoke(ctx, Service_GetShippingRate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateShippingRate(ctx context.Context, in *ShippingRate, opts ...grpc.CallOption) (*ShippingRate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShippingRate)
	err := c.cc.Invoke(ctx, Service_CreateShippingRate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListTaxRule(ctx context.Context, in *LisTaxRequest, opts ...grpc.CallOption) (*ListTaxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTaxResponse)
	err := c.cc.Invoke(ctx, Service_ListTaxRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTaxRule(ctx context.Context, in *TaxRule, opts ...grpc.CallOption) (*TaxRule, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaxRule)
	err := c.cc.Invoke(ctx, Service_GetTaxRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateTaxRule(ctx context.Context, in *TaxRule, opts ...grpc.CallOption) (*TaxRule, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaxRule)
	err := c.cc.Invoke(ctx, Service_CreateTaxRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateTaxRule(ctx context.Context, in *TaxRule, opts ...grpc.CallOption) (*TaxRule, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaxRule)
	err := c.cc.Invoke(ctx, Service_UpdateTaxRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListOrg(ctx context.Context, in *ListOrganizationRequest, opts ...grpc.CallOption) (*ListOrganizationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrganizationResponse)
	err := c.cc.Invoke(ctx, Service_ListOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetOrg(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Organization)
	err := c.cc.Invoke(ctx, Service_GetOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateOrg(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Organization)
	err := c.cc.Invoke(ctx, Service_CreateOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateOrg(ctx context.Context, in *Organization, opts ...grpc.CallOption) (*Organization, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Organization)
	err := c.cc.Invoke(ctx, Service_UpdateOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteOrg(ctx context.Context, in *Organization, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListLocation(ctx context.Context, in *ListLocationRequest, opts ...grpc.CallOption) (*ListLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLocationResponse)
	err := c.cc.Invoke(ctx, Service_ListLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Location)
	err := c.cc.Invoke(ctx, Service_GetLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CrateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Location)
	err := c.cc.Invoke(ctx, Service_CrateLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Location)
	err := c.cc.Invoke(ctx, Service_UpdateLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetSubscription(ctx context.Context, in *v1.Subscription, opts ...grpc.CallOption) (*v1.Subscription, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Subscription)
	err := c.cc.Invoke(ctx, Service_GetSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateBillingPortal(ctx context.Context, in *CreateBillingRequest, opts ...grpc.CallOption) (*v1.BillingPortalSession, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.BillingPortalSession)
	err := c.cc.Invoke(ctx, Service_CreateBillingPortal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	ListShippingRate(context.Context, *ListShippingRateRequest) (*ListShippingRateResponse, error)
	GetShippingRate(context.Context, *ShippingRate) (*ShippingRate, error)
	CreateShippingRate(context.Context, *ShippingRate) (*ShippingRate, error)
	ListTaxRule(context.Context, *LisTaxRequest) (*ListTaxResponse, error)
	GetTaxRule(context.Context, *TaxRule) (*TaxRule, error)
	CreateTaxRule(context.Context, *TaxRule) (*TaxRule, error)
	UpdateTaxRule(context.Context, *TaxRule) (*TaxRule, error)
	ListOrg(context.Context, *ListOrganizationRequest) (*ListOrganizationResponse, error)
	GetOrg(context.Context, *GetOrganizationRequest) (*Organization, error)
	CreateOrg(context.Context, *CreateOrganizationRequest) (*Organization, error)
	UpdateOrg(context.Context, *Organization) (*Organization, error)
	DeleteOrg(context.Context, *Organization) (*emptypb.Empty, error)
	ListLocation(context.Context, *ListLocationRequest) (*ListLocationResponse, error)
	GetLocation(context.Context, *Location) (*Location, error)
	CrateLocation(context.Context, *Location) (*Location, error)
	UpdateLocation(context.Context, *Location) (*Location, error)
	GetSubscription(context.Context, *v1.Subscription) (*v1.Subscription, error)
	CreateBillingPortal(context.Context, *CreateBillingRequest) (*v1.BillingPortalSession, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) ListShippingRate(context.Context, *ListShippingRateRequest) (*ListShippingRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShippingRate not implemented")
}
func (UnimplementedServiceServer) GetShippingRate(context.Context, *ShippingRate) (*ShippingRate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShippingRate not implemented")
}
func (UnimplementedServiceServer) CreateShippingRate(context.Context, *ShippingRate) (*ShippingRate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShippingRate not implemented")
}
func (UnimplementedServiceServer) ListTaxRule(context.Context, *LisTaxRequest) (*ListTaxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaxRule not implemented")
}
func (UnimplementedServiceServer) GetTaxRule(context.Context, *TaxRule) (*TaxRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaxRule not implemented")
}
func (UnimplementedServiceServer) CreateTaxRule(context.Context, *TaxRule) (*TaxRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaxRule not implemented")
}
func (UnimplementedServiceServer) UpdateTaxRule(context.Context, *TaxRule) (*TaxRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaxRule not implemented")
}
func (UnimplementedServiceServer) ListOrg(context.Context, *ListOrganizationRequest) (*ListOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrg not implemented")
}
func (UnimplementedServiceServer) GetOrg(context.Context, *GetOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedServiceServer) CreateOrg(context.Context, *CreateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (UnimplementedServiceServer) UpdateOrg(context.Context, *Organization) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrg not implemented")
}
func (UnimplementedServiceServer) DeleteOrg(context.Context, *Organization) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrg not implemented")
}
func (UnimplementedServiceServer) ListLocation(context.Context, *ListLocationRequest) (*ListLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocation not implemented")
}
func (UnimplementedServiceServer) GetLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedServiceServer) CrateLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrateLocation not implemented")
}
func (UnimplementedServiceServer) UpdateLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedServiceServer) GetSubscription(context.Context, *v1.Subscription) (*v1.Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedServiceServer) CreateBillingPortal(context.Context, *CreateBillingRequest) (*v1.BillingPortalSession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBillingPortal not implemented")
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

func _Service_ListShippingRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShippingRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListShippingRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListShippingRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListShippingRate(ctx, req.(*ListShippingRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetShippingRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingRate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetShippingRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetShippingRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetShippingRate(ctx, req.(*ShippingRate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateShippingRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingRate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateShippingRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateShippingRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateShippingRate(ctx, req.(*ShippingRate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListTaxRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LisTaxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListTaxRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListTaxRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListTaxRule(ctx, req.(*LisTaxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTaxRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaxRule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTaxRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetTaxRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTaxRule(ctx, req.(*TaxRule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateTaxRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaxRule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateTaxRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateTaxRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateTaxRule(ctx, req.(*TaxRule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateTaxRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaxRule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateTaxRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateTaxRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateTaxRule(ctx, req.(*TaxRule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListOrg(ctx, req.(*ListOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetOrg(ctx, req.(*GetOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateOrg(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateOrg(ctx, req.(*Organization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteOrg(ctx, req.(*Organization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListLocation(ctx, req.(*ListLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CrateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CrateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CrateLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CrateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetSubscription(ctx, req.(*v1.Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateBillingPortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateBillingPortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateBillingPortal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateBillingPortal(ctx, req.(*CreateBillingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.organization.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShippingRate",
			Handler:    _Service_ListShippingRate_Handler,
		},
		{
			MethodName: "GetShippingRate",
			Handler:    _Service_GetShippingRate_Handler,
		},
		{
			MethodName: "CreateShippingRate",
			Handler:    _Service_CreateShippingRate_Handler,
		},
		{
			MethodName: "ListTaxRule",
			Handler:    _Service_ListTaxRule_Handler,
		},
		{
			MethodName: "GetTaxRule",
			Handler:    _Service_GetTaxRule_Handler,
		},
		{
			MethodName: "CreateTaxRule",
			Handler:    _Service_CreateTaxRule_Handler,
		},
		{
			MethodName: "UpdateTaxRule",
			Handler:    _Service_UpdateTaxRule_Handler,
		},
		{
			MethodName: "ListOrg",
			Handler:    _Service_ListOrg_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _Service_GetOrg_Handler,
		},
		{
			MethodName: "CreateOrg",
			Handler:    _Service_CreateOrg_Handler,
		},
		{
			MethodName: "UpdateOrg",
			Handler:    _Service_UpdateOrg_Handler,
		},
		{
			MethodName: "DeleteOrg",
			Handler:    _Service_DeleteOrg_Handler,
		},
		{
			MethodName: "ListLocation",
			Handler:    _Service_ListLocation_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Service_GetLocation_Handler,
		},
		{
			MethodName: "CrateLocation",
			Handler:    _Service_CrateLocation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _Service_UpdateLocation_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _Service_GetSubscription_Handler,
		},
		{
			MethodName: "CreateBillingPortal",
			Handler:    _Service_CreateBillingPortal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/organization/v1/organization.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: smallbiznis/subscription/v1/subscription.proto

package subscription

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
	SubscriptionService_CreateBillingPortal_FullMethodName = "/smallbiznis.subscription.v1.SubscriptionService/CreateBillingPortal"
	SubscriptionService_CreateSession_FullMethodName       = "/smallbiznis.subscription.v1.SubscriptionService/CreateSession"
	SubscriptionService_ConfirmationSession_FullMethodName = "/smallbiznis.subscription.v1.SubscriptionService/ConfirmationSession"
	SubscriptionService_CreateCustomer_FullMethodName      = "/smallbiznis.subscription.v1.SubscriptionService/CreateCustomer"
	SubscriptionService_GetCustomer_FullMethodName         = "/smallbiznis.subscription.v1.SubscriptionService/GetCustomer"
	SubscriptionService_ListProduct_FullMethodName         = "/smallbiznis.subscription.v1.SubscriptionService/ListProduct"
	SubscriptionService_GetProduct_FullMethodName          = "/smallbiznis.subscription.v1.SubscriptionService/GetProduct"
	SubscriptionService_UpdateProduct_FullMethodName       = "/smallbiznis.subscription.v1.SubscriptionService/UpdateProduct"
	SubscriptionService_ListPrice_FullMethodName           = "/smallbiznis.subscription.v1.SubscriptionService/ListPrice"
	SubscriptionService_GetPrice_FullMethodName            = "/smallbiznis.subscription.v1.SubscriptionService/GetPrice"
	SubscriptionService_ListSubscription_FullMethodName    = "/smallbiznis.subscription.v1.SubscriptionService/ListSubscription"
	SubscriptionService_CreateSubscription_FullMethodName  = "/smallbiznis.subscription.v1.SubscriptionService/CreateSubscription"
	SubscriptionService_GetSubscription_FullMethodName     = "/smallbiznis.subscription.v1.SubscriptionService/GetSubscription"
	SubscriptionService_UpdateSubscription_FullMethodName  = "/smallbiznis.subscription.v1.SubscriptionService/UpdateSubscription"
	SubscriptionService_DeleteSubscription_FullMethodName  = "/smallbiznis.subscription.v1.SubscriptionService/DeleteSubscription"
	SubscriptionService_ListInvoice_FullMethodName         = "/smallbiznis.subscription.v1.SubscriptionService/ListInvoice"
	SubscriptionService_GetInvoice_FullMethodName          = "/smallbiznis.subscription.v1.SubscriptionService/GetInvoice"
)

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	CreateBillingPortal(ctx context.Context, in *CreateBillingRequest, opts ...grpc.CallOption) (*BillingPortalSession, error)
	CreateSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*Session, error)
	ConfirmationSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*PaymentMethod, error)
	CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	GetCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	GetProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	ListPrice(ctx context.Context, in *ListPriceRequest, opts ...grpc.CallOption) (*ListPriceResponse, error)
	GetPrice(ctx context.Context, in *Price, opts ...grpc.CallOption) (*Price, error)
	ListSubscription(ctx context.Context, in *ListSubscriptionRequest, opts ...grpc.CallOption) (*ListSubscriptionResponse, error)
	CreateSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error)
	GetSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error)
	UpdateSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error)
	DeleteSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error)
	ListInvoice(ctx context.Context, in *ListInvoiceRequest, opts ...grpc.CallOption) (*ListInvoiceResponse, error)
	GetInvoice(ctx context.Context, in *Invoice, opts ...grpc.CallOption) (*Invoice, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) CreateBillingPortal(ctx context.Context, in *CreateBillingRequest, opts ...grpc.CallOption) (*BillingPortalSession, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BillingPortalSession)
	err := c.cc.Invoke(ctx, SubscriptionService_CreateBillingPortal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CreateSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*Session, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Session)
	err := c.cc.Invoke(ctx, SubscriptionService_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ConfirmationSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*PaymentMethod, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentMethod)
	err := c.cc.Invoke(ctx, SubscriptionService_ConfirmationSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Customer)
	err := c.cc.Invoke(ctx, SubscriptionService_CreateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Customer)
	err := c.cc.Invoke(ctx, SubscriptionService_GetCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_ListProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, SubscriptionService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Product)
	err := c.cc.Invoke(ctx, SubscriptionService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListPrice(ctx context.Context, in *ListPriceRequest, opts ...grpc.CallOption) (*ListPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPriceResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_ListPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetPrice(ctx context.Context, in *Price, opts ...grpc.CallOption) (*Price, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Price)
	err := c.cc.Invoke(ctx, SubscriptionService_GetPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListSubscription(ctx context.Context, in *ListSubscriptionRequest, opts ...grpc.CallOption) (*ListSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_ListSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CreateSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Subscription)
	err := c.cc.Invoke(ctx, SubscriptionService_CreateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Subscription)
	err := c.cc.Invoke(ctx, SubscriptionService_GetSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) UpdateSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Subscription)
	err := c.cc.Invoke(ctx, SubscriptionService_UpdateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) DeleteSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Subscription)
	err := c.cc.Invoke(ctx, SubscriptionService_DeleteSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ListInvoice(ctx context.Context, in *ListInvoiceRequest, opts ...grpc.CallOption) (*ListInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInvoiceResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_ListInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetInvoice(ctx context.Context, in *Invoice, opts ...grpc.CallOption) (*Invoice, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Invoice)
	err := c.cc.Invoke(ctx, SubscriptionService_GetInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility.
type SubscriptionServiceServer interface {
	CreateBillingPortal(context.Context, *CreateBillingRequest) (*BillingPortalSession, error)
	CreateSession(context.Context, *SessionRequest) (*Session, error)
	ConfirmationSession(context.Context, *Session) (*PaymentMethod, error)
	CreateCustomer(context.Context, *Customer) (*Customer, error)
	GetCustomer(context.Context, *Customer) (*Customer, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
	GetProduct(context.Context, *Product) (*Product, error)
	UpdateProduct(context.Context, *Product) (*Product, error)
	ListPrice(context.Context, *ListPriceRequest) (*ListPriceResponse, error)
	GetPrice(context.Context, *Price) (*Price, error)
	ListSubscription(context.Context, *ListSubscriptionRequest) (*ListSubscriptionResponse, error)
	CreateSubscription(context.Context, *Subscription) (*Subscription, error)
	GetSubscription(context.Context, *Subscription) (*Subscription, error)
	UpdateSubscription(context.Context, *Subscription) (*Subscription, error)
	DeleteSubscription(context.Context, *Subscription) (*Subscription, error)
	ListInvoice(context.Context, *ListInvoiceRequest) (*ListInvoiceResponse, error)
	GetInvoice(context.Context, *Invoice) (*Invoice, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubscriptionServiceServer struct{}

func (UnimplementedSubscriptionServiceServer) CreateBillingPortal(context.Context, *CreateBillingRequest) (*BillingPortalSession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBillingPortal not implemented")
}
func (UnimplementedSubscriptionServiceServer) CreateSession(context.Context, *SessionRequest) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedSubscriptionServiceServer) ConfirmationSession(context.Context, *Session) (*PaymentMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmationSession not implemented")
}
func (UnimplementedSubscriptionServiceServer) CreateCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedSubscriptionServiceServer) ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetProduct(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedSubscriptionServiceServer) UpdateProduct(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedSubscriptionServiceServer) ListPrice(context.Context, *ListPriceRequest) (*ListPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrice not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetPrice(context.Context, *Price) (*Price, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}
func (UnimplementedSubscriptionServiceServer) ListSubscription(context.Context, *ListSubscriptionRequest) (*ListSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) CreateSubscription(context.Context, *Subscription) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetSubscription(context.Context, *Subscription) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) UpdateSubscription(context.Context, *Subscription) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) DeleteSubscription(context.Context, *Subscription) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) ListInvoice(context.Context, *ListInvoiceRequest) (*ListInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvoice not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetInvoice(context.Context, *Invoice) (*Invoice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}
func (UnimplementedSubscriptionServiceServer) testEmbeddedByValue()                             {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubscriptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_CreateBillingPortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateBillingPortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CreateBillingPortal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateBillingPortal(ctx, req.(*CreateBillingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ConfirmationSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ConfirmationSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ConfirmationSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ConfirmationSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CreateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ListProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ListPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListPrice(ctx, req.(*ListPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Price)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetPrice(ctx, req.(*Price))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ListSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListSubscription(ctx, req.(*ListSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_UpdateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).UpdateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_UpdateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).UpdateSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_DeleteSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ListInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ListInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListInvoice(ctx, req.(*ListInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Invoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetInvoice(ctx, req.(*Invoice))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smallbiznis.subscription.v1.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBillingPortal",
			Handler:    _SubscriptionService_CreateBillingPortal_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _SubscriptionService_CreateSession_Handler,
		},
		{
			MethodName: "ConfirmationSession",
			Handler:    _SubscriptionService_ConfirmationSession_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _SubscriptionService_CreateCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _SubscriptionService_GetCustomer_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _SubscriptionService_ListProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _SubscriptionService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _SubscriptionService_UpdateProduct_Handler,
		},
		{
			MethodName: "ListPrice",
			Handler:    _SubscriptionService_ListPrice_Handler,
		},
		{
			MethodName: "GetPrice",
			Handler:    _SubscriptionService_GetPrice_Handler,
		},
		{
			MethodName: "ListSubscription",
			Handler:    _SubscriptionService_ListSubscription_Handler,
		},
		{
			MethodName: "CreateSubscription",
			Handler:    _SubscriptionService_CreateSubscription_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _SubscriptionService_GetSubscription_Handler,
		},
		{
			MethodName: "UpdateSubscription",
			Handler:    _SubscriptionService_UpdateSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _SubscriptionService_DeleteSubscription_Handler,
		},
		{
			MethodName: "ListInvoice",
			Handler:    _SubscriptionService_ListInvoice_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _SubscriptionService_GetInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smallbiznis/subscription/v1/subscription.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/subscription/v1/subscription_resource.proto

package subscription

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlanInterval int32

const (
	PlanInterval_day   PlanInterval = 0
	PlanInterval_week  PlanInterval = 1
	PlanInterval_month PlanInterval = 2
	PlanInterval_year  PlanInterval = 3
)

// Enum value maps for PlanInterval.
var (
	PlanInterval_name = map[int32]string{
		0: "day",
		1: "week",
		2: "month",
		3: "year",
	}
	PlanInterval_value = map[string]int32{
		"day":   0,
		"week":  1,
		"month": 2,
		"year":  3,
	}
)

func (x PlanInterval) Enum() *PlanInterval {
	p := new(PlanInterval)
	*p = x
	return p
}

func (x PlanInterval) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanInterval) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[0].Descriptor()
}

func (PlanInterval) Type() protoreflect.EnumType {
	return &file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[0]
}

func (x PlanInterval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanInterval.Descriptor instead.
func (PlanInterval) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Currency     string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Session) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City       string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Country    string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Line1      string `protobuf:"bytes,3,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2      string `protobuf:"bytes,4,opt,name=line2,proto3" json:"line2,omitempty"`
	PostalCode string `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetLine1() string {
	if x != nil {
		return x.Line1
	}
	return ""
}

func (x *Address) GetLine2() string {
	if x != nil {
		return x.Line2
	}
	return ""
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type BillingDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Email   string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone   string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *BillingDetail) Reset() {
	*x = BillingDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingDetail) ProtoMessage() {}

func (x *BillingDetail) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingDetail.ProtoReflect.Descriptor instead.
func (*BillingDetail) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{3}
}

func (x *BillingDetail) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *BillingDetail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BillingDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BillingDetail) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CardDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number   string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ExpMonth int32  `protobuf:"varint,2,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear  int32  `protobuf:"varint,3,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvv      string `protobuf:"bytes,4,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *CardDetail) Reset() {
	*x = CardDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardDetail) ProtoMessage() {}

func (x *CardDetail) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardDetail.ProtoReflect.Descriptor instead.
func (*CardDetail) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{4}
}

func (x *CardDetail) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CardDetail) GetExpMonth() int32 {
	if x != nil {
		return x.ExpMonth
	}
	return 0
}

func (x *CardDetail) GetExpYear() int32 {
	if x != nil {
		return x.ExpYear
	}
	return 0
}

func (x *CardDetail) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

type PaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId    string         `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	BillingDetail *BillingDetail `protobuf:"bytes,2,opt,name=billing_detail,json=billingDetail,proto3" json:"billing_detail,omitempty"`
	Card          *CardDetail    `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	Token         string         `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PaymentMethod) Reset() {
	*x = PaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod) ProtoMessage() {}

func (x *PaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod.ProtoReflect.Descriptor instead.
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentMethod) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *PaymentMethod) GetBillingDetail() *BillingDetail {
	if x != nil {
		return x.BillingDetail
	}
	return nil
}

func (x *PaymentMethod) GetCard() *CardDetail {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *PaymentMethod) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Images      []string               `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Active      bool                   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{6}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product       *Product               `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Interval      string                 `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	IntervalCount int32                  `protobuf:"varint,4,opt,name=interval_count,json=intervalCount,proto3" json:"interval_count,omitempty"`
	Amount        float32                `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{7}
}

func (x *Plan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plan) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *Plan) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *Plan) GetIntervalCount() int32 {
	if x != nil {
		return x.IntervalCount
	}
	return 0
}

func (x *Plan) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Plan) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string                 `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PlanId     string                 `protobuf:"bytes,3,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	StartDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Status     string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{8}
}

func (x *Subscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subscription) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Subscription) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *Subscription) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Subscription) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Subscription) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Subscription) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_smallbiznis_subscription_v1_subscription_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_subscription_v1_subscription_resource_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x44, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x32,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0d, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3e, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x6e, 0x0a, 0x0a, 0x43, 0x61, 0x72,
	0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x78, 0x70, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x78, 0x70, 0x59, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x76, 0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x0e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x3b, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xf5, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x04, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9d, 0x02, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x36, 0x0a, 0x0c, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x10,
	0x03, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescOnce sync.Once
	file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescData = file_smallbiznis_subscription_v1_subscription_resource_proto_rawDesc
)

func file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP() []byte {
	file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescOnce.Do(func() {
		file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescData)
	})
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescData
}

var file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_smallbiznis_subscription_v1_subscription_resource_proto_goTypes = []interface{}{
	(PlanInterval)(0),             // 0: smallbiznis.subscription.v1.PlanInterval
	(*Session)(nil),               // 1: smallbiznis.subscription.v1.Session
	(*Customer)(nil),              // 2: smallbiznis.subscription.v1.Customer
	(*Address)(nil),               // 3: smallbiznis.subscription.v1.Address
	(*BillingDetail)(nil),         // 4: smallbiznis.subscription.v1.BillingDetail
	(*CardDetail)(nil),            // 5: smallbiznis.subscription.v1.CardDetail
	(*PaymentMethod)(nil),         // 6: smallbiznis.subscription.v1.PaymentMethod
	(*Product)(nil),               // 7: smallbiznis.subscription.v1.Product
	(*Plan)(nil),                  // 8: smallbiznis.subscription.v1.Plan
	(*Subscription)(nil),          // 9: smallbiznis.subscription.v1.Subscription
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_smallbiznis_subscription_v1_subscription_resource_proto_depIdxs = []int32{
	3,  // 0: smallbiznis.subscription.v1.BillingDetail.address:type_name -> smallbiznis.subscription.v1.Address
	4,  // 1: smallbiznis.subscription.v1.PaymentMethod.billing_detail:type_name -> smallbiznis.subscription.v1.BillingDetail
	5,  // 2: smallbiznis.subscription.v1.PaymentMethod.card:type_name -> smallbiznis.subscription.v1.CardDetail
	10, // 3: smallbiznis.subscription.v1.Product.created_at:type_name -> google.protobuf.Timestamp
	10, // 4: smallbiznis.subscription.v1.Product.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 5: smallbiznis.subscription.v1.Plan.product:type_name -> smallbiznis.subscription.v1.Product
	10, // 6: smallbiznis.subscription.v1.Plan.created_at:type_name -> google.protobuf.Timestamp
	10, // 7: smallbiznis.subscription.v1.Subscription.start_date:type_name -> google.protobuf.Timestamp
	10, // 8: smallbiznis.subscription.v1.Subscription.end_date:type_name -> google.protobuf.Timestamp
	10, // 9: smallbiznis.subscription.v1.Subscription.created_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_smallbiznis_subscription_v1_subscription_resource_proto_init() }
func file_smallbiznis_subscription_v1_subscription_resource_proto_init() {
	if File_smallbiznis_subscription_v1_subscription_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_smallbiznis_subscription_v1_subscription_resource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_subscription_v1_subscription_resource_proto_goTypes,
		DependencyIndexes: file_smallbiznis_subscription_v1_subscription_resource_proto_depIdxs,
		EnumInfos:         file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes,
		MessageInfos:      file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes,
	}.Build()
	File_smallbiznis_subscription_v1_subscription_resource_proto = out.File
	file_smallbiznis_subscription_v1_subscription_resource_proto_rawDesc = nil
	file_smallbiznis_subscription_v1_subscription_resource_proto_goTypes = nil
	file_smallbiznis_subscription_v1_subscription_resource_proto_depIdxs = nil
}
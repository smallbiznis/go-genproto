// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/subscription/v1/subscription_resource.proto

package subscription

import (
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

type PricingPlan_Frequency int32

const (
	PricingPlan_MONTHLY PricingPlan_Frequency = 0
	PricingPlan_YEARLY  PricingPlan_Frequency = 1
)

// Enum value maps for PricingPlan_Frequency.
var (
	PricingPlan_Frequency_name = map[int32]string{
		0: "MONTHLY",
		1: "YEARLY",
	}
	PricingPlan_Frequency_value = map[string]int32{
		"MONTHLY": 0,
		"YEARLY":  1,
	}
)

func (x PricingPlan_Frequency) Enum() *PricingPlan_Frequency {
	p := new(PricingPlan_Frequency)
	*p = x
	return p
}

func (x PricingPlan_Frequency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PricingPlan_Frequency) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[0].Descriptor()
}

func (PricingPlan_Frequency) Type() protoreflect.EnumType {
	return &file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[0]
}

func (x PricingPlan_Frequency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PricingPlan_Frequency.Descriptor instead.
func (PricingPlan_Frequency) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{1, 0}
}

type PricingPlan_PlanCost_Name int32

const (
	PricingPlan_PlanCost_USD PricingPlan_PlanCost_Name = 0
	PricingPlan_PlanCost_IDR PricingPlan_PlanCost_Name = 1
)

// Enum value maps for PricingPlan_PlanCost_Name.
var (
	PricingPlan_PlanCost_Name_name = map[int32]string{
		0: "USD",
		1: "IDR",
	}
	PricingPlan_PlanCost_Name_value = map[string]int32{
		"USD": 0,
		"IDR": 1,
	}
)

func (x PricingPlan_PlanCost_Name) Enum() *PricingPlan_PlanCost_Name {
	p := new(PricingPlan_PlanCost_Name)
	*p = x
	return p
}

func (x PricingPlan_PlanCost_Name) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PricingPlan_PlanCost_Name) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[1].Descriptor()
}

func (PricingPlan_PlanCost_Name) Type() protoreflect.EnumType {
	return &file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[1]
}

func (x PricingPlan_PlanCost_Name) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PricingPlan_PlanCost_Name.Descriptor instead.
func (PricingPlan_PlanCost_Name) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{1, 0, 0}
}

type Subscription_Status int32

const (
	Subscription_INITIALIZED        Subscription_Status = 0
	Subscription_FREE_TRIAL         Subscription_Status = 1
	Subscription_FREE_TRIAL_EXPIRED Subscription_Status = 2
	Subscription_ACTIVE             Subscription_Status = 3
	Subscription_SUSPENDED          Subscription_Status = 4
	Subscription_CANCELLED          Subscription_Status = 5
)

// Enum value maps for Subscription_Status.
var (
	Subscription_Status_name = map[int32]string{
		0: "INITIALIZED",
		1: "FREE_TRIAL",
		2: "FREE_TRIAL_EXPIRED",
		3: "ACTIVE",
		4: "SUSPENDED",
		5: "CANCELLED",
	}
	Subscription_Status_value = map[string]int32{
		"INITIALIZED":        0,
		"FREE_TRIAL":         1,
		"FREE_TRIAL_EXPIRED": 2,
		"ACTIVE":             3,
		"SUSPENDED":          4,
		"CANCELLED":          5,
	}
)

func (x Subscription_Status) Enum() *Subscription_Status {
	p := new(Subscription_Status)
	*p = x
	return p
}

func (x Subscription_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subscription_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[2].Descriptor()
}

func (Subscription_Status) Type() protoreflect.EnumType {
	return &file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes[2]
}

func (x Subscription_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subscription_Status.Descriptor instead.
func (Subscription_Status) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{2, 0}
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId            string         `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Name              string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description       string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DefaultFreeTrial  bool           `protobuf:"varint,4,opt,name=default_free_trial,json=defaultFreeTrial,proto3" json:"default_free_trial,omitempty"`
	FreeTrialDuration int64          `protobuf:"varint,6,opt,name=free_trial_duration,json=freeTrialDuration,proto3" json:"free_trial_duration,omitempty"`
	FreeTrialUnit     string         `protobuf:"bytes,7,opt,name=free_trial_unit,json=freeTrialUnit,proto3" json:"free_trial_unit,omitempty"`
	Prices            []*PricingPlan `protobuf:"bytes,8,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Plan) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Plan) GetDefaultFreeTrial() bool {
	if x != nil {
		return x.DefaultFreeTrial
	}
	return false
}

func (x *Plan) GetFreeTrialDuration() int64 {
	if x != nil {
		return x.FreeTrialDuration
	}
	return 0
}

func (x *Plan) GetFreeTrialUnit() string {
	if x != nil {
		return x.FreeTrialUnit
	}
	return ""
}

func (x *Plan) GetPrices() []*PricingPlan {
	if x != nil {
		return x.Prices
	}
	return nil
}

type PricingPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PricingPlanId string                  `protobuf:"bytes,1,opt,name=pricing_plan_id,json=pricingPlanId,proto3" json:"pricing_plan_id,omitempty"`
	Plan          *Plan                   `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Frequency     PricingPlan_Frequency   `protobuf:"varint,3,opt,name=frequency,proto3,enum=smallbiznis.subscription.v1.PricingPlan_Frequency" json:"frequency,omitempty"`
	Costs         []*PricingPlan_PlanCost `protobuf:"bytes,5,rep,name=costs,proto3" json:"costs,omitempty"`
	CreatedAt     *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PricingPlan) Reset() {
	*x = PricingPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingPlan) ProtoMessage() {}

func (x *PricingPlan) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PricingPlan.ProtoReflect.Descriptor instead.
func (*PricingPlan) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{1}
}

func (x *PricingPlan) GetPricingPlanId() string {
	if x != nil {
		return x.PricingPlanId
	}
	return ""
}

func (x *PricingPlan) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *PricingPlan) GetFrequency() PricingPlan_Frequency {
	if x != nil {
		return x.Frequency
	}
	return PricingPlan_MONTHLY
}

func (x *PricingPlan) GetCosts() []*PricingPlan_PlanCost {
	if x != nil {
		return x.Costs
	}
	return nil
}

func (x *PricingPlan) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PricingPlan) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string                 `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	StoreId        string                 `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	PlanId         string                 `protobuf:"bytes,3,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	EndDate        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Status         Subscription_Status    `protobuf:"varint,5,opt,name=status,proto3,enum=smallbiznis.subscription.v1.Subscription_Status" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{2}
}

func (x *Subscription) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *Subscription) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Subscription) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *Subscription) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Subscription) GetStatus() Subscription_Status {
	if x != nil {
		return x.Status
	}
	return Subscription_INITIALIZED
}

func (x *Subscription) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Subscription) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PricingPlan_PlanCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   PricingPlan_PlanCost_Name `protobuf:"varint,1,opt,name=name,proto3,enum=smallbiznis.subscription.v1.PricingPlan_PlanCost_Name" json:"name,omitempty"`
	Amount float32                   `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PricingPlan_PlanCost) Reset() {
	*x = PricingPlan_PlanCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingPlan_PlanCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingPlan_PlanCost) ProtoMessage() {}

func (x *PricingPlan_PlanCost) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PricingPlan_PlanCost.ProtoReflect.Descriptor instead.
func (*PricingPlan_PlanCost) Descriptor() ([]byte, []int) {
	return file_smallbiznis_subscription_v1_subscription_resource_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PricingPlan_PlanCost) GetName() PricingPlan_PlanCost_Name {
	if x != nil {
		return x.Name
	}
	return PricingPlan_PlanCost_USD
}

func (x *PricingPlan_PlanCost) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_smallbiznis_subscription_v1_subscription_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_subscription_v1_subscription_resource_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x2e, 0x0a,
	0x13, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x72, 0x65, 0x65,
	0x54, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x54, 0x72, 0x69, 0x61,
	0x6c, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0xae, 0x04, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x50, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x6e, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x66,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x47, 0x0a, 0x05, 0x63, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61,
	0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x88, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x43, 0x6f, 0x73, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x55, 0x53, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x44, 0x52,
	0x10, 0x01, 0x22, 0x24, 0x0a, 0x09, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x59, 0x45, 0x41, 0x52, 0x4c, 0x59, 0x10, 0x01, 0x22, 0xcf, 0x03, 0x0a, 0x0c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6b, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x49, 0x54, 0x49,
	0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x52, 0x45, 0x45,
	0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x52, 0x45, 0x45,
	0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69,
	0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_smallbiznis_subscription_v1_subscription_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_smallbiznis_subscription_v1_subscription_resource_proto_goTypes = []interface{}{
	(PricingPlan_Frequency)(0),     // 0: smallbiznis.subscription.v1.PricingPlan.Frequency
	(PricingPlan_PlanCost_Name)(0), // 1: smallbiznis.subscription.v1.PricingPlan.PlanCost.Name
	(Subscription_Status)(0),       // 2: smallbiznis.subscription.v1.Subscription.Status
	(*Plan)(nil),                   // 3: smallbiznis.subscription.v1.Plan
	(*PricingPlan)(nil),            // 4: smallbiznis.subscription.v1.PricingPlan
	(*Subscription)(nil),           // 5: smallbiznis.subscription.v1.Subscription
	(*PricingPlan_PlanCost)(nil),   // 6: smallbiznis.subscription.v1.PricingPlan.PlanCost
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
}
var file_smallbiznis_subscription_v1_subscription_resource_proto_depIdxs = []int32{
	4,  // 0: smallbiznis.subscription.v1.Plan.prices:type_name -> smallbiznis.subscription.v1.PricingPlan
	3,  // 1: smallbiznis.subscription.v1.PricingPlan.plan:type_name -> smallbiznis.subscription.v1.Plan
	0,  // 2: smallbiznis.subscription.v1.PricingPlan.frequency:type_name -> smallbiznis.subscription.v1.PricingPlan.Frequency
	6,  // 3: smallbiznis.subscription.v1.PricingPlan.costs:type_name -> smallbiznis.subscription.v1.PricingPlan.PlanCost
	7,  // 4: smallbiznis.subscription.v1.PricingPlan.created_at:type_name -> google.protobuf.Timestamp
	7,  // 5: smallbiznis.subscription.v1.PricingPlan.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 6: smallbiznis.subscription.v1.Subscription.end_date:type_name -> google.protobuf.Timestamp
	2,  // 7: smallbiznis.subscription.v1.Subscription.status:type_name -> smallbiznis.subscription.v1.Subscription.Status
	7,  // 8: smallbiznis.subscription.v1.Subscription.created_at:type_name -> google.protobuf.Timestamp
	7,  // 9: smallbiznis.subscription.v1.Subscription.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 10: smallbiznis.subscription.v1.PricingPlan.PlanCost.name:type_name -> smallbiznis.subscription.v1.PricingPlan.PlanCost.Name
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_smallbiznis_subscription_v1_subscription_resource_proto_init() }
func file_smallbiznis_subscription_v1_subscription_resource_proto_init() {
	if File_smallbiznis_subscription_v1_subscription_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingPlan); i {
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
		file_smallbiznis_subscription_v1_subscription_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingPlan_PlanCost); i {
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
			NumEnums:      3,
			NumMessages:   4,
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

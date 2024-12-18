// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/discount/v1/discount_resource.proto

package discount

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
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

type Type int32

const (
	Type_fixed      Type = 0
	Type_percentage Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "fixed",
		1: "percentage",
	}
	Type_value = map[string]int32{
		"fixed":      0,
		"percentage": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_discount_v1_discount_resource_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_smallbiznis_discount_v1_discount_resource_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_discount_v1_discount_resource_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_active   Status = 0
	Status_draft    Status = 1
	Status_archived Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "active",
		1: "draft",
		2: "archived",
	}
	Status_value = map[string]int32{
		"active":   0,
		"draft":    1,
		"archived": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_discount_v1_discount_resource_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_smallbiznis_discount_v1_discount_resource_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_discount_v1_discount_resource_proto_rawDescGZIP(), []int{1}
}

type Discount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountId     string                 `protobuf:"bytes,1,opt,name=discount_id,json=discountId,proto3" json:"discount_id,omitempty"`
	OrganizationId string                 `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type           Type                   `protobuf:"varint,4,opt,name=type,proto3,enum=smallbiznis.discount.v1.Type" json:"type,omitempty"`
	Value          float32                `protobuf:"fixed32,5,opt,name=value,proto3" json:"value,omitempty"`
	Status         Status                 `protobuf:"varint,6,opt,name=status,proto3,enum=smallbiznis.discount.v1.Status" json:"status,omitempty"`
	StartDate      *date.Date             `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate        *date.Date             `protobuf:"bytes,8,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Rules          []*DiscountRule        `protobuf:"bytes,9,rep,name=rules,proto3" json:"rules,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Discount) Reset() {
	*x = Discount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discount) ProtoMessage() {}

func (x *Discount) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discount.ProtoReflect.Descriptor instead.
func (*Discount) Descriptor() ([]byte, []int) {
	return file_smallbiznis_discount_v1_discount_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Discount) GetDiscountId() string {
	if x != nil {
		return x.DiscountId
	}
	return ""
}

func (x *Discount) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Discount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Discount) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_fixed
}

func (x *Discount) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Discount) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_active
}

func (x *Discount) GetStartDate() *date.Date {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Discount) GetEndDate() *date.Date {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Discount) GetRules() []*DiscountRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Discount) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Discount) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DiscountRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountRuleId string                 `protobuf:"bytes,1,opt,name=discount_rule_id,json=discountRuleId,proto3" json:"discount_rule_id,omitempty"`
	ParentId       string                 `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	RuleType       string                 `protobuf:"bytes,3,opt,name=rule_type,json=ruleType,proto3" json:"rule_type,omitempty"`
	RuleValue      string                 `protobuf:"bytes,4,opt,name=rule_value,json=ruleValue,proto3" json:"rule_value,omitempty"`
	Operator       string                 `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	AdditionalInfo string                 `protobuf:"bytes,6,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Conditions     []*DiscountConditions  `protobuf:"bytes,9,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *DiscountRule) Reset() {
	*x = DiscountRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountRule) ProtoMessage() {}

func (x *DiscountRule) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountRule.ProtoReflect.Descriptor instead.
func (*DiscountRule) Descriptor() ([]byte, []int) {
	return file_smallbiznis_discount_v1_discount_resource_proto_rawDescGZIP(), []int{1}
}

func (x *DiscountRule) GetDiscountRuleId() string {
	if x != nil {
		return x.DiscountRuleId
	}
	return ""
}

func (x *DiscountRule) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *DiscountRule) GetRuleType() string {
	if x != nil {
		return x.RuleType
	}
	return ""
}

func (x *DiscountRule) GetRuleValue() string {
	if x != nil {
		return x.RuleValue
	}
	return ""
}

func (x *DiscountRule) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *DiscountRule) GetAdditionalInfo() string {
	if x != nil {
		return x.AdditionalInfo
	}
	return ""
}

func (x *DiscountRule) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DiscountRule) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DiscountRule) GetConditions() []*DiscountConditions {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type DiscountConditions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountConditionId string                 `protobuf:"bytes,1,opt,name=discount_condition_id,json=discountConditionId,proto3" json:"discount_condition_id,omitempty"`
	RuleId              string                 `protobuf:"bytes,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	ConditionKey        string                 `protobuf:"bytes,3,opt,name=condition_key,json=conditionKey,proto3" json:"condition_key,omitempty"`
	ConditionValue      string                 `protobuf:"bytes,4,opt,name=condition_value,json=conditionValue,proto3" json:"condition_value,omitempty"`
	CreatedAt           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt           *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DiscountConditions) Reset() {
	*x = DiscountConditions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountConditions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountConditions) ProtoMessage() {}

func (x *DiscountConditions) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountConditions.ProtoReflect.Descriptor instead.
func (*DiscountConditions) Descriptor() ([]byte, []int) {
	return file_smallbiznis_discount_v1_discount_resource_proto_rawDescGZIP(), []int{2}
}

func (x *DiscountConditions) GetDiscountConditionId() string {
	if x != nil {
		return x.DiscountConditionId
	}
	return ""
}

func (x *DiscountConditions) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *DiscountConditions) GetConditionKey() string {
	if x != nil {
		return x.ConditionKey
	}
	return ""
}

func (x *DiscountConditions) GetConditionValue() string {
	if x != nil {
		return x.ConditionValue
	}
	return ""
}

func (x *DiscountConditions) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DiscountConditions) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_smallbiznis_discount_v1_discount_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_discount_v1_discount_resource_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x99, 0x03, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xa5, 0x02, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x10, 0x01, 0x2a, 0x2d, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x10, 0x02, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69,
	0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_discount_v1_discount_resource_proto_rawDescOnce sync.Once
	file_smallbiznis_discount_v1_discount_resource_proto_rawDescData = file_smallbiznis_discount_v1_discount_resource_proto_rawDesc
)

func file_smallbiznis_discount_v1_discount_resource_proto_rawDescGZIP() []byte {
	file_smallbiznis_discount_v1_discount_resource_proto_rawDescOnce.Do(func() {
		file_smallbiznis_discount_v1_discount_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_discount_v1_discount_resource_proto_rawDescData)
	})
	return file_smallbiznis_discount_v1_discount_resource_proto_rawDescData
}

var file_smallbiznis_discount_v1_discount_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_smallbiznis_discount_v1_discount_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_smallbiznis_discount_v1_discount_resource_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: smallbiznis.discount.v1.Type
	(Status)(0),                   // 1: smallbiznis.discount.v1.Status
	(*Discount)(nil),              // 2: smallbiznis.discount.v1.Discount
	(*DiscountRule)(nil),          // 3: smallbiznis.discount.v1.DiscountRule
	(*DiscountConditions)(nil),    // 4: smallbiznis.discount.v1.DiscountConditions
	(*date.Date)(nil),             // 5: google.type.Date
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_smallbiznis_discount_v1_discount_resource_proto_depIdxs = []int32{
	0,  // 0: smallbiznis.discount.v1.Discount.type:type_name -> smallbiznis.discount.v1.Type
	1,  // 1: smallbiznis.discount.v1.Discount.status:type_name -> smallbiznis.discount.v1.Status
	5,  // 2: smallbiznis.discount.v1.Discount.start_date:type_name -> google.type.Date
	5,  // 3: smallbiznis.discount.v1.Discount.end_date:type_name -> google.type.Date
	3,  // 4: smallbiznis.discount.v1.Discount.rules:type_name -> smallbiznis.discount.v1.DiscountRule
	6,  // 5: smallbiznis.discount.v1.Discount.created_at:type_name -> google.protobuf.Timestamp
	6,  // 6: smallbiznis.discount.v1.Discount.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 7: smallbiznis.discount.v1.DiscountRule.created_at:type_name -> google.protobuf.Timestamp
	6,  // 8: smallbiznis.discount.v1.DiscountRule.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 9: smallbiznis.discount.v1.DiscountRule.conditions:type_name -> smallbiznis.discount.v1.DiscountConditions
	6,  // 10: smallbiznis.discount.v1.DiscountConditions.created_at:type_name -> google.protobuf.Timestamp
	6,  // 11: smallbiznis.discount.v1.DiscountConditions.updated_at:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_smallbiznis_discount_v1_discount_resource_proto_init() }
func file_smallbiznis_discount_v1_discount_resource_proto_init() {
	if File_smallbiznis_discount_v1_discount_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Discount); i {
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
		file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountRule); i {
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
		file_smallbiznis_discount_v1_discount_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountConditions); i {
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
			RawDescriptor: file_smallbiznis_discount_v1_discount_resource_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_discount_v1_discount_resource_proto_goTypes,
		DependencyIndexes: file_smallbiznis_discount_v1_discount_resource_proto_depIdxs,
		EnumInfos:         file_smallbiznis_discount_v1_discount_resource_proto_enumTypes,
		MessageInfos:      file_smallbiznis_discount_v1_discount_resource_proto_msgTypes,
	}.Build()
	File_smallbiznis_discount_v1_discount_resource_proto = out.File
	file_smallbiznis_discount_v1_discount_resource_proto_rawDesc = nil
	file_smallbiznis_discount_v1_discount_resource_proto_goTypes = nil
	file_smallbiznis_discount_v1_discount_resource_proto_depIdxs = nil
}

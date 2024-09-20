// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/rules/v1/rules.proto

package rules

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

type LisTaxRequest_OrderBy int32

const (
	LisTaxRequest_ASC  LisTaxRequest_OrderBy = 0
	LisTaxRequest_DESC LisTaxRequest_OrderBy = 1
)

// Enum value maps for LisTaxRequest_OrderBy.
var (
	LisTaxRequest_OrderBy_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	LisTaxRequest_OrderBy_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x LisTaxRequest_OrderBy) Enum() *LisTaxRequest_OrderBy {
	p := new(LisTaxRequest_OrderBy)
	*p = x
	return p
}

func (x LisTaxRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LisTaxRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_rules_v1_rules_proto_enumTypes[0].Descriptor()
}

func (LisTaxRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_smallbiznis_rules_v1_rules_proto_enumTypes[0]
}

func (x LisTaxRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LisTaxRequest_OrderBy.Descriptor instead.
func (LisTaxRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_rules_v1_rules_proto_rawDescGZIP(), []int{0, 0}
}

type ListServiceFeeRequest_OrderBy int32

const (
	ListServiceFeeRequest_ASC  ListServiceFeeRequest_OrderBy = 0
	ListServiceFeeRequest_DESC ListServiceFeeRequest_OrderBy = 1
)

// Enum value maps for ListServiceFeeRequest_OrderBy.
var (
	ListServiceFeeRequest_OrderBy_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	ListServiceFeeRequest_OrderBy_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x ListServiceFeeRequest_OrderBy) Enum() *ListServiceFeeRequest_OrderBy {
	p := new(ListServiceFeeRequest_OrderBy)
	*p = x
	return p
}

func (x ListServiceFeeRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListServiceFeeRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_rules_v1_rules_proto_enumTypes[1].Descriptor()
}

func (ListServiceFeeRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_smallbiznis_rules_v1_rules_proto_enumTypes[1]
}

func (x ListServiceFeeRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListServiceFeeRequest_OrderBy.Descriptor instead.
func (ListServiceFeeRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_rules_v1_rules_proto_rawDescGZIP(), []int{2, 0}
}

type LisTaxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size       int32                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	SortBy     string                 `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy    LisTaxRequest_OrderBy  `protobuf:"varint,4,opt,name=order_by,json=orderBy,proto3,enum=smallbiznis.rules.v1.LisTaxRequest_OrderBy" json:"order_by,omitempty"`
	CountryId  string                 `protobuf:"bytes,5,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	ValidFrom  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ValidUntil *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
}

func (x *LisTaxRequest) Reset() {
	*x = LisTaxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LisTaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LisTaxRequest) ProtoMessage() {}

func (x *LisTaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LisTaxRequest.ProtoReflect.Descriptor instead.
func (*LisTaxRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_rules_v1_rules_proto_rawDescGZIP(), []int{0}
}

func (x *LisTaxRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *LisTaxRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *LisTaxRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *LisTaxRequest) GetOrderBy() LisTaxRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return LisTaxRequest_ASC
}

func (x *LisTaxRequest) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

func (x *LisTaxRequest) GetValidFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

func (x *LisTaxRequest) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

type ListTaxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalData int32       `protobuf:"varint,1,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
	Data      []*TaxRules `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListTaxResponse) Reset() {
	*x = ListTaxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaxResponse) ProtoMessage() {}

func (x *ListTaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaxResponse.ProtoReflect.Descriptor instead.
func (*ListTaxResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_rules_v1_rules_proto_rawDescGZIP(), []int{1}
}

func (x *ListTaxResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

func (x *ListTaxResponse) GetData() []*TaxRules {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListServiceFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32                         `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size       int32                         `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	SortBy     string                        `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy    ListServiceFeeRequest_OrderBy `protobuf:"varint,4,opt,name=order_by,json=orderBy,proto3,enum=smallbiznis.rules.v1.ListServiceFeeRequest_OrderBy" json:"order_by,omitempty"`
	CountryId  string                        `protobuf:"bytes,5,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	ValidFrom  *timestamppb.Timestamp        `protobuf:"bytes,6,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ValidUntil *timestamppb.Timestamp        `protobuf:"bytes,7,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
}

func (x *ListServiceFeeRequest) Reset() {
	*x = ListServiceFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceFeeRequest) ProtoMessage() {}

func (x *ListServiceFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceFeeRequest.ProtoReflect.Descriptor instead.
func (*ListServiceFeeRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_rules_v1_rules_proto_rawDescGZIP(), []int{2}
}

func (x *ListServiceFeeRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListServiceFeeRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListServiceFeeRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListServiceFeeRequest) GetOrderBy() ListServiceFeeRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return ListServiceFeeRequest_ASC
}

func (x *ListServiceFeeRequest) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

func (x *ListServiceFeeRequest) GetValidFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

func (x *ListServiceFeeRequest) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

type ListServiceFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalData int32              `protobuf:"varint,1,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
	Data      []*ServiceFeeRules `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListServiceFeeResponse) Reset() {
	*x = ListServiceFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceFeeResponse) ProtoMessage() {}

func (x *ListServiceFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_rules_v1_rules_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceFeeResponse.ProtoReflect.Descriptor instead.
func (*ListServiceFeeResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_rules_v1_rules_proto_rawDescGZIP(), []int{3}
}

func (x *ListServiceFeeResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

func (x *ListServiceFeeResponse) GetData() []*ServiceFeeRules {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_smallbiznis_rules_v1_rules_proto protoreflect.FileDescriptor

var file_smallbiznis_rules_v1_rules_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x54, 0x61, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x46, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x54, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x1c, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x01, 0x22, 0x64, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xdd, 0x02, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x12, 0x4e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69,
	0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x3b,
	0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x1c, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x22, 0x72, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46,
	0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xfe, 0x07,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x54, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x78, 0x12, 0x66, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x78, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x1a, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x61, 0x78, 0x2f, 0x7b, 0x74, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1e,
	0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x1e,
	0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x78, 0x2f, 0x7b, 0x74, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x1e, 0x2e, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x78,
	0x2f, 0x7b, 0x74, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x89, 0x01, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x2b, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x66, 0x65, 0x65, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x1a, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66,
	0x65, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x25, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x1a, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x2c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x1a, 0x21, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_rules_v1_rules_proto_rawDescOnce sync.Once
	file_smallbiznis_rules_v1_rules_proto_rawDescData = file_smallbiznis_rules_v1_rules_proto_rawDesc
)

func file_smallbiznis_rules_v1_rules_proto_rawDescGZIP() []byte {
	file_smallbiznis_rules_v1_rules_proto_rawDescOnce.Do(func() {
		file_smallbiznis_rules_v1_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_rules_v1_rules_proto_rawDescData)
	})
	return file_smallbiznis_rules_v1_rules_proto_rawDescData
}

var file_smallbiznis_rules_v1_rules_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_smallbiznis_rules_v1_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_smallbiznis_rules_v1_rules_proto_goTypes = []interface{}{
	(LisTaxRequest_OrderBy)(0),         // 0: smallbiznis.rules.v1.LisTaxRequest.OrderBy
	(ListServiceFeeRequest_OrderBy)(0), // 1: smallbiznis.rules.v1.ListServiceFeeRequest.OrderBy
	(*LisTaxRequest)(nil),              // 2: smallbiznis.rules.v1.LisTaxRequest
	(*ListTaxResponse)(nil),            // 3: smallbiznis.rules.v1.ListTaxResponse
	(*ListServiceFeeRequest)(nil),      // 4: smallbiznis.rules.v1.ListServiceFeeRequest
	(*ListServiceFeeResponse)(nil),     // 5: smallbiznis.rules.v1.ListServiceFeeResponse
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
	(*TaxRules)(nil),                   // 7: smallbiznis.rules.v1.TaxRules
	(*ServiceFeeRules)(nil),            // 8: smallbiznis.rules.v1.ServiceFeeRules
}
var file_smallbiznis_rules_v1_rules_proto_depIdxs = []int32{
	0,  // 0: smallbiznis.rules.v1.LisTaxRequest.order_by:type_name -> smallbiznis.rules.v1.LisTaxRequest.OrderBy
	6,  // 1: smallbiznis.rules.v1.LisTaxRequest.valid_from:type_name -> google.protobuf.Timestamp
	6,  // 2: smallbiznis.rules.v1.LisTaxRequest.valid_until:type_name -> google.protobuf.Timestamp
	7,  // 3: smallbiznis.rules.v1.ListTaxResponse.data:type_name -> smallbiznis.rules.v1.TaxRules
	1,  // 4: smallbiznis.rules.v1.ListServiceFeeRequest.order_by:type_name -> smallbiznis.rules.v1.ListServiceFeeRequest.OrderBy
	6,  // 5: smallbiznis.rules.v1.ListServiceFeeRequest.valid_from:type_name -> google.protobuf.Timestamp
	6,  // 6: smallbiznis.rules.v1.ListServiceFeeRequest.valid_until:type_name -> google.protobuf.Timestamp
	8,  // 7: smallbiznis.rules.v1.ListServiceFeeResponse.data:type_name -> smallbiznis.rules.v1.ServiceFeeRules
	2,  // 8: smallbiznis.rules.v1.Service.ListTaxRule:input_type -> smallbiznis.rules.v1.LisTaxRequest
	7,  // 9: smallbiznis.rules.v1.Service.GetTaxRule:input_type -> smallbiznis.rules.v1.TaxRules
	7,  // 10: smallbiznis.rules.v1.Service.CreateTaxRule:input_type -> smallbiznis.rules.v1.TaxRules
	7,  // 11: smallbiznis.rules.v1.Service.UpdateTaxRule:input_type -> smallbiznis.rules.v1.TaxRules
	4,  // 12: smallbiznis.rules.v1.Service.ListServiceFeeRule:input_type -> smallbiznis.rules.v1.ListServiceFeeRequest
	8,  // 13: smallbiznis.rules.v1.Service.GetServiceFeeRule:input_type -> smallbiznis.rules.v1.ServiceFeeRules
	8,  // 14: smallbiznis.rules.v1.Service.CreateServiceFeeRule:input_type -> smallbiznis.rules.v1.ServiceFeeRules
	8,  // 15: smallbiznis.rules.v1.Service.UpdateServiceFeeRule:input_type -> smallbiznis.rules.v1.ServiceFeeRules
	3,  // 16: smallbiznis.rules.v1.Service.ListTaxRule:output_type -> smallbiznis.rules.v1.ListTaxResponse
	7,  // 17: smallbiznis.rules.v1.Service.GetTaxRule:output_type -> smallbiznis.rules.v1.TaxRules
	7,  // 18: smallbiznis.rules.v1.Service.CreateTaxRule:output_type -> smallbiznis.rules.v1.TaxRules
	7,  // 19: smallbiznis.rules.v1.Service.UpdateTaxRule:output_type -> smallbiznis.rules.v1.TaxRules
	5,  // 20: smallbiznis.rules.v1.Service.ListServiceFeeRule:output_type -> smallbiznis.rules.v1.ListServiceFeeResponse
	8,  // 21: smallbiznis.rules.v1.Service.GetServiceFeeRule:output_type -> smallbiznis.rules.v1.ServiceFeeRules
	8,  // 22: smallbiznis.rules.v1.Service.CreateServiceFeeRule:output_type -> smallbiznis.rules.v1.ServiceFeeRules
	8,  // 23: smallbiznis.rules.v1.Service.UpdateServiceFeeRule:output_type -> smallbiznis.rules.v1.ServiceFeeRules
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_smallbiznis_rules_v1_rules_proto_init() }
func file_smallbiznis_rules_v1_rules_proto_init() {
	if File_smallbiznis_rules_v1_rules_proto != nil {
		return
	}
	file_smallbiznis_rules_v1_rules_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_rules_v1_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LisTaxRequest); i {
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
		file_smallbiznis_rules_v1_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaxResponse); i {
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
		file_smallbiznis_rules_v1_rules_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceFeeRequest); i {
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
		file_smallbiznis_rules_v1_rules_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceFeeResponse); i {
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
			RawDescriptor: file_smallbiznis_rules_v1_rules_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smallbiznis_rules_v1_rules_proto_goTypes,
		DependencyIndexes: file_smallbiznis_rules_v1_rules_proto_depIdxs,
		EnumInfos:         file_smallbiznis_rules_v1_rules_proto_enumTypes,
		MessageInfos:      file_smallbiznis_rules_v1_rules_proto_msgTypes,
	}.Build()
	File_smallbiznis_rules_v1_rules_proto = out.File
	file_smallbiznis_rules_v1_rules_proto_rawDesc = nil
	file_smallbiznis_rules_v1_rules_proto_goTypes = nil
	file_smallbiznis_rules_v1_rules_proto_depIdxs = nil
}

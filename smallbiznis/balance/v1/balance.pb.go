// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/balance/v1/balance.proto

package balance

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListTransactionRequest_OrderBy int32

const (
	ListTransactionRequest_ASC  ListTransactionRequest_OrderBy = 0
	ListTransactionRequest_DESC ListTransactionRequest_OrderBy = 1
)

// Enum value maps for ListTransactionRequest_OrderBy.
var (
	ListTransactionRequest_OrderBy_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	ListTransactionRequest_OrderBy_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x ListTransactionRequest_OrderBy) Enum() *ListTransactionRequest_OrderBy {
	p := new(ListTransactionRequest_OrderBy)
	*p = x
	return p
}

func (x ListTransactionRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListTransactionRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_balance_v1_balance_proto_enumTypes[0].Descriptor()
}

func (ListTransactionRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_smallbiznis_balance_v1_balance_proto_enumTypes[0]
}

func (x ListTransactionRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListTransactionRequest_OrderBy.Descriptor instead.
func (ListTransactionRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{4, 0}
}

type CreateBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *CreateBalanceRequest) Reset() {
	*x = CreateBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBalanceRequest) ProtoMessage() {}

func (x *CreateBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBalanceRequest.ProtoReflect.Descriptor instead.
func (*CreateBalanceRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBalanceRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float32 `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BalanceId   string          `protobuf:"bytes,1,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	Type        TransactionType `protobuf:"varint,2,opt,name=type,proto3,enum=smallbiznis.balance.v1.TransactionType" json:"type,omitempty"`
	ReferenceId string          `protobuf:"bytes,3,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	TotalAmount float32         `protobuf:"fixed32,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTransactionRequest) GetBalanceId() string {
	if x != nil {
		return x.BalanceId
	}
	return ""
}

func (x *CreateTransactionRequest) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_TRANSACTION
}

func (x *CreateTransactionRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *CreateTransactionRequest) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type ListTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int32                          `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size        int32                          `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	SortBy      string                         `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy     ListTransactionRequest_OrderBy `protobuf:"varint,4,opt,name=order_by,json=orderBy,proto3,enum=smallbiznis.balance.v1.ListTransactionRequest_OrderBy" json:"order_by,omitempty"`
	BalanceId   string                         `protobuf:"bytes,5,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	FromDate    string                         `protobuf:"bytes,6,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate      string                         `protobuf:"bytes,7,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	ReferenceId string                         `protobuf:"bytes,8,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Type        TransactionType                `protobuf:"varint,9,opt,name=type,proto3,enum=smallbiznis.balance.v1.TransactionType" json:"type,omitempty"`
}

func (x *ListTransactionRequest) Reset() {
	*x = ListTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionRequest) ProtoMessage() {}

func (x *ListTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionRequest.ProtoReflect.Descriptor instead.
func (*ListTransactionRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{4}
}

func (x *ListTransactionRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTransactionRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListTransactionRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListTransactionRequest) GetOrderBy() ListTransactionRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return ListTransactionRequest_ASC
}

func (x *ListTransactionRequest) GetBalanceId() string {
	if x != nil {
		return x.BalanceId
	}
	return ""
}

func (x *ListTransactionRequest) GetFromDate() string {
	if x != nil {
		return x.FromDate
	}
	return ""
}

func (x *ListTransactionRequest) GetToDate() string {
	if x != nil {
		return x.ToDate
	}
	return ""
}

func (x *ListTransactionRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *ListTransactionRequest) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_TRANSACTION
}

type ListTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*Transaction `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalData int32          `protobuf:"varint,2,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
}

func (x *ListTransactionResponse) Reset() {
	*x = ListTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactionResponse) ProtoMessage() {}

func (x *ListTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactionResponse.ProtoReflect.Descriptor instead.
func (*ListTransactionResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{5}
}

func (x *ListTransactionResponse) GetData() []*Transaction {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListTransactionResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_balance_v1_balance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_balance_v1_balance_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_smallbiznis_balance_v1_balance_proto protoreflect.FileDescriptor

var file_smallbiznis_balance_v1_balance_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2d,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3b,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x84, 0x03, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x51, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x22, 0x0a, 0x0a, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1c, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x22, 0x71, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd8, 0x04, 0x0a, 0x0e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x6c, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x30, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x7e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69,
	0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_smallbiznis_balance_v1_balance_proto_rawDescOnce sync.Once
	file_smallbiznis_balance_v1_balance_proto_rawDescData = file_smallbiznis_balance_v1_balance_proto_rawDesc
)

func file_smallbiznis_balance_v1_balance_proto_rawDescGZIP() []byte {
	file_smallbiznis_balance_v1_balance_proto_rawDescOnce.Do(func() {
		file_smallbiznis_balance_v1_balance_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_balance_v1_balance_proto_rawDescData)
	})
	return file_smallbiznis_balance_v1_balance_proto_rawDescData
}

var file_smallbiznis_balance_v1_balance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_smallbiznis_balance_v1_balance_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_smallbiznis_balance_v1_balance_proto_goTypes = []interface{}{
	(ListTransactionRequest_OrderBy)(0), // 0: smallbiznis.balance.v1.ListTransactionRequest.OrderBy
	(*CreateBalanceRequest)(nil),        // 1: smallbiznis.balance.v1.CreateBalanceRequest
	(*GetBalanceRequest)(nil),           // 2: smallbiznis.balance.v1.GetBalanceRequest
	(*GetBalanceResponse)(nil),          // 3: smallbiznis.balance.v1.GetBalanceResponse
	(*CreateTransactionRequest)(nil),    // 4: smallbiznis.balance.v1.CreateTransactionRequest
	(*ListTransactionRequest)(nil),      // 5: smallbiznis.balance.v1.ListTransactionRequest
	(*ListTransactionResponse)(nil),     // 6: smallbiznis.balance.v1.ListTransactionResponse
	(*GetTransactionRequest)(nil),       // 7: smallbiznis.balance.v1.GetTransactionRequest
	(TransactionType)(0),                // 8: smallbiznis.balance.v1.TransactionType
	(*Transaction)(nil),                 // 9: smallbiznis.balance.v1.Transaction
	(*Balance)(nil),                     // 10: smallbiznis.balance.v1.Balance
}
var file_smallbiznis_balance_v1_balance_proto_depIdxs = []int32{
	8,  // 0: smallbiznis.balance.v1.CreateTransactionRequest.type:type_name -> smallbiznis.balance.v1.TransactionType
	0,  // 1: smallbiznis.balance.v1.ListTransactionRequest.order_by:type_name -> smallbiznis.balance.v1.ListTransactionRequest.OrderBy
	8,  // 2: smallbiznis.balance.v1.ListTransactionRequest.type:type_name -> smallbiznis.balance.v1.TransactionType
	9,  // 3: smallbiznis.balance.v1.ListTransactionResponse.data:type_name -> smallbiznis.balance.v1.Transaction
	1,  // 4: smallbiznis.balance.v1.BalanceService.CreateBalance:input_type -> smallbiznis.balance.v1.CreateBalanceRequest
	2,  // 5: smallbiznis.balance.v1.BalanceService.GetBalance:input_type -> smallbiznis.balance.v1.GetBalanceRequest
	4,  // 6: smallbiznis.balance.v1.BalanceService.CreateTransaction:input_type -> smallbiznis.balance.v1.CreateTransactionRequest
	5,  // 7: smallbiznis.balance.v1.BalanceService.ListTransaction:input_type -> smallbiznis.balance.v1.ListTransactionRequest
	7,  // 8: smallbiznis.balance.v1.BalanceService.GetTransaction:input_type -> smallbiznis.balance.v1.GetTransactionRequest
	10, // 9: smallbiznis.balance.v1.BalanceService.CreateBalance:output_type -> smallbiznis.balance.v1.Balance
	10, // 10: smallbiznis.balance.v1.BalanceService.GetBalance:output_type -> smallbiznis.balance.v1.Balance
	9,  // 11: smallbiznis.balance.v1.BalanceService.CreateTransaction:output_type -> smallbiznis.balance.v1.Transaction
	6,  // 12: smallbiznis.balance.v1.BalanceService.ListTransaction:output_type -> smallbiznis.balance.v1.ListTransactionResponse
	9,  // 13: smallbiznis.balance.v1.BalanceService.GetTransaction:output_type -> smallbiznis.balance.v1.Transaction
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_smallbiznis_balance_v1_balance_proto_init() }
func file_smallbiznis_balance_v1_balance_proto_init() {
	if File_smallbiznis_balance_v1_balance_proto != nil {
		return
	}
	file_smallbiznis_balance_v1_balance_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_balance_v1_balance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBalanceRequest); i {
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
		file_smallbiznis_balance_v1_balance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_smallbiznis_balance_v1_balance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
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
		file_smallbiznis_balance_v1_balance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionRequest); i {
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
		file_smallbiznis_balance_v1_balance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransactionRequest); i {
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
		file_smallbiznis_balance_v1_balance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransactionResponse); i {
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
		file_smallbiznis_balance_v1_balance_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRequest); i {
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
			RawDescriptor: file_smallbiznis_balance_v1_balance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smallbiznis_balance_v1_balance_proto_goTypes,
		DependencyIndexes: file_smallbiznis_balance_v1_balance_proto_depIdxs,
		EnumInfos:         file_smallbiznis_balance_v1_balance_proto_enumTypes,
		MessageInfos:      file_smallbiznis_balance_v1_balance_proto_msgTypes,
	}.Build()
	File_smallbiznis_balance_v1_balance_proto = out.File
	file_smallbiznis_balance_v1_balance_proto_rawDesc = nil
	file_smallbiznis_balance_v1_balance_proto_goTypes = nil
	file_smallbiznis_balance_v1_balance_proto_depIdxs = nil
}
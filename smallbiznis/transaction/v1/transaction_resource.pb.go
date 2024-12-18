// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/transaction/v1/transaction_resource.proto

package transaction

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/type/date"
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

type OrderStatus int32

const (
	OrderStatus_draft     OrderStatus = 0
	OrderStatus_created   OrderStatus = 1
	OrderStatus_confirmed OrderStatus = 2
	OrderStatus_cancelled OrderStatus = 3
	OrderStatus_completed OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "draft",
		1: "created",
		2: "confirmed",
		3: "cancelled",
		4: "completed",
	}
	OrderStatus_value = map[string]int32{
		"draft":     0,
		"created":   1,
		"confirmed": 2,
		"cancelled": 3,
		"completed": 4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_transaction_v1_transaction_resource_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_smallbiznis_transaction_v1_transaction_resource_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescGZIP(), []int{0}
}

type FulfillmentStatus int32

const (
	FulfillmentStatus_not_fulfilled     FulfillmentStatus = 0
	FulfillmentStatus_partial_fulfilled FulfillmentStatus = 1
	FulfillmentStatus_fulfilled         FulfillmentStatus = 2
	FulfillmentStatus_failed            FulfillmentStatus = 3
)

// Enum value maps for FulfillmentStatus.
var (
	FulfillmentStatus_name = map[int32]string{
		0: "not_fulfilled",
		1: "partial_fulfilled",
		2: "fulfilled",
		3: "failed",
	}
	FulfillmentStatus_value = map[string]int32{
		"not_fulfilled":     0,
		"partial_fulfilled": 1,
		"fulfilled":         2,
		"failed":            3,
	}
)

func (x FulfillmentStatus) Enum() *FulfillmentStatus {
	p := new(FulfillmentStatus)
	*p = x
	return p
}

func (x FulfillmentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FulfillmentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_transaction_v1_transaction_resource_proto_enumTypes[1].Descriptor()
}

func (FulfillmentStatus) Type() protoreflect.EnumType {
	return &file_smallbiznis_transaction_v1_transaction_resource_proto_enumTypes[1]
}

func (x FulfillmentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FulfillmentStatus.Descriptor instead.
func (FulfillmentStatus) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescGZIP(), []int{1}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId        string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrganizationId string                 `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	CustomerId     string                 `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SalesChannelId string                 `protobuf:"bytes,4,opt,name=sales_channel_id,json=salesChannelId,proto3" json:"sales_channel_id,omitempty"`
	OrderNo        string                 `protobuf:"bytes,5,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	SubTotal       float32                `protobuf:"fixed32,6,opt,name=sub_total,json=subTotal,proto3" json:"sub_total,omitempty"`
	TaxAmount      float32                `protobuf:"fixed32,7,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TotalAmount    float32                `protobuf:"fixed32,8,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Status         OrderStatus            `protobuf:"varint,12,opt,name=status,proto3,enum=smallbiznis.transaction.v1.OrderStatus" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	OrderItems     []*OrderItem           `protobuf:"bytes,15,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Order) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Order) GetSalesChannelId() string {
	if x != nil {
		return x.SalesChannelId
	}
	return ""
}

func (x *Order) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *Order) GetSubTotal() float32 {
	if x != nil {
		return x.SubTotal
	}
	return 0
}

func (x *Order) GetTaxAmount() float32 {
	if x != nil {
		return x.TaxAmount
	}
	return 0
}

func (x *Order) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_draft
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Order) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderItemId string                 `protobuf:"bytes,1,opt,name=order_item_id,json=orderItemId,proto3" json:"order_item_id,omitempty"`
	OrderId     string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ItemId      string                 `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemName    string                 `protobuf:"bytes,4,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Quantity    int32                  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitPrice   float32                `protobuf:"fixed32,6,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	TaxAmount   float32                `protobuf:"fixed32,7,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TotalPrice  float32                `protobuf:"fixed32,8,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItem) GetOrderItemId() string {
	if x != nil {
		return x.OrderItemId
	}
	return ""
}

func (x *OrderItem) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderItem) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *OrderItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetUnitPrice() float32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *OrderItem) GetTaxAmount() float32 {
	if x != nil {
		return x.TaxAmount
	}
	return 0
}

func (x *OrderItem) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OrderItem) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_smallbiznis_transaction_v1_transaction_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_transaction_v1_transaction_resource_proto_rawDesc = []byte{
	0x0a, 0x35, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69,
	0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x04,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x74, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0xf1, 0x02, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a,
	0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x78, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x74, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x2a, 0x52, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x04, 0x2a, 0x58, 0x0a, 0x11, 0x46, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d,
	0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10,
	0x03, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescOnce sync.Once
	file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescData = file_smallbiznis_transaction_v1_transaction_resource_proto_rawDesc
)

func file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescGZIP() []byte {
	file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescOnce.Do(func() {
		file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescData)
	})
	return file_smallbiznis_transaction_v1_transaction_resource_proto_rawDescData
}

var file_smallbiznis_transaction_v1_transaction_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_smallbiznis_transaction_v1_transaction_resource_proto_goTypes = []interface{}{
	(OrderStatus)(0),              // 0: smallbiznis.transaction.v1.OrderStatus
	(FulfillmentStatus)(0),        // 1: smallbiznis.transaction.v1.FulfillmentStatus
	(*Order)(nil),                 // 2: smallbiznis.transaction.v1.Order
	(*OrderItem)(nil),             // 3: smallbiznis.transaction.v1.OrderItem
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_smallbiznis_transaction_v1_transaction_resource_proto_depIdxs = []int32{
	0, // 0: smallbiznis.transaction.v1.Order.status:type_name -> smallbiznis.transaction.v1.OrderStatus
	4, // 1: smallbiznis.transaction.v1.Order.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: smallbiznis.transaction.v1.Order.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: smallbiznis.transaction.v1.Order.order_items:type_name -> smallbiznis.transaction.v1.OrderItem
	4, // 4: smallbiznis.transaction.v1.OrderItem.created_at:type_name -> google.protobuf.Timestamp
	4, // 5: smallbiznis.transaction.v1.OrderItem.updated_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_smallbiznis_transaction_v1_transaction_resource_proto_init() }
func file_smallbiznis_transaction_v1_transaction_resource_proto_init() {
	if File_smallbiznis_transaction_v1_transaction_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
			RawDescriptor: file_smallbiznis_transaction_v1_transaction_resource_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_transaction_v1_transaction_resource_proto_goTypes,
		DependencyIndexes: file_smallbiznis_transaction_v1_transaction_resource_proto_depIdxs,
		EnumInfos:         file_smallbiznis_transaction_v1_transaction_resource_proto_enumTypes,
		MessageInfos:      file_smallbiznis_transaction_v1_transaction_resource_proto_msgTypes,
	}.Build()
	File_smallbiznis_transaction_v1_transaction_resource_proto = out.File
	file_smallbiznis_transaction_v1_transaction_resource_proto_rawDesc = nil
	file_smallbiznis_transaction_v1_transaction_resource_proto_goTypes = nil
	file_smallbiznis_transaction_v1_transaction_resource_proto_depIdxs = nil
}

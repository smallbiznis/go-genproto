// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/hosting/v1/hosting.proto

package hosting

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListDomainRequest_OrderBy int32

const (
	ListDomainRequest_ASC  ListDomainRequest_OrderBy = 0
	ListDomainRequest_DESC ListDomainRequest_OrderBy = 1
)

// Enum value maps for ListDomainRequest_OrderBy.
var (
	ListDomainRequest_OrderBy_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	ListDomainRequest_OrderBy_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x ListDomainRequest_OrderBy) Enum() *ListDomainRequest_OrderBy {
	p := new(ListDomainRequest_OrderBy)
	*p = x
	return p
}

func (x ListDomainRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListDomainRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_hosting_v1_hosting_proto_enumTypes[0].Descriptor()
}

func (ListDomainRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_smallbiznis_hosting_v1_hosting_proto_enumTypes[0]
}

func (x ListDomainRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListDomainRequest_OrderBy.Descriptor instead.
func (ListDomainRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_hosting_v1_hosting_proto_rawDescGZIP(), []int{0, 0}
}

type ListDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           int32                     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size           int32                     `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	SortBy         string                    `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy        ListDomainRequest_OrderBy `protobuf:"varint,4,opt,name=order_by,json=orderBy,proto3,enum=smallbiznis.hosting.v1.ListDomainRequest_OrderBy" json:"order_by,omitempty"`
	OrganizationId string                    `protobuf:"bytes,5,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *ListDomainRequest) Reset() {
	*x = ListDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_hosting_v1_hosting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDomainRequest) ProtoMessage() {}

func (x *ListDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_hosting_v1_hosting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDomainRequest.ProtoReflect.Descriptor instead.
func (*ListDomainRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_hosting_v1_hosting_proto_rawDescGZIP(), []int{0}
}

func (x *ListDomainRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListDomainRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListDomainRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListDomainRequest) GetOrderBy() ListDomainRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return ListDomainRequest_ASC
}

func (x *ListDomainRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type ListDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*Domain `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalData int32     `protobuf:"varint,2,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
}

func (x *ListDomainResponse) Reset() {
	*x = ListDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_hosting_v1_hosting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDomainResponse) ProtoMessage() {}

func (x *ListDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_hosting_v1_hosting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDomainResponse.ProtoReflect.Descriptor instead.
func (*ListDomainResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_hosting_v1_hosting_proto_rawDescGZIP(), []int{1}
}

func (x *ListDomainResponse) GetData() []*Domain {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListDomainResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

type GetDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *GetDomainRequest) Reset() {
	*x = GetDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_hosting_v1_hosting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainRequest) ProtoMessage() {}

func (x *GetDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_hosting_v1_hosting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainRequest.ProtoReflect.Descriptor instead.
func (*GetDomainRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_hosting_v1_hosting_proto_rawDescGZIP(), []int{2}
}

func (x *GetDomainRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_smallbiznis_hosting_v1_hosting_proto protoreflect.FileDescriptor

var file_smallbiznis_hosting_v1_hosting_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x2d,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x12, 0x4c, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x1c, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53,
	0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x22, 0x67, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x32, 0xdb, 0x03, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x12, 0x76, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x2f, 0x7b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x72, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x1a, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x1a, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x3b, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_smallbiznis_hosting_v1_hosting_proto_rawDescOnce sync.Once
	file_smallbiznis_hosting_v1_hosting_proto_rawDescData = file_smallbiznis_hosting_v1_hosting_proto_rawDesc
)

func file_smallbiznis_hosting_v1_hosting_proto_rawDescGZIP() []byte {
	file_smallbiznis_hosting_v1_hosting_proto_rawDescOnce.Do(func() {
		file_smallbiznis_hosting_v1_hosting_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_hosting_v1_hosting_proto_rawDescData)
	})
	return file_smallbiznis_hosting_v1_hosting_proto_rawDescData
}

var file_smallbiznis_hosting_v1_hosting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_smallbiznis_hosting_v1_hosting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_smallbiznis_hosting_v1_hosting_proto_goTypes = []interface{}{
	(ListDomainRequest_OrderBy)(0), // 0: smallbiznis.hosting.v1.ListDomainRequest.OrderBy
	(*ListDomainRequest)(nil),      // 1: smallbiznis.hosting.v1.ListDomainRequest
	(*ListDomainResponse)(nil),     // 2: smallbiznis.hosting.v1.ListDomainResponse
	(*GetDomainRequest)(nil),       // 3: smallbiznis.hosting.v1.GetDomainRequest
	(*Domain)(nil),                 // 4: smallbiznis.hosting.v1.Domain
}
var file_smallbiznis_hosting_v1_hosting_proto_depIdxs = []int32{
	0, // 0: smallbiznis.hosting.v1.ListDomainRequest.order_by:type_name -> smallbiznis.hosting.v1.ListDomainRequest.OrderBy
	4, // 1: smallbiznis.hosting.v1.ListDomainResponse.data:type_name -> smallbiznis.hosting.v1.Domain
	1, // 2: smallbiznis.hosting.v1.HostingService.ListDomain:input_type -> smallbiznis.hosting.v1.ListDomainRequest
	3, // 3: smallbiznis.hosting.v1.HostingService.GetDomain:input_type -> smallbiznis.hosting.v1.GetDomainRequest
	4, // 4: smallbiznis.hosting.v1.HostingService.AddDomain:input_type -> smallbiznis.hosting.v1.Domain
	4, // 5: smallbiznis.hosting.v1.HostingService.UpdateDomain:input_type -> smallbiznis.hosting.v1.Domain
	2, // 6: smallbiznis.hosting.v1.HostingService.ListDomain:output_type -> smallbiznis.hosting.v1.ListDomainResponse
	4, // 7: smallbiznis.hosting.v1.HostingService.GetDomain:output_type -> smallbiznis.hosting.v1.Domain
	4, // 8: smallbiznis.hosting.v1.HostingService.AddDomain:output_type -> smallbiznis.hosting.v1.Domain
	4, // 9: smallbiznis.hosting.v1.HostingService.UpdateDomain:output_type -> smallbiznis.hosting.v1.Domain
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_smallbiznis_hosting_v1_hosting_proto_init() }
func file_smallbiznis_hosting_v1_hosting_proto_init() {
	if File_smallbiznis_hosting_v1_hosting_proto != nil {
		return
	}
	file_smallbiznis_hosting_v1_hosting_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_hosting_v1_hosting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDomainRequest); i {
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
		file_smallbiznis_hosting_v1_hosting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDomainResponse); i {
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
		file_smallbiznis_hosting_v1_hosting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDomainRequest); i {
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
			RawDescriptor: file_smallbiznis_hosting_v1_hosting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smallbiznis_hosting_v1_hosting_proto_goTypes,
		DependencyIndexes: file_smallbiznis_hosting_v1_hosting_proto_depIdxs,
		EnumInfos:         file_smallbiznis_hosting_v1_hosting_proto_enumTypes,
		MessageInfos:      file_smallbiznis_hosting_v1_hosting_proto_msgTypes,
	}.Build()
	File_smallbiznis_hosting_v1_hosting_proto = out.File
	file_smallbiznis_hosting_v1_hosting_proto_rawDesc = nil
	file_smallbiznis_hosting_v1_hosting_proto_goTypes = nil
	file_smallbiznis_hosting_v1_hosting_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/member/v1/member.proto

package member

import (
	protobuf "github.com/smallbiznis/go-genproto/smallbiznis/protobuf"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListMemberRequest_OrderBy int32

const (
	ListMemberRequest_ASC  ListMemberRequest_OrderBy = 0
	ListMemberRequest_DESC ListMemberRequest_OrderBy = 1
)

// Enum value maps for ListMemberRequest_OrderBy.
var (
	ListMemberRequest_OrderBy_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	ListMemberRequest_OrderBy_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x ListMemberRequest_OrderBy) Enum() *ListMemberRequest_OrderBy {
	p := new(ListMemberRequest_OrderBy)
	*p = x
	return p
}

func (x ListMemberRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListMemberRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_member_v1_member_proto_enumTypes[0].Descriptor()
}

func (ListMemberRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_smallbiznis_member_v1_member_proto_enumTypes[0]
}

func (x ListMemberRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListMemberRequest_OrderBy.Descriptor instead.
func (ListMemberRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{0, 0}
}

type ListMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           int32                     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size           int32                     `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	SortBy         string                    `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy        ListMemberRequest_OrderBy `protobuf:"varint,4,opt,name=order_by,json=orderBy,proto3,enum=smallbiznis.member.v1.ListMemberRequest_OrderBy" json:"order_by,omitempty"`
	UserId         string                    `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email          string                    `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	OrganizationId string                    `protobuf:"bytes,7,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *ListMemberRequest) Reset() {
	*x = ListMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemberRequest) ProtoMessage() {}

func (x *ListMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemberRequest.ProtoReflect.Descriptor instead.
func (*ListMemberRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{0}
}

func (x *ListMemberRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListMemberRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListMemberRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListMemberRequest) GetOrderBy() ListMemberRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return ListMemberRequest_ASC
}

func (x *ListMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListMemberRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ListMemberRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

type ListMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*Member `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalData int32     `protobuf:"varint,2,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
}

func (x *ListMemberResponse) Reset() {
	*x = ListMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemberResponse) ProtoMessage() {}

func (x *ListMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemberResponse.ProtoReflect.Descriptor instead.
func (*ListMemberResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{1}
}

func (x *ListMemberResponse) GetData() []*Member {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListMemberResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

type GetMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *GetMemberRequest) Reset() {
	*x = GetMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberRequest) ProtoMessage() {}

func (x *GetMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberRequest.ProtoReflect.Descriptor instead.
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{2}
}

func (x *GetMemberRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type AddMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Roles  []Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=smallbiznis.member.v1.Role" json:"roles,omitempty"`
}

func (x *AddMemberRequest) Reset() {
	*x = AddMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMemberRequest) ProtoMessage() {}

func (x *AddMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMemberRequest.ProtoReflect.Descriptor instead.
func (*AddMemberRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{3}
}

func (x *AddMemberRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddMemberRequest) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UpdateMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Roles    []Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=smallbiznis.member.v1.Role" json:"roles,omitempty"`
}

func (x *UpdateMemberRequest) Reset() {
	*x = UpdateMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemberRequest) ProtoMessage() {}

func (x *UpdateMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemberRequest.ProtoReflect.Descriptor instead.
func (*UpdateMemberRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMemberRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *UpdateMemberRequest) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type DeleteMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *DeleteMemberRequest) Reset() {
	*x = DeleteMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemberRequest) ProtoMessage() {}

func (x *DeleteMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemberRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemberRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMemberRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

var File_smallbiznis_member_v1_member_proto protoreflect.FileDescriptor

var file_smallbiznis_member_v1_member_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2b, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x12, 0x4b, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x01, 0x22, 0x66, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5e, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x65, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x32, 0xe3, 0x04, 0x0a, 0x0d, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x74, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x27, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x7d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x1a, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x78, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_member_v1_member_proto_rawDescOnce sync.Once
	file_smallbiznis_member_v1_member_proto_rawDescData = file_smallbiznis_member_v1_member_proto_rawDesc
)

func file_smallbiznis_member_v1_member_proto_rawDescGZIP() []byte {
	file_smallbiznis_member_v1_member_proto_rawDescOnce.Do(func() {
		file_smallbiznis_member_v1_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_member_v1_member_proto_rawDescData)
	})
	return file_smallbiznis_member_v1_member_proto_rawDescData
}

var file_smallbiznis_member_v1_member_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_smallbiznis_member_v1_member_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_smallbiznis_member_v1_member_proto_goTypes = []interface{}{
	(ListMemberRequest_OrderBy)(0), // 0: smallbiznis.member.v1.ListMemberRequest.OrderBy
	(*ListMemberRequest)(nil),      // 1: smallbiznis.member.v1.ListMemberRequest
	(*ListMemberResponse)(nil),     // 2: smallbiznis.member.v1.ListMemberResponse
	(*GetMemberRequest)(nil),       // 3: smallbiznis.member.v1.GetMemberRequest
	(*AddMemberRequest)(nil),       // 4: smallbiznis.member.v1.AddMemberRequest
	(*UpdateMemberRequest)(nil),    // 5: smallbiznis.member.v1.UpdateMemberRequest
	(*DeleteMemberRequest)(nil),    // 6: smallbiznis.member.v1.DeleteMemberRequest
	(*Member)(nil),                 // 7: smallbiznis.member.v1.Member
	(Role)(0),                      // 8: smallbiznis.member.v1.Role
	(*protobuf.Empty)(nil),         // 9: smallbiznis.protobuf.Empty
}
var file_smallbiznis_member_v1_member_proto_depIdxs = []int32{
	0, // 0: smallbiznis.member.v1.ListMemberRequest.order_by:type_name -> smallbiznis.member.v1.ListMemberRequest.OrderBy
	7, // 1: smallbiznis.member.v1.ListMemberResponse.data:type_name -> smallbiznis.member.v1.Member
	8, // 2: smallbiznis.member.v1.AddMemberRequest.roles:type_name -> smallbiznis.member.v1.Role
	8, // 3: smallbiznis.member.v1.UpdateMemberRequest.roles:type_name -> smallbiznis.member.v1.Role
	1, // 4: smallbiznis.member.v1.MemberService.ListMember:input_type -> smallbiznis.member.v1.ListMemberRequest
	3, // 5: smallbiznis.member.v1.MemberService.GetMember:input_type -> smallbiznis.member.v1.GetMemberRequest
	4, // 6: smallbiznis.member.v1.MemberService.AddMember:input_type -> smallbiznis.member.v1.AddMemberRequest
	5, // 7: smallbiznis.member.v1.MemberService.UpdateMember:input_type -> smallbiznis.member.v1.UpdateMemberRequest
	6, // 8: smallbiznis.member.v1.MemberService.DeleteMember:input_type -> smallbiznis.member.v1.DeleteMemberRequest
	2, // 9: smallbiznis.member.v1.MemberService.ListMember:output_type -> smallbiznis.member.v1.ListMemberResponse
	7, // 10: smallbiznis.member.v1.MemberService.GetMember:output_type -> smallbiznis.member.v1.Member
	7, // 11: smallbiznis.member.v1.MemberService.AddMember:output_type -> smallbiznis.member.v1.Member
	7, // 12: smallbiznis.member.v1.MemberService.UpdateMember:output_type -> smallbiznis.member.v1.Member
	9, // 13: smallbiznis.member.v1.MemberService.DeleteMember:output_type -> smallbiznis.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_smallbiznis_member_v1_member_proto_init() }
func file_smallbiznis_member_v1_member_proto_init() {
	if File_smallbiznis_member_v1_member_proto != nil {
		return
	}
	file_smallbiznis_member_v1_member_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_member_v1_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMemberRequest); i {
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
		file_smallbiznis_member_v1_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMemberResponse); i {
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
		file_smallbiznis_member_v1_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberRequest); i {
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
		file_smallbiznis_member_v1_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMemberRequest); i {
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
		file_smallbiznis_member_v1_member_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemberRequest); i {
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
		file_smallbiznis_member_v1_member_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemberRequest); i {
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
			RawDescriptor: file_smallbiznis_member_v1_member_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smallbiznis_member_v1_member_proto_goTypes,
		DependencyIndexes: file_smallbiznis_member_v1_member_proto_depIdxs,
		EnumInfos:         file_smallbiznis_member_v1_member_proto_enumTypes,
		MessageInfos:      file_smallbiznis_member_v1_member_proto_msgTypes,
	}.Build()
	File_smallbiznis_member_v1_member_proto = out.File
	file_smallbiznis_member_v1_member_proto_rawDesc = nil
	file_smallbiznis_member_v1_member_proto_goTypes = nil
	file_smallbiznis_member_v1_member_proto_depIdxs = nil
}

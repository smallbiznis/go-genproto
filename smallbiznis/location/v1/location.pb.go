// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: smallbiznis/location/v1/location.proto

package location

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type ListProvinceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProvinceRequest) Reset() {
	*x = ListProvinceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvinceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvinceRequest) ProtoMessage() {}

func (x *ListProvinceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvinceRequest.ProtoReflect.Descriptor instead.
func (*ListProvinceRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_proto_rawDescGZIP(), []int{0}
}

type ListProvinceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*Province `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalData int32       `protobuf:"varint,2,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
}

func (x *ListProvinceResponse) Reset() {
	*x = ListProvinceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvinceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvinceResponse) ProtoMessage() {}

func (x *ListProvinceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvinceResponse.ProtoReflect.Descriptor instead.
func (*ListProvinceResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_proto_rawDescGZIP(), []int{1}
}

func (x *ListProvinceResponse) GetData() []*Province {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListProvinceResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

type ListCityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceId string `protobuf:"bytes,1,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
}

func (x *ListCityRequest) Reset() {
	*x = ListCityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCityRequest) ProtoMessage() {}

func (x *ListCityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCityRequest.ProtoReflect.Descriptor instead.
func (*ListCityRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_proto_rawDescGZIP(), []int{2}
}

func (x *ListCityRequest) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

type ListCityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*City `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalData int32   `protobuf:"varint,2,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
}

func (x *ListCityResponse) Reset() {
	*x = ListCityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCityResponse) ProtoMessage() {}

func (x *ListCityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCityResponse.ProtoReflect.Descriptor instead.
func (*ListCityResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_proto_rawDescGZIP(), []int{3}
}

func (x *ListCityResponse) GetData() []*City {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListCityResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

type ListDistrictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviceId string `protobuf:"bytes,1,opt,name=provice_id,json=proviceId,proto3" json:"provice_id,omitempty"`
}

func (x *ListDistrictRequest) Reset() {
	*x = ListDistrictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDistrictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDistrictRequest) ProtoMessage() {}

func (x *ListDistrictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDistrictRequest.ProtoReflect.Descriptor instead.
func (*ListDistrictRequest) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_proto_rawDescGZIP(), []int{4}
}

func (x *ListDistrictRequest) GetProviceId() string {
	if x != nil {
		return x.ProviceId
	}
	return ""
}

type ListDistrictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*District `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalData int32       `protobuf:"varint,2,opt,name=total_data,json=totalData,proto3" json:"total_data,omitempty"`
}

func (x *ListDistrictResponse) Reset() {
	*x = ListDistrictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDistrictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDistrictResponse) ProtoMessage() {}

func (x *ListDistrictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDistrictResponse.ProtoReflect.Descriptor instead.
func (*ListDistrictResponse) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_proto_rawDescGZIP(), []int{5}
}

func (x *ListDistrictResponse) GetData() []*District {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListDistrictResponse) GetTotalData() int32 {
	if x != nil {
		return x.TotalData
	}
	return 0
}

var File_smallbiznis_location_v1_location_proto protoreflect.FileDescriptor

var file_smallbiznis_location_v1_location_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x34,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x32, 0x93, 0x03, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x76, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_location_v1_location_proto_rawDescOnce sync.Once
	file_smallbiznis_location_v1_location_proto_rawDescData = file_smallbiznis_location_v1_location_proto_rawDesc
)

func file_smallbiznis_location_v1_location_proto_rawDescGZIP() []byte {
	file_smallbiznis_location_v1_location_proto_rawDescOnce.Do(func() {
		file_smallbiznis_location_v1_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_location_v1_location_proto_rawDescData)
	})
	return file_smallbiznis_location_v1_location_proto_rawDescData
}

var file_smallbiznis_location_v1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_smallbiznis_location_v1_location_proto_goTypes = []any{
	(*ListProvinceRequest)(nil),  // 0: smallbiznis.location.v1.ListProvinceRequest
	(*ListProvinceResponse)(nil), // 1: smallbiznis.location.v1.ListProvinceResponse
	(*ListCityRequest)(nil),      // 2: smallbiznis.location.v1.ListCityRequest
	(*ListCityResponse)(nil),     // 3: smallbiznis.location.v1.ListCityResponse
	(*ListDistrictRequest)(nil),  // 4: smallbiznis.location.v1.ListDistrictRequest
	(*ListDistrictResponse)(nil), // 5: smallbiznis.location.v1.ListDistrictResponse
	(*Province)(nil),             // 6: smallbiznis.location.v1.Province
	(*City)(nil),                 // 7: smallbiznis.location.v1.City
	(*District)(nil),             // 8: smallbiznis.location.v1.District
}
var file_smallbiznis_location_v1_location_proto_depIdxs = []int32{
	6, // 0: smallbiznis.location.v1.ListProvinceResponse.data:type_name -> smallbiznis.location.v1.Province
	7, // 1: smallbiznis.location.v1.ListCityResponse.data:type_name -> smallbiznis.location.v1.City
	8, // 2: smallbiznis.location.v1.ListDistrictResponse.data:type_name -> smallbiznis.location.v1.District
	0, // 3: smallbiznis.location.v1.LocationService.ListProvince:input_type -> smallbiznis.location.v1.ListProvinceRequest
	2, // 4: smallbiznis.location.v1.LocationService.ListCity:input_type -> smallbiznis.location.v1.ListCityRequest
	4, // 5: smallbiznis.location.v1.LocationService.ListDistrict:input_type -> smallbiznis.location.v1.ListDistrictRequest
	1, // 6: smallbiznis.location.v1.LocationService.ListProvince:output_type -> smallbiznis.location.v1.ListProvinceResponse
	3, // 7: smallbiznis.location.v1.LocationService.ListCity:output_type -> smallbiznis.location.v1.ListCityResponse
	5, // 8: smallbiznis.location.v1.LocationService.ListDistrict:output_type -> smallbiznis.location.v1.ListDistrictResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_smallbiznis_location_v1_location_proto_init() }
func file_smallbiznis_location_v1_location_proto_init() {
	if File_smallbiznis_location_v1_location_proto != nil {
		return
	}
	file_smallbiznis_location_v1_location_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_location_v1_location_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListProvinceRequest); i {
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
		file_smallbiznis_location_v1_location_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListProvinceResponse); i {
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
		file_smallbiznis_location_v1_location_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListCityRequest); i {
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
		file_smallbiznis_location_v1_location_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListCityResponse); i {
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
		file_smallbiznis_location_v1_location_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListDistrictRequest); i {
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
		file_smallbiznis_location_v1_location_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListDistrictResponse); i {
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
			RawDescriptor: file_smallbiznis_location_v1_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smallbiznis_location_v1_location_proto_goTypes,
		DependencyIndexes: file_smallbiznis_location_v1_location_proto_depIdxs,
		MessageInfos:      file_smallbiznis_location_v1_location_proto_msgTypes,
	}.Build()
	File_smallbiznis_location_v1_location_proto = out.File
	file_smallbiznis_location_v1_location_proto_rawDesc = nil
	file_smallbiznis_location_v1_location_proto_goTypes = nil
	file_smallbiznis_location_v1_location_proto_depIdxs = nil
}

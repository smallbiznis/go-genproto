// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: smallbiznis/location/v1/location_resource.proto

package location

import (
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

type Province struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceId string `protobuf:"bytes,1,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Province) Reset() {
	*x = Province{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Province) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Province) ProtoMessage() {}

func (x *Province) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Province.ProtoReflect.Descriptor instead.
func (*Province) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Province) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

func (x *Province) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId     string `protobuf:"bytes,1,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	ProvinceId string `protobuf:"bytes,2,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_resource_proto_rawDescGZIP(), []int{1}
}

func (x *City) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *City) GetProvinceId() string {
	if x != nil {
		return x.ProvinceId
	}
	return ""
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type District struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictId string `protobuf:"bytes,1,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	CityId     string `protobuf:"bytes,2,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *District) Reset() {
	*x = District{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_location_v1_location_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *District) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*District) ProtoMessage() {}

func (x *District) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_location_v1_location_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use District.ProtoReflect.Descriptor instead.
func (*District) Descriptor() ([]byte, []int) {
	return file_smallbiznis_location_v1_location_resource_proto_rawDescGZIP(), []int{2}
}

func (x *District) GetDistrictId() string {
	if x != nil {
		return x.DistrictId
	}
	return ""
}

func (x *District) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *District) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_smallbiznis_location_v1_location_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_location_v1_location_resource_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58,
	0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e,
	0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_location_v1_location_resource_proto_rawDescOnce sync.Once
	file_smallbiznis_location_v1_location_resource_proto_rawDescData = file_smallbiznis_location_v1_location_resource_proto_rawDesc
)

func file_smallbiznis_location_v1_location_resource_proto_rawDescGZIP() []byte {
	file_smallbiznis_location_v1_location_resource_proto_rawDescOnce.Do(func() {
		file_smallbiznis_location_v1_location_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_location_v1_location_resource_proto_rawDescData)
	})
	return file_smallbiznis_location_v1_location_resource_proto_rawDescData
}

var file_smallbiznis_location_v1_location_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_smallbiznis_location_v1_location_resource_proto_goTypes = []any{
	(*Province)(nil), // 0: smallbiznis.location.v1.Province
	(*City)(nil),     // 1: smallbiznis.location.v1.City
	(*District)(nil), // 2: smallbiznis.location.v1.District
}
var file_smallbiznis_location_v1_location_resource_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_smallbiznis_location_v1_location_resource_proto_init() }
func file_smallbiznis_location_v1_location_resource_proto_init() {
	if File_smallbiznis_location_v1_location_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_location_v1_location_resource_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Province); i {
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
		file_smallbiznis_location_v1_location_resource_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*City); i {
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
		file_smallbiznis_location_v1_location_resource_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*District); i {
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
			RawDescriptor: file_smallbiznis_location_v1_location_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_location_v1_location_resource_proto_goTypes,
		DependencyIndexes: file_smallbiznis_location_v1_location_resource_proto_depIdxs,
		MessageInfos:      file_smallbiznis_location_v1_location_resource_proto_msgTypes,
	}.Build()
	File_smallbiznis_location_v1_location_resource_proto = out.File
	file_smallbiznis_location_v1_location_resource_proto_rawDesc = nil
	file_smallbiznis_location_v1_location_resource_proto_goTypes = nil
	file_smallbiznis_location_v1_location_resource_proto_depIdxs = nil
}

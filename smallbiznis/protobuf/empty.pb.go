// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/protobuf/empty.proto

package protobuf

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_protobuf_empty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_protobuf_empty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_smallbiznis_protobuf_empty_proto_rawDescGZIP(), []int{0}
}

var File_smallbiznis_protobuf_empty_proto protoreflect.FileDescriptor

var file_smallbiznis_protobuf_empty_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_protobuf_empty_proto_rawDescOnce sync.Once
	file_smallbiznis_protobuf_empty_proto_rawDescData = file_smallbiznis_protobuf_empty_proto_rawDesc
)

func file_smallbiznis_protobuf_empty_proto_rawDescGZIP() []byte {
	file_smallbiznis_protobuf_empty_proto_rawDescOnce.Do(func() {
		file_smallbiznis_protobuf_empty_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_protobuf_empty_proto_rawDescData)
	})
	return file_smallbiznis_protobuf_empty_proto_rawDescData
}

var file_smallbiznis_protobuf_empty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_smallbiznis_protobuf_empty_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: smallbiznis.protobuf.Empty
}
var file_smallbiznis_protobuf_empty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_smallbiznis_protobuf_empty_proto_init() }
func file_smallbiznis_protobuf_empty_proto_init() {
	if File_smallbiznis_protobuf_empty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_protobuf_empty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_smallbiznis_protobuf_empty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_protobuf_empty_proto_goTypes,
		DependencyIndexes: file_smallbiznis_protobuf_empty_proto_depIdxs,
		MessageInfos:      file_smallbiznis_protobuf_empty_proto_msgTypes,
	}.Build()
	File_smallbiznis_protobuf_empty_proto = out.File
	file_smallbiznis_protobuf_empty_proto_rawDesc = nil
	file_smallbiznis_protobuf_empty_proto_goTypes = nil
	file_smallbiznis_protobuf_empty_proto_depIdxs = nil
}

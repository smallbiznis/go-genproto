// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/common/enums.proto

package common

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

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_CREDIT                       TransactionType = 1
	TransactionType_DEBIT                        TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSACTION_TYPE_UNSPECIFIED",
		1: "CREDIT",
		2: "DEBIT",
	}
	TransactionType_value = map[string]int32{
		"TRANSACTION_TYPE_UNSPECIFIED": 0,
		"CREDIT":                       1,
		"DEBIT":                        2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_common_enums_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_smallbiznis_common_enums_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_common_enums_proto_rawDescGZIP(), []int{0}
}

type TransactionSubType int32

const (
	TransactionSubType_TRANSACTION_SUB_TYPE_UNSPECIFIED TransactionSubType = 0
	TransactionSubType_EARNING                          TransactionSubType = 1
	TransactionSubType_REDEEM                           TransactionSubType = 2
	TransactionSubType_EXPIRY                           TransactionSubType = 3
	TransactionSubType_ADJUSTMENT                       TransactionSubType = 4
	TransactionSubType_TRANSFER_IN                      TransactionSubType = 5
	TransactionSubType_TRANSFER_OUT                     TransactionSubType = 6
)

// Enum value maps for TransactionSubType.
var (
	TransactionSubType_name = map[int32]string{
		0: "TRANSACTION_SUB_TYPE_UNSPECIFIED",
		1: "EARNING",
		2: "REDEEM",
		3: "EXPIRY",
		4: "ADJUSTMENT",
		5: "TRANSFER_IN",
		6: "TRANSFER_OUT",
	}
	TransactionSubType_value = map[string]int32{
		"TRANSACTION_SUB_TYPE_UNSPECIFIED": 0,
		"EARNING":                          1,
		"REDEEM":                           2,
		"EXPIRY":                           3,
		"ADJUSTMENT":                       4,
		"TRANSFER_IN":                      5,
		"TRANSFER_OUT":                     6,
	}
)

func (x TransactionSubType) Enum() *TransactionSubType {
	p := new(TransactionSubType)
	*p = x
	return p
}

func (x TransactionSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_common_enums_proto_enumTypes[1].Descriptor()
}

func (TransactionSubType) Type() protoreflect.EnumType {
	return &file_smallbiznis_common_enums_proto_enumTypes[1]
}

func (x TransactionSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionSubType.Descriptor instead.
func (TransactionSubType) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_common_enums_proto_rawDescGZIP(), []int{1}
}

type EarnRuleType int32

const (
	EarnRuleType_EARN_RULE_TYPE_UNSPECIFIED EarnRuleType = 0
	EarnRuleType_TRANSACTION                EarnRuleType = 1 // default belanja
	EarnRuleType_SIGNUP                     EarnRuleType = 2
	EarnRuleType_REFERRAL                   EarnRuleType = 3
	EarnRuleType_BIRTHDAY                   EarnRuleType = 4
	EarnRuleType_CUSTOM_EVENT               EarnRuleType = 5
)

// Enum value maps for EarnRuleType.
var (
	EarnRuleType_name = map[int32]string{
		0: "EARN_RULE_TYPE_UNSPECIFIED",
		1: "TRANSACTION",
		2: "SIGNUP",
		3: "REFERRAL",
		4: "BIRTHDAY",
		5: "CUSTOM_EVENT",
	}
	EarnRuleType_value = map[string]int32{
		"EARN_RULE_TYPE_UNSPECIFIED": 0,
		"TRANSACTION":                1,
		"SIGNUP":                     2,
		"REFERRAL":                   3,
		"BIRTHDAY":                   4,
		"CUSTOM_EVENT":               5,
	}
)

func (x EarnRuleType) Enum() *EarnRuleType {
	p := new(EarnRuleType)
	*p = x
	return p
}

func (x EarnRuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EarnRuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_common_enums_proto_enumTypes[2].Descriptor()
}

func (EarnRuleType) Type() protoreflect.EnumType {
	return &file_smallbiznis_common_enums_proto_enumTypes[2]
}

func (x EarnRuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EarnRuleType.Descriptor instead.
func (EarnRuleType) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_common_enums_proto_rawDescGZIP(), []int{2}
}

type RewardType int32

const (
	RewardType_REWARD_TYPE_UNSPECIFIED RewardType = 0
	RewardType_FIXED_POINT             RewardType = 1 // Mendapatkan point tetap
	RewardType_MULTIPLIER              RewardType = 2 // Dikalikan dengan amount (misal: 0.01 * 100_000)
)

// Enum value maps for RewardType.
var (
	RewardType_name = map[int32]string{
		0: "REWARD_TYPE_UNSPECIFIED",
		1: "FIXED_POINT",
		2: "MULTIPLIER",
	}
	RewardType_value = map[string]int32{
		"REWARD_TYPE_UNSPECIFIED": 0,
		"FIXED_POINT":             1,
		"MULTIPLIER":              2,
	}
)

func (x RewardType) Enum() *RewardType {
	p := new(RewardType)
	*p = x
	return p
}

func (x RewardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RewardType) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_common_enums_proto_enumTypes[3].Descriptor()
}

func (RewardType) Type() protoreflect.EnumType {
	return &file_smallbiznis_common_enums_proto_enumTypes[3]
}

func (x RewardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RewardType.Descriptor instead.
func (RewardType) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_common_enums_proto_rawDescGZIP(), []int{3}
}

type RewardUnit int32

const (
	RewardUnit_REWARD_UNIT_UNSPECIFIED RewardUnit = 0
	RewardUnit_POINT                   RewardUnit = 1 // Default, satuan poin tetap
	RewardUnit_PERCENTAGE              RewardUnit = 2 // Digunakan untuk multiplier berbasis persen (misal: 1% = 0.01)
)

// Enum value maps for RewardUnit.
var (
	RewardUnit_name = map[int32]string{
		0: "REWARD_UNIT_UNSPECIFIED",
		1: "POINT",
		2: "PERCENTAGE",
	}
	RewardUnit_value = map[string]int32{
		"REWARD_UNIT_UNSPECIFIED": 0,
		"POINT":                   1,
		"PERCENTAGE":              2,
	}
)

func (x RewardUnit) Enum() *RewardUnit {
	p := new(RewardUnit)
	*p = x
	return p
}

func (x RewardUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RewardUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_common_enums_proto_enumTypes[4].Descriptor()
}

func (RewardUnit) Type() protoreflect.EnumType {
	return &file_smallbiznis_common_enums_proto_enumTypes[4]
}

func (x RewardUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RewardUnit.Descriptor instead.
func (RewardUnit) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_common_enums_proto_rawDescGZIP(), []int{4}
}

var File_smallbiznis_common_enums_proto protoreflect.FileDescriptor

var file_smallbiznis_common_enums_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x4a, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45,
	0x44, 0x49, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x49, 0x54, 0x10, 0x02,
	0x2a, 0x92, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45,
	0x44, 0x45, 0x45, 0x4d, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59,
	0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f,
	0x4f, 0x55, 0x54, 0x10, 0x06, 0x2a, 0x79, 0x0a, 0x0c, 0x45, 0x61, 0x72, 0x6e, 0x52, 0x75, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x41, 0x52, 0x4e, 0x5f, 0x52, 0x55,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x46, 0x45, 0x52, 0x52, 0x41, 0x4c, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x49, 0x52, 0x54, 0x48, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x10,
	0x0a, 0x0c, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x05,
	0x2a, 0x4a, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46,
	0x49, 0x58, 0x45, 0x44, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x49, 0x45, 0x52, 0x10, 0x02, 0x2a, 0x44, 0x0a, 0x0a,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45,
	0x57, 0x41, 0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x4f, 0x49, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41, 0x47, 0x45,
	0x10, 0x02, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69,
	0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_common_enums_proto_rawDescOnce sync.Once
	file_smallbiznis_common_enums_proto_rawDescData = file_smallbiznis_common_enums_proto_rawDesc
)

func file_smallbiznis_common_enums_proto_rawDescGZIP() []byte {
	file_smallbiznis_common_enums_proto_rawDescOnce.Do(func() {
		file_smallbiznis_common_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_common_enums_proto_rawDescData)
	})
	return file_smallbiznis_common_enums_proto_rawDescData
}

var file_smallbiznis_common_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_smallbiznis_common_enums_proto_goTypes = []interface{}{
	(TransactionType)(0),    // 0: smallbiznis.common.TransactionType
	(TransactionSubType)(0), // 1: smallbiznis.common.TransactionSubType
	(EarnRuleType)(0),       // 2: smallbiznis.common.EarnRuleType
	(RewardType)(0),         // 3: smallbiznis.common.RewardType
	(RewardUnit)(0),         // 4: smallbiznis.common.RewardUnit
}
var file_smallbiznis_common_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_smallbiznis_common_enums_proto_init() }
func file_smallbiznis_common_enums_proto_init() {
	if File_smallbiznis_common_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_smallbiznis_common_enums_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_common_enums_proto_goTypes,
		DependencyIndexes: file_smallbiznis_common_enums_proto_depIdxs,
		EnumInfos:         file_smallbiznis_common_enums_proto_enumTypes,
	}.Build()
	File_smallbiznis_common_enums_proto = out.File
	file_smallbiznis_common_enums_proto_rawDesc = nil
	file_smallbiznis_common_enums_proto_goTypes = nil
	file_smallbiznis_common_enums_proto_depIdxs = nil
}

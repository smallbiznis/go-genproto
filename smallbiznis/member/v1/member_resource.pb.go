// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: smallbiznis/member/v1/member_resource.proto

package member

import (
	v1 "github.com/smallbiznis/go-genproto/smallbiznis/organization/v1"
	v11 "github.com/smallbiznis/go-genproto/smallbiznis/user/v1"
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

type Role int32

const (
	Role_ROLE_ADMIN Role = 0
	// E-Commerce Management
	Role_ROLE_PRODUCT_MANAGER  Role = 1
	Role_ROLE_CATEGORY_MANAGER Role = 2
	// Marketing and Sales
	Role_ROLE_MARKETING_MANAGER Role = 3
	// Customer Service and Support
	Role_ROLE_CUSTOMER_SERVICE_MANAGER        Role = 6
	Role_ROLE_CUSTOMER_SUPPORT_REPRESENTATIVE Role = 7
	// Finance and Accounting
	Role_ROLE_ACCOUNTANT      Role = 8
	Role_ROLE_FINANCE_ANALYST Role = 9
	// IT
	Role_ROLE_DEVELOPER Role = 10
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0:  "ROLE_ADMIN",
		1:  "ROLE_PRODUCT_MANAGER",
		2:  "ROLE_CATEGORY_MANAGER",
		3:  "ROLE_MARKETING_MANAGER",
		6:  "ROLE_CUSTOMER_SERVICE_MANAGER",
		7:  "ROLE_CUSTOMER_SUPPORT_REPRESENTATIVE",
		8:  "ROLE_ACCOUNTANT",
		9:  "ROLE_FINANCE_ANALYST",
		10: "ROLE_DEVELOPER",
	}
	Role_value = map[string]int32{
		"ROLE_ADMIN":                           0,
		"ROLE_PRODUCT_MANAGER":                 1,
		"ROLE_CATEGORY_MANAGER":                2,
		"ROLE_MARKETING_MANAGER":               3,
		"ROLE_CUSTOMER_SERVICE_MANAGER":        6,
		"ROLE_CUSTOMER_SUPPORT_REPRESENTATIVE": 7,
		"ROLE_ACCOUNTANT":                      8,
		"ROLE_FINANCE_ANALYST":                 9,
		"ROLE_DEVELOPER":                       10,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_member_v1_member_resource_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_smallbiznis_member_v1_member_resource_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_resource_proto_rawDescGZIP(), []int{0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId     string                 `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Organization *v1.Organization       `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	User         *v11.User              `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Roles        []Role                 `protobuf:"varint,4,rep,packed,name=roles,proto3,enum=smallbiznis.member.v1.Role" json:"roles,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_member_v1_member_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_member_v1_member_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_smallbiznis_member_v1_member_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Member) GetOrganization() *v1.Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *Member) GetUser() *v11.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Member) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Member) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Member) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_smallbiznis_member_v1_member_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_member_v1_member_resource_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4d, 0x0a,
	0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x2a, 0xf7, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4d, 0x41,
	0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x21,
	0x0a, 0x1d, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10,
	0x06, 0x12, 0x28, 0x0a, 0x24, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x52, 0x45,
	0x53, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x52,
	0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x08,
	0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4e, 0x43, 0x45,
	0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x53, 0x54, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x10, 0x0a, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61,
	0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_member_v1_member_resource_proto_rawDescOnce sync.Once
	file_smallbiznis_member_v1_member_resource_proto_rawDescData = file_smallbiznis_member_v1_member_resource_proto_rawDesc
)

func file_smallbiznis_member_v1_member_resource_proto_rawDescGZIP() []byte {
	file_smallbiznis_member_v1_member_resource_proto_rawDescOnce.Do(func() {
		file_smallbiznis_member_v1_member_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_member_v1_member_resource_proto_rawDescData)
	})
	return file_smallbiznis_member_v1_member_resource_proto_rawDescData
}

var file_smallbiznis_member_v1_member_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_smallbiznis_member_v1_member_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_smallbiznis_member_v1_member_resource_proto_goTypes = []any{
	(Role)(0),                     // 0: smallbiznis.member.v1.Role
	(*Member)(nil),                // 1: smallbiznis.member.v1.Member
	(*v1.Organization)(nil),       // 2: smallbiznis.organization.v1.Organization
	(*v11.User)(nil),              // 3: smallbiznis.user.v1.User
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_smallbiznis_member_v1_member_resource_proto_depIdxs = []int32{
	2, // 0: smallbiznis.member.v1.Member.organization:type_name -> smallbiznis.organization.v1.Organization
	3, // 1: smallbiznis.member.v1.Member.user:type_name -> smallbiznis.user.v1.User
	0, // 2: smallbiznis.member.v1.Member.roles:type_name -> smallbiznis.member.v1.Role
	4, // 3: smallbiznis.member.v1.Member.created_at:type_name -> google.protobuf.Timestamp
	4, // 4: smallbiznis.member.v1.Member.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_smallbiznis_member_v1_member_resource_proto_init() }
func file_smallbiznis_member_v1_member_resource_proto_init() {
	if File_smallbiznis_member_v1_member_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_member_v1_member_resource_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_smallbiznis_member_v1_member_resource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_member_v1_member_resource_proto_goTypes,
		DependencyIndexes: file_smallbiznis_member_v1_member_resource_proto_depIdxs,
		EnumInfos:         file_smallbiznis_member_v1_member_resource_proto_enumTypes,
		MessageInfos:      file_smallbiznis_member_v1_member_resource_proto_msgTypes,
	}.Build()
	File_smallbiznis_member_v1_member_resource_proto = out.File
	file_smallbiznis_member_v1_member_resource_proto_rawDesc = nil
	file_smallbiznis_member_v1_member_resource_proto_goTypes = nil
	file_smallbiznis_member_v1_member_resource_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.23.4
// source: smallbiznis/item/v1/item_resource.proto

package item

import (
	v1 "github.com/smallbiznis/go-genproto/smallbiznis/inventory/v1"
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

type Type int32

const (
	Type_physical Type = 0
	Type_menu     Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "physical",
		1: "menu",
	}
	Type_value = map[string]int32{
		"physical": 0,
		"menu":     1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_item_v1_item_resource_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_smallbiznis_item_v1_item_resource_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_active   Status = 0
	Status_draft    Status = 1
	Status_archived Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "active",
		1: "draft",
		2: "archived",
	}
	Status_value = map[string]int32{
		"active":   0,
		"draft":    1,
		"archived": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_item_v1_item_resource_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_smallbiznis_item_v1_item_resource_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{1}
}

type WeightUnit int32

const (
	WeightUnit_kg WeightUnit = 0
	WeightUnit_g  WeightUnit = 1
)

// Enum value maps for WeightUnit.
var (
	WeightUnit_name = map[int32]string{
		0: "kg",
		1: "g",
	}
	WeightUnit_value = map[string]int32{
		"kg": 0,
		"g":  1,
	}
)

func (x WeightUnit) Enum() *WeightUnit {
	p := new(WeightUnit)
	*p = x
	return p
}

func (x WeightUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeightUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_smallbiznis_item_v1_item_resource_proto_enumTypes[2].Descriptor()
}

func (WeightUnit) Type() protoreflect.EnumType {
	return &file_smallbiznis_item_v1_item_resource_proto_enumTypes[2]
}

func (x WeightUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeightUnit.Descriptor instead.
func (WeightUnit) EnumDescriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{2}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId     string      `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	OrganizationId string      `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	ParentId       string      `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name           string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Children       []*Category `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Category) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Category) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetChildren() []*Category {
	if x != nil {
		return x.Children
	}
	return nil
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionId       string   `protobuf:"bytes,1,opt,name=option_id,json=optionId,proto3" json:"option_id,omitempty"`
	OrganizationId string   `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	OptionName     string   `protobuf:"bytes,3,opt,name=option_name,json=optionName,proto3" json:"option_name,omitempty"`
	OptionValues   []string `protobuf:"bytes,4,rep,name=option_values,json=optionValues,proto3" json:"option_values,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Option) GetOptionId() string {
	if x != nil {
		return x.OptionId
	}
	return ""
}

func (x *Option) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Option) GetOptionName() string {
	if x != nil {
		return x.OptionName
	}
	return ""
}

func (x *Option) GetOptionValues() []string {
	if x != nil {
		return x.OptionValues
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId          string `protobuf:"bytes,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	OrganizationId string `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Value          string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{2}
}

func (x *Tag) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *Tag) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId         string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	OrganizationId string                 `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Type           Type                   `protobuf:"varint,3,opt,name=type,proto3,enum=smallbiznis.item.v1.Type" json:"type,omitempty"`
	Slug           string                 `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Title          string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	BodyHtml       string                 `protobuf:"bytes,6,opt,name=body_html,json=bodyHtml,proto3" json:"body_html,omitempty"`
	Status         Status                 `protobuf:"varint,7,opt,name=status,proto3,enum=smallbiznis.item.v1.Status" json:"status,omitempty"`
	Options        []*Item_Option         `protobuf:"bytes,8,rep,name=options,proto3" json:"options,omitempty"`
	Variants       []*Variant             `protobuf:"bytes,9,rep,name=variants,proto3" json:"variants,omitempty"`
	Tags           []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Item) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Item) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_physical
}

func (x *Item) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetBodyHtml() string {
	if x != nil {
		return x.BodyHtml
	}
	return ""
}

func (x *Item) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_active
}

func (x *Item) GetOptions() []*Item_Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Item) GetVariants() []*Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *Item) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Item) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Item) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VariantId       string                 `protobuf:"bytes,1,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	ItemId          string                 `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Sku             string                 `protobuf:"bytes,4,opt,name=sku,proto3" json:"sku,omitempty"`
	Title           string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Taxable         bool                   `protobuf:"varint,6,opt,name=taxable,proto3" json:"taxable,omitempty"`
	Price           float32                `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	CompareAtPrice  float32                `protobuf:"fixed32,8,opt,name=compare_at_price,json=compareAtPrice,proto3" json:"compare_at_price,omitempty"`
	Cost            float32                `protobuf:"fixed32,9,opt,name=cost,proto3" json:"cost,omitempty"`
	Profit          float32                `protobuf:"fixed32,10,opt,name=profit,proto3" json:"profit,omitempty"`
	Margin          float32                `protobuf:"fixed32,12,opt,name=margin,proto3" json:"margin,omitempty"`
	Weight          float32                `protobuf:"fixed32,13,opt,name=weight,proto3" json:"weight,omitempty"`
	WeightUnit      WeightUnit             `protobuf:"varint,14,opt,name=weight_unit,json=weightUnit,proto3,enum=smallbiznis.item.v1.WeightUnit" json:"weight_unit,omitempty"`
	Attributes      []string               `protobuf:"bytes,15,rep,name=attributes,proto3" json:"attributes,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Barcode         string                 `protobuf:"bytes,19,opt,name=barcode,proto3" json:"barcode,omitempty"`
	PreparationTime int32                  `protobuf:"varint,20,opt,name=preparation_time,json=preparationTime,proto3" json:"preparation_time,omitempty"`
	Inventories     []*v1.Inventory        `protobuf:"bytes,21,rep,name=inventories,proto3" json:"inventories,omitempty"`
}

func (x *Variant) Reset() {
	*x = Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{4}
}

func (x *Variant) GetVariantId() string {
	if x != nil {
		return x.VariantId
	}
	return ""
}

func (x *Variant) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Variant) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Variant) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Variant) GetTaxable() bool {
	if x != nil {
		return x.Taxable
	}
	return false
}

func (x *Variant) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Variant) GetCompareAtPrice() float32 {
	if x != nil {
		return x.CompareAtPrice
	}
	return 0
}

func (x *Variant) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Variant) GetProfit() float32 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *Variant) GetMargin() float32 {
	if x != nil {
		return x.Margin
	}
	return 0
}

func (x *Variant) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Variant) GetWeightUnit() WeightUnit {
	if x != nil {
		return x.WeightUnit
	}
	return WeightUnit_kg
}

func (x *Variant) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Variant) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Variant) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Variant) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *Variant) GetPreparationTime() int32 {
	if x != nil {
		return x.PreparationTime
	}
	return 0
}

func (x *Variant) GetInventories() []*v1.Inventory {
	if x != nil {
		return x.Inventories
	}
	return nil
}

type Item_Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionId   string   `protobuf:"bytes,1,opt,name=option_id,json=optionId,proto3" json:"option_id,omitempty"`
	OptionName string   `protobuf:"bytes,2,opt,name=option_name,json=optionName,proto3" json:"option_name,omitempty"`
	Position   int32    `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Values     []string `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Item_Option) Reset() {
	*x = Item_Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item_Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item_Option) ProtoMessage() {}

func (x *Item_Option) ProtoReflect() protoreflect.Message {
	mi := &file_smallbiznis_item_v1_item_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item_Option.ProtoReflect.Descriptor instead.
func (*Item_Option) Descriptor() ([]byte, []int) {
	return file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Item_Option) GetOptionId() string {
	if x != nil {
		return x.OptionId
	}
	return ""
}

func (x *Item_Option) GetOptionName() string {
	if x != nil {
		return x.OptionName
	}
	return ""
}

func (x *Item_Option) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Item_Option) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_smallbiznis_item_v1_item_resource_proto protoreflect.FileDescriptor

var file_smallbiznis_item_v1_item_resource_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x69, 0x74,
	0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x31,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x5b, 0x0a,
	0x03, 0x54, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xef, 0x04, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69,
	0x73, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x33, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x08,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x1a, 0x7a, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x88, 0x05, 0x0a,
	0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x41, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x0a, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x45, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a,
	0x6e, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2a, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x6d, 0x65, 0x6e, 0x75, 0x10, 0x01, 0x2a, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x1b, 0x0a, 0x0a, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x6b, 0x67, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01,
	0x67, 0x10, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x69, 0x7a, 0x6e, 0x69, 0x73, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x74,
	0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smallbiznis_item_v1_item_resource_proto_rawDescOnce sync.Once
	file_smallbiznis_item_v1_item_resource_proto_rawDescData = file_smallbiznis_item_v1_item_resource_proto_rawDesc
)

func file_smallbiznis_item_v1_item_resource_proto_rawDescGZIP() []byte {
	file_smallbiznis_item_v1_item_resource_proto_rawDescOnce.Do(func() {
		file_smallbiznis_item_v1_item_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_smallbiznis_item_v1_item_resource_proto_rawDescData)
	})
	return file_smallbiznis_item_v1_item_resource_proto_rawDescData
}

var file_smallbiznis_item_v1_item_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_smallbiznis_item_v1_item_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_smallbiznis_item_v1_item_resource_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: smallbiznis.item.v1.Type
	(Status)(0),                   // 1: smallbiznis.item.v1.Status
	(WeightUnit)(0),               // 2: smallbiznis.item.v1.WeightUnit
	(*Category)(nil),              // 3: smallbiznis.item.v1.Category
	(*Option)(nil),                // 4: smallbiznis.item.v1.Option
	(*Tag)(nil),                   // 5: smallbiznis.item.v1.Tag
	(*Item)(nil),                  // 6: smallbiznis.item.v1.Item
	(*Variant)(nil),               // 7: smallbiznis.item.v1.Variant
	(*Item_Option)(nil),           // 8: smallbiznis.item.v1.Item.Option
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*v1.Inventory)(nil),          // 10: smallbiznis.inventory.v1.Inventory
}
var file_smallbiznis_item_v1_item_resource_proto_depIdxs = []int32{
	3,  // 0: smallbiznis.item.v1.Category.children:type_name -> smallbiznis.item.v1.Category
	0,  // 1: smallbiznis.item.v1.Item.type:type_name -> smallbiznis.item.v1.Type
	1,  // 2: smallbiznis.item.v1.Item.status:type_name -> smallbiznis.item.v1.Status
	8,  // 3: smallbiznis.item.v1.Item.options:type_name -> smallbiznis.item.v1.Item.Option
	7,  // 4: smallbiznis.item.v1.Item.variants:type_name -> smallbiznis.item.v1.Variant
	9,  // 5: smallbiznis.item.v1.Item.created_at:type_name -> google.protobuf.Timestamp
	9,  // 6: smallbiznis.item.v1.Item.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 7: smallbiznis.item.v1.Variant.weight_unit:type_name -> smallbiznis.item.v1.WeightUnit
	9,  // 8: smallbiznis.item.v1.Variant.created_at:type_name -> google.protobuf.Timestamp
	9,  // 9: smallbiznis.item.v1.Variant.updated_at:type_name -> google.protobuf.Timestamp
	10, // 10: smallbiznis.item.v1.Variant.inventories:type_name -> smallbiznis.inventory.v1.Inventory
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_smallbiznis_item_v1_item_resource_proto_init() }
func file_smallbiznis_item_v1_item_resource_proto_init() {
	if File_smallbiznis_item_v1_item_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smallbiznis_item_v1_item_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_smallbiznis_item_v1_item_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_smallbiznis_item_v1_item_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_smallbiznis_item_v1_item_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_smallbiznis_item_v1_item_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant); i {
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
		file_smallbiznis_item_v1_item_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item_Option); i {
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
			RawDescriptor: file_smallbiznis_item_v1_item_resource_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smallbiznis_item_v1_item_resource_proto_goTypes,
		DependencyIndexes: file_smallbiznis_item_v1_item_resource_proto_depIdxs,
		EnumInfos:         file_smallbiznis_item_v1_item_resource_proto_enumTypes,
		MessageInfos:      file_smallbiznis_item_v1_item_resource_proto_msgTypes,
	}.Build()
	File_smallbiznis_item_v1_item_resource_proto = out.File
	file_smallbiznis_item_v1_item_resource_proto_rawDesc = nil
	file_smallbiznis_item_v1_item_resource_proto_goTypes = nil
	file_smallbiznis_item_v1_item_resource_proto_depIdxs = nil
}

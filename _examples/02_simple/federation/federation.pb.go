// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: federation/federation.proto

package federation

import (
	_ "example/post"
	_ "example/user"
	_ "github.com/mercari/grpc-federation/grpc/federation"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Item_ItemType int32

const (
	Item_ITEM_TYPE_UNSPECIFIED Item_ItemType = 0
	Item_ITEM_TYPE_1           Item_ItemType = 1
	Item_ITEM_TYPE_2           Item_ItemType = 2
	Item_ITEM_TYPE_3           Item_ItemType = 3
)

// Enum value maps for Item_ItemType.
var (
	Item_ItemType_name = map[int32]string{
		0: "ITEM_TYPE_UNSPECIFIED",
		1: "ITEM_TYPE_1",
		2: "ITEM_TYPE_2",
		3: "ITEM_TYPE_3",
	}
	Item_ItemType_value = map[string]int32{
		"ITEM_TYPE_UNSPECIFIED": 0,
		"ITEM_TYPE_1":           1,
		"ITEM_TYPE_2":           2,
		"ITEM_TYPE_3":           3,
	}
)

func (x Item_ItemType) Enum() *Item_ItemType {
	p := new(Item_ItemType)
	*p = x
	return p
}

func (x Item_ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Item_ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_federation_federation_proto_enumTypes[0].Descriptor()
}

func (Item_ItemType) Type() protoreflect.EnumType {
	return &file_federation_federation_proto_enumTypes[0]
}

func (x Item_ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Item_ItemType.Descriptor instead.
func (Item_ItemType) EnumDescriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{4, 0}
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	Str  string `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *GetPostResponse) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	User    *User  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{2}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Items   []*Item               `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Profile map[string]*anypb.Any `protobuf:"bytes,4,rep,name=profile,proto3" json:"profile,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Attr:
	//
	//	*User_AttrA_
	//	*User_B
	Attr isUser_Attr `protobuf_oneof:"attr"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *User) GetProfile() map[string]*anypb.Any {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (m *User) GetAttr() isUser_Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (x *User) GetAttrA() *User_AttrA {
	if x, ok := x.GetAttr().(*User_AttrA_); ok {
		return x.AttrA
	}
	return nil
}

func (x *User) GetB() *User_AttrB {
	if x, ok := x.GetAttr().(*User_B); ok {
		return x.B
	}
	return nil
}

type isUser_Attr interface {
	isUser_Attr()
}

type User_AttrA_ struct {
	AttrA *User_AttrA `protobuf:"bytes,5,opt,name=attr_a,json=attrA,proto3,oneof"`
}

type User_B struct {
	B *User_AttrB `protobuf:"bytes,6,opt,name=b,proto3,oneof"`
}

func (*User_AttrA_) isUser_Attr() {}

func (*User_B) isUser_Attr() {}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     Item_ItemType  `protobuf:"varint,2,opt,name=type,proto3,enum=federation.Item_ItemType" json:"type,omitempty"`
	Value    int64          `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	Location *Item_Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[4]
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
	return file_federation_federation_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetType() Item_ItemType {
	if x != nil {
		return x.Type
	}
	return Item_ITEM_TYPE_UNSPECIFIED
}

func (x *Item) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Item) GetLocation() *Item_Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type User_AttrA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *User_AttrA) Reset() {
	*x = User_AttrA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_AttrA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_AttrA) ProtoMessage() {}

func (x *User_AttrA) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_AttrA.ProtoReflect.Descriptor instead.
func (*User_AttrA) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{3, 1}
}

func (x *User_AttrA) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

type User_AttrB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bar bool `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *User_AttrB) Reset() {
	*x = User_AttrB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_AttrB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_AttrB) ProtoMessage() {}

func (x *User_AttrB) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_AttrB.ProtoReflect.Descriptor instead.
func (*User_AttrB) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{3, 2}
}

func (x *User_AttrB) GetBar() bool {
	if x != nil {
		return x.Bar
	}
	return false
}

type Item_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr1 string `protobuf:"bytes,1,opt,name=addr1,proto3" json:"addr1,omitempty"`
	Addr2 string `protobuf:"bytes,2,opt,name=addr2,proto3" json:"addr2,omitempty"`
	// Types that are assignable to Addr3:
	//
	//	*Item_Location_AddrA_
	//	*Item_Location_B
	Addr3 isItem_Location_Addr3 `protobuf_oneof:"addr3"`
}

func (x *Item_Location) Reset() {
	*x = Item_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item_Location) ProtoMessage() {}

func (x *Item_Location) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item_Location.ProtoReflect.Descriptor instead.
func (*Item_Location) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Item_Location) GetAddr1() string {
	if x != nil {
		return x.Addr1
	}
	return ""
}

func (x *Item_Location) GetAddr2() string {
	if x != nil {
		return x.Addr2
	}
	return ""
}

func (m *Item_Location) GetAddr3() isItem_Location_Addr3 {
	if m != nil {
		return m.Addr3
	}
	return nil
}

func (x *Item_Location) GetAddrA() *Item_Location_AddrA {
	if x, ok := x.GetAddr3().(*Item_Location_AddrA_); ok {
		return x.AddrA
	}
	return nil
}

func (x *Item_Location) GetB() *Item_Location_AddrB {
	if x, ok := x.GetAddr3().(*Item_Location_B); ok {
		return x.B
	}
	return nil
}

type isItem_Location_Addr3 interface {
	isItem_Location_Addr3()
}

type Item_Location_AddrA_ struct {
	AddrA *Item_Location_AddrA `protobuf:"bytes,3,opt,name=addr_a,json=addrA,proto3,oneof"`
}

type Item_Location_B struct {
	B *Item_Location_AddrB `protobuf:"bytes,4,opt,name=b,proto3,oneof"`
}

func (*Item_Location_AddrA_) isItem_Location_Addr3() {}

func (*Item_Location_B) isItem_Location_Addr3() {}

type Item_Location_AddrA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *Item_Location_AddrA) Reset() {
	*x = Item_Location_AddrA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item_Location_AddrA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item_Location_AddrA) ProtoMessage() {}

func (x *Item_Location_AddrA) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item_Location_AddrA.ProtoReflect.Descriptor instead.
func (*Item_Location_AddrA) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{4, 0, 0}
}

func (x *Item_Location_AddrA) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

type Item_Location_AddrB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bar int64 `protobuf:"varint,1,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *Item_Location_AddrB) Reset() {
	*x = Item_Location_AddrB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item_Location_AddrB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item_Location_AddrB) ProtoMessage() {}

func (x *Item_Location_AddrB) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item_Location_AddrB.ProtoReflect.Descriptor instead.
func (*Item_Location_AddrB) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{4, 0, 1}
}

func (x *Item_Location_AddrB) GetBar() int64 {
	if x != nil {
		return x.Bar
	}
	return 0
}

var File_federation_federation_proto protoreflect.FileDescriptor

var file_federation_federation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x0a,
	0x82, 0x97, 0x22, 0x06, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x82,
	0x97, 0x22, 0x08, 0x82, 0x01, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x03, 0x73, 0x74, 0x72,
	0x3a, 0x1e, 0x82, 0x97, 0x22, 0x1a, 0x12, 0x18, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x22, 0x0a, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x04, 0x24, 0x2e, 0x69, 0x64,
	0x22, 0xdb, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0a, 0x82, 0x97, 0x22, 0x06, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x61, 0x82, 0x97, 0x22,
	0x5d, 0x0a, 0x45, 0x0a, 0x18, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0a, 0x0a,
	0x02, 0x69, 0x64, 0x12, 0x04, 0x24, 0x2e, 0x69, 0x64, 0x1a, 0x0e, 0x0a, 0x04, 0x70, 0x6f, 0x73,
	0x74, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x22, 0x03, 0x31, 0x30, 0x73, 0x2a, 0x08,
	0x0a, 0x06, 0x0a, 0x02, 0x32, 0x73, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x06, 0x1a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x8b,
	0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x61, 0x74, 0x74, 0x72, 0x5f, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x41, 0x48, 0x00, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x41, 0x12, 0x26, 0x0a,
	0x01, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x42,
	0x48, 0x00, 0x52, 0x01, 0x62, 0x1a, 0x50, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x30, 0x0a, 0x05, 0x41, 0x74, 0x74, 0x72, 0x41,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66,
	0x6f, 0x6f, 0x3a, 0x15, 0x82, 0x97, 0x22, 0x11, 0x22, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x41, 0x1a, 0x30, 0x0a, 0x05, 0x41, 0x74, 0x74,
	0x72, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x62, 0x61, 0x72, 0x3a, 0x15, 0x82, 0x97, 0x22, 0x11, 0x22, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x42, 0x3a, 0x67, 0x82, 0x97, 0x22,
	0x63, 0x0a, 0x61, 0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x0a,
	0x02, 0x69, 0x64, 0x12, 0x09, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x1a, 0x0e,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x22, 0x03,
	0x32, 0x30, 0x73, 0x2a, 0x1f, 0x12, 0x1d, 0x0a, 0x02, 0x31, 0x73, 0x11, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0xe6, 0x3f, 0x19, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xfb, 0x3f, 0x22, 0x03, 0x33,
	0x30, 0x73, 0x28, 0x03, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x22, 0xd8, 0x04, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xba, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x32, 0x12, 0x38,
	0x0a, 0x06, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x41, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x41, 0x12, 0x2f, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x42, 0x48, 0x00, 0x52, 0x01, 0x62, 0x1a, 0x39, 0x0a, 0x05, 0x41, 0x64, 0x64,
	0x72, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x66, 0x6f, 0x6f, 0x3a, 0x1e, 0x82, 0x97, 0x22, 0x1a, 0x22, 0x18, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x41, 0x1a, 0x39, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x72, 0x42, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x62, 0x61, 0x72, 0x3a,
	0x1e, 0x82, 0x97, 0x22, 0x1a, 0x22, 0x18, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x42, 0x3a,
	0x18, 0x82, 0x97, 0x22, 0x14, 0x22, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x72, 0x33, 0x22, 0x72, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x54, 0x45,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x32, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x54, 0x45, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x33, 0x10, 0x03, 0x1a, 0x18, 0x82, 0x97,
	0x22, 0x14, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x0f, 0x82, 0x97, 0x22, 0x0b, 0x22, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x32, 0xa3, 0x01, 0x0a, 0x11, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x1a, 0x48, 0x82, 0x97, 0x22, 0x44, 0x0a, 0x20, 0x0a, 0x0c, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x0a, 0x20, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x88, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0xe2, 0x02, 0x16, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_federation_federation_proto_rawDescOnce sync.Once
	file_federation_federation_proto_rawDescData = file_federation_federation_proto_rawDesc
)

func file_federation_federation_proto_rawDescGZIP() []byte {
	file_federation_federation_proto_rawDescOnce.Do(func() {
		file_federation_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_federation_federation_proto_rawDescData)
	})
	return file_federation_federation_proto_rawDescData
}

var file_federation_federation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_federation_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_federation_federation_proto_goTypes = []interface{}{
	(Item_ItemType)(0),          // 0: federation.Item.ItemType
	(*GetPostRequest)(nil),      // 1: federation.GetPostRequest
	(*GetPostResponse)(nil),     // 2: federation.GetPostResponse
	(*Post)(nil),                // 3: federation.Post
	(*User)(nil),                // 4: federation.User
	(*Item)(nil),                // 5: federation.Item
	nil,                         // 6: federation.User.ProfileEntry
	(*User_AttrA)(nil),          // 7: federation.User.AttrA
	(*User_AttrB)(nil),          // 8: federation.User.AttrB
	(*Item_Location)(nil),       // 9: federation.Item.Location
	(*Item_Location_AddrA)(nil), // 10: federation.Item.Location.AddrA
	(*Item_Location_AddrB)(nil), // 11: federation.Item.Location.AddrB
	(*anypb.Any)(nil),           // 12: google.protobuf.Any
}
var file_federation_federation_proto_depIdxs = []int32{
	3,  // 0: federation.GetPostResponse.post:type_name -> federation.Post
	4,  // 1: federation.Post.user:type_name -> federation.User
	5,  // 2: federation.User.items:type_name -> federation.Item
	6,  // 3: federation.User.profile:type_name -> federation.User.ProfileEntry
	7,  // 4: federation.User.attr_a:type_name -> federation.User.AttrA
	8,  // 5: federation.User.b:type_name -> federation.User.AttrB
	0,  // 6: federation.Item.type:type_name -> federation.Item.ItemType
	9,  // 7: federation.Item.location:type_name -> federation.Item.Location
	12, // 8: federation.User.ProfileEntry.value:type_name -> google.protobuf.Any
	10, // 9: federation.Item.Location.addr_a:type_name -> federation.Item.Location.AddrA
	11, // 10: federation.Item.Location.b:type_name -> federation.Item.Location.AddrB
	1,  // 11: federation.FederationService.GetPost:input_type -> federation.GetPostRequest
	2,  // 12: federation.FederationService.GetPost:output_type -> federation.GetPostResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_federation_federation_proto_init() }
func file_federation_federation_proto_init() {
	if File_federation_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_federation_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_federation_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResponse); i {
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
		file_federation_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_federation_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_federation_federation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_federation_federation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_AttrA); i {
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
		file_federation_federation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_AttrB); i {
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
		file_federation_federation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item_Location); i {
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
		file_federation_federation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item_Location_AddrA); i {
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
		file_federation_federation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item_Location_AddrB); i {
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
	file_federation_federation_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*User_AttrA_)(nil),
		(*User_B)(nil),
	}
	file_federation_federation_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Item_Location_AddrA_)(nil),
		(*Item_Location_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_federation_federation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_federation_federation_proto_goTypes,
		DependencyIndexes: file_federation_federation_proto_depIdxs,
		EnumInfos:         file_federation_federation_proto_enumTypes,
		MessageInfos:      file_federation_federation_proto_msgTypes,
	}.Build()
	File_federation_federation_proto = out.File
	file_federation_federation_proto_rawDesc = nil
	file_federation_federation_proto_goTypes = nil
	file_federation_federation_proto_depIdxs = nil
}

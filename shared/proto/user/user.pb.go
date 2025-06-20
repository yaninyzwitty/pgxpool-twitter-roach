// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0--rc3
// source: user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
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
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetUserByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIdRequest) Reset() {
	*x = GetUserByIdRequest{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRequest) ProtoMessage() {}

func (x *GetUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIdResponse) Reset() {
	*x = GetUserByIdResponse{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdResponse) ProtoMessage() {}

func (x *GetUserByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIdResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIdResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByEmailRequest) Reset() {
	*x = GetUserByEmailRequest{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailRequest) ProtoMessage() {}

func (x *GetUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*GetUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUserByEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByEmailResponse) Reset() {
	*x = GetUserByEmailResponse{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailResponse) ProtoMessage() {}

func (x *GetUserByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailResponse.ProtoReflect.Descriptor instead.
func (*GetUserByEmailResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // FIXED
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type StreamUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamUsersRequest) Reset() {
	*x = StreamUsersRequest{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUsersRequest) ProtoMessage() {}

func (x *StreamUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamUsersRequest.ProtoReflect.Descriptor instead.
func (*StreamUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *StreamUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type StreamUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamUsersResponse) Reset() {
	*x = StreamUsersResponse{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUsersResponse) ProtoMessage() {}

func (x *StreamUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamUsersResponse.ProtoReflect.Descriptor instead.
func (*StreamUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *StreamUsersResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\x04user\x1a\x1fgoogle/protobuf/timestamp.proto\"\x83\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"$\n" +
	"\x12GetUserByIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"5\n" +
	"\x13GetUserByIdResponse\x12\x1e\n" +
	"\x04user\x18\x01 \x01(\v2\n" +
	".user.UserR\x04user\"-\n" +
	"\x15GetUserByEmailRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"8\n" +
	"\x16GetUserByEmailResponse\x12\x1e\n" +
	"\x04user\x18\x01 \x01(\v2\n" +
	".user.UserR\x04user\"?\n" +
	"\x0fGetUsersRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x05R\x06offset\"4\n" +
	"\x10GetUsersResponse\x12 \n" +
	"\x05users\x18\x01 \x03(\v2\n" +
	".user.UserR\x05users\"*\n" +
	"\x12StreamUsersRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\"5\n" +
	"\x13StreamUsersResponse\x12\x1e\n" +
	"\x04user\x18\x01 \x01(\v2\n" +
	".user.UserR\x04user2\x9f\x02\n" +
	"\vUserService\x12B\n" +
	"\vGetUserById\x12\x18.user.GetUserByIdRequest\x1a\x19.user.GetUserByIdResponse\x12K\n" +
	"\x0eGetUserByEmail\x12\x1b.user.GetUserByEmailRequest\x1a\x1c.user.GetUserByEmailResponse\x129\n" +
	"\bGetUsers\x12\x15.user.GetUsersRequest\x1a\x16.user.GetUsersResponse\x12D\n" +
	"\vStreamUsers\x12\x18.user.StreamUsersRequest\x1a\x19.user.StreamUsersResponse0\x01B\x18Z\x16shared/proto/user;userb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_proto_goTypes = []any{
	(*User)(nil),                   // 0: user.User
	(*GetUserByIdRequest)(nil),     // 1: user.GetUserByIdRequest
	(*GetUserByIdResponse)(nil),    // 2: user.GetUserByIdResponse
	(*GetUserByEmailRequest)(nil),  // 3: user.GetUserByEmailRequest
	(*GetUserByEmailResponse)(nil), // 4: user.GetUserByEmailResponse
	(*GetUsersRequest)(nil),        // 5: user.GetUsersRequest
	(*GetUsersResponse)(nil),       // 6: user.GetUsersResponse
	(*StreamUsersRequest)(nil),     // 7: user.StreamUsersRequest
	(*StreamUsersResponse)(nil),    // 8: user.StreamUsersResponse
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	9, // 0: user.User.updated_at:type_name -> google.protobuf.Timestamp
	0, // 1: user.GetUserByIdResponse.user:type_name -> user.User
	0, // 2: user.GetUserByEmailResponse.user:type_name -> user.User
	0, // 3: user.GetUsersResponse.users:type_name -> user.User
	0, // 4: user.StreamUsersResponse.user:type_name -> user.User
	1, // 5: user.UserService.GetUserById:input_type -> user.GetUserByIdRequest
	3, // 6: user.UserService.GetUserByEmail:input_type -> user.GetUserByEmailRequest
	5, // 7: user.UserService.GetUsers:input_type -> user.GetUsersRequest
	7, // 8: user.UserService.StreamUsers:input_type -> user.StreamUsersRequest
	2, // 9: user.UserService.GetUserById:output_type -> user.GetUserByIdResponse
	4, // 10: user.UserService.GetUserByEmail:output_type -> user.GetUserByEmailResponse
	6, // 11: user.UserService.GetUsers:output_type -> user.GetUsersResponse
	8, // 12: user.UserService.StreamUsers:output_type -> user.StreamUsersResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: protos/v1/user/user.proto

package user

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

type UserMeesage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Age         int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserMeesage) Reset() {
	*x = UserMeesage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeesage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeesage) ProtoMessage() {}

func (x *UserMeesage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeesage.ProtoReflect.Descriptor instead.
func (*UserMeesage) Descriptor() ([]byte, []int) {
	return file_protos_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserMeesage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserMeesage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserMeesage) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserMeesage) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMessage *UserMeesage `protobuf:"bytes,1,opt,name=user_message,json=userMessage,proto3" json:"user_message,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserResponse) GetUserMessage() *UserMeesage {
	if x != nil {
		return x.UserMessage
	}
	return nil
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_user_user_proto_rawDescGZIP(), []int{3}
}

type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMessages []*UserMeesage `protobuf:"bytes,1,rep,name=user_messages,json=userMessages,proto3" json:"user_messages,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserResponse) GetUserMessages() []*UserMeesage {
	if x != nil {
		return x.UserMessages
	}
	return nil
}

var File_protos_v1_user_user_proto protoreflect.FileDescriptor

var file_protos_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x22, 0x6f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x65, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x65, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x11, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x65, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x85,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6e, 0x67, 0x67, 0x79, 0x75, 0x4c, 0x69, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x61, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_user_user_proto_rawDescOnce sync.Once
	file_protos_v1_user_user_proto_rawDescData = file_protos_v1_user_user_proto_rawDesc
)

func file_protos_v1_user_user_proto_rawDescGZIP() []byte {
	file_protos_v1_user_user_proto_rawDescOnce.Do(func() {
		file_protos_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_user_user_proto_rawDescData)
	})
	return file_protos_v1_user_user_proto_rawDescData
}

var file_protos_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_v1_user_user_proto_goTypes = []interface{}{
	(*UserMeesage)(nil),      // 0: grpcapi.UserMeesage
	(*GetUserRequest)(nil),   // 1: grpcapi.GetUserRequest
	(*GetUserResponse)(nil),  // 2: grpcapi.GetUserResponse
	(*ListUserRequest)(nil),  // 3: grpcapi.ListUserRequest
	(*ListUserResponse)(nil), // 4: grpcapi.ListUserResponse
}
var file_protos_v1_user_user_proto_depIdxs = []int32{
	0, // 0: grpcapi.GetUserResponse.user_message:type_name -> grpcapi.UserMeesage
	0, // 1: grpcapi.ListUserResponse.user_messages:type_name -> grpcapi.UserMeesage
	1, // 2: grpcapi.User.GetUser:input_type -> grpcapi.GetUserRequest
	3, // 3: grpcapi.User.ListUser:input_type -> grpcapi.ListUserRequest
	2, // 4: grpcapi.User.GetUser:output_type -> grpcapi.GetUserResponse
	4, // 5: grpcapi.User.ListUser:output_type -> grpcapi.ListUserResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_v1_user_user_proto_init() }
func file_protos_v1_user_user_proto_init() {
	if File_protos_v1_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeesage); i {
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
		file_protos_v1_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_protos_v1_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_protos_v1_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_protos_v1_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponse); i {
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
			RawDescriptor: file_protos_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_user_user_proto_goTypes,
		DependencyIndexes: file_protos_v1_user_user_proto_depIdxs,
		MessageInfos:      file_protos_v1_user_user_proto_msgTypes,
	}.Build()
	File_protos_v1_user_user_proto = out.File
	file_protos_v1_user_user_proto_rawDesc = nil
	file_protos_v1_user_user_proto_goTypes = nil
	file_protos_v1_user_user_proto_depIdxs = nil
}
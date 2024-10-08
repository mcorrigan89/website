// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: serviceapis/messaging/v1/identity.proto

package messagingv1

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

type SendVerificationEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VerificationLink string `protobuf:"bytes,2,opt,name=verification_link,json=verificationLink,proto3" json:"verification_link,omitempty"`
}

func (x *SendVerificationEmailRequest) Reset() {
	*x = SendVerificationEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationEmailRequest) ProtoMessage() {}

func (x *SendVerificationEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationEmailRequest.ProtoReflect.Descriptor instead.
func (*SendVerificationEmailRequest) Descriptor() ([]byte, []int) {
	return file_serviceapis_messaging_v1_identity_proto_rawDescGZIP(), []int{0}
}

func (x *SendVerificationEmailRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendVerificationEmailRequest) GetVerificationLink() string {
	if x != nil {
		return x.VerificationLink
	}
	return ""
}

type SendVerificationEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendVerificationEmailResponse) Reset() {
	*x = SendVerificationEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationEmailResponse) ProtoMessage() {}

func (x *SendVerificationEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationEmailResponse.ProtoReflect.Descriptor instead.
func (*SendVerificationEmailResponse) Descriptor() ([]byte, []int) {
	return file_serviceapis_messaging_v1_identity_proto_rawDescGZIP(), []int{1}
}

func (x *SendVerificationEmailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendPasswordResetEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PasswordResetLink string `protobuf:"bytes,2,opt,name=password_reset_link,json=passwordResetLink,proto3" json:"password_reset_link,omitempty"`
}

func (x *SendPasswordResetEmailRequest) Reset() {
	*x = SendPasswordResetEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPasswordResetEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPasswordResetEmailRequest) ProtoMessage() {}

func (x *SendPasswordResetEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPasswordResetEmailRequest.ProtoReflect.Descriptor instead.
func (*SendPasswordResetEmailRequest) Descriptor() ([]byte, []int) {
	return file_serviceapis_messaging_v1_identity_proto_rawDescGZIP(), []int{2}
}

func (x *SendPasswordResetEmailRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendPasswordResetEmailRequest) GetPasswordResetLink() string {
	if x != nil {
		return x.PasswordResetLink
	}
	return ""
}

type SendPasswordResetEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendPasswordResetEmailResponse) Reset() {
	*x = SendPasswordResetEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPasswordResetEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPasswordResetEmailResponse) ProtoMessage() {}

func (x *SendPasswordResetEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serviceapis_messaging_v1_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPasswordResetEmailResponse.ProtoReflect.Descriptor instead.
func (*SendPasswordResetEmailResponse) Descriptor() ([]byte, []int) {
	return file_serviceapis_messaging_v1_identity_proto_rawDescGZIP(), []int{3}
}

func (x *SendPasswordResetEmailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_serviceapis_messaging_v1_identity_proto protoreflect.FileDescriptor

var file_serviceapis_messaging_v1_identity_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x22, 0x64, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x39, 0x0a, 0x1d, 0x53, 0x65, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x68, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e,
	0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x3a,
	0x0a, 0x1e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb7, 0x02, 0x0a, 0x18, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x8d, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xf8, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x6f, 0x72, 0x72, 0x69, 0x67, 0x61, 0x6e, 0x38, 0x39, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x53, 0x4d, 0x58, 0xaa, 0x02, 0x18, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x18, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceapis_messaging_v1_identity_proto_rawDescOnce sync.Once
	file_serviceapis_messaging_v1_identity_proto_rawDescData = file_serviceapis_messaging_v1_identity_proto_rawDesc
)

func file_serviceapis_messaging_v1_identity_proto_rawDescGZIP() []byte {
	file_serviceapis_messaging_v1_identity_proto_rawDescOnce.Do(func() {
		file_serviceapis_messaging_v1_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceapis_messaging_v1_identity_proto_rawDescData)
	})
	return file_serviceapis_messaging_v1_identity_proto_rawDescData
}

var file_serviceapis_messaging_v1_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_serviceapis_messaging_v1_identity_proto_goTypes = []any{
	(*SendVerificationEmailRequest)(nil),   // 0: serviceapis.messaging.v1.SendVerificationEmailRequest
	(*SendVerificationEmailResponse)(nil),  // 1: serviceapis.messaging.v1.SendVerificationEmailResponse
	(*SendPasswordResetEmailRequest)(nil),  // 2: serviceapis.messaging.v1.SendPasswordResetEmailRequest
	(*SendPasswordResetEmailResponse)(nil), // 3: serviceapis.messaging.v1.SendPasswordResetEmailResponse
}
var file_serviceapis_messaging_v1_identity_proto_depIdxs = []int32{
	0, // 0: serviceapis.messaging.v1.IdentityMessagingService.SendVerificationEmail:input_type -> serviceapis.messaging.v1.SendVerificationEmailRequest
	2, // 1: serviceapis.messaging.v1.IdentityMessagingService.SendPasswordResetEmail:input_type -> serviceapis.messaging.v1.SendPasswordResetEmailRequest
	1, // 2: serviceapis.messaging.v1.IdentityMessagingService.SendVerificationEmail:output_type -> serviceapis.messaging.v1.SendVerificationEmailResponse
	3, // 3: serviceapis.messaging.v1.IdentityMessagingService.SendPasswordResetEmail:output_type -> serviceapis.messaging.v1.SendPasswordResetEmailResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serviceapis_messaging_v1_identity_proto_init() }
func file_serviceapis_messaging_v1_identity_proto_init() {
	if File_serviceapis_messaging_v1_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceapis_messaging_v1_identity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendVerificationEmailRequest); i {
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
		file_serviceapis_messaging_v1_identity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SendVerificationEmailResponse); i {
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
		file_serviceapis_messaging_v1_identity_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SendPasswordResetEmailRequest); i {
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
		file_serviceapis_messaging_v1_identity_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendPasswordResetEmailResponse); i {
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
			RawDescriptor: file_serviceapis_messaging_v1_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serviceapis_messaging_v1_identity_proto_goTypes,
		DependencyIndexes: file_serviceapis_messaging_v1_identity_proto_depIdxs,
		MessageInfos:      file_serviceapis_messaging_v1_identity_proto_msgTypes,
	}.Build()
	File_serviceapis_messaging_v1_identity_proto = out.File
	file_serviceapis_messaging_v1_identity_proto_rawDesc = nil
	file_serviceapis_messaging_v1_identity_proto_goTypes = nil
	file_serviceapis_messaging_v1_identity_proto_depIdxs = nil
}

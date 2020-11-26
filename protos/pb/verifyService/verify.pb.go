// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.7.1
// source: verify.proto

package verifyService

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ValidateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ValidateEmailRequest) Reset() {
	*x = ValidateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateEmailRequest) ProtoMessage() {}

func (x *ValidateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateEmailRequest.ProtoReflect.Descriptor instead.
func (*ValidateEmailRequest) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateEmailRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ValidateEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ValidateEmailRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{1}
}

func (x *SuccessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OtpVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match   bool   `protobuf:"varint,1,opt,name=match,proto3" json:"match,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OtpVerificationResponse) Reset() {
	*x = OtpVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpVerificationResponse) ProtoMessage() {}

func (x *OtpVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpVerificationResponse.ProtoReflect.Descriptor instead.
func (*OtpVerificationResponse) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{2}
}

func (x *OtpVerificationResponse) GetMatch() bool {
	if x != nil {
		return x.Match
	}
	return false
}

func (x *OtpVerificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OtpVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *OtpVerificationRequest) Reset() {
	*x = OtpVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpVerificationRequest) ProtoMessage() {}

func (x *OtpVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpVerificationRequest.ProtoReflect.Descriptor instead.
func (*OtpVerificationRequest) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{3}
}

func (x *OtpVerificationRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OtpVerificationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type OtpVerificationByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *OtpVerificationByEmailRequest) Reset() {
	*x = OtpVerificationByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpVerificationByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpVerificationByEmailRequest) ProtoMessage() {}

func (x *OtpVerificationByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpVerificationByEmailRequest.ProtoReflect.Descriptor instead.
func (*OtpVerificationByEmailRequest) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{4}
}

func (x *OtpVerificationByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OtpVerificationByEmailRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ResendOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ResendOtpRequest) Reset() {
	*x = ResendOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendOtpRequest) ProtoMessage() {}

func (x *ResendOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendOtpRequest.ProtoReflect.Descriptor instead.
func (*ResendOtpRequest) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{5}
}

func (x *ResendOtpRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_verify_proto protoreflect.FileDescriptor

var file_verify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x5c, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x49, 0x0a, 0x17, 0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a, 0x16,
	0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x49, 0x0a, 0x1d, 0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x52,
	0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0xca, 0x02, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x6d, 0x73, 0x4f, 0x74,
	0x70, 0x12, 0x1e, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x4f, 0x74, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x4f, 0x74, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x4f, 0x74, 0x70, 0x12, 0x25, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x4f, 0x74, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x12, 0x18, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verify_proto_rawDescOnce sync.Once
	file_verify_proto_rawDescData = file_verify_proto_rawDesc
)

func file_verify_proto_rawDescGZIP() []byte {
	file_verify_proto_rawDescOnce.Do(func() {
		file_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_verify_proto_rawDescData)
	})
	return file_verify_proto_rawDescData
}

var file_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_verify_proto_goTypes = []interface{}{
	(*ValidateEmailRequest)(nil),          // 0: verify.ValidateEmailRequest
	(*SuccessResponse)(nil),               // 1: verify.SuccessResponse
	(*OtpVerificationResponse)(nil),       // 2: verify.OtpVerificationResponse
	(*OtpVerificationRequest)(nil),        // 3: verify.OtpVerificationRequest
	(*OtpVerificationByEmailRequest)(nil), // 4: verify.OtpVerificationByEmailRequest
	(*ResendOtpRequest)(nil),              // 5: verify.ResendOtpRequest
}
var file_verify_proto_depIdxs = []int32{
	0, // 0: verify.VerifyService.ValidateEmail:input_type -> verify.ValidateEmailRequest
	3, // 1: verify.VerifyService.VerifySmsOtp:input_type -> verify.OtpVerificationRequest
	4, // 2: verify.VerifyService.VerifyEmailOtp:input_type -> verify.OtpVerificationByEmailRequest
	5, // 3: verify.VerifyService.ResendOtp:input_type -> verify.ResendOtpRequest
	1, // 4: verify.VerifyService.ValidateEmail:output_type -> verify.SuccessResponse
	2, // 5: verify.VerifyService.VerifySmsOtp:output_type -> verify.OtpVerificationResponse
	2, // 6: verify.VerifyService.VerifyEmailOtp:output_type -> verify.OtpVerificationResponse
	1, // 7: verify.VerifyService.ResendOtp:output_type -> verify.SuccessResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_verify_proto_init() }
func file_verify_proto_init() {
	if File_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateEmailRequest); i {
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
		file_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_verify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpVerificationResponse); i {
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
		file_verify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpVerificationRequest); i {
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
		file_verify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpVerificationByEmailRequest); i {
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
		file_verify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendOtpRequest); i {
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
			RawDescriptor: file_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verify_proto_goTypes,
		DependencyIndexes: file_verify_proto_depIdxs,
		MessageInfos:      file_verify_proto_msgTypes,
	}.Build()
	File_verify_proto = out.File
	file_verify_proto_rawDesc = nil
	file_verify_proto_goTypes = nil
	file_verify_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VerifyServiceClient is the client API for VerifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifyServiceClient interface {
	ValidateEmail(ctx context.Context, in *ValidateEmailRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	VerifySmsOtp(ctx context.Context, in *OtpVerificationRequest, opts ...grpc.CallOption) (*OtpVerificationResponse, error)
	VerifyEmailOtp(ctx context.Context, in *OtpVerificationByEmailRequest, opts ...grpc.CallOption) (*OtpVerificationResponse, error)
	ResendOtp(ctx context.Context, in *ResendOtpRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type verifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyServiceClient(cc grpc.ClientConnInterface) VerifyServiceClient {
	return &verifyServiceClient{cc}
}

func (c *verifyServiceClient) ValidateEmail(ctx context.Context, in *ValidateEmailRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/verify.VerifyService/ValidateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyServiceClient) VerifySmsOtp(ctx context.Context, in *OtpVerificationRequest, opts ...grpc.CallOption) (*OtpVerificationResponse, error) {
	out := new(OtpVerificationResponse)
	err := c.cc.Invoke(ctx, "/verify.VerifyService/VerifySmsOtp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyServiceClient) VerifyEmailOtp(ctx context.Context, in *OtpVerificationByEmailRequest, opts ...grpc.CallOption) (*OtpVerificationResponse, error) {
	out := new(OtpVerificationResponse)
	err := c.cc.Invoke(ctx, "/verify.VerifyService/VerifyEmailOtp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyServiceClient) ResendOtp(ctx context.Context, in *ResendOtpRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/verify.VerifyService/ResendOtp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyServiceServer is the server API for VerifyService service.
type VerifyServiceServer interface {
	ValidateEmail(context.Context, *ValidateEmailRequest) (*SuccessResponse, error)
	VerifySmsOtp(context.Context, *OtpVerificationRequest) (*OtpVerificationResponse, error)
	VerifyEmailOtp(context.Context, *OtpVerificationByEmailRequest) (*OtpVerificationResponse, error)
	ResendOtp(context.Context, *ResendOtpRequest) (*SuccessResponse, error)
}

// UnimplementedVerifyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVerifyServiceServer struct {
}

func (*UnimplementedVerifyServiceServer) ValidateEmail(context.Context, *ValidateEmailRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateEmail not implemented")
}
func (*UnimplementedVerifyServiceServer) VerifySmsOtp(context.Context, *OtpVerificationRequest) (*OtpVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySmsOtp not implemented")
}
func (*UnimplementedVerifyServiceServer) VerifyEmailOtp(context.Context, *OtpVerificationByEmailRequest) (*OtpVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmailOtp not implemented")
}
func (*UnimplementedVerifyServiceServer) ResendOtp(context.Context, *ResendOtpRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendOtp not implemented")
}

func RegisterVerifyServiceServer(s *grpc.Server, srv VerifyServiceServer) {
	s.RegisterService(&_VerifyService_serviceDesc, srv)
}

func _VerifyService_ValidateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServiceServer).ValidateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verify.VerifyService/ValidateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServiceServer).ValidateEmail(ctx, req.(*ValidateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyService_VerifySmsOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OtpVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServiceServer).VerifySmsOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verify.VerifyService/VerifySmsOtp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServiceServer).VerifySmsOtp(ctx, req.(*OtpVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyService_VerifyEmailOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OtpVerificationByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServiceServer).VerifyEmailOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verify.VerifyService/VerifyEmailOtp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServiceServer).VerifyEmailOtp(ctx, req.(*OtpVerificationByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyService_ResendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServiceServer).ResendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verify.VerifyService/ResendOtp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServiceServer).ResendOtp(ctx, req.(*ResendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "verify.VerifyService",
	HandlerType: (*VerifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateEmail",
			Handler:    _VerifyService_ValidateEmail_Handler,
		},
		{
			MethodName: "VerifySmsOtp",
			Handler:    _VerifyService_VerifySmsOtp_Handler,
		},
		{
			MethodName: "VerifyEmailOtp",
			Handler:    _VerifyService_VerifyEmailOtp_Handler,
		},
		{
			MethodName: "ResendOtp",
			Handler:    _VerifyService_ResendOtp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verify.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: onboarding.proto

package onboardingService

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

type CreatePhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string  `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Device      *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreatePhoneRequest) Reset() {
	*x = CreatePhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhoneRequest) ProtoMessage() {}

func (x *CreatePhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhoneRequest.ProtoReflect.Descriptor instead.
func (*CreatePhoneRequest) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePhoneRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreatePhoneRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os          string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Brand       string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	DeviceToken string `protobuf:"bytes,3,opt,name=deviceToken,proto3" json:"deviceToken,omitempty"`
	DeviceId    string `protobuf:"bytes,4,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Device) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Device) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
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
		mi := &file_onboarding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[2]
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
	return file_onboarding_proto_rawDescGZIP(), []int{2}
}

func (x *SuccessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreatePhoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	EmailToken string `protobuf:"bytes,2,opt,name=email_token,json=emailToken,proto3" json:"email_token,omitempty"`
}

func (x *CreatePhoneResponse) Reset() {
	*x = CreatePhoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhoneResponse) ProtoMessage() {}

func (x *CreatePhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhoneResponse.ProtoReflect.Descriptor instead.
func (*CreatePhoneResponse) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePhoneResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreatePhoneResponse) GetEmailToken() string {
	if x != nil {
		return x.EmailToken
	}
	return ""
}

type CreatePasscodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId string `protobuf:"bytes,1,opt,name=personId,proto3" json:"personId,omitempty"`
	Passcode string `protobuf:"bytes,2,opt,name=passcode,proto3" json:"passcode,omitempty"`
}

func (x *CreatePasscodeRequest) Reset() {
	*x = CreatePasscodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePasscodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePasscodeRequest) ProtoMessage() {}

func (x *CreatePasscodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePasscodeRequest.ProtoReflect.Descriptor instead.
func (*CreatePasscodeRequest) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePasscodeRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *CreatePasscodeRequest) GetPasscode() string {
	if x != nil {
		return x.Passcode
	}
	return ""
}

type CreateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value      string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EmailToken string `protobuf:"bytes,3,opt,name=emailToken,proto3" json:"emailToken,omitempty"`
	Primary    bool   `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Verified   bool   `protobuf:"varint,5,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *CreateEmailRequest) Reset() {
	*x = CreateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailRequest) ProtoMessage() {}

func (x *CreateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailRequest.ProtoReflect.Descriptor instead.
func (*CreateEmailRequest) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{5}
}

func (x *CreateEmailRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEmailRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CreateEmailRequest) GetEmailToken() string {
	if x != nil {
		return x.EmailToken
	}
	return ""
}

func (x *CreateEmailRequest) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *CreateEmailRequest) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type UpdatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId  string `protobuf:"bytes,1,opt,name=personId,proto3" json:"personId,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Dob       string `protobuf:"bytes,5,opt,name=dob,proto3" json:"dob,omitempty"`
}

func (x *UpdatePersonRequest) Reset() {
	*x = UpdatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonRequest) ProtoMessage() {}

func (x *UpdatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersonRequest) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePersonRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *UpdatePersonRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdatePersonRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdatePersonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdatePersonRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

type RoavaReasonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId string `protobuf:"bytes,1,opt,name=personId,proto3" json:"personId,omitempty"`
	Reasons  string `protobuf:"bytes,2,opt,name=reasons,proto3" json:"reasons,omitempty"`
}

func (x *RoavaReasonsRequest) Reset() {
	*x = RoavaReasonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoavaReasonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoavaReasonsRequest) ProtoMessage() {}

func (x *RoavaReasonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoavaReasonsRequest.ProtoReflect.Descriptor instead.
func (*RoavaReasonsRequest) Descriptor() ([]byte, []int) {
	return file_onboarding_proto_rawDescGZIP(), []int{7}
}

func (x *RoavaReasonsRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *RoavaReasonsRequest) GetReasons() string {
	if x != nil {
		return x.Reasons
	}
	return ""
}

var File_onboarding_proto protoreflect.FileDescriptor

var file_onboarding_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x62,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x6c, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x4f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x94, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f,
	0x62, 0x22, 0x4b, 0x0a, 0x13, 0x52, 0x6f, 0x61, 0x76, 0x61, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x32, 0xb9,
	0x03, 0x0a, 0x11, 0x4f, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x69, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1f, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x59, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72,
	0x55, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x61, 0x76, 0x61, 0x12, 0x1f, 0x2e, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x61, 0x76, 0x61, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x62,
	0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_onboarding_proto_rawDescOnce sync.Once
	file_onboarding_proto_rawDescData = file_onboarding_proto_rawDesc
)

func file_onboarding_proto_rawDescGZIP() []byte {
	file_onboarding_proto_rawDescOnce.Do(func() {
		file_onboarding_proto_rawDescData = protoimpl.X.CompressGZIP(file_onboarding_proto_rawDescData)
	})
	return file_onboarding_proto_rawDescData
}

var file_onboarding_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_onboarding_proto_goTypes = []interface{}{
	(*CreatePhoneRequest)(nil),    // 0: onboarding.CreatePhoneRequest
	(*Device)(nil),                // 1: onboarding.Device
	(*SuccessResponse)(nil),       // 2: onboarding.SuccessResponse
	(*CreatePhoneResponse)(nil),   // 3: onboarding.CreatePhoneResponse
	(*CreatePasscodeRequest)(nil), // 4: onboarding.CreatePasscodeRequest
	(*CreateEmailRequest)(nil),    // 5: onboarding.CreateEmailRequest
	(*UpdatePersonRequest)(nil),   // 6: onboarding.UpdatePersonRequest
	(*RoavaReasonsRequest)(nil),   // 7: onboarding.RoavaReasonsRequest
}
var file_onboarding_proto_depIdxs = []int32{
	1, // 0: onboarding.CreatePhoneRequest.device:type_name -> onboarding.Device
	0, // 1: onboarding.OnBoardingService.CreatePhone:input_type -> onboarding.CreatePhoneRequest
	5, // 2: onboarding.OnBoardingService.CreateEmail:input_type -> onboarding.CreateEmailRequest
	4, // 3: onboarding.OnBoardingService.CreatePasscode:input_type -> onboarding.CreatePasscodeRequest
	6, // 4: onboarding.OnBoardingService.UpdatePersonBiodata:input_type -> onboarding.UpdatePersonRequest
	7, // 5: onboarding.OnBoardingService.AddReasonsForUsingRoava:input_type -> onboarding.RoavaReasonsRequest
	3, // 6: onboarding.OnBoardingService.CreatePhone:output_type -> onboarding.CreatePhoneResponse
	2, // 7: onboarding.OnBoardingService.CreateEmail:output_type -> onboarding.SuccessResponse
	2, // 8: onboarding.OnBoardingService.CreatePasscode:output_type -> onboarding.SuccessResponse
	2, // 9: onboarding.OnBoardingService.UpdatePersonBiodata:output_type -> onboarding.SuccessResponse
	2, // 10: onboarding.OnBoardingService.AddReasonsForUsingRoava:output_type -> onboarding.SuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_onboarding_proto_init() }
func file_onboarding_proto_init() {
	if File_onboarding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_onboarding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhoneRequest); i {
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
		file_onboarding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_onboarding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_onboarding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhoneResponse); i {
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
		file_onboarding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePasscodeRequest); i {
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
		file_onboarding_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailRequest); i {
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
		file_onboarding_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersonRequest); i {
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
		file_onboarding_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoavaReasonsRequest); i {
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
			RawDescriptor: file_onboarding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_onboarding_proto_goTypes,
		DependencyIndexes: file_onboarding_proto_depIdxs,
		MessageInfos:      file_onboarding_proto_msgTypes,
	}.Build()
	File_onboarding_proto = out.File
	file_onboarding_proto_rawDesc = nil
	file_onboarding_proto_goTypes = nil
	file_onboarding_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OnBoardingServiceClient is the client API for OnBoardingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnBoardingServiceClient interface {
	CreatePhone(ctx context.Context, in *CreatePhoneRequest, opts ...grpc.CallOption) (*CreatePhoneResponse, error)
	CreateEmail(ctx context.Context, in *CreateEmailRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	CreatePasscode(ctx context.Context, in *CreatePasscodeRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	UpdatePersonBiodata(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	AddReasonsForUsingRoava(ctx context.Context, in *RoavaReasonsRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type onBoardingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnBoardingServiceClient(cc grpc.ClientConnInterface) OnBoardingServiceClient {
	return &onBoardingServiceClient{cc}
}

func (c *onBoardingServiceClient) CreatePhone(ctx context.Context, in *CreatePhoneRequest, opts ...grpc.CallOption) (*CreatePhoneResponse, error) {
	out := new(CreatePhoneResponse)
	err := c.cc.Invoke(ctx, "/onboarding.OnBoardingService/CreatePhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onBoardingServiceClient) CreateEmail(ctx context.Context, in *CreateEmailRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/onboarding.OnBoardingService/CreateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onBoardingServiceClient) CreatePasscode(ctx context.Context, in *CreatePasscodeRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/onboarding.OnBoardingService/CreatePasscode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onBoardingServiceClient) UpdatePersonBiodata(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/onboarding.OnBoardingService/UpdatePersonBiodata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onBoardingServiceClient) AddReasonsForUsingRoava(ctx context.Context, in *RoavaReasonsRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/onboarding.OnBoardingService/AddReasonsForUsingRoava", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnBoardingServiceServer is the server API for OnBoardingService service.
type OnBoardingServiceServer interface {
	CreatePhone(context.Context, *CreatePhoneRequest) (*CreatePhoneResponse, error)
	CreateEmail(context.Context, *CreateEmailRequest) (*SuccessResponse, error)
	CreatePasscode(context.Context, *CreatePasscodeRequest) (*SuccessResponse, error)
	UpdatePersonBiodata(context.Context, *UpdatePersonRequest) (*SuccessResponse, error)
	AddReasonsForUsingRoava(context.Context, *RoavaReasonsRequest) (*SuccessResponse, error)
}

// UnimplementedOnBoardingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnBoardingServiceServer struct {
}

func (*UnimplementedOnBoardingServiceServer) CreatePhone(context.Context, *CreatePhoneRequest) (*CreatePhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePhone not implemented")
}
func (*UnimplementedOnBoardingServiceServer) CreateEmail(context.Context, *CreateEmailRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmail not implemented")
}
func (*UnimplementedOnBoardingServiceServer) CreatePasscode(context.Context, *CreatePasscodeRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePasscode not implemented")
}
func (*UnimplementedOnBoardingServiceServer) UpdatePersonBiodata(context.Context, *UpdatePersonRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonBiodata not implemented")
}
func (*UnimplementedOnBoardingServiceServer) AddReasonsForUsingRoava(context.Context, *RoavaReasonsRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReasonsForUsingRoava not implemented")
}

func RegisterOnBoardingServiceServer(s *grpc.Server, srv OnBoardingServiceServer) {
	s.RegisterService(&_OnBoardingService_serviceDesc, srv)
}

func _OnBoardingService_CreatePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnBoardingServiceServer).CreatePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboarding.OnBoardingService/CreatePhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnBoardingServiceServer).CreatePhone(ctx, req.(*CreatePhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnBoardingService_CreateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnBoardingServiceServer).CreateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboarding.OnBoardingService/CreateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnBoardingServiceServer).CreateEmail(ctx, req.(*CreateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnBoardingService_CreatePasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnBoardingServiceServer).CreatePasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboarding.OnBoardingService/CreatePasscode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnBoardingServiceServer).CreatePasscode(ctx, req.(*CreatePasscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnBoardingService_UpdatePersonBiodata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnBoardingServiceServer).UpdatePersonBiodata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboarding.OnBoardingService/UpdatePersonBiodata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnBoardingServiceServer).UpdatePersonBiodata(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnBoardingService_AddReasonsForUsingRoava_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoavaReasonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnBoardingServiceServer).AddReasonsForUsingRoava(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onboarding.OnBoardingService/AddReasonsForUsingRoava",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnBoardingServiceServer).AddReasonsForUsingRoava(ctx, req.(*RoavaReasonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnBoardingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onboarding.OnBoardingService",
	HandlerType: (*OnBoardingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePhone",
			Handler:    _OnBoardingService_CreatePhone_Handler,
		},
		{
			MethodName: "CreateEmail",
			Handler:    _OnBoardingService_CreateEmail_Handler,
		},
		{
			MethodName: "CreatePasscode",
			Handler:    _OnBoardingService_CreatePasscode_Handler,
		},
		{
			MethodName: "UpdatePersonBiodata",
			Handler:    _OnBoardingService_UpdatePersonBiodata_Handler,
		},
		{
			MethodName: "AddReasonsForUsingRoava",
			Handler:    _OnBoardingService_AddReasonsForUsingRoava_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onboarding.proto",
}

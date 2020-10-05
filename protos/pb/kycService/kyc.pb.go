// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: kyc.proto

package kycService

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{0}
}

type ApplicationIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *ApplicationIdRequest) Reset() {
	*x = ApplicationIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationIdRequest) ProtoMessage() {}

func (x *ApplicationIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationIdRequest.ProtoReflect.Descriptor instead.
func (*ApplicationIdRequest) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationIdRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type PersonIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId string `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *PersonIdRequest) Reset() {
	*x = PersonIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonIdRequest) ProtoMessage() {}

func (x *PersonIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonIdRequest.ProtoReflect.Descriptor instead.
func (*PersonIdRequest) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{2}
}

func (x *PersonIdRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

type Cdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner       string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Details     string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Kyc         *Kyc   `protobuf:"bytes,5,opt,name=kyc,proto3" json:"kyc,omitempty"`
	Aml         *Aml   `protobuf:"bytes,6,opt,name=aml,proto3" json:"aml,omitempty"`
	Roava       *Roava `protobuf:"bytes,7,opt,name=roava,proto3" json:"roava,omitempty"`
	TimeCreated int64  `protobuf:"varint,8,opt,name=time_created,json=timeCreated,proto3" json:"time_created,omitempty"`
	TimeUpdated int64  `protobuf:"varint,9,opt,name=time_updated,json=timeUpdated,proto3" json:"time_updated,omitempty"`
}

func (x *Cdd) Reset() {
	*x = Cdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cdd) ProtoMessage() {}

func (x *Cdd) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cdd.ProtoReflect.Descriptor instead.
func (*Cdd) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{3}
}

func (x *Cdd) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cdd) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Cdd) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Cdd) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Cdd) GetKyc() *Kyc {
	if x != nil {
		return x.Kyc
	}
	return nil
}

func (x *Cdd) GetAml() *Aml {
	if x != nil {
		return x.Aml
	}
	return nil
}

func (x *Cdd) GetRoava() *Roava {
	if x != nil {
		return x.Roava
	}
	return nil
}

func (x *Cdd) GetTimeCreated() int64 {
	if x != nil {
		return x.TimeCreated
	}
	return 0
}

func (x *Cdd) GetTimeUpdated() int64 {
	if x != nil {
		return x.TimeUpdated
	}
	return 0
}

type Kyc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Applicant   *Applicant `protobuf:"bytes,2,opt,name=applicant,proto3" json:"applicant,omitempty"`
	Status      string     `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Vendor      string     `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	TimeCreated int64      `protobuf:"varint,5,opt,name=time_created,json=timeCreated,proto3" json:"time_created,omitempty"`
	TimeUpdated int64      `protobuf:"varint,6,opt,name=time_updated,json=timeUpdated,proto3" json:"time_updated,omitempty"`
}

func (x *Kyc) Reset() {
	*x = Kyc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kyc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kyc) ProtoMessage() {}

func (x *Kyc) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kyc.ProtoReflect.Descriptor instead.
func (*Kyc) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{4}
}

func (x *Kyc) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Kyc) GetApplicant() *Applicant {
	if x != nil {
		return x.Applicant
	}
	return nil
}

func (x *Kyc) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Kyc) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *Kyc) GetTimeCreated() int64 {
	if x != nil {
		return x.TimeCreated
	}
	return 0
}

func (x *Kyc) GetTimeUpdated() int64 {
	if x != nil {
		return x.TimeUpdated
	}
	return 0
}

type Aml struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Aml) Reset() {
	*x = Aml{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aml) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aml) ProtoMessage() {}

func (x *Aml) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aml.ProtoReflect.Descriptor instead.
func (*Aml) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{5}
}

type Roava struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Roava) Reset() {
	*x = Roava{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roava) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roava) ProtoMessage() {}

func (x *Roava) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roava.ProtoReflect.Descriptor instead.
func (*Roava) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{6}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlatNumber     string `protobuf:"bytes,1,opt,name=flat_number,json=flatNumber,proto3" json:"flat_number,omitempty"`
	BuildingNumber string `protobuf:"bytes,2,opt,name=building_number,json=buildingNumber,proto3" json:"building_number,omitempty"`
	BuildingName   string `protobuf:"bytes,3,opt,name=building_name,json=buildingName,proto3" json:"building_name,omitempty"`
	Street         string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	SubStreet      string `protobuf:"bytes,5,opt,name=sub_street,json=subStreet,proto3" json:"sub_street,omitempty"`
	Town           string `protobuf:"bytes,6,opt,name=town,proto3" json:"town,omitempty"`
	State          string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Postcode       string `protobuf:"bytes,8,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Country        string `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{7}
}

func (x *Address) GetFlatNumber() string {
	if x != nil {
		return x.FlatNumber
	}
	return ""
}

func (x *Address) GetBuildingNumber() string {
	if x != nil {
		return x.BuildingNumber
	}
	return ""
}

func (x *Address) GetBuildingName() string {
	if x != nil {
		return x.BuildingName
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetSubStreet() string {
	if x != nil {
		return x.SubStreet
	}
	return ""
}

func (x *Address) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Applicant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId string   `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	ApplicantId   string   `protobuf:"bytes,2,opt,name=applicant_id,json=applicantId,proto3" json:"applicant_id,omitempty"`
	FirstName     string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Dob           string   `protobuf:"bytes,6,opt,name=dob,proto3" json:"dob,omitempty"`
	Address       *Address `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Vendor        string   `protobuf:"bytes,8,opt,name=vendor,proto3" json:"vendor,omitempty"`
}

func (x *Applicant) Reset() {
	*x = Applicant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Applicant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Applicant) ProtoMessage() {}

func (x *Applicant) ProtoReflect() protoreflect.Message {
	mi := &file_kyc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Applicant.ProtoReflect.Descriptor instead.
func (*Applicant) Descriptor() ([]byte, []int) {
	return file_kyc_proto_rawDescGZIP(), []int{8}
}

func (x *Applicant) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Applicant) GetApplicantId() string {
	if x != nil {
		return x.ApplicantId
	}
	return ""
}

func (x *Applicant) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Applicant) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Applicant) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Applicant) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Applicant) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Applicant) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

var File_kyc_proto protoreflect.FileDescriptor

var file_kyc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6b, 0x79, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6b, 0x79, 0x63,
	0x22, 0x06, 0x0a, 0x04, 0x76, 0x6f, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xfd, 0x01, 0x0a, 0x03, 0x63, 0x64, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x03, 0x6b, 0x79, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x6b, 0x79, 0x63, 0x52, 0x03,
	0x6b, 0x79, 0x63, 0x12, 0x1a, 0x0a, 0x03, 0x61, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x61, 0x6d, 0x6c, 0x52, 0x03, 0x61, 0x6d, 0x6c, 0x12,
	0x20, 0x0a, 0x05, 0x72, 0x6f, 0x61, 0x76, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x72, 0x6f, 0x61, 0x76, 0x61, 0x52, 0x05, 0x72, 0x6f, 0x61, 0x76,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x03, 0x6b, 0x79, 0x63, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x6e, 0x74, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x05, 0x0a, 0x03, 0x61, 0x6d, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x72, 0x6f,
	0x61, 0x76, 0x61, 0x22, 0x8f, 0x02, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xf9, 0x01, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f,
	0x62, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x32, 0xd1, 0x01, 0x0a, 0x0a, 0x4b, 0x79, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x6e, 0x74, 0x42, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x1e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4b,
	0x79, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x61, 0x77,
	0x61, 0x69, 0x74, 0x43, 0x44, 0x44, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x6b,
	0x79, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x63, 0x64,
	0x64, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x62, 0x2f, 0x6b, 0x79, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kyc_proto_rawDescOnce sync.Once
	file_kyc_proto_rawDescData = file_kyc_proto_rawDesc
)

func file_kyc_proto_rawDescGZIP() []byte {
	file_kyc_proto_rawDescOnce.Do(func() {
		file_kyc_proto_rawDescData = protoimpl.X.CompressGZIP(file_kyc_proto_rawDescData)
	})
	return file_kyc_proto_rawDescData
}

var file_kyc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kyc_proto_goTypes = []interface{}{
	(*Void)(nil),                 // 0: kyc.void
	(*ApplicationIdRequest)(nil), // 1: kyc.ApplicationIdRequest
	(*PersonIdRequest)(nil),      // 2: kyc.PersonIdRequest
	(*Cdd)(nil),                  // 3: kyc.cdd
	(*Kyc)(nil),                  // 4: kyc.kyc
	(*Aml)(nil),                  // 5: kyc.aml
	(*Roava)(nil),                // 6: kyc.roava
	(*Address)(nil),              // 7: kyc.address
	(*Applicant)(nil),            // 8: kyc.applicant
}
var file_kyc_proto_depIdxs = []int32{
	4, // 0: kyc.cdd.kyc:type_name -> kyc.kyc
	5, // 1: kyc.cdd.aml:type_name -> kyc.aml
	6, // 2: kyc.cdd.roava:type_name -> kyc.roava
	8, // 3: kyc.kyc.applicant:type_name -> kyc.applicant
	7, // 4: kyc.applicant.address:type_name -> kyc.address
	2, // 5: kyc.KycService.getKycApplicantByPersonId:input_type -> kyc.PersonIdRequest
	2, // 6: kyc.KycService.submitKycApplicationByPersonId:input_type -> kyc.PersonIdRequest
	1, // 7: kyc.KycService.awaitCDDReport:input_type -> kyc.ApplicationIdRequest
	8, // 8: kyc.KycService.getKycApplicantByPersonId:output_type -> kyc.applicant
	0, // 9: kyc.KycService.submitKycApplicationByPersonId:output_type -> kyc.void
	3, // 10: kyc.KycService.awaitCDDReport:output_type -> kyc.cdd
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kyc_proto_init() }
func file_kyc_proto_init() {
	if File_kyc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kyc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_kyc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationIdRequest); i {
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
		file_kyc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonIdRequest); i {
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
		file_kyc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cdd); i {
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
		file_kyc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kyc); i {
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
		file_kyc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aml); i {
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
		file_kyc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roava); i {
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
		file_kyc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_kyc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Applicant); i {
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
			RawDescriptor: file_kyc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kyc_proto_goTypes,
		DependencyIndexes: file_kyc_proto_depIdxs,
		MessageInfos:      file_kyc_proto_msgTypes,
	}.Build()
	File_kyc_proto = out.File
	file_kyc_proto_rawDesc = nil
	file_kyc_proto_goTypes = nil
	file_kyc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KycServiceClient is the client API for KycService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KycServiceClient interface {
	GetKycApplicantByPersonId(ctx context.Context, in *PersonIdRequest, opts ...grpc.CallOption) (*Applicant, error)
	SubmitKycApplicationByPersonId(ctx context.Context, in *PersonIdRequest, opts ...grpc.CallOption) (*Void, error)
	AwaitCDDReport(ctx context.Context, in *ApplicationIdRequest, opts ...grpc.CallOption) (KycService_AwaitCDDReportClient, error)
}

type kycServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKycServiceClient(cc grpc.ClientConnInterface) KycServiceClient {
	return &kycServiceClient{cc}
}

func (c *kycServiceClient) GetKycApplicantByPersonId(ctx context.Context, in *PersonIdRequest, opts ...grpc.CallOption) (*Applicant, error) {
	out := new(Applicant)
	err := c.cc.Invoke(ctx, "/kyc.KycService/getKycApplicantByPersonId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycServiceClient) SubmitKycApplicationByPersonId(ctx context.Context, in *PersonIdRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/kyc.KycService/submitKycApplicationByPersonId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycServiceClient) AwaitCDDReport(ctx context.Context, in *ApplicationIdRequest, opts ...grpc.CallOption) (KycService_AwaitCDDReportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KycService_serviceDesc.Streams[0], "/kyc.KycService/awaitCDDReport", opts...)
	if err != nil {
		return nil, err
	}
	x := &kycServiceAwaitCDDReportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KycService_AwaitCDDReportClient interface {
	Recv() (*Cdd, error)
	grpc.ClientStream
}

type kycServiceAwaitCDDReportClient struct {
	grpc.ClientStream
}

func (x *kycServiceAwaitCDDReportClient) Recv() (*Cdd, error) {
	m := new(Cdd)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KycServiceServer is the server API for KycService service.
type KycServiceServer interface {
	GetKycApplicantByPersonId(context.Context, *PersonIdRequest) (*Applicant, error)
	SubmitKycApplicationByPersonId(context.Context, *PersonIdRequest) (*Void, error)
	AwaitCDDReport(*ApplicationIdRequest, KycService_AwaitCDDReportServer) error
}

// UnimplementedKycServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKycServiceServer struct {
}

func (*UnimplementedKycServiceServer) GetKycApplicantByPersonId(context.Context, *PersonIdRequest) (*Applicant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycApplicantByPersonId not implemented")
}
func (*UnimplementedKycServiceServer) SubmitKycApplicationByPersonId(context.Context, *PersonIdRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitKycApplicationByPersonId not implemented")
}
func (*UnimplementedKycServiceServer) AwaitCDDReport(*ApplicationIdRequest, KycService_AwaitCDDReportServer) error {
	return status.Errorf(codes.Unimplemented, "method AwaitCDDReport not implemented")
}

func RegisterKycServiceServer(s *grpc.Server, srv KycServiceServer) {
	s.RegisterService(&_KycService_serviceDesc, srv)
}

func _KycService_GetKycApplicantByPersonId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycServiceServer).GetKycApplicantByPersonId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.KycService/GetKycApplicantByPersonId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycServiceServer).GetKycApplicantByPersonId(ctx, req.(*PersonIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycService_SubmitKycApplicationByPersonId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycServiceServer).SubmitKycApplicationByPersonId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.KycService/SubmitKycApplicationByPersonId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycServiceServer).SubmitKycApplicationByPersonId(ctx, req.(*PersonIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycService_AwaitCDDReport_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ApplicationIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KycServiceServer).AwaitCDDReport(m, &kycServiceAwaitCDDReportServer{stream})
}

type KycService_AwaitCDDReportServer interface {
	Send(*Cdd) error
	grpc.ServerStream
}

type kycServiceAwaitCDDReportServer struct {
	grpc.ServerStream
}

func (x *kycServiceAwaitCDDReportServer) Send(m *Cdd) error {
	return x.ServerStream.SendMsg(m)
}

var _KycService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kyc.KycService",
	HandlerType: (*KycServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getKycApplicantByPersonId",
			Handler:    _KycService_GetKycApplicantByPersonId_Handler,
		},
		{
			MethodName: "submitKycApplicationByPersonId",
			Handler:    _KycService_SubmitKycApplicationByPersonId_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "awaitCDDReport",
			Handler:       _KycService_AwaitCDDReport_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kyc.proto",
}
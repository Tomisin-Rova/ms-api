// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cdd.proto

package cddService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SuccessResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessResponse) Reset()         { *m = SuccessResponse{} }
func (m *SuccessResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessResponse) ProtoMessage()    {}
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{0}
}

func (m *SuccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessResponse.Unmarshal(m, b)
}
func (m *SuccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessResponse.Marshal(b, m, deterministic)
}
func (m *SuccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessResponse.Merge(m, src)
}
func (m *SuccessResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessResponse.Size(m)
}
func (m *SuccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessResponse proto.InternalMessageInfo

func (m *SuccessResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{2}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type PersonIdRequest struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonIdRequest) Reset()         { *m = PersonIdRequest{} }
func (m *PersonIdRequest) String() string { return proto.CompactTextString(m) }
func (*PersonIdRequest) ProtoMessage()    {}
func (*PersonIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{3}
}

func (m *PersonIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonIdRequest.Unmarshal(m, b)
}
func (m *PersonIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonIdRequest.Marshal(b, m, deterministic)
}
func (m *PersonIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonIdRequest.Merge(m, src)
}
func (m *PersonIdRequest) XXX_Size() int {
	return xxx_messageInfo_PersonIdRequest.Size(m)
}
func (m *PersonIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonIdRequest proto.InternalMessageInfo

func (m *PersonIdRequest) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

type CddIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CddIdRequest) Reset()         { *m = CddIdRequest{} }
func (m *CddIdRequest) String() string { return proto.CompactTextString(m) }
func (*CddIdRequest) ProtoMessage()    {}
func (*CddIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{4}
}

func (m *CddIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CddIdRequest.Unmarshal(m, b)
}
func (m *CddIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CddIdRequest.Marshal(b, m, deterministic)
}
func (m *CddIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CddIdRequest.Merge(m, src)
}
func (m *CddIdRequest) XXX_Size() int {
	return xxx_messageInfo_CddIdRequest.Size(m)
}
func (m *CddIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CddIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CddIdRequest proto.InternalMessageInfo

func (m *CddIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Cdd struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Details              string   `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Kyc                  *Kyc     `protobuf:"bytes,5,opt,name=kyc,proto3" json:"kyc,omitempty"`
	Aml                  *Aml     `protobuf:"bytes,6,opt,name=aml,proto3" json:"aml,omitempty"`
	Roava                *Roava   `protobuf:"bytes,7,opt,name=roava,proto3" json:"roava,omitempty"`
	Cra                  *Cra     `protobuf:"bytes,8,opt,name=cra,proto3" json:"cra,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cdd) Reset()         { *m = Cdd{} }
func (m *Cdd) String() string { return proto.CompactTextString(m) }
func (*Cdd) ProtoMessage()    {}
func (*Cdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{5}
}

func (m *Cdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cdd.Unmarshal(m, b)
}
func (m *Cdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cdd.Marshal(b, m, deterministic)
}
func (m *Cdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cdd.Merge(m, src)
}
func (m *Cdd) XXX_Size() int {
	return xxx_messageInfo_Cdd.Size(m)
}
func (m *Cdd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cdd.DiscardUnknown(m)
}

var xxx_messageInfo_Cdd proto.InternalMessageInfo

func (m *Cdd) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Cdd) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Cdd) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Cdd) GetKyc() *Kyc {
	if m != nil {
		return m.Kyc
	}
	return nil
}

func (m *Cdd) GetAml() *Aml {
	if m != nil {
		return m.Aml
	}
	return nil
}

func (m *Cdd) GetRoava() *Roava {
	if m != nil {
		return m.Roava
	}
	return nil
}

func (m *Cdd) GetCra() *Cra {
	if m != nil {
		return m.Cra
	}
	return nil
}

func (m *Cdd) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Cdd) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Cddsummary struct {
	Status               string                `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Documents            []*CddSummaryDocument `protobuf:"bytes,2,rep,name=documents,proto3" json:"documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Cddsummary) Reset()         { *m = Cddsummary{} }
func (m *Cddsummary) String() string { return proto.CompactTextString(m) }
func (*Cddsummary) ProtoMessage()    {}
func (*Cddsummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{6}
}

func (m *Cddsummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cddsummary.Unmarshal(m, b)
}
func (m *Cddsummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cddsummary.Marshal(b, m, deterministic)
}
func (m *Cddsummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cddsummary.Merge(m, src)
}
func (m *Cddsummary) XXX_Size() int {
	return xxx_messageInfo_Cddsummary.Size(m)
}
func (m *Cddsummary) XXX_DiscardUnknown() {
	xxx_messageInfo_Cddsummary.DiscardUnknown(m)
}

var xxx_messageInfo_Cddsummary proto.InternalMessageInfo

func (m *Cddsummary) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Cddsummary) GetDocuments() []*CddSummaryDocument {
	if m != nil {
		return m.Documents
	}
	return nil
}

type CddSummaryDocument struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Reasons              []string `protobuf:"bytes,3,rep,name=reasons,proto3" json:"reasons,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CddSummaryDocument) Reset()         { *m = CddSummaryDocument{} }
func (m *CddSummaryDocument) String() string { return proto.CompactTextString(m) }
func (*CddSummaryDocument) ProtoMessage()    {}
func (*CddSummaryDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{7}
}

func (m *CddSummaryDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CddSummaryDocument.Unmarshal(m, b)
}
func (m *CddSummaryDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CddSummaryDocument.Marshal(b, m, deterministic)
}
func (m *CddSummaryDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CddSummaryDocument.Merge(m, src)
}
func (m *CddSummaryDocument) XXX_Size() int {
	return xxx_messageInfo_CddSummaryDocument.Size(m)
}
func (m *CddSummaryDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_CddSummaryDocument.DiscardUnknown(m)
}

var xxx_messageInfo_CddSummaryDocument proto.InternalMessageInfo

func (m *CddSummaryDocument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CddSummaryDocument) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CddSummaryDocument) GetReasons() []string {
	if m != nil {
		return m.Reasons
	}
	return nil
}

type Kyc struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Applicant            *Applicant `protobuf:"bytes,2,opt,name=applicant,proto3" json:"applicant,omitempty"`
	Status               string     `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Vendor               string     `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	CreatedAt            string     `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string     `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Kyc) Reset()         { *m = Kyc{} }
func (m *Kyc) String() string { return proto.CompactTextString(m) }
func (*Kyc) ProtoMessage()    {}
func (*Kyc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{8}
}

func (m *Kyc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Kyc.Unmarshal(m, b)
}
func (m *Kyc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Kyc.Marshal(b, m, deterministic)
}
func (m *Kyc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Kyc.Merge(m, src)
}
func (m *Kyc) XXX_Size() int {
	return xxx_messageInfo_Kyc.Size(m)
}
func (m *Kyc) XXX_DiscardUnknown() {
	xxx_messageInfo_Kyc.DiscardUnknown(m)
}

var xxx_messageInfo_Kyc proto.InternalMessageInfo

func (m *Kyc) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Kyc) GetApplicant() *Applicant {
	if m != nil {
		return m.Applicant
	}
	return nil
}

func (m *Kyc) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Kyc) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Kyc) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Kyc) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Aml struct {
	CreatedAt            string   `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	PersonId             string   `protobuf:"bytes,4,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	FirstName            string   `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MiddleName           string   `protobuf:"bytes,7,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,8,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Email                string   `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	Nationality          string   `protobuf:"bytes,10,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Reports              *any.Any `protobuf:"bytes,11,opt,name=reports,proto3" json:"reports,omitempty"`
	Status               string   `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Aml) Reset()         { *m = Aml{} }
func (m *Aml) String() string { return proto.CompactTextString(m) }
func (*Aml) ProtoMessage()    {}
func (*Aml) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{9}
}

func (m *Aml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Aml.Unmarshal(m, b)
}
func (m *Aml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Aml.Marshal(b, m, deterministic)
}
func (m *Aml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aml.Merge(m, src)
}
func (m *Aml) XXX_Size() int {
	return xxx_messageInfo_Aml.Size(m)
}
func (m *Aml) XXX_DiscardUnknown() {
	xxx_messageInfo_Aml.DiscardUnknown(m)
}

var xxx_messageInfo_Aml proto.InternalMessageInfo

func (m *Aml) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Aml) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Aml) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Aml) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

func (m *Aml) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Aml) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Aml) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *Aml) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *Aml) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Aml) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *Aml) GetReports() *any.Any {
	if m != nil {
		return m.Reports
	}
	return nil
}

func (m *Aml) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Roava struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Roava) Reset()         { *m = Roava{} }
func (m *Roava) String() string { return proto.CompactTextString(m) }
func (*Roava) ProtoMessage()    {}
func (*Roava) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{10}
}

func (m *Roava) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Roava.Unmarshal(m, b)
}
func (m *Roava) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Roava.Marshal(b, m, deterministic)
}
func (m *Roava) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Roava.Merge(m, src)
}
func (m *Roava) XXX_Size() int {
	return xxx_messageInfo_Roava.Size(m)
}
func (m *Roava) XXX_DiscardUnknown() {
	xxx_messageInfo_Roava.DiscardUnknown(m)
}

var xxx_messageInfo_Roava proto.InternalMessageInfo

type Cra struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cra) Reset()         { *m = Cra{} }
func (m *Cra) String() string { return proto.CompactTextString(m) }
func (*Cra) ProtoMessage()    {}
func (*Cra) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{11}
}

func (m *Cra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cra.Unmarshal(m, b)
}
func (m *Cra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cra.Marshal(b, m, deterministic)
}
func (m *Cra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cra.Merge(m, src)
}
func (m *Cra) XXX_Size() int {
	return xxx_messageInfo_Cra.Size(m)
}
func (m *Cra) XXX_DiscardUnknown() {
	xxx_messageInfo_Cra.DiscardUnknown(m)
}

var xxx_messageInfo_Cra proto.InternalMessageInfo

type Address struct {
	FlatNumber           string   `protobuf:"bytes,1,opt,name=flat_number,json=flatNumber,proto3" json:"flat_number,omitempty"`
	BuildingNumber       string   `protobuf:"bytes,2,opt,name=building_number,json=buildingNumber,proto3" json:"building_number,omitempty"`
	BuildingName         string   `protobuf:"bytes,3,opt,name=building_name,json=buildingName,proto3" json:"building_name,omitempty"`
	Street               string   `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	SubStreet            string   `protobuf:"bytes,5,opt,name=sub_street,json=subStreet,proto3" json:"sub_street,omitempty"`
	Town                 string   `protobuf:"bytes,6,opt,name=town,proto3" json:"town,omitempty"`
	State                string   `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Postcode             string   `protobuf:"bytes,8,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Country              string   `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{12}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetFlatNumber() string {
	if m != nil {
		return m.FlatNumber
	}
	return ""
}

func (m *Address) GetBuildingNumber() string {
	if m != nil {
		return m.BuildingNumber
	}
	return ""
}

func (m *Address) GetBuildingName() string {
	if m != nil {
		return m.BuildingName
	}
	return ""
}

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetSubStreet() string {
	if m != nil {
		return m.SubStreet
	}
	return ""
}

func (m *Address) GetTown() string {
	if m != nil {
		return m.Town
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetPostcode() string {
	if m != nil {
		return m.Postcode
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type Applicant struct {
	ApplicationId        string   `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	ApplicantId          string   `protobuf:"bytes,2,opt,name=applicant_id,json=applicantId,proto3" json:"applicant_id,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Dob                  string   `protobuf:"bytes,6,opt,name=dob,proto3" json:"dob,omitempty"`
	Address              *Address `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Vendor               string   `protobuf:"bytes,8,opt,name=vendor,proto3" json:"vendor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Applicant) Reset()         { *m = Applicant{} }
func (m *Applicant) String() string { return proto.CompactTextString(m) }
func (*Applicant) ProtoMessage()    {}
func (*Applicant) Descriptor() ([]byte, []int) {
	return fileDescriptor_c22245ed07c5f7ff, []int{13}
}

func (m *Applicant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Applicant.Unmarshal(m, b)
}
func (m *Applicant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Applicant.Marshal(b, m, deterministic)
}
func (m *Applicant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Applicant.Merge(m, src)
}
func (m *Applicant) XXX_Size() int {
	return xxx_messageInfo_Applicant.Size(m)
}
func (m *Applicant) XXX_DiscardUnknown() {
	xxx_messageInfo_Applicant.DiscardUnknown(m)
}

var xxx_messageInfo_Applicant proto.InternalMessageInfo

func (m *Applicant) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *Applicant) GetApplicantId() string {
	if m != nil {
		return m.ApplicantId
	}
	return ""
}

func (m *Applicant) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Applicant) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Applicant) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Applicant) GetDob() string {
	if m != nil {
		return m.Dob
	}
	return ""
}

func (m *Applicant) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Applicant) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func init() {
	proto.RegisterType((*SuccessResponse)(nil), "cdd.SuccessResponse")
	proto.RegisterType((*Empty)(nil), "cdd.Empty")
	proto.RegisterType((*Void)(nil), "cdd.void")
	proto.RegisterType((*PersonIdRequest)(nil), "cdd.PersonIdRequest")
	proto.RegisterType((*CddIdRequest)(nil), "cdd.CddIdRequest")
	proto.RegisterType((*Cdd)(nil), "cdd.Cdd")
	proto.RegisterType((*Cddsummary)(nil), "cdd.Cddsummary")
	proto.RegisterType((*CddSummaryDocument)(nil), "cdd.cdd_summary_document")
	proto.RegisterType((*Kyc)(nil), "cdd.kyc")
	proto.RegisterType((*Aml)(nil), "cdd.aml")
	proto.RegisterType((*Roava)(nil), "cdd.roava")
	proto.RegisterType((*Cra)(nil), "cdd.cra")
	proto.RegisterType((*Address)(nil), "cdd.address")
	proto.RegisterType((*Applicant)(nil), "cdd.applicant")
}

func init() {
	proto.RegisterFile("cdd.proto", fileDescriptor_c22245ed07c5f7ff)
}

var fileDescriptor_c22245ed07c5f7ff = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0xdd, 0xd8, 0xf9, 0xf2, 0x75, 0xda, 0xc2, 0x50, 0xad, 0xbc, 0x5d, 0x01, 0xc1, 0x08, 0xa8,
	0xb4, 0x28, 0x2b, 0x95, 0x07, 0xde, 0x90, 0xda, 0x2e, 0x42, 0x7d, 0x59, 0x90, 0xf3, 0x86, 0x40,
	0xd6, 0xd8, 0x33, 0x09, 0xd6, 0xda, 0x33, 0x66, 0x66, 0x9c, 0x95, 0x7f, 0x05, 0xe2, 0xa7, 0xf0,
	0xeb, 0x10, 0x6f, 0x68, 0x3e, 0xfc, 0x91, 0x74, 0xd5, 0xb7, 0xdc, 0x73, 0xce, 0xcc, 0x5c, 0xdf,
	0x73, 0xef, 0x0d, 0x04, 0x39, 0x21, 0x9b, 0x5a, 0x70, 0xc5, 0x91, 0x9f, 0x13, 0x72, 0xf5, 0x62,
	0xcf, 0xf9, 0xbe, 0xa4, 0xaf, 0x0d, 0x94, 0x35, 0xbb, 0xd7, 0x98, 0xb5, 0x96, 0x8f, 0x5f, 0xc1,
	0xc5, 0xb6, 0xc9, 0x73, 0x2a, 0x65, 0x42, 0x65, 0xcd, 0x99, 0xa4, 0x28, 0x82, 0x45, 0x45, 0xa5,
	0xc4, 0x7b, 0x1a, 0x4d, 0xd6, 0x93, 0xeb, 0x20, 0xe9, 0xc2, 0x78, 0x01, 0xb3, 0x1f, 0xab, 0x5a,
	0xb5, 0xf1, 0x1c, 0xa6, 0x07, 0x5e, 0x90, 0x78, 0x03, 0x17, 0xbf, 0x50, 0x21, 0x39, 0x7b, 0x20,
	0x09, 0xfd, 0xb3, 0xa1, 0x52, 0xa1, 0x97, 0x10, 0xd4, 0x06, 0x4a, 0x0b, 0xe2, 0xce, 0x2f, 0x6b,
	0xa7, 0x89, 0x3f, 0x83, 0xd5, 0x3d, 0x21, 0x83, 0xf8, 0x1c, 0xbc, 0x5e, 0xe5, 0x15, 0x24, 0xfe,
	0xdb, 0x03, 0xff, 0x9e, 0x90, 0x53, 0x1c, 0x5d, 0xc2, 0x8c, 0xbf, 0x67, 0x54, 0x44, 0x9e, 0x81,
	0x6c, 0xa0, 0x13, 0x25, 0x54, 0xe1, 0xa2, 0x94, 0x91, 0x6f, 0x13, 0x75, 0x21, 0x7a, 0x0e, 0x73,
	0xa9, 0xb0, 0x6a, 0x64, 0x34, 0x35, 0x84, 0x8b, 0xd0, 0x15, 0xf8, 0xef, 0xda, 0x3c, 0x9a, 0xad,
	0x27, 0xd7, 0xe1, 0xcd, 0x72, 0xa3, 0xcb, 0xf4, 0xae, 0xcd, 0x13, 0x0d, 0x6a, 0x0e, 0x57, 0x65,
	0x34, 0x1f, 0x71, 0xb8, 0x2a, 0x13, 0x0d, 0xa2, 0x35, 0xcc, 0x04, 0xc7, 0x07, 0x1c, 0x2d, 0x0c,
	0x0b, 0x86, 0x35, 0x48, 0x62, 0x09, 0x7d, 0x3a, 0x17, 0x38, 0x5a, 0x8e, 0x4e, 0xe7, 0x02, 0x27,
	0x1a, 0x44, 0x9f, 0x02, 0xe4, 0x82, 0x62, 0x45, 0x49, 0x8a, 0x55, 0x14, 0x98, 0x8c, 0x02, 0x87,
	0xdc, 0x2a, 0x4d, 0x37, 0x35, 0xe9, 0x68, 0xb0, 0xb4, 0x43, 0x6e, 0x55, 0xfc, 0x3b, 0xc0, 0x3d,
	0x21, 0xb2, 0xa9, 0x2a, 0x2c, 0xda, 0xd1, 0x97, 0x4d, 0x8e, 0xbe, 0xec, 0x7b, 0x08, 0x08, 0xcf,
	0x9b, 0x8a, 0x32, 0x25, 0x23, 0x6f, 0xed, 0x5f, 0x87, 0x37, 0x2f, 0x6c, 0x16, 0x84, 0xa4, 0xee,
	0x70, 0xda, 0x29, 0x92, 0x41, 0x1b, 0xff, 0x06, 0x97, 0x1f, 0x92, 0x20, 0x04, 0x53, 0x86, 0xab,
	0xae, 0x05, 0xcc, 0xef, 0xd1, 0xe3, 0xde, 0xd1, 0xe3, 0x11, 0x2c, 0x04, 0xc5, 0x92, 0x33, 0x6d,
	0x84, 0xaf, 0x8d, 0x70, 0x61, 0xfc, 0xcf, 0xc4, 0x54, 0xfc, 0x91, 0xa1, 0xdf, 0x42, 0x80, 0xeb,
	0xba, 0x2c, 0x72, 0xcc, 0x94, 0xb9, 0x2c, 0xbc, 0x39, 0xb7, 0x25, 0xef, 0xd0, 0x64, 0x10, 0x8c,
	0xde, 0xf5, 0x8f, 0xde, 0x7d, 0x0e, 0xf3, 0x03, 0x65, 0x84, 0x8b, 0xce, 0x66, 0x1b, 0x9d, 0x14,
	0x7c, 0xf6, 0x74, 0xc1, 0xe7, 0xa7, 0x05, 0xff, 0xd7, 0x33, 0x9d, 0x70, 0x72, 0xcb, 0xe4, 0xe9,
	0x5b, 0xbc, 0x93, 0x5b, 0x74, 0xcb, 0xaa, 0x42, 0x95, 0xd4, 0xa5, 0x6c, 0x83, 0xe3, 0xe9, 0x98,
	0x1e, 0x4f, 0x87, 0xbe, 0x71, 0x57, 0x08, 0xa9, 0x52, 0x53, 0x78, 0x97, 0xb6, 0x41, 0xde, 0xea,
	0xea, 0xbf, 0x84, 0xa0, 0xc4, 0x1d, 0x6b, 0xb3, 0x5e, 0x6a, 0xc0, 0x90, 0x9f, 0x43, 0x58, 0x15,
	0x84, 0x94, 0xd4, 0xd2, 0x0b, 0x43, 0x83, 0x85, 0x8c, 0x20, 0x86, 0x33, 0x9d, 0x5a, 0xca, 0x77,
	0x69, 0x56, 0x08, 0xf5, 0x87, 0x69, 0xd5, 0x20, 0x09, 0x35, 0xf8, 0xf3, 0xee, 0x4e, 0x43, 0x3a,
	0x67, 0x5a, 0xe1, 0xa2, 0x74, 0x3d, 0x6a, 0x03, 0xb4, 0x86, 0x90, 0x61, 0x55, 0x70, 0x86, 0xcb,
	0x42, 0xb5, 0xae, 0x41, 0xc7, 0x10, 0xda, 0x68, 0xff, 0x6b, 0x2e, 0x94, 0x8c, 0x42, 0xe3, 0xe5,
	0xe5, 0xc6, 0x6e, 0x9c, 0x4d, 0xb7, 0x71, 0x36, 0xb7, 0xac, 0x4d, 0x3a, 0xd1, 0xc8, 0xcf, 0xd5,
	0xd8, 0x4f, 0xbd, 0x5f, 0xcc, 0x34, 0xc5, 0x33, 0x33, 0x4d, 0xf1, 0x5f, 0x1e, 0x2c, 0x30, 0x21,
	0x82, 0x4a, 0xa9, 0x3f, 0x70, 0x57, 0x62, 0x95, 0xb2, 0xa6, 0xca, 0xa8, 0x70, 0x76, 0x80, 0x86,
	0xde, 0x1a, 0x04, 0x7d, 0x03, 0x17, 0x59, 0x53, 0x94, 0xa4, 0x60, 0xfb, 0x4e, 0x64, 0x4d, 0x39,
	0xef, 0x60, 0x27, 0xfc, 0x12, 0xce, 0x06, 0xa1, 0x2e, 0x96, 0x75, 0x68, 0xd5, 0xcb, 0xfa, 0x56,
	0x17, 0x94, 0xaa, 0x61, 0x83, 0xe8, 0x48, 0x7b, 0x24, 0x9b, 0x2c, 0x75, 0x9c, 0xf3, 0x48, 0x36,
	0xd9, 0xd6, 0xd2, 0x08, 0xa6, 0x8a, 0xbf, 0x67, 0xce, 0x1e, 0xf3, 0x5b, 0x57, 0x55, 0x7f, 0x5f,
	0x67, 0x8a, 0x0d, 0xd0, 0x15, 0x2c, 0x6b, 0x2e, 0x55, 0xce, 0x09, 0x75, 0x56, 0xf4, 0xb1, 0x9e,
	0xa7, 0x9c, 0x37, 0x4c, 0x89, 0xd6, 0x39, 0xd1, 0x85, 0xf1, 0x7f, 0x93, 0xd1, 0xe0, 0xa0, 0xaf,
	0xe0, 0xdc, 0x05, 0xda, 0x8b, 0x61, 0xe1, 0x9e, 0x8d, 0xd0, 0x07, 0x82, 0xbe, 0x80, 0x55, 0x7f,
	0x46, 0x8b, 0x6c, 0x59, 0xc2, 0x1e, 0x7b, 0xd4, 0x7a, 0xfe, 0x93, 0xad, 0x37, 0x3d, 0x69, 0xbd,
	0xbe, 0x6b, 0x66, 0xe3, 0xae, 0xf9, 0x08, 0x7c, 0xc2, 0x33, 0x57, 0x08, 0xfd, 0x13, 0x7d, 0xdd,
	0x9b, 0xe9, 0xd6, 0xe8, 0xca, 0x4e, 0xbc, 0xc5, 0x92, 0xde, 0xe9, 0x61, 0xaa, 0x97, 0xe3, 0xa9,
	0xbe, 0x69, 0xcd, 0x22, 0xdc, 0x52, 0x71, 0x28, 0x72, 0x8a, 0x7e, 0x80, 0x4f, 0x7e, 0xa2, 0xea,
	0xfe, 0xcd, 0x9b, 0xad, 0xdd, 0x5c, 0x89, 0xe9, 0x2d, 0x74, 0x69, 0xee, 0x3c, 0xf9, 0x53, 0xba,
	0xba, 0x30, 0xe8, 0xb0, 0x46, 0xe3, 0x67, 0xe8, 0x15, 0x80, 0x3d, 0x7f, 0xd7, 0x3e, 0x10, 0xf4,
	0x71, 0x27, 0x18, 0xce, 0x2c, 0x3b, 0x28, 0x7e, 0x76, 0xb7, 0xfa, 0x15, 0xf2, 0xfe, 0xe9, 0x6c,
	0x6e, 0xba, 0xfa, 0xbb, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x67, 0x6b, 0x0c, 0x2e, 0x67, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CddServiceClient is the client API for CddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CddServiceClient interface {
	GetCDDSummaryReport(ctx context.Context, in *PersonIdRequest, opts ...grpc.CallOption) (*Cddsummary, error)
	GetCDDById(ctx context.Context, in *CddIdRequest, opts ...grpc.CallOption) (*Cdd, error)
}

type cddServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCddServiceClient(cc grpc.ClientConnInterface) CddServiceClient {
	return &cddServiceClient{cc}
}

func (c *cddServiceClient) GetCDDSummaryReport(ctx context.Context, in *PersonIdRequest, opts ...grpc.CallOption) (*Cddsummary, error) {
	out := new(Cddsummary)
	err := c.cc.Invoke(ctx, "/cdd.CddService/GetCDDSummaryReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cddServiceClient) GetCDDById(ctx context.Context, in *CddIdRequest, opts ...grpc.CallOption) (*Cdd, error) {
	out := new(Cdd)
	err := c.cc.Invoke(ctx, "/cdd.CddService/GetCDDById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CddServiceServer is the server API for CddService service.
type CddServiceServer interface {
	GetCDDSummaryReport(context.Context, *PersonIdRequest) (*Cddsummary, error)
	GetCDDById(context.Context, *CddIdRequest) (*Cdd, error)
}

// UnimplementedCddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCddServiceServer struct {
}

func (*UnimplementedCddServiceServer) GetCDDSummaryReport(ctx context.Context, req *PersonIdRequest) (*Cddsummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCDDSummaryReport not implemented")
}
func (*UnimplementedCddServiceServer) GetCDDById(ctx context.Context, req *CddIdRequest) (*Cdd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCDDById not implemented")
}

func RegisterCddServiceServer(s *grpc.Server, srv CddServiceServer) {
	s.RegisterService(&_CddService_serviceDesc, srv)
}

func _CddService_GetCDDSummaryReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CddServiceServer).GetCDDSummaryReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cdd.CddService/GetCDDSummaryReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CddServiceServer).GetCDDSummaryReport(ctx, req.(*PersonIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CddService_GetCDDById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CddIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CddServiceServer).GetCDDById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cdd.CddService/GetCDDById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CddServiceServer).GetCDDById(ctx, req.(*CddIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cdd.CddService",
	HandlerType: (*CddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCDDSummaryReport",
			Handler:    _CddService_GetCDDSummaryReport_Handler,
		},
		{
			MethodName: "GetCDDById",
			Handler:    _CddService_GetCDDById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cdd.proto",
}

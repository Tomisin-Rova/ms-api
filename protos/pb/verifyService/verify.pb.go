// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verify.proto

package verifyService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ValidateEmailRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateEmailRequest) Reset()         { *m = ValidateEmailRequest{} }
func (m *ValidateEmailRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateEmailRequest) ProtoMessage()    {}
func (*ValidateEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{0}
}

func (m *ValidateEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateEmailRequest.Unmarshal(m, b)
}
func (m *ValidateEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateEmailRequest.Marshal(b, m, deterministic)
}
func (m *ValidateEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateEmailRequest.Merge(m, src)
}
func (m *ValidateEmailRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateEmailRequest.Size(m)
}
func (m *ValidateEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateEmailRequest proto.InternalMessageInfo

func (m *ValidateEmailRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ValidateEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ValidateEmailRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

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
	return fileDescriptor_a87721bac73e05a3, []int{1}
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

type OtpVerificationResponse struct {
	Match                bool     `protobuf:"varint,1,opt,name=match,proto3" json:"match,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtpVerificationResponse) Reset()         { *m = OtpVerificationResponse{} }
func (m *OtpVerificationResponse) String() string { return proto.CompactTextString(m) }
func (*OtpVerificationResponse) ProtoMessage()    {}
func (*OtpVerificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{2}
}

func (m *OtpVerificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtpVerificationResponse.Unmarshal(m, b)
}
func (m *OtpVerificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtpVerificationResponse.Marshal(b, m, deterministic)
}
func (m *OtpVerificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtpVerificationResponse.Merge(m, src)
}
func (m *OtpVerificationResponse) XXX_Size() int {
	return xxx_messageInfo_OtpVerificationResponse.Size(m)
}
func (m *OtpVerificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OtpVerificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OtpVerificationResponse proto.InternalMessageInfo

func (m *OtpVerificationResponse) GetMatch() bool {
	if m != nil {
		return m.Match
	}
	return false
}

func (m *OtpVerificationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type OtpVerificationRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtpVerificationRequest) Reset()         { *m = OtpVerificationRequest{} }
func (m *OtpVerificationRequest) String() string { return proto.CompactTextString(m) }
func (*OtpVerificationRequest) ProtoMessage()    {}
func (*OtpVerificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{3}
}

func (m *OtpVerificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtpVerificationRequest.Unmarshal(m, b)
}
func (m *OtpVerificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtpVerificationRequest.Marshal(b, m, deterministic)
}
func (m *OtpVerificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtpVerificationRequest.Merge(m, src)
}
func (m *OtpVerificationRequest) XXX_Size() int {
	return xxx_messageInfo_OtpVerificationRequest.Size(m)
}
func (m *OtpVerificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OtpVerificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OtpVerificationRequest proto.InternalMessageInfo

func (m *OtpVerificationRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *OtpVerificationRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OtpVerificationByEmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtpVerificationByEmailRequest) Reset()         { *m = OtpVerificationByEmailRequest{} }
func (m *OtpVerificationByEmailRequest) String() string { return proto.CompactTextString(m) }
func (*OtpVerificationByEmailRequest) ProtoMessage()    {}
func (*OtpVerificationByEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{4}
}

func (m *OtpVerificationByEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtpVerificationByEmailRequest.Unmarshal(m, b)
}
func (m *OtpVerificationByEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtpVerificationByEmailRequest.Marshal(b, m, deterministic)
}
func (m *OtpVerificationByEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtpVerificationByEmailRequest.Merge(m, src)
}
func (m *OtpVerificationByEmailRequest) XXX_Size() int {
	return xxx_messageInfo_OtpVerificationByEmailRequest.Size(m)
}
func (m *OtpVerificationByEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OtpVerificationByEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OtpVerificationByEmailRequest proto.InternalMessageInfo

func (m *OtpVerificationByEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *OtpVerificationByEmailRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ResendOtpRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResendOtpRequest) Reset()         { *m = ResendOtpRequest{} }
func (m *ResendOtpRequest) String() string { return proto.CompactTextString(m) }
func (*ResendOtpRequest) ProtoMessage()    {}
func (*ResendOtpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87721bac73e05a3, []int{5}
}

func (m *ResendOtpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResendOtpRequest.Unmarshal(m, b)
}
func (m *ResendOtpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResendOtpRequest.Marshal(b, m, deterministic)
}
func (m *ResendOtpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResendOtpRequest.Merge(m, src)
}
func (m *ResendOtpRequest) XXX_Size() int {
	return xxx_messageInfo_ResendOtpRequest.Size(m)
}
func (m *ResendOtpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResendOtpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResendOtpRequest proto.InternalMessageInfo

func (m *ResendOtpRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*ValidateEmailRequest)(nil), "verify.ValidateEmailRequest")
	proto.RegisterType((*SuccessResponse)(nil), "verify.SuccessResponse")
	proto.RegisterType((*OtpVerificationResponse)(nil), "verify.OtpVerificationResponse")
	proto.RegisterType((*OtpVerificationRequest)(nil), "verify.OtpVerificationRequest")
	proto.RegisterType((*OtpVerificationByEmailRequest)(nil), "verify.OtpVerificationByEmailRequest")
	proto.RegisterType((*ResendOtpRequest)(nil), "verify.ResendOtpRequest")
}

func init() {
	proto.RegisterFile("verify.proto", fileDescriptor_a87721bac73e05a3)
}

var fileDescriptor_a87721bac73e05a3 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xc2, 0x40,
	0x10, 0xc6, 0xab, 0xad, 0x56, 0x07, 0xad, 0x65, 0x09, 0x35, 0x48, 0xff, 0x11, 0x28, 0x08, 0x05,
	0x0f, 0xed, 0x0b, 0x14, 0xa1, 0x50, 0x4f, 0xd2, 0x08, 0x1e, 0xa4, 0x97, 0xed, 0x66, 0x5a, 0x17,
	0x4c, 0x76, 0x9b, 0x5d, 0x2d, 0xbe, 0x62, 0x9f, 0xaa, 0x64, 0xd7, 0x0d, 0xc6, 0x46, 0xe9, 0x6d,
	0xbf, 0xcc, 0xcc, 0x6f, 0x66, 0xbe, 0x09, 0xb4, 0x56, 0x98, 0xf2, 0x8f, 0xf5, 0x40, 0xa6, 0x42,
	0x0b, 0x52, 0xb7, 0x2a, 0x78, 0x03, 0x6f, 0x4a, 0x17, 0x3c, 0xa2, 0x1a, 0x9f, 0x63, 0xca, 0x17,
	0x21, 0x7e, 0x2d, 0x51, 0x69, 0x42, 0xe0, 0x24, 0xa1, 0x31, 0xfa, 0x95, 0xdb, 0x4a, 0xbf, 0x19,
	0x9a, 0x37, 0xf1, 0xa0, 0x86, 0x59, 0x8e, 0x5f, 0x35, 0x1f, 0xad, 0x20, 0x3d, 0x68, 0x48, 0xaa,
	0xd4, 0xb7, 0x48, 0x23, 0xff, 0xd8, 0x04, 0x72, 0x1d, 0xdc, 0x43, 0x67, 0xb2, 0x64, 0x0c, 0x95,
	0x0a, 0x51, 0x49, 0x91, 0x28, 0x24, 0x3e, 0x9c, 0xc6, 0xa8, 0x14, 0xfd, 0x74, 0x6c, 0x27, 0x83,
	0x11, 0x74, 0xc7, 0x5a, 0x4e, 0xb3, 0xb9, 0x38, 0xa3, 0x9a, 0x8b, 0x24, 0x2f, 0xf2, 0xa0, 0x16,
	0x53, 0xcd, 0xe6, 0xa6, 0xa4, 0x11, 0x5a, 0xb1, 0x8d, 0xaa, 0x16, 0x51, 0x43, 0xb8, 0xf8, 0x83,
	0xb2, 0x7b, 0x79, 0x50, 0x93, 0x73, 0x91, 0xb8, 0xe6, 0x56, 0x64, 0xdb, 0x32, 0x11, 0x39, 0x8c,
	0x79, 0x07, 0x23, 0xb8, 0xda, 0x61, 0x0c, 0xd7, 0x05, 0x8b, 0x72, 0x3b, 0x2a, 0xdb, 0x76, 0x94,
	0xa1, 0xfa, 0x70, 0x1e, 0xa2, 0xc2, 0x24, 0x1a, 0x6b, 0x79, 0x70, 0x90, 0x87, 0x9f, 0x2a, 0xb4,
	0x4d, 0xcb, 0xf5, 0x04, 0xd3, 0x15, 0x67, 0x48, 0x5e, 0xa0, 0x5d, 0x38, 0x10, 0xb9, 0x1c, 0x6c,
	0x0e, 0x59, 0x76, 0xb7, 0x5e, 0xd7, 0x45, 0x77, 0x7c, 0x0f, 0x8e, 0xc8, 0x2b, 0xb4, 0x36, 0xe8,
	0x58, 0x8d, 0xb5, 0x24, 0xd7, 0x2e, 0xb5, 0xdc, 0xaa, 0xde, 0xcd, 0xde, 0x78, 0x8e, 0x9c, 0xc1,
	0x99, 0x45, 0x9a, 0x19, 0x32, 0xe8, 0xdd, 0x9e, 0xa2, 0xa2, 0x77, 0xff, 0x61, 0x3f, 0x41, 0x33,
	0x37, 0x8d, 0xf8, 0x2e, 0x7f, 0xd7, 0xc7, 0x03, 0x0b, 0x0f, 0x3b, 0xb3, 0xf6, 0x6a, 0xdb, 0xcb,
	0xf7, 0xba, 0xf9, 0xf7, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xce, 0xbf, 0x49, 0x0b,
	0x03, 0x00, 0x00,
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

func (*UnimplementedVerifyServiceServer) ValidateEmail(ctx context.Context, req *ValidateEmailRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateEmail not implemented")
}
func (*UnimplementedVerifyServiceServer) VerifySmsOtp(ctx context.Context, req *OtpVerificationRequest) (*OtpVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySmsOtp not implemented")
}
func (*UnimplementedVerifyServiceServer) VerifyEmailOtp(ctx context.Context, req *OtpVerificationByEmailRequest) (*OtpVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmailOtp not implemented")
}
func (*UnimplementedVerifyServiceServer) ResendOtp(ctx context.Context, req *ResendOtpRequest) (*SuccessResponse, error) {
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

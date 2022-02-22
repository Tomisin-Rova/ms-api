// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: verification.proto

package verification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	types "ms.api/protos/pb/types"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestOTPRequest_DeliveryMode int32

const (
	RequestOTPRequest_EMAIL RequestOTPRequest_DeliveryMode = 0
	RequestOTPRequest_SMS   RequestOTPRequest_DeliveryMode = 1
	RequestOTPRequest_PUSH  RequestOTPRequest_DeliveryMode = 2
)

// Enum value maps for RequestOTPRequest_DeliveryMode.
var (
	RequestOTPRequest_DeliveryMode_name = map[int32]string{
		0: "EMAIL",
		1: "SMS",
		2: "PUSH",
	}
	RequestOTPRequest_DeliveryMode_value = map[string]int32{
		"EMAIL": 0,
		"SMS":   1,
		"PUSH":  2,
	}
)

func (x RequestOTPRequest_DeliveryMode) Enum() *RequestOTPRequest_DeliveryMode {
	p := new(RequestOTPRequest_DeliveryMode)
	*p = x
	return p
}

func (x RequestOTPRequest_DeliveryMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestOTPRequest_DeliveryMode) Descriptor() protoreflect.EnumDescriptor {
	return file_verification_proto_enumTypes[0].Descriptor()
}

func (RequestOTPRequest_DeliveryMode) Type() protoreflect.EnumType {
	return &file_verification_proto_enumTypes[0]
}

func (x RequestOTPRequest_DeliveryMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestOTPRequest_DeliveryMode.Descriptor instead.
func (RequestOTPRequest_DeliveryMode) EnumDescriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{0, 0}
}

type RequestOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                RequestOTPRequest_DeliveryMode `protobuf:"varint,1,opt,name=type,proto3,enum=verification.RequestOTPRequest_DeliveryMode" json:"type,omitempty"`
	Target              string                         `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	ExpireTimeInSeconds int32                          `protobuf:"varint,3,opt,name=expireTimeInSeconds,proto3" json:"expireTimeInSeconds,omitempty"`
}

func (x *RequestOTPRequest) Reset() {
	*x = RequestOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestOTPRequest) ProtoMessage() {}

func (x *RequestOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestOTPRequest.ProtoReflect.Descriptor instead.
func (*RequestOTPRequest) Descriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{0}
}

func (x *RequestOTPRequest) GetType() RequestOTPRequest_DeliveryMode {
	if x != nil {
		return x.Type
	}
	return RequestOTPRequest_EMAIL
}

func (x *RequestOTPRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *RequestOTPRequest) GetExpireTimeInSeconds() int32 {
	if x != nil {
		return x.ExpireTimeInSeconds
	}
	return 0
}

type VerifyOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	OtpToken string `protobuf:"bytes,2,opt,name=otp_token,json=otpToken,proto3" json:"otp_token,omitempty"`
}

func (x *VerifyOTPRequest) Reset() {
	*x = VerifyOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPRequest) ProtoMessage() {}

func (x *VerifyOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPRequest.ProtoReflect.Descriptor instead.
func (*VerifyOTPRequest) Descriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyOTPRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *VerifyOTPRequest) GetOtpToken() string {
	if x != nil {
		return x.OtpToken
	}
	return ""
}

var File_verification_proto protoreflect.FileDescriptor

var file_verification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcd, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x54, 0x50, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x54, 0x50, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x30, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x22, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x4d, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x02,
	0x22, 0x47, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x74, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa1, 0x01, 0x0a, 0x13, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x54, 0x50, 0x12,
	0x1f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x1e, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a,
	0x1d, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verification_proto_rawDescOnce sync.Once
	file_verification_proto_rawDescData = file_verification_proto_rawDesc
)

func file_verification_proto_rawDescGZIP() []byte {
	file_verification_proto_rawDescOnce.Do(func() {
		file_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_verification_proto_rawDescData)
	})
	return file_verification_proto_rawDescData
}

var file_verification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_verification_proto_goTypes = []interface{}{
	(RequestOTPRequest_DeliveryMode)(0), // 0: verification.RequestOTPRequest.DeliveryMode
	(*RequestOTPRequest)(nil),           // 1: verification.RequestOTPRequest
	(*VerifyOTPRequest)(nil),            // 2: verification.VerifyOTPRequest
	(*types.DefaultResponse)(nil),       // 3: types.DefaultResponse
}
var file_verification_proto_depIdxs = []int32{
	0, // 0: verification.RequestOTPRequest.type:type_name -> verification.RequestOTPRequest.DeliveryMode
	1, // 1: verification.VerificationService.RequestOTP:input_type -> verification.RequestOTPRequest
	2, // 2: verification.VerificationService.VerifyOTP:input_type -> verification.VerifyOTPRequest
	3, // 3: verification.VerificationService.RequestOTP:output_type -> types.DefaultResponse
	3, // 4: verification.VerificationService.VerifyOTP:output_type -> types.DefaultResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_verification_proto_init() }
func file_verification_proto_init() {
	if File_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestOTPRequest); i {
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
		file_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPRequest); i {
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
			RawDescriptor: file_verification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verification_proto_goTypes,
		DependencyIndexes: file_verification_proto_depIdxs,
		EnumInfos:         file_verification_proto_enumTypes,
		MessageInfos:      file_verification_proto_msgTypes,
	}.Build()
	File_verification_proto = out.File
	file_verification_proto_rawDesc = nil
	file_verification_proto_goTypes = nil
	file_verification_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VerificationServiceClient is the client API for VerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerificationServiceClient interface {
	RequestOTP(ctx context.Context, in *RequestOTPRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error)
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error)
}

type verificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationServiceClient(cc grpc.ClientConnInterface) VerificationServiceClient {
	return &verificationServiceClient{cc}
}

func (c *verificationServiceClient) RequestOTP(ctx context.Context, in *RequestOTPRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error) {
	out := new(types.DefaultResponse)
	err := c.cc.Invoke(ctx, "/verification.VerificationService/RequestOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationServiceClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*types.DefaultResponse, error) {
	out := new(types.DefaultResponse)
	err := c.cc.Invoke(ctx, "/verification.VerificationService/VerifyOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationServiceServer is the server API for VerificationService service.
type VerificationServiceServer interface {
	RequestOTP(context.Context, *RequestOTPRequest) (*types.DefaultResponse, error)
	VerifyOTP(context.Context, *VerifyOTPRequest) (*types.DefaultResponse, error)
}

// UnimplementedVerificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVerificationServiceServer struct {
}

func (*UnimplementedVerificationServiceServer) RequestOTP(context.Context, *RequestOTPRequest) (*types.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestOTP not implemented")
}
func (*UnimplementedVerificationServiceServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*types.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}

func RegisterVerificationServiceServer(s *grpc.Server, srv VerificationServiceServer) {
	s.RegisterService(&_VerificationService_serviceDesc, srv)
}

func _VerificationService_RequestOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).RequestOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.VerificationService/RequestOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).RequestOTP(ctx, req.(*RequestOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationService_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.VerificationService/VerifyOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "verification.VerificationService",
	HandlerType: (*VerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestOTP",
			Handler:    _VerificationService_RequestOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _VerificationService_VerifyOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verification.proto",
}

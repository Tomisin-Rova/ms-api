// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: identity.proto

package identityService

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
		mi := &file_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[0]
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
	return file_identity_proto_rawDescGZIP(), []int{0}
}

type UpdateDeviceTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId   string          `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	IdentityId string          `protobuf:"bytes,2,opt,name=identityId,proto3" json:"identityId,omitempty"`
	Tokens     []*DeviceTokens `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *UpdateDeviceTokensRequest) Reset() {
	*x = UpdateDeviceTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceTokensRequest) ProtoMessage() {}

func (x *UpdateDeviceTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceTokensRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceTokensRequest) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDeviceTokensRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UpdateDeviceTokensRequest) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *UpdateDeviceTokensRequest) GetTokens() []*DeviceTokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type DeviceTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *DeviceTokens) Reset() {
	*x = DeviceTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceTokens) ProtoMessage() {}

func (x *DeviceTokens) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceTokens.ProtoReflect.Descriptor instead.
func (*DeviceTokens) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceTokens) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeviceTokens) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_identity_proto protoreflect.FileDescriptor

var file_identity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x76, 0x6f, 0x69, 0x64, 0x22,
	0x87, 0x01, 0x0a, 0x19, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x38, 0x0a, 0x0c, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x5f, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_proto_rawDescOnce sync.Once
	file_identity_proto_rawDescData = file_identity_proto_rawDesc
)

func file_identity_proto_rawDescGZIP() []byte {
	file_identity_proto_rawDescOnce.Do(func() {
		file_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_proto_rawDescData)
	})
	return file_identity_proto_rawDescData
}

var file_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_identity_proto_goTypes = []interface{}{
	(*Void)(nil),                      // 0: identity.void
	(*UpdateDeviceTokensRequest)(nil), // 1: identity.updateDeviceTokensRequest
	(*DeviceTokens)(nil),              // 2: identity.DeviceTokens
	(*types.Response)(nil),            // 3: types.Response
}
var file_identity_proto_depIdxs = []int32{
	2, // 0: identity.updateDeviceTokensRequest.tokens:type_name -> identity.DeviceTokens
	1, // 1: identity.identityService.updateDeviceTokens:input_type -> identity.updateDeviceTokensRequest
	3, // 2: identity.identityService.updateDeviceTokens:output_type -> types.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_identity_proto_init() }
func file_identity_proto_init() {
	if File_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceTokensRequest); i {
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
		file_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceTokens); i {
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
			RawDescriptor: file_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_proto_goTypes,
		DependencyIndexes: file_identity_proto_depIdxs,
		MessageInfos:      file_identity_proto_msgTypes,
	}.Build()
	File_identity_proto = out.File
	file_identity_proto_rawDesc = nil
	file_identity_proto_goTypes = nil
	file_identity_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityServiceClient interface {
	UpdateDeviceTokens(ctx context.Context, in *UpdateDeviceTokensRequest, opts ...grpc.CallOption) (*types.Response, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) UpdateDeviceTokens(ctx context.Context, in *UpdateDeviceTokensRequest, opts ...grpc.CallOption) (*types.Response, error) {
	out := new(types.Response)
	err := c.cc.Invoke(ctx, "/identity.identityService/updateDeviceTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
type IdentityServiceServer interface {
	UpdateDeviceTokens(context.Context, *UpdateDeviceTokensRequest) (*types.Response, error)
}

// UnimplementedIdentityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (*UnimplementedIdentityServiceServer) UpdateDeviceTokens(context.Context, *UpdateDeviceTokensRequest) (*types.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceTokens not implemented")
}

func RegisterIdentityServiceServer(s *grpc.Server, srv IdentityServiceServer) {
	s.RegisterService(&_IdentityService_serviceDesc, srv)
}

func _IdentityService_UpdateDeviceTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).UpdateDeviceTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.identityService/UpdateDeviceTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).UpdateDeviceTokens(ctx, req.(*UpdateDeviceTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identity.identityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "updateDeviceTokens",
			Handler:    _IdentityService_UpdateDeviceTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

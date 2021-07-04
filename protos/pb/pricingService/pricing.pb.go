// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: pricing.proto

package pricingService

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

type FxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency     string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	BaseCurrency string `protobuf:"bytes,2,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
}

func (x *FxRequest) Reset() {
	*x = FxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FxRequest) ProtoMessage() {}

func (x *FxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FxRequest.ProtoReflect.Descriptor instead.
func (*FxRequest) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{0}
}

func (x *FxRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *FxRequest) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

type TransferFeesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency     string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	BaseCurrency string `protobuf:"bytes,2,opt,name=baseCurrency,proto3" json:"baseCurrency,omitempty"`
}

func (x *TransferFeesRequest) Reset() {
	*x = TransferFeesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFeesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFeesRequest) ProtoMessage() {}

func (x *TransferFeesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFeesRequest.ProtoReflect.Descriptor instead.
func (*TransferFeesRequest) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{1}
}

func (x *TransferFeesRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferFeesRequest) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

var File_pricing_proto protoreflect.FileDescriptor

var file_pricing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x09, 0x46, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x22, 0x55, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x65,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x32, 0x87, 0x01, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x46, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x78, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x65, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x46, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x65, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pricing_proto_rawDescOnce sync.Once
	file_pricing_proto_rawDescData = file_pricing_proto_rawDesc
)

func file_pricing_proto_rawDescGZIP() []byte {
	file_pricing_proto_rawDescOnce.Do(func() {
		file_pricing_proto_rawDescData = protoimpl.X.CompressGZIP(file_pricing_proto_rawDescData)
	})
	return file_pricing_proto_rawDescData
}

var file_pricing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pricing_proto_goTypes = []interface{}{
	(*FxRequest)(nil),           // 0: pricing.FxRequest
	(*TransferFeesRequest)(nil), // 1: pricing.TransferFeesRequest
	(*types.Fx)(nil),            // 2: types.Fx
	(*types.TransferFees)(nil),  // 3: types.TransferFees
}
var file_pricing_proto_depIdxs = []int32{
	0, // 0: pricing.pricingService.GetFxRates:input_type -> pricing.FxRequest
	1, // 1: pricing.pricingService.GetTransferFees:input_type -> pricing.TransferFeesRequest
	2, // 2: pricing.pricingService.GetFxRates:output_type -> types.Fx
	3, // 3: pricing.pricingService.GetTransferFees:output_type -> types.TransferFees
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pricing_proto_init() }
func file_pricing_proto_init() {
	if File_pricing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pricing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FxRequest); i {
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
		file_pricing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFeesRequest); i {
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
			RawDescriptor: file_pricing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pricing_proto_goTypes,
		DependencyIndexes: file_pricing_proto_depIdxs,
		MessageInfos:      file_pricing_proto_msgTypes,
	}.Build()
	File_pricing_proto = out.File
	file_pricing_proto_rawDesc = nil
	file_pricing_proto_goTypes = nil
	file_pricing_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PricingServiceClient is the client API for PricingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PricingServiceClient interface {
	GetFxRates(ctx context.Context, in *FxRequest, opts ...grpc.CallOption) (*types.Fx, error)
	GetTransferFees(ctx context.Context, in *TransferFeesRequest, opts ...grpc.CallOption) (*types.TransferFees, error)
}

type pricingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingServiceClient(cc grpc.ClientConnInterface) PricingServiceClient {
	return &pricingServiceClient{cc}
}

func (c *pricingServiceClient) GetFxRates(ctx context.Context, in *FxRequest, opts ...grpc.CallOption) (*types.Fx, error) {
	out := new(types.Fx)
	err := c.cc.Invoke(ctx, "/pricing.pricingService/GetFxRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) GetTransferFees(ctx context.Context, in *TransferFeesRequest, opts ...grpc.CallOption) (*types.TransferFees, error) {
	out := new(types.TransferFees)
	err := c.cc.Invoke(ctx, "/pricing.pricingService/GetTransferFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricingServiceServer is the server API for PricingService service.
type PricingServiceServer interface {
	GetFxRates(context.Context, *FxRequest) (*types.Fx, error)
	GetTransferFees(context.Context, *TransferFeesRequest) (*types.TransferFees, error)
}

// UnimplementedPricingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPricingServiceServer struct {
}

func (*UnimplementedPricingServiceServer) GetFxRates(context.Context, *FxRequest) (*types.Fx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFxRates not implemented")
}
func (*UnimplementedPricingServiceServer) GetTransferFees(context.Context, *TransferFeesRequest) (*types.TransferFees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferFees not implemented")
}

func RegisterPricingServiceServer(s *grpc.Server, srv PricingServiceServer) {
	s.RegisterService(&_PricingService_serviceDesc, srv)
}

func _PricingService_GetFxRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetFxRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricing.pricingService/GetFxRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetFxRates(ctx, req.(*FxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_GetTransferFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetTransferFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricing.pricingService/GetTransferFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetTransferFees(ctx, req.(*TransferFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PricingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pricing.pricingService",
	HandlerType: (*PricingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFxRates",
			Handler:    _PricingService_GetFxRates_Handler,
		},
		{
			MethodName: "GetTransferFees",
			Handler:    _PricingService_GetTransferFees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricing.proto",
}

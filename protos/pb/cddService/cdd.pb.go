// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: cdd.proto

package cddService

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

type PersonIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId string `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *PersonIdRequest) Reset() {
	*x = PersonIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonIdRequest) ProtoMessage() {}

func (x *PersonIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cdd_proto_msgTypes[0]
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
	return file_cdd_proto_rawDescGZIP(), []int{0}
}

func (x *PersonIdRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

type CddIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CddIdRequest) Reset() {
	*x = CddIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CddIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CddIdRequest) ProtoMessage() {}

func (x *CddIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cdd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CddIdRequest.ProtoReflect.Descriptor instead.
func (*CddIdRequest) Descriptor() ([]byte, []int) {
	return file_cdd_proto_rawDescGZIP(), []int{1}
}

func (x *CddIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Cddsummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string                `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Documents []*CddSummaryDocument `protobuf:"bytes,2,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *Cddsummary) Reset() {
	*x = Cddsummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cddsummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cddsummary) ProtoMessage() {}

func (x *Cddsummary) ProtoReflect() protoreflect.Message {
	mi := &file_cdd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cddsummary.ProtoReflect.Descriptor instead.
func (*Cddsummary) Descriptor() ([]byte, []int) {
	return file_cdd_proto_rawDescGZIP(), []int{2}
}

func (x *Cddsummary) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Cddsummary) GetDocuments() []*CddSummaryDocument {
	if x != nil {
		return x.Documents
	}
	return nil
}

type CddSummaryDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status  string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Reasons []string `protobuf:"bytes,3,rep,name=reasons,proto3" json:"reasons,omitempty"`
}

func (x *CddSummaryDocument) Reset() {
	*x = CddSummaryDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CddSummaryDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CddSummaryDocument) ProtoMessage() {}

func (x *CddSummaryDocument) ProtoReflect() protoreflect.Message {
	mi := &file_cdd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CddSummaryDocument.ProtoReflect.Descriptor instead.
func (*CddSummaryDocument) Descriptor() ([]byte, []int) {
	return file_cdd_proto_rawDescGZIP(), []int{3}
}

func (x *CddSummaryDocument) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CddSummaryDocument) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CddSummaryDocument) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

type GetCDDByOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId string `protobuf:"bytes,1,opt,name=personId,proto3" json:"personId,omitempty"`
}

func (x *GetCDDByOwnerRequest) Reset() {
	*x = GetCDDByOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cdd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCDDByOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCDDByOwnerRequest) ProtoMessage() {}

func (x *GetCDDByOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cdd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCDDByOwnerRequest.ProtoReflect.Descriptor instead.
func (*GetCDDByOwnerRequest) Descriptor() ([]byte, []int) {
	return file_cdd_proto_rawDescGZIP(), []int{4}
}

func (x *GetCDDByOwnerRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

var File_cdd_proto protoreflect.FileDescriptor

var file_cdd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x64, 0x64,
	0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x0f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1e, 0x0a,
	0x0c, 0x43, 0x64, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a,
	0x0a, 0x43, 0x64, 0x64, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x64, 0x64, 0x2e, 0x63, 0x64, 0x64,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x14,
	0x63, 0x64, 0x64, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x43, 0x44, 0x44, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xb5,
	0x01, 0x0a, 0x0a, 0x43, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x44, 0x44, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x64, 0x64, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x64, 0x64,
	0x2e, 0x43, 0x64, 0x64, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x44, 0x44, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x63, 0x64,
	0x64, 0x2e, 0x43, 0x64, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x64, 0x64, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x43, 0x44, 0x44, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x63, 0x64, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x44, 0x44, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x64, 0x64, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x64, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cdd_proto_rawDescOnce sync.Once
	file_cdd_proto_rawDescData = file_cdd_proto_rawDesc
)

func file_cdd_proto_rawDescGZIP() []byte {
	file_cdd_proto_rawDescOnce.Do(func() {
		file_cdd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cdd_proto_rawDescData)
	})
	return file_cdd_proto_rawDescData
}

var file_cdd_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cdd_proto_goTypes = []interface{}{
	(*PersonIdRequest)(nil),      // 0: cdd.PersonIdRequest
	(*CddIdRequest)(nil),         // 1: cdd.CddIdRequest
	(*Cddsummary)(nil),           // 2: cdd.Cddsummary
	(*CddSummaryDocument)(nil),   // 3: cdd.cdd_summary_document
	(*GetCDDByOwnerRequest)(nil), // 4: cdd.GetCDDByOwnerRequest
	(*types.Cdd)(nil),            // 5: types.Cdd
}
var file_cdd_proto_depIdxs = []int32{
	3, // 0: cdd.Cddsummary.documents:type_name -> cdd.cdd_summary_document
	0, // 1: cdd.CddService.GetCDDSummaryReport:input_type -> cdd.PersonIdRequest
	1, // 2: cdd.CddService.GetCDDById:input_type -> cdd.CddIdRequest
	4, // 3: cdd.CddService.GetCDDByOwner:input_type -> cdd.GetCDDByOwnerRequest
	2, // 4: cdd.CddService.GetCDDSummaryReport:output_type -> cdd.Cddsummary
	5, // 5: cdd.CddService.GetCDDById:output_type -> types.Cdd
	5, // 6: cdd.CddService.GetCDDByOwner:output_type -> types.Cdd
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cdd_proto_init() }
func file_cdd_proto_init() {
	if File_cdd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cdd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_cdd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CddIdRequest); i {
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
		file_cdd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cddsummary); i {
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
		file_cdd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CddSummaryDocument); i {
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
		file_cdd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCDDByOwnerRequest); i {
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
			RawDescriptor: file_cdd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cdd_proto_goTypes,
		DependencyIndexes: file_cdd_proto_depIdxs,
		MessageInfos:      file_cdd_proto_msgTypes,
	}.Build()
	File_cdd_proto = out.File
	file_cdd_proto_rawDesc = nil
	file_cdd_proto_goTypes = nil
	file_cdd_proto_depIdxs = nil
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
	GetCDDById(ctx context.Context, in *CddIdRequest, opts ...grpc.CallOption) (*types.Cdd, error)
	GetCDDByOwner(ctx context.Context, in *GetCDDByOwnerRequest, opts ...grpc.CallOption) (*types.Cdd, error)
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

func (c *cddServiceClient) GetCDDById(ctx context.Context, in *CddIdRequest, opts ...grpc.CallOption) (*types.Cdd, error) {
	out := new(types.Cdd)
	err := c.cc.Invoke(ctx, "/cdd.CddService/GetCDDById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cddServiceClient) GetCDDByOwner(ctx context.Context, in *GetCDDByOwnerRequest, opts ...grpc.CallOption) (*types.Cdd, error) {
	out := new(types.Cdd)
	err := c.cc.Invoke(ctx, "/cdd.CddService/GetCDDByOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CddServiceServer is the server API for CddService service.
type CddServiceServer interface {
	GetCDDSummaryReport(context.Context, *PersonIdRequest) (*Cddsummary, error)
	GetCDDById(context.Context, *CddIdRequest) (*types.Cdd, error)
	GetCDDByOwner(context.Context, *GetCDDByOwnerRequest) (*types.Cdd, error)
}

// UnimplementedCddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCddServiceServer struct {
}

func (*UnimplementedCddServiceServer) GetCDDSummaryReport(context.Context, *PersonIdRequest) (*Cddsummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCDDSummaryReport not implemented")
}
func (*UnimplementedCddServiceServer) GetCDDById(context.Context, *CddIdRequest) (*types.Cdd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCDDById not implemented")
}
func (*UnimplementedCddServiceServer) GetCDDByOwner(context.Context, *GetCDDByOwnerRequest) (*types.Cdd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCDDByOwner not implemented")
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

func _CddService_GetCDDByOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCDDByOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CddServiceServer).GetCDDByOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cdd.CddService/GetCDDByOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CddServiceServer).GetCDDByOwner(ctx, req.(*GetCDDByOwnerRequest))
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
		{
			MethodName: "GetCDDByOwner",
			Handler:    _CddService_GetCDDByOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cdd.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: pricing.proto

package pricing

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

type GetCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCurrencyRequest) Reset() {
	*x = GetCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyRequest) ProtoMessage() {}

func (x *GetCurrencyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCurrencyRequest.ProtoReflect.Descriptor instead.
func (*GetCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrencyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCurrenciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords string `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"`
	First    int32  `protobuf:"varint,2,opt,name=first,proto3" json:"first,omitempty"`
	After    string `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	Last     int32  `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
	Before   string `protobuf:"bytes,5,opt,name=before,proto3" json:"before,omitempty"`
}

func (x *GetCurrenciesRequest) Reset() {
	*x = GetCurrenciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesRequest) ProtoMessage() {}

func (x *GetCurrenciesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCurrenciesRequest.ProtoReflect.Descriptor instead.
func (*GetCurrenciesRequest) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrenciesRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *GetCurrenciesRequest) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *GetCurrenciesRequest) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *GetCurrenciesRequest) GetLast() int32 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *GetCurrenciesRequest) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

type GetCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes          []*types.Currency     `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	PaginationInfo *types.PaginationInfo `protobuf:"bytes,2,opt,name=pagination_info,json=paginationInfo,proto3" json:"pagination_info,omitempty"`
	TotalCount     int32                 `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetCurrenciesResponse) Reset() {
	*x = GetCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesResponse) ProtoMessage() {}

func (x *GetCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrenciesResponse) GetNodes() []*types.Currency {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *GetCurrenciesResponse) GetPaginationInfo() *types.PaginationInfo {
	if x != nil {
		return x.PaginationInfo
	}
	return nil
}

func (x *GetCurrenciesResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetFeesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionTypeId string `protobuf:"bytes,1,opt,name=transaction_type_id,json=transactionTypeId,proto3" json:"transaction_type_id,omitempty"`
}

func (x *GetFeesRequest) Reset() {
	*x = GetFeesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeesRequest) ProtoMessage() {}

func (x *GetFeesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeesRequest.ProtoReflect.Descriptor instead.
func (*GetFeesRequest) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeesRequest) GetTransactionTypeId() string {
	if x != nil {
		return x.TransactionTypeId
	}
	return ""
}

type GetFeesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fees []*types.Fee `protobuf:"bytes,1,rep,name=fees,proto3" json:"fees,omitempty"`
}

func (x *GetFeesResponse) Reset() {
	*x = GetFeesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeesResponse) ProtoMessage() {}

func (x *GetFeesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeesResponse.ProtoReflect.Descriptor instead.
func (*GetFeesResponse) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{4}
}

func (x *GetFeesResponse) GetFees() []*types.Fee {
	if x != nil {
		return x.Fees
	}
	return nil
}

type GetExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionTypeId string `protobuf:"bytes,1,opt,name=transaction_type_id,json=transactionTypeId,proto3" json:"transaction_type_id,omitempty"`
}

func (x *GetExchangeRateRequest) Reset() {
	*x = GetExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateRequest) ProtoMessage() {}

func (x *GetExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*GetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{5}
}

func (x *GetExchangeRateRequest) GetTransactionTypeId() string {
	if x != nil {
		return x.TransactionTypeId
	}
	return ""
}

type GetExchangeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeRate *types.ExchangeRate `protobuf:"bytes,1,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`
}

func (x *GetExchangeRateResponse) Reset() {
	*x = GetExchangeRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateResponse) ProtoMessage() {}

func (x *GetExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*GetExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_pricing_proto_rawDescGZIP(), []int{6}
}

func (x *GetExchangeRateResponse) GetExchangeRate() *types.ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

var File_pricing_proto protoreflect.FileDescriptor

var file_pricing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x04, 0x66, 0x65, 0x65, 0x73,
	0x22, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x32,
	0xb1, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x12, 0x1d, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pricing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pricing_proto_goTypes = []interface{}{
	(*GetCurrencyRequest)(nil),      // 0: pricing.GetCurrencyRequest
	(*GetCurrenciesRequest)(nil),    // 1: pricing.GetCurrenciesRequest
	(*GetCurrenciesResponse)(nil),   // 2: pricing.GetCurrenciesResponse
	(*GetFeesRequest)(nil),          // 3: pricing.GetFeesRequest
	(*GetFeesResponse)(nil),         // 4: pricing.GetFeesResponse
	(*GetExchangeRateRequest)(nil),  // 5: pricing.GetExchangeRateRequest
	(*GetExchangeRateResponse)(nil), // 6: pricing.GetExchangeRateResponse
	(*types.Currency)(nil),          // 7: types.Currency
	(*types.PaginationInfo)(nil),    // 8: types.PaginationInfo
	(*types.Fee)(nil),               // 9: types.Fee
	(*types.ExchangeRate)(nil),      // 10: types.ExchangeRate
}
var file_pricing_proto_depIdxs = []int32{
	7,  // 0: pricing.GetCurrenciesResponse.nodes:type_name -> types.Currency
	8,  // 1: pricing.GetCurrenciesResponse.pagination_info:type_name -> types.PaginationInfo
	9,  // 2: pricing.GetFeesResponse.fees:type_name -> types.Fee
	10, // 3: pricing.GetExchangeRateResponse.exchange_rate:type_name -> types.ExchangeRate
	0,  // 4: pricing.PricingService.GetCurrency:input_type -> pricing.GetCurrencyRequest
	1,  // 5: pricing.PricingService.GetCurrencies:input_type -> pricing.GetCurrenciesRequest
	3,  // 6: pricing.PricingService.GetFees:input_type -> pricing.GetFeesRequest
	5,  // 7: pricing.PricingService.GetExchangeRate:input_type -> pricing.GetExchangeRateRequest
	7,  // 8: pricing.PricingService.GetCurrency:output_type -> types.Currency
	2,  // 9: pricing.PricingService.GetCurrencies:output_type -> pricing.GetCurrenciesResponse
	4,  // 10: pricing.PricingService.GetFees:output_type -> pricing.GetFeesResponse
	6,  // 11: pricing.PricingService.GetExchangeRate:output_type -> pricing.GetExchangeRateResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pricing_proto_init() }
func file_pricing_proto_init() {
	if File_pricing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pricing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyRequest); i {
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
			switch v := v.(*GetCurrenciesRequest); i {
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
		file_pricing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesResponse); i {
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
		file_pricing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeesRequest); i {
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
		file_pricing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeesResponse); i {
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
		file_pricing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeRateRequest); i {
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
		file_pricing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeRateResponse); i {
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
			NumMessages:   7,
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
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*types.Currency, error)
	GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
	GetFees(ctx context.Context, in *GetFeesRequest, opts ...grpc.CallOption) (*GetFeesResponse, error)
	GetExchangeRate(ctx context.Context, in *GetExchangeRateRequest, opts ...grpc.CallOption) (*GetExchangeRateResponse, error)
}

type pricingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingServiceClient(cc grpc.ClientConnInterface) PricingServiceClient {
	return &pricingServiceClient{cc}
}

func (c *pricingServiceClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*types.Currency, error) {
	out := new(types.Currency)
	err := c.cc.Invoke(ctx, "/pricing.PricingService/GetCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/pricing.PricingService/GetCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) GetFees(ctx context.Context, in *GetFeesRequest, opts ...grpc.CallOption) (*GetFeesResponse, error) {
	out := new(GetFeesResponse)
	err := c.cc.Invoke(ctx, "/pricing.PricingService/GetFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) GetExchangeRate(ctx context.Context, in *GetExchangeRateRequest, opts ...grpc.CallOption) (*GetExchangeRateResponse, error) {
	out := new(GetExchangeRateResponse)
	err := c.cc.Invoke(ctx, "/pricing.PricingService/GetExchangeRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricingServiceServer is the server API for PricingService service.
type PricingServiceServer interface {
	GetCurrency(context.Context, *GetCurrencyRequest) (*types.Currency, error)
	GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	GetFees(context.Context, *GetFeesRequest) (*GetFeesResponse, error)
	GetExchangeRate(context.Context, *GetExchangeRateRequest) (*GetExchangeRateResponse, error)
}

// UnimplementedPricingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPricingServiceServer struct {
}

func (*UnimplementedPricingServiceServer) GetCurrency(context.Context, *GetCurrencyRequest) (*types.Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (*UnimplementedPricingServiceServer) GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
}
func (*UnimplementedPricingServiceServer) GetFees(context.Context, *GetFeesRequest) (*GetFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFees not implemented")
}
func (*UnimplementedPricingServiceServer) GetExchangeRate(context.Context, *GetExchangeRateRequest) (*GetExchangeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRate not implemented")
}

func RegisterPricingServiceServer(s *grpc.Server, srv PricingServiceServer) {
	s.RegisterService(&_PricingService_serviceDesc, srv)
}

func _PricingService_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricing.PricingService/GetCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricing.PricingService/GetCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetCurrencies(ctx, req.(*GetCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_GetFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricing.PricingService/GetFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetFees(ctx, req.(*GetFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_GetExchangeRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExchangeRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetExchangeRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricing.PricingService/GetExchangeRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetExchangeRate(ctx, req.(*GetExchangeRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PricingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pricing.PricingService",
	HandlerType: (*PricingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrency",
			Handler:    _PricingService_GetCurrency_Handler,
		},
		{
			MethodName: "GetCurrencies",
			Handler:    _PricingService_GetCurrencies_Handler,
		},
		{
			MethodName: "GetFees",
			Handler:    _PricingService_GetFees_Handler,
		},
		{
			MethodName: "GetExchangeRate",
			Handler:    _PricingService_GetExchangeRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricing.proto",
}

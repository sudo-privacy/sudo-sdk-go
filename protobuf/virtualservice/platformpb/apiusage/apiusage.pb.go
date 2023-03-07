// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: protobuf/mpc_virtual_service/platform/apiusage.proto

package apiusage

import (
	api_usage "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/api_usage"
	paginator "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/paginator"
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

type UpsertAPIUsageReceiverResponse_Status int32

const (
	UpsertAPIUsageReceiverResponse_UNKNOWN_STATUS UpsertAPIUsageReceiverResponse_Status = 0
	UpsertAPIUsageReceiverResponse_SUCCESS        UpsertAPIUsageReceiverResponse_Status = 1
	UpsertAPIUsageReceiverResponse_FAILED         UpsertAPIUsageReceiverResponse_Status = 2
)

// Enum value maps for UpsertAPIUsageReceiverResponse_Status.
var (
	UpsertAPIUsageReceiverResponse_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "SUCCESS",
		2: "FAILED",
	}
	UpsertAPIUsageReceiverResponse_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"SUCCESS":        1,
		"FAILED":         2,
	}
)

func (x UpsertAPIUsageReceiverResponse_Status) Enum() *UpsertAPIUsageReceiverResponse_Status {
	p := new(UpsertAPIUsageReceiverResponse_Status)
	*p = x
	return p
}

func (x UpsertAPIUsageReceiverResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpsertAPIUsageReceiverResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_enumTypes[0].Descriptor()
}

func (UpsertAPIUsageReceiverResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_mpc_virtual_service_platform_apiusage_proto_enumTypes[0]
}

func (x UpsertAPIUsageReceiverResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpsertAPIUsageReceiverResponse_Status.Descriptor instead.
func (UpsertAPIUsageReceiverResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{1, 0}
}

type APIUsageNotifyResponse_Status int32

const (
	APIUsageNotifyResponse_UNKNOWN_STATUS APIUsageNotifyResponse_Status = 0
	APIUsageNotifyResponse_SUCCESS        APIUsageNotifyResponse_Status = 1
	APIUsageNotifyResponse_FAILED         APIUsageNotifyResponse_Status = 2
)

// Enum value maps for APIUsageNotifyResponse_Status.
var (
	APIUsageNotifyResponse_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "SUCCESS",
		2: "FAILED",
	}
	APIUsageNotifyResponse_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"SUCCESS":        1,
		"FAILED":         2,
	}
)

func (x APIUsageNotifyResponse_Status) Enum() *APIUsageNotifyResponse_Status {
	p := new(APIUsageNotifyResponse_Status)
	*p = x
	return p
}

func (x APIUsageNotifyResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIUsageNotifyResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_enumTypes[1].Descriptor()
}

func (APIUsageNotifyResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_mpc_virtual_service_platform_apiusage_proto_enumTypes[1]
}

func (x APIUsageNotifyResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIUsageNotifyResponse_Status.Descriptor instead.
func (APIUsageNotifyResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{2, 0}
}

type UpsertAPIUsageReceiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Disabled bool   `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *UpsertAPIUsageReceiverRequest) Reset() {
	*x = UpsertAPIUsageReceiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAPIUsageReceiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAPIUsageReceiverRequest) ProtoMessage() {}

func (x *UpsertAPIUsageReceiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAPIUsageReceiverRequest.ProtoReflect.Descriptor instead.
func (*UpsertAPIUsageReceiverRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertAPIUsageReceiverRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpsertAPIUsageReceiverRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpsertAPIUsageReceiverRequest) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

type UpsertAPIUsageReceiverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status UpsertAPIUsageReceiverResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=protobuf.platform.apiusage.UpsertAPIUsageReceiverResponse_Status" json:"status,omitempty"`
	Msg    string                                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 异常时返回原因
}

func (x *UpsertAPIUsageReceiverResponse) Reset() {
	*x = UpsertAPIUsageReceiverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAPIUsageReceiverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAPIUsageReceiverResponse) ProtoMessage() {}

func (x *UpsertAPIUsageReceiverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAPIUsageReceiverResponse.ProtoReflect.Descriptor instead.
func (*UpsertAPIUsageReceiverResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertAPIUsageReceiverResponse) GetStatus() UpsertAPIUsageReceiverResponse_Status {
	if x != nil {
		return x.Status
	}
	return UpsertAPIUsageReceiverResponse_UNKNOWN_STATUS
}

func (x *UpsertAPIUsageReceiverResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type APIUsageNotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status APIUsageNotifyResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=protobuf.platform.apiusage.APIUsageNotifyResponse_Status" json:"status,omitempty"`
	Msg    string                        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 异常时返回原因
}

func (x *APIUsageNotifyResponse) Reset() {
	*x = APIUsageNotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIUsageNotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIUsageNotifyResponse) ProtoMessage() {}

func (x *APIUsageNotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIUsageNotifyResponse.ProtoReflect.Descriptor instead.
func (*APIUsageNotifyResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{2}
}

func (x *APIUsageNotifyResponse) GetStatus() APIUsageNotifyResponse_Status {
	if x != nil {
		return x.Status
	}
	return APIUsageNotifyResponse_UNKNOWN_STATUS
}

func (x *APIUsageNotifyResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListAPIUsagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginator      *paginator.Paginator      `protobuf:"bytes,1,opt,name=paginator,proto3" json:"paginator,omitempty"`
	Type           api_usage.APIUsage_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=sudo.protobuf.api_usage.APIUsage_Type" json:"type,omitempty"`
	Status         api_usage.APIUsage_Status `protobuf:"varint,3,opt,name=status,proto3,enum=sudo.protobuf.api_usage.APIUsage_Status" json:"status,omitempty"`
	CreateTimeFrom int64                     `protobuf:"varint,4,opt,name=create_time_from,json=createTimeFrom,proto3" json:"create_time_from,omitempty"`
	CreateTimeTo   int64                     `protobuf:"varint,5,opt,name=create_time_to,json=createTimeTo,proto3" json:"create_time_to,omitempty"`
	PartyId        string                    `protobuf:"bytes,6,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Caller         string                    `protobuf:"bytes,7,opt,name=caller,proto3" json:"caller,omitempty"`
	ApiId          string                    `protobuf:"bytes,8,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	IsClient       bool                      `protobuf:"varint,9,opt,name=is_client,json=isClient,proto3" json:"is_client,omitempty"`
}

func (x *ListAPIUsagesRequest) Reset() {
	*x = ListAPIUsagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAPIUsagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAPIUsagesRequest) ProtoMessage() {}

func (x *ListAPIUsagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAPIUsagesRequest.ProtoReflect.Descriptor instead.
func (*ListAPIUsagesRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{3}
}

func (x *ListAPIUsagesRequest) GetPaginator() *paginator.Paginator {
	if x != nil {
		return x.Paginator
	}
	return nil
}

func (x *ListAPIUsagesRequest) GetType() api_usage.APIUsage_Type {
	if x != nil {
		return x.Type
	}
	return api_usage.APIUsage_UNKNOWN_TYPE
}

func (x *ListAPIUsagesRequest) GetStatus() api_usage.APIUsage_Status {
	if x != nil {
		return x.Status
	}
	return api_usage.APIUsage_UNKNOWN_STATUS
}

func (x *ListAPIUsagesRequest) GetCreateTimeFrom() int64 {
	if x != nil {
		return x.CreateTimeFrom
	}
	return 0
}

func (x *ListAPIUsagesRequest) GetCreateTimeTo() int64 {
	if x != nil {
		return x.CreateTimeTo
	}
	return 0
}

func (x *ListAPIUsagesRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *ListAPIUsagesRequest) GetCaller() string {
	if x != nil {
		return x.Caller
	}
	return ""
}

func (x *ListAPIUsagesRequest) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *ListAPIUsagesRequest) GetIsClient() bool {
	if x != nil {
		return x.IsClient
	}
	return false
}

type ListAPIUsagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*api_usage.APIUsage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int64                 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListAPIUsagesResponse) Reset() {
	*x = ListAPIUsagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAPIUsagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAPIUsagesResponse) ProtoMessage() {}

func (x *ListAPIUsagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAPIUsagesResponse.ProtoReflect.Descriptor instead.
func (*ListAPIUsagesResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP(), []int{4}
}

func (x *ListAPIUsagesResponse) GetData() []*api_usage.APIUsage {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListAPIUsagesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_protobuf_mpc_virtual_service_platform_apiusage_proto protoreflect.FileDescriptor

var file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x27, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x73, 0x75, 0x64,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x1d, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xc4, 0x01, 0x0a,
	0x1e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x22, 0xb4, 0x01, 0x0a, 0x16, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0x96, 0x03, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x75, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x15, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x50, 0x49, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x75, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescOnce sync.Once
	file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescData = file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDesc
)

func file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescGZIP() []byte {
	file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescOnce.Do(func() {
		file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescData)
	})
	return file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDescData
}

var file_protobuf_mpc_virtual_service_platform_apiusage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_mpc_virtual_service_platform_apiusage_proto_goTypes = []interface{}{
	(UpsertAPIUsageReceiverResponse_Status)(0), // 0: protobuf.platform.apiusage.UpsertAPIUsageReceiverResponse.Status
	(APIUsageNotifyResponse_Status)(0),         // 1: protobuf.platform.apiusage.APIUsageNotifyResponse.Status
	(*UpsertAPIUsageReceiverRequest)(nil),      // 2: protobuf.platform.apiusage.UpsertAPIUsageReceiverRequest
	(*UpsertAPIUsageReceiverResponse)(nil),     // 3: protobuf.platform.apiusage.UpsertAPIUsageReceiverResponse
	(*APIUsageNotifyResponse)(nil),             // 4: protobuf.platform.apiusage.APIUsageNotifyResponse
	(*ListAPIUsagesRequest)(nil),               // 5: protobuf.platform.apiusage.ListAPIUsagesRequest
	(*ListAPIUsagesResponse)(nil),              // 6: protobuf.platform.apiusage.ListAPIUsagesResponse
	(*paginator.Paginator)(nil),                // 7: sudo.protobuf.platform.paginator.Paginator
	(api_usage.APIUsage_Type)(0),               // 8: sudo.protobuf.api_usage.APIUsage.Type
	(api_usage.APIUsage_Status)(0),             // 9: sudo.protobuf.api_usage.APIUsage.Status
	(*api_usage.APIUsage)(nil),                 // 10: sudo.protobuf.api_usage.APIUsage
}
var file_protobuf_mpc_virtual_service_platform_apiusage_proto_depIdxs = []int32{
	0,  // 0: protobuf.platform.apiusage.UpsertAPIUsageReceiverResponse.status:type_name -> protobuf.platform.apiusage.UpsertAPIUsageReceiverResponse.Status
	1,  // 1: protobuf.platform.apiusage.APIUsageNotifyResponse.status:type_name -> protobuf.platform.apiusage.APIUsageNotifyResponse.Status
	7,  // 2: protobuf.platform.apiusage.ListAPIUsagesRequest.paginator:type_name -> sudo.protobuf.platform.paginator.Paginator
	8,  // 3: protobuf.platform.apiusage.ListAPIUsagesRequest.type:type_name -> sudo.protobuf.api_usage.APIUsage.Type
	9,  // 4: protobuf.platform.apiusage.ListAPIUsagesRequest.status:type_name -> sudo.protobuf.api_usage.APIUsage.Status
	10, // 5: protobuf.platform.apiusage.ListAPIUsagesResponse.data:type_name -> sudo.protobuf.api_usage.APIUsage
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_protobuf_mpc_virtual_service_platform_apiusage_proto_init() }
func file_protobuf_mpc_virtual_service_platform_apiusage_proto_init() {
	if File_protobuf_mpc_virtual_service_platform_apiusage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertAPIUsageReceiverRequest); i {
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
		file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertAPIUsageReceiverResponse); i {
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
		file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIUsageNotifyResponse); i {
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
		file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAPIUsagesRequest); i {
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
		file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAPIUsagesResponse); i {
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
			RawDescriptor: file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_mpc_virtual_service_platform_apiusage_proto_goTypes,
		DependencyIndexes: file_protobuf_mpc_virtual_service_platform_apiusage_proto_depIdxs,
		EnumInfos:         file_protobuf_mpc_virtual_service_platform_apiusage_proto_enumTypes,
		MessageInfos:      file_protobuf_mpc_virtual_service_platform_apiusage_proto_msgTypes,
	}.Build()
	File_protobuf_mpc_virtual_service_platform_apiusage_proto = out.File
	file_protobuf_mpc_virtual_service_platform_apiusage_proto_rawDesc = nil
	file_protobuf_mpc_virtual_service_platform_apiusage_proto_goTypes = nil
	file_protobuf_mpc_virtual_service_platform_apiusage_proto_depIdxs = nil
}

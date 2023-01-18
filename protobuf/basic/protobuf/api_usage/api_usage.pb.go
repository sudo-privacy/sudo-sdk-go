// Copyright 2022 Sudo Technology Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: sudo/protobuf/api_usage/api_usage.proto

package apiusage

import (
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

type APIUsage_Type int32

const (
	APIUsage_UNKNOWN_TYPE APIUsage_Type = 0
	APIUsage_FACTOR3s     APIUsage_Type = 1
	APIUsage_PIR          APIUsage_Type = 2
)

// Enum value maps for APIUsage_Type.
var (
	APIUsage_Type_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "FACTOR3s",
		2: "PIR",
	}
	APIUsage_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"FACTOR3s":     1,
		"PIR":          2,
	}
)

func (x APIUsage_Type) Enum() *APIUsage_Type {
	p := new(APIUsage_Type)
	*p = x
	return p
}

func (x APIUsage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIUsage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sudo_protobuf_api_usage_api_usage_proto_enumTypes[0].Descriptor()
}

func (APIUsage_Type) Type() protoreflect.EnumType {
	return &file_sudo_protobuf_api_usage_api_usage_proto_enumTypes[0]
}

func (x APIUsage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIUsage_Type.Descriptor instead.
func (APIUsage_Type) EnumDescriptor() ([]byte, []int) {
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP(), []int{0, 0}
}

type APIUsage_Status int32

const (
	APIUsage_UNKNOWN_STATUS APIUsage_Status = 0
	APIUsage_SUCCESS        APIUsage_Status = 1
	APIUsage_FAILED         APIUsage_Status = 2
)

// Enum value maps for APIUsage_Status.
var (
	APIUsage_Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "SUCCESS",
		2: "FAILED",
	}
	APIUsage_Status_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"SUCCESS":        1,
		"FAILED":         2,
	}
)

func (x APIUsage_Status) Enum() *APIUsage_Status {
	p := new(APIUsage_Status)
	*p = x
	return p
}

func (x APIUsage_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIUsage_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_sudo_protobuf_api_usage_api_usage_proto_enumTypes[1].Descriptor()
}

func (APIUsage_Status) Type() protoreflect.EnumType {
	return &file_sudo_protobuf_api_usage_api_usage_proto_enumTypes[1]
}

func (x APIUsage_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIUsage_Status.Descriptor instead.
func (APIUsage_Status) EnumDescriptor() ([]byte, []int) {
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP(), []int{0, 1}
}

type APIUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestId      string          `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"` // 与innerID共同组成unique key
	InnerId        int32           `protobuf:"varint,3,opt,name=inner_id,json=innerId,proto3" json:"inner_id,omitempty"`      // client端的innerID是0， server是1、2以此类推
	PartyId        string          `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Caller         string          `protobuf:"bytes,5,opt,name=caller,proto3" json:"caller,omitempty"`            // 发起调用者，client端记录userID，server端记录clientPartyID
	ApiId          string          `protobuf:"bytes,6,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"` // 对于三要素，APIID的值是ServiceID
	ApiPath        string          `protobuf:"bytes,7,opt,name=api_path,json=apiPath,proto3" json:"api_path,omitempty"`
	Type           APIUsage_Type   `protobuf:"varint,8,opt,name=type,proto3,enum=sudo.protobuf.api_usage.APIUsage_Type" json:"type,omitempty"`
	AdditionalInfo *AdditionalInfo `protobuf:"bytes,9,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"` // 对于不同type的API, 个性化属性保存在AdditionalInfo
	Status         APIUsage_Status `protobuf:"varint,10,opt,name=status,proto3,enum=sudo.protobuf.api_usage.APIUsage_Status" json:"status,omitempty"`
	Message        string          `protobuf:"bytes,11,opt,name=message,proto3" json:"message,omitempty"` // 如果失败，期待可以上报失败原因
	CreatedTs      int64           `protobuf:"varint,12,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	StartTime      int64           `protobuf:"varint,13,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` // 服务端收到请求的时间（精确到毫秒
	EndTime        int64           `protobuf:"varint,14,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`       // 服务端收到请求的时间（精确到毫秒
}

func (x *APIUsage) Reset() {
	*x = APIUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIUsage) ProtoMessage() {}

func (x *APIUsage) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIUsage.ProtoReflect.Descriptor instead.
func (*APIUsage) Descriptor() ([]byte, []int) {
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP(), []int{0}
}

func (x *APIUsage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *APIUsage) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *APIUsage) GetInnerId() int32 {
	if x != nil {
		return x.InnerId
	}
	return 0
}

func (x *APIUsage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *APIUsage) GetCaller() string {
	if x != nil {
		return x.Caller
	}
	return ""
}

func (x *APIUsage) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *APIUsage) GetApiPath() string {
	if x != nil {
		return x.ApiPath
	}
	return ""
}

func (x *APIUsage) GetType() APIUsage_Type {
	if x != nil {
		return x.Type
	}
	return APIUsage_UNKNOWN_TYPE
}

func (x *APIUsage) GetAdditionalInfo() *AdditionalInfo {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *APIUsage) GetStatus() APIUsage_Status {
	if x != nil {
		return x.Status
	}
	return APIUsage_UNKNOWN_STATUS
}

func (x *APIUsage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *APIUsage) GetCreatedTs() int64 {
	if x != nil {
		return x.CreatedTs
	}
	return 0
}

func (x *APIUsage) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *APIUsage) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type AdditionalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Info:
	//	*AdditionalInfo_Factor3S
	//	*AdditionalInfo_Pir
	Info isAdditionalInfo_Info `protobuf_oneof:"info"`
}

func (x *AdditionalInfo) Reset() {
	*x = AdditionalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionalInfo) ProtoMessage() {}

func (x *AdditionalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionalInfo.ProtoReflect.Descriptor instead.
func (*AdditionalInfo) Descriptor() ([]byte, []int) {
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP(), []int{1}
}

func (m *AdditionalInfo) GetInfo() isAdditionalInfo_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (x *AdditionalInfo) GetFactor3S() *Factor3SAdditionalInfo {
	if x, ok := x.GetInfo().(*AdditionalInfo_Factor3S); ok {
		return x.Factor3S
	}
	return nil
}

func (x *AdditionalInfo) GetPir() *PIRAdditionalInfo {
	if x, ok := x.GetInfo().(*AdditionalInfo_Pir); ok {
		return x.Pir
	}
	return nil
}

type isAdditionalInfo_Info interface {
	isAdditionalInfo_Info()
}

type AdditionalInfo_Factor3S struct {
	Factor3S *Factor3SAdditionalInfo `protobuf:"bytes,1,opt,name=factor3s,proto3,oneof"`
}

type AdditionalInfo_Pir struct {
	Pir *PIRAdditionalInfo `protobuf:"bytes,2,opt,name=pir,proto3,oneof"`
}

func (*AdditionalInfo_Factor3S) isAdditionalInfo_Info() {}

func (*AdditionalInfo_Pir) isAdditionalInfo_Info() {}

type Factor3SAdditionalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestCount int32 `protobuf:"varint,1,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	UsageCount   int32 `protobuf:"varint,2,opt,name=usage_count,json=usageCount,proto3" json:"usage_count,omitempty"`
}

func (x *Factor3SAdditionalInfo) Reset() {
	*x = Factor3SAdditionalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Factor3SAdditionalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Factor3SAdditionalInfo) ProtoMessage() {}

func (x *Factor3SAdditionalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Factor3SAdditionalInfo.ProtoReflect.Descriptor instead.
func (*Factor3SAdditionalInfo) Descriptor() ([]byte, []int) {
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP(), []int{2}
}

func (x *Factor3SAdditionalInfo) GetRequestCount() int32 {
	if x != nil {
		return x.RequestCount
	}
	return 0
}

func (x *Factor3SAdditionalInfo) GetUsageCount() int32 {
	if x != nil {
		return x.UsageCount
	}
	return 0
}

type PIRAdditionalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestCount int32 `protobuf:"varint,1,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	UsageCount   int32 `protobuf:"varint,2,opt,name=usage_count,json=usageCount,proto3" json:"usage_count,omitempty"`
}

func (x *PIRAdditionalInfo) Reset() {
	*x = PIRAdditionalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIRAdditionalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIRAdditionalInfo) ProtoMessage() {}

func (x *PIRAdditionalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIRAdditionalInfo.ProtoReflect.Descriptor instead.
func (*PIRAdditionalInfo) Descriptor() ([]byte, []int) {
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP(), []int{3}
}

func (x *PIRAdditionalInfo) GetRequestCount() int32 {
	if x != nil {
		return x.RequestCount
	}
	return 0
}

func (x *PIRAdditionalInfo) GetUsageCount() int32 {
	if x != nil {
		return x.UsageCount
	}
	return 0
}

var File_sudo_protobuf_api_usage_api_usage_proto protoreflect.FileDescriptor

var file_sudo_protobuf_api_usage_api_usage_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x75, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xe4, 0x04, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x69, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x50, 0x49, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x41, 0x43, 0x54, 0x4f, 0x52, 0x33, 0x73, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x49, 0x52,
	0x10, 0x02, 0x22, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x08,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x33,
	0x73, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x08, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x73, 0x12, 0x3e, 0x0a, 0x03, 0x70,
	0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x49, 0x52, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x03, 0x70, 0x69, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x5e, 0x0a, 0x16, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x73, 0x41,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x50, 0x49, 0x52, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1a,
	0x5a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x70, 0x69, 0x75, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sudo_protobuf_api_usage_api_usage_proto_rawDescOnce sync.Once
	file_sudo_protobuf_api_usage_api_usage_proto_rawDescData = file_sudo_protobuf_api_usage_api_usage_proto_rawDesc
)

func file_sudo_protobuf_api_usage_api_usage_proto_rawDescGZIP() []byte {
	file_sudo_protobuf_api_usage_api_usage_proto_rawDescOnce.Do(func() {
		file_sudo_protobuf_api_usage_api_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudo_protobuf_api_usage_api_usage_proto_rawDescData)
	})
	return file_sudo_protobuf_api_usage_api_usage_proto_rawDescData
}

var file_sudo_protobuf_api_usage_api_usage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sudo_protobuf_api_usage_api_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sudo_protobuf_api_usage_api_usage_proto_goTypes = []interface{}{
	(APIUsage_Type)(0),             // 0: sudo.protobuf.api_usage.APIUsage.Type
	(APIUsage_Status)(0),           // 1: sudo.protobuf.api_usage.APIUsage.Status
	(*APIUsage)(nil),               // 2: sudo.protobuf.api_usage.APIUsage
	(*AdditionalInfo)(nil),         // 3: sudo.protobuf.api_usage.AdditionalInfo
	(*Factor3SAdditionalInfo)(nil), // 4: sudo.protobuf.api_usage.Factor3sAdditionalInfo
	(*PIRAdditionalInfo)(nil),      // 5: sudo.protobuf.api_usage.PIRAdditionalInfo
}
var file_sudo_protobuf_api_usage_api_usage_proto_depIdxs = []int32{
	0, // 0: sudo.protobuf.api_usage.APIUsage.type:type_name -> sudo.protobuf.api_usage.APIUsage.Type
	3, // 1: sudo.protobuf.api_usage.APIUsage.additional_info:type_name -> sudo.protobuf.api_usage.AdditionalInfo
	1, // 2: sudo.protobuf.api_usage.APIUsage.status:type_name -> sudo.protobuf.api_usage.APIUsage.Status
	4, // 3: sudo.protobuf.api_usage.AdditionalInfo.factor3s:type_name -> sudo.protobuf.api_usage.Factor3sAdditionalInfo
	5, // 4: sudo.protobuf.api_usage.AdditionalInfo.pir:type_name -> sudo.protobuf.api_usage.PIRAdditionalInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sudo_protobuf_api_usage_api_usage_proto_init() }
func file_sudo_protobuf_api_usage_api_usage_proto_init() {
	if File_sudo_protobuf_api_usage_api_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIUsage); i {
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
		file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionalInfo); i {
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
		file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Factor3SAdditionalInfo); i {
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
		file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIRAdditionalInfo); i {
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
	file_sudo_protobuf_api_usage_api_usage_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AdditionalInfo_Factor3S)(nil),
		(*AdditionalInfo_Pir)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sudo_protobuf_api_usage_api_usage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sudo_protobuf_api_usage_api_usage_proto_goTypes,
		DependencyIndexes: file_sudo_protobuf_api_usage_api_usage_proto_depIdxs,
		EnumInfos:         file_sudo_protobuf_api_usage_api_usage_proto_enumTypes,
		MessageInfos:      file_sudo_protobuf_api_usage_api_usage_proto_msgTypes,
	}.Build()
	File_sudo_protobuf_api_usage_api_usage_proto = out.File
	file_sudo_protobuf_api_usage_api_usage_proto_rawDesc = nil
	file_sudo_protobuf_api_usage_api_usage_proto_goTypes = nil
	file_sudo_protobuf_api_usage_api_usage_proto_depIdxs = nil
}

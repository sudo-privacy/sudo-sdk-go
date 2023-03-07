// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: protobuf/mpc_virtual_service/platform/token.proto

package token

import (
	enums "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/enums"
	_ "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/options"
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

type GenerateTokenForPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartyId   string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	ServiceId uint64 `protobuf:"varint,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *GenerateTokenForPartyRequest) Reset() {
	*x = GenerateTokenForPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenForPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenForPartyRequest) ProtoMessage() {}

func (x *GenerateTokenForPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenForPartyRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenForPartyRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenForPartyRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *GenerateTokenForPartyRequest) GetServiceId() uint64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

type GenerateTokenForPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GenerateTokenForPartyResponse) Reset() {
	*x = GenerateTokenForPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenForPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenForPartyResponse) ProtoMessage() {}

func (x *GenerateTokenForPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenForPartyResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenForPartyResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenForPartyResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetPartyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginator *paginator.Paginator `protobuf:"bytes,1,opt,name=paginator,proto3" json:"paginator,omitempty"`
	TokenType uint32               `protobuf:"varint,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ServiceId uint64               `protobuf:"varint,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *GetPartyTokenRequest) Reset() {
	*x = GetPartyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyTokenRequest) ProtoMessage() {}

func (x *GetPartyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyTokenRequest.ProtoReflect.Descriptor instead.
func (*GetPartyTokenRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP(), []int{2}
}

func (x *GetPartyTokenRequest) GetPaginator() *paginator.Paginator {
	if x != nil {
		return x.Paginator
	}
	return nil
}

func (x *GetPartyTokenRequest) GetTokenType() uint32 {
	if x != nil {
		return x.TokenType
	}
	return 0
}

func (x *GetPartyTokenRequest) GetServiceId() uint64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

type PirServerToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string             `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PartyId      string             `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	IssuerUserId uint64             `protobuf:"varint,3,opt,name=issuer_user_id,json=issuerUserId,proto3" json:"issuer_user_id,omitempty"`
	ServiceId    uint64             `protobuf:"varint,4,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Status       enums.Token_Status `protobuf:"varint,5,opt,name=status,proto3,enum=sudo.protobuf.enums.Token_Status" json:"status,omitempty"`
	CreatedTime  uint64             `protobuf:"varint,6,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	LastUsedTime uint64             `protobuf:"varint,7,opt,name=last_used_time,json=lastUsedTime,proto3" json:"last_used_time,omitempty"`
}

func (x *PirServerToken) Reset() {
	*x = PirServerToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PirServerToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirServerToken) ProtoMessage() {}

func (x *PirServerToken) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirServerToken.ProtoReflect.Descriptor instead.
func (*PirServerToken) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP(), []int{3}
}

func (x *PirServerToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PirServerToken) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PirServerToken) GetIssuerUserId() uint64 {
	if x != nil {
		return x.IssuerUserId
	}
	return 0
}

func (x *PirServerToken) GetServiceId() uint64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *PirServerToken) GetStatus() enums.Token_Status {
	if x != nil {
		return x.Status
	}
	return enums.Token_UNKNOWN_STATUS
}

func (x *PirServerToken) GetCreatedTime() uint64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *PirServerToken) GetLastUsedTime() uint64 {
	if x != nil {
		return x.LastUsedTime
	}
	return 0
}

type GetPartyTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*PirServerToken `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Count uint64            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetPartyTokenResponse) Reset() {
	*x = GetPartyTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyTokenResponse) ProtoMessage() {}

func (x *GetPartyTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyTokenResponse.ProtoReflect.Descriptor instead.
func (*GetPartyTokenResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP(), []int{4}
}

func (x *GetPartyTokenResponse) GetData() []*PirServerToken {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetPartyTokenResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ActionOnPirServerToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string             `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Action enums.Token_Action `protobuf:"varint,2,opt,name=action,proto3,enum=sudo.protobuf.enums.Token_Action" json:"action,omitempty"`
}

func (x *ActionOnPirServerToken) Reset() {
	*x = ActionOnPirServerToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionOnPirServerToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionOnPirServerToken) ProtoMessage() {}

func (x *ActionOnPirServerToken) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionOnPirServerToken.ProtoReflect.Descriptor instead.
func (*ActionOnPirServerToken) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP(), []int{5}
}

func (x *ActionOnPirServerToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ActionOnPirServerToken) GetAction() enums.Token_Action {
	if x != nil {
		return x.Action
	}
	return enums.Token_UNKNOWN_ACTION
}

var File_protobuf_mpc_virtual_service_platform_token_proto protoreflect.FileDescriptor

var file_protobuf_mpc_virtual_service_platform_token_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1f, 0x73, 0x75,
	0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x73,
	0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63,
	0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x1c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a,
	0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x6f,
	0x72, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x03, 0xc0,
	0x3e, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x0e,
	0x50, 0x69, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x69, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e,
	0x50, 0x69, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x36, 0x5a, 0x34, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_mpc_virtual_service_platform_token_proto_rawDescOnce sync.Once
	file_protobuf_mpc_virtual_service_platform_token_proto_rawDescData = file_protobuf_mpc_virtual_service_platform_token_proto_rawDesc
)

func file_protobuf_mpc_virtual_service_platform_token_proto_rawDescGZIP() []byte {
	file_protobuf_mpc_virtual_service_platform_token_proto_rawDescOnce.Do(func() {
		file_protobuf_mpc_virtual_service_platform_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_mpc_virtual_service_platform_token_proto_rawDescData)
	})
	return file_protobuf_mpc_virtual_service_platform_token_proto_rawDescData
}

var file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protobuf_mpc_virtual_service_platform_token_proto_goTypes = []interface{}{
	(*GenerateTokenForPartyRequest)(nil),  // 0: protobuf.platform.token.GenerateTokenForPartyRequest
	(*GenerateTokenForPartyResponse)(nil), // 1: protobuf.platform.token.GenerateTokenForPartyResponse
	(*GetPartyTokenRequest)(nil),          // 2: protobuf.platform.token.GetPartyTokenRequest
	(*PirServerToken)(nil),                // 3: protobuf.platform.token.PirServerToken
	(*GetPartyTokenResponse)(nil),         // 4: protobuf.platform.token.GetPartyTokenResponse
	(*ActionOnPirServerToken)(nil),        // 5: protobuf.platform.token.ActionOnPirServerToken
	(*paginator.Paginator)(nil),           // 6: sudo.protobuf.platform.paginator.Paginator
	(enums.Token_Status)(0),               // 7: sudo.protobuf.enums.Token.Status
	(enums.Token_Action)(0),               // 8: sudo.protobuf.enums.Token.Action
}
var file_protobuf_mpc_virtual_service_platform_token_proto_depIdxs = []int32{
	6, // 0: protobuf.platform.token.GetPartyTokenRequest.paginator:type_name -> sudo.protobuf.platform.paginator.Paginator
	7, // 1: protobuf.platform.token.PirServerToken.status:type_name -> sudo.protobuf.enums.Token.Status
	3, // 2: protobuf.platform.token.GetPartyTokenResponse.data:type_name -> protobuf.platform.token.PirServerToken
	8, // 3: protobuf.platform.token.ActionOnPirServerToken.action:type_name -> sudo.protobuf.enums.Token.Action
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protobuf_mpc_virtual_service_platform_token_proto_init() }
func file_protobuf_mpc_virtual_service_platform_token_proto_init() {
	if File_protobuf_mpc_virtual_service_platform_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenForPartyRequest); i {
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
		file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenForPartyResponse); i {
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
		file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyTokenRequest); i {
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
		file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PirServerToken); i {
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
		file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyTokenResponse); i {
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
		file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionOnPirServerToken); i {
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
			RawDescriptor: file_protobuf_mpc_virtual_service_platform_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_mpc_virtual_service_platform_token_proto_goTypes,
		DependencyIndexes: file_protobuf_mpc_virtual_service_platform_token_proto_depIdxs,
		MessageInfos:      file_protobuf_mpc_virtual_service_platform_token_proto_msgTypes,
	}.Build()
	File_protobuf_mpc_virtual_service_platform_token_proto = out.File
	file_protobuf_mpc_virtual_service_platform_token_proto_rawDesc = nil
	file_protobuf_mpc_virtual_service_platform_token_proto_goTypes = nil
	file_protobuf_mpc_virtual_service_platform_token_proto_depIdxs = nil
}

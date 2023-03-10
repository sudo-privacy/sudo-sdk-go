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
// source: sudo/service/enums/online_svc_enums.proto

package enums

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

type DataMode int32

const (
	DataMode_UNKNOWN_DATA_MODE DataMode = 0
	DataMode_VTABLE            DataMode = 1
	DataMode_API               DataMode = 2
	DataMode_MYSQL             DataMode = 3
)

// Enum value maps for DataMode.
var (
	DataMode_name = map[int32]string{
		0: "UNKNOWN_DATA_MODE",
		1: "VTABLE",
		2: "API",
		3: "MYSQL",
	}
	DataMode_value = map[string]int32{
		"UNKNOWN_DATA_MODE": 0,
		"VTABLE":            1,
		"API":               2,
		"MYSQL":             3,
	}
)

func (x DataMode) Enum() *DataMode {
	p := new(DataMode)
	*p = x
	return p
}

func (x DataMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataMode) Descriptor() protoreflect.EnumDescriptor {
	return file_sudo_service_enums_online_svc_enums_proto_enumTypes[0].Descriptor()
}

func (DataMode) Type() protoreflect.EnumType {
	return &file_sudo_service_enums_online_svc_enums_proto_enumTypes[0]
}

func (x DataMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataMode.Descriptor instead.
func (DataMode) EnumDescriptor() ([]byte, []int) {
	return file_sudo_service_enums_online_svc_enums_proto_rawDescGZIP(), []int{0}
}

type SVCType int32

const (
	SVCType_UNKNOWN_TYPE SVCType = 0
	SVCType_FACTOR3s     SVCType = 1
	SVCType_PIR          SVCType = 2
)

// Enum value maps for SVCType.
var (
	SVCType_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "FACTOR3s",
		2: "PIR",
	}
	SVCType_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"FACTOR3s":     1,
		"PIR":          2,
	}
)

func (x SVCType) Enum() *SVCType {
	p := new(SVCType)
	*p = x
	return p
}

func (x SVCType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SVCType) Descriptor() protoreflect.EnumDescriptor {
	return file_sudo_service_enums_online_svc_enums_proto_enumTypes[1].Descriptor()
}

func (SVCType) Type() protoreflect.EnumType {
	return &file_sudo_service_enums_online_svc_enums_proto_enumTypes[1]
}

func (x SVCType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SVCType.Descriptor instead.
func (SVCType) EnumDescriptor() ([]byte, []int) {
	return file_sudo_service_enums_online_svc_enums_proto_rawDescGZIP(), []int{1}
}

type AuthMethod int32

const (
	AuthMethod_UNKNOWN_AUTH_METHOD AuthMethod = 0
	AuthMethod_NO_AUTH             AuthMethod = 1
	AuthMethod_BASIC_AUTH          AuthMethod = 2
)

// Enum value maps for AuthMethod.
var (
	AuthMethod_name = map[int32]string{
		0: "UNKNOWN_AUTH_METHOD",
		1: "NO_AUTH",
		2: "BASIC_AUTH",
	}
	AuthMethod_value = map[string]int32{
		"UNKNOWN_AUTH_METHOD": 0,
		"NO_AUTH":             1,
		"BASIC_AUTH":          2,
	}
)

func (x AuthMethod) Enum() *AuthMethod {
	p := new(AuthMethod)
	*p = x
	return p
}

func (x AuthMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_sudo_service_enums_online_svc_enums_proto_enumTypes[2].Descriptor()
}

func (AuthMethod) Type() protoreflect.EnumType {
	return &file_sudo_service_enums_online_svc_enums_proto_enumTypes[2]
}

func (x AuthMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMethod.Descriptor instead.
func (AuthMethod) EnumDescriptor() ([]byte, []int) {
	return file_sudo_service_enums_online_svc_enums_proto_rawDescGZIP(), []int{2}
}

var File_sudo_service_enums_online_svc_enums_proto protoreflect.FileDescriptor

var file_sudo_service_enums_online_svc_enums_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x5f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x75, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x41, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x56,
	0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x03, 0x2a, 0x32, 0x0a, 0x07, 0x53,
	0x56, 0x43, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x43, 0x54,
	0x4f, 0x52, 0x33, 0x73, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x49, 0x52, 0x10, 0x02, 0x2a,
	0x42, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x17, 0x0a,
	0x13, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x41, 0x53, 0x49, 0x43, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sudo_service_enums_online_svc_enums_proto_rawDescOnce sync.Once
	file_sudo_service_enums_online_svc_enums_proto_rawDescData = file_sudo_service_enums_online_svc_enums_proto_rawDesc
)

func file_sudo_service_enums_online_svc_enums_proto_rawDescGZIP() []byte {
	file_sudo_service_enums_online_svc_enums_proto_rawDescOnce.Do(func() {
		file_sudo_service_enums_online_svc_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudo_service_enums_online_svc_enums_proto_rawDescData)
	})
	return file_sudo_service_enums_online_svc_enums_proto_rawDescData
}

var file_sudo_service_enums_online_svc_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sudo_service_enums_online_svc_enums_proto_goTypes = []interface{}{
	(DataMode)(0),   // 0: sudo.protobuf.service.enums.DataMode
	(SVCType)(0),    // 1: sudo.protobuf.service.enums.SVCType
	(AuthMethod)(0), // 2: sudo.protobuf.service.enums.AuthMethod
}
var file_sudo_service_enums_online_svc_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sudo_service_enums_online_svc_enums_proto_init() }
func file_sudo_service_enums_online_svc_enums_proto_init() {
	if File_sudo_service_enums_online_svc_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sudo_service_enums_online_svc_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sudo_service_enums_online_svc_enums_proto_goTypes,
		DependencyIndexes: file_sudo_service_enums_online_svc_enums_proto_depIdxs,
		EnumInfos:         file_sudo_service_enums_online_svc_enums_proto_enumTypes,
	}.Build()
	File_sudo_service_enums_online_svc_enums_proto = out.File
	file_sudo_service_enums_online_svc_enums_proto_rawDesc = nil
	file_sudo_service_enums_online_svc_enums_proto_goTypes = nil
	file_sudo_service_enums_online_svc_enums_proto_depIdxs = nil
}

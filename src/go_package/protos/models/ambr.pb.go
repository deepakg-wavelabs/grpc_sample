//
//Nudm_SDM
//
//Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
//
//The version of the OpenAPI document: 2.1.7
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.6.1
// source: ambr.proto

package models

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

type Ambr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uplink   string `protobuf:"bytes,301729516,opt,name=uplink,proto3" json:"uplink,omitempty"`
	Downlink string `protobuf:"bytes,354071454,opt,name=downlink,proto3" json:"downlink,omitempty"`
}

func (x *Ambr) Reset() {
	*x = Ambr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ambr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ambr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ambr) ProtoMessage() {}

func (x *Ambr) ProtoReflect() protoreflect.Message {
	mi := &file_ambr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ambr.ProtoReflect.Descriptor instead.
func (*Ambr) Descriptor() ([]byte, []int) {
	return file_ambr_proto_rawDescGZIP(), []int{0}
}

func (x *Ambr) GetUplink() string {
	if x != nil {
		return x.Uplink
	}
	return ""
}

func (x *Ambr) GetDownlink() string {
	if x != nil {
		return x.Downlink
	}
	return ""
}

var File_ambr_proto protoreflect.FileDescriptor

var file_ambr_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x6d, 0x62, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22,
	0x42, 0x0a, 0x04, 0x41, 0x6d, 0x62, 0x72, 0x12, 0x1a, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0xec, 0x8d, 0xf0, 0x8f, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x9e, 0xe7, 0xea, 0xa8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ambr_proto_rawDescOnce sync.Once
	file_ambr_proto_rawDescData = file_ambr_proto_rawDesc
)

func file_ambr_proto_rawDescGZIP() []byte {
	file_ambr_proto_rawDescOnce.Do(func() {
		file_ambr_proto_rawDescData = protoimpl.X.CompressGZIP(file_ambr_proto_rawDescData)
	})
	return file_ambr_proto_rawDescData
}

var file_ambr_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ambr_proto_goTypes = []interface{}{
	(*Ambr)(nil), // 0: go_package.protos.Ambr
}
var file_ambr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ambr_proto_init() }
func file_ambr_proto_init() {
	if File_ambr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ambr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ambr); i {
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
			RawDescriptor: file_ambr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ambr_proto_goTypes,
		DependencyIndexes: file_ambr_proto_depIdxs,
		MessageInfos:      file_ambr_proto_msgTypes,
	}.Build()
	File_ambr_proto = out.File
	file_ambr_proto_rawDesc = nil
	file_ambr_proto_goTypes = nil
	file_ambr_proto_depIdxs = nil
}

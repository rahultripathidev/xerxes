// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: containerManager.proto

package containerManager

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ContainerMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XCid    string `protobuf:"bytes,1,opt,name=_cid,json=Cid,proto3" json:"_cid,omitempty"`
	XNodeId string `protobuf:"bytes,2,opt,name=_nodeId,json=NodeId,proto3" json:"_nodeId,omitempty"`
	Force   bool   `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	Timeout int64  `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *ContainerMeta) Reset() {
	*x = ContainerMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_containerManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerMeta) ProtoMessage() {}

func (x *ContainerMeta) ProtoReflect() protoreflect.Message {
	mi := &file_containerManager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerMeta.ProtoReflect.Descriptor instead.
func (*ContainerMeta) Descriptor() ([]byte, []int) {
	return file_containerManager_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerMeta) GetXCid() string {
	if x != nil {
		return x.XCid
	}
	return ""
}

func (x *ContainerMeta) GetXNodeId() string {
	if x != nil {
		return x.XNodeId
	}
	return ""
}

func (x *ContainerMeta) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *ContainerMeta) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type DeleteContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *ContainerMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *DeleteContainerRequest) Reset() {
	*x = DeleteContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_containerManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContainerRequest) ProtoMessage() {}

func (x *DeleteContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_containerManager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContainerRequest.ProtoReflect.Descriptor instead.
func (*DeleteContainerRequest) Descriptor() ([]byte, []int) {
	return file_containerManager_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteContainerRequest) GetMeta() *ContainerMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type DeleteContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteContainerResponse) Reset() {
	*x = DeleteContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_containerManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContainerResponse) ProtoMessage() {}

func (x *DeleteContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_containerManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContainerResponse.ProtoReflect.Descriptor instead.
func (*DeleteContainerResponse) Descriptor() ([]byte, []int) {
	return file_containerManager_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteContainerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_containerManager_proto protoreflect.FileDescriptor

var file_containerManager_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x11, 0x0a, 0x04, 0x5f, 0x63, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x3c, 0x0a, 0x16, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x58, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_containerManager_proto_rawDescOnce sync.Once
	file_containerManager_proto_rawDescData = file_containerManager_proto_rawDesc
)

func file_containerManager_proto_rawDescGZIP() []byte {
	file_containerManager_proto_rawDescOnce.Do(func() {
		file_containerManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_containerManager_proto_rawDescData)
	})
	return file_containerManager_proto_rawDescData
}

var file_containerManager_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_containerManager_proto_goTypes = []interface{}{
	(*ContainerMeta)(nil),           // 0: containerMeta
	(*DeleteContainerRequest)(nil),  // 1: deleteContainerRequest
	(*DeleteContainerResponse)(nil), // 2: deleteContainerResponse
}
var file_containerManager_proto_depIdxs = []int32{
	0, // 0: deleteContainerRequest.meta:type_name -> containerMeta
	1, // 1: containerManager.deleteContainer:input_type -> deleteContainerRequest
	2, // 2: containerManager.deleteContainer:output_type -> deleteContainerResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_containerManager_proto_init() }
func file_containerManager_proto_init() {
	if File_containerManager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_containerManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerMeta); i {
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
		file_containerManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteContainerRequest); i {
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
		file_containerManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteContainerResponse); i {
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
			RawDescriptor: file_containerManager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_containerManager_proto_goTypes,
		DependencyIndexes: file_containerManager_proto_depIdxs,
		MessageInfos:      file_containerManager_proto_msgTypes,
	}.Build()
	File_containerManager_proto = out.File
	file_containerManager_proto_rawDesc = nil
	file_containerManager_proto_goTypes = nil
	file_containerManager_proto_depIdxs = nil
}

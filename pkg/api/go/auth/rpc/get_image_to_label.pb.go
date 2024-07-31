// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: auth/rpc/get_image_to_label.proto

package rpc

import (
	model "github.com/manhrev/labeler/pkg/api/go/auth/model"
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

type GetImageToLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetImageToLabelRequest) Reset() {
	*x = GetImageToLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_rpc_get_image_to_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageToLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageToLabelRequest) ProtoMessage() {}

func (x *GetImageToLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_rpc_get_image_to_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageToLabelRequest.ProtoReflect.Descriptor instead.
func (*GetImageToLabelRequest) Descriptor() ([]byte, []int) {
	return file_auth_rpc_get_image_to_label_proto_rawDescGZIP(), []int{0}
}

type GetImageToLabelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *model.Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetImageToLabelResponse) Reset() {
	*x = GetImageToLabelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_rpc_get_image_to_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageToLabelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageToLabelResponse) ProtoMessage() {}

func (x *GetImageToLabelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_rpc_get_image_to_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageToLabelResponse.ProtoReflect.Descriptor instead.
func (*GetImageToLabelResponse) Descriptor() ([]byte, []int) {
	return file_auth_rpc_get_image_to_label_proto_rawDescGZIP(), []int{1}
}

func (x *GetImageToLabelResponse) GetImage() *model.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_auth_rpc_get_image_to_label_proto protoreflect.FileDescriptor

var file_auth_rpc_get_image_to_label_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x6e, 0x68, 0x72, 0x65, 0x76, 0x2f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_rpc_get_image_to_label_proto_rawDescOnce sync.Once
	file_auth_rpc_get_image_to_label_proto_rawDescData = file_auth_rpc_get_image_to_label_proto_rawDesc
)

func file_auth_rpc_get_image_to_label_proto_rawDescGZIP() []byte {
	file_auth_rpc_get_image_to_label_proto_rawDescOnce.Do(func() {
		file_auth_rpc_get_image_to_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_rpc_get_image_to_label_proto_rawDescData)
	})
	return file_auth_rpc_get_image_to_label_proto_rawDescData
}

var file_auth_rpc_get_image_to_label_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_rpc_get_image_to_label_proto_goTypes = []interface{}{
	(*GetImageToLabelRequest)(nil),  // 0: rpc.GetImageToLabelRequest
	(*GetImageToLabelResponse)(nil), // 1: rpc.GetImageToLabelResponse
	(*model.Image)(nil),             // 2: model.Image
}
var file_auth_rpc_get_image_to_label_proto_depIdxs = []int32{
	2, // 0: rpc.GetImageToLabelResponse.image:type_name -> model.Image
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_rpc_get_image_to_label_proto_init() }
func file_auth_rpc_get_image_to_label_proto_init() {
	if File_auth_rpc_get_image_to_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_rpc_get_image_to_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageToLabelRequest); i {
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
		file_auth_rpc_get_image_to_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageToLabelResponse); i {
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
			RawDescriptor: file_auth_rpc_get_image_to_label_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_rpc_get_image_to_label_proto_goTypes,
		DependencyIndexes: file_auth_rpc_get_image_to_label_proto_depIdxs,
		MessageInfos:      file_auth_rpc_get_image_to_label_proto_msgTypes,
	}.Build()
	File_auth_rpc_get_image_to_label_proto = out.File
	file_auth_rpc_get_image_to_label_proto_rawDesc = nil
	file_auth_rpc_get_image_to_label_proto_goTypes = nil
	file_auth_rpc_get_image_to_label_proto_depIdxs = nil
}

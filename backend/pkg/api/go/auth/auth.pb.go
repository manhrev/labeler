// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: auth/auth.proto

package auth

import (
	rpc "github.com/manhrev/labeler/pkg/api/go/auth/rpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65,
	0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x79, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6d, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x05,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x0e, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x12, 0x23, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x14, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x14, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x79,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x6e, 0x68, 0x72,
	0x65, 0x76, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_auth_proto_goTypes = []interface{}{
	(*rpc.SignUpRequest)(nil),                   // 0: rpc.SignUpRequest
	(*rpc.LoginRequest)(nil),                    // 1: rpc.LoginRequest
	(*rpc.MeRequest)(nil),                       // 2: rpc.MeRequest
	(*rpc.GetImageToLabelRequest)(nil),          // 3: rpc.GetImageToLabelRequest
	(*rpc.UpdateImageAfterLabeledRequest)(nil),  // 4: rpc.UpdateImageAfterLabeledRequest
	(*rpc.RollbackLabeledImageRequest)(nil),     // 5: rpc.RollbackLabeledImageRequest
	(*rpc.GetMyLabeledImageRequest)(nil),        // 6: rpc.GetMyLabeledImageRequest
	(*rpc.CountMyLabeledImagesRequest)(nil),     // 7: rpc.CountMyLabeledImagesRequest
	(*rpc.GetLabeledImagesRequest)(nil),         // 8: rpc.GetLabeledImagesRequest
	(*rpc.SignUpResponse)(nil),                  // 9: rpc.SignUpResponse
	(*rpc.LoginResponse)(nil),                   // 10: rpc.LoginResponse
	(*rpc.MeResponse)(nil),                      // 11: rpc.MeResponse
	(*rpc.GetImageToLabelResponse)(nil),         // 12: rpc.GetImageToLabelResponse
	(*rpc.UpdateImageAfterLabeledResponse)(nil), // 13: rpc.UpdateImageAfterLabeledResponse
	(*rpc.RollbackLabeledImageResponse)(nil),    // 14: rpc.RollbackLabeledImageResponse
	(*rpc.GetMyLabeledImageResponse)(nil),       // 15: rpc.GetMyLabeledImageResponse
	(*rpc.CountMyLabeledImagesResponse)(nil),    // 16: rpc.CountMyLabeledImagesResponse
	(*rpc.GetLabeledImagesResponse)(nil),        // 17: rpc.GetLabeledImagesResponse
}
var file_auth_auth_proto_depIdxs = []int32{
	0,  // 0: auth.AuthService.SignUp:input_type -> rpc.SignUpRequest
	1,  // 1: auth.AuthService.Login:input_type -> rpc.LoginRequest
	2,  // 2: auth.AuthService.Me:input_type -> rpc.MeRequest
	3,  // 3: auth.AuthService.GetImageToLabel:input_type -> rpc.GetImageToLabelRequest
	4,  // 4: auth.AuthService.UpdateImageAfterLabeled:input_type -> rpc.UpdateImageAfterLabeledRequest
	5,  // 5: auth.AuthService.RollbackLabeledImage:input_type -> rpc.RollbackLabeledImageRequest
	6,  // 6: auth.AuthService.GetMyLabeledImage:input_type -> rpc.GetMyLabeledImageRequest
	7,  // 7: auth.AuthService.CountMyLabeledImages:input_type -> rpc.CountMyLabeledImagesRequest
	8,  // 8: auth.AuthService.GetLabeledImages:input_type -> rpc.GetLabeledImagesRequest
	9,  // 9: auth.AuthService.SignUp:output_type -> rpc.SignUpResponse
	10, // 10: auth.AuthService.Login:output_type -> rpc.LoginResponse
	11, // 11: auth.AuthService.Me:output_type -> rpc.MeResponse
	12, // 12: auth.AuthService.GetImageToLabel:output_type -> rpc.GetImageToLabelResponse
	13, // 13: auth.AuthService.UpdateImageAfterLabeled:output_type -> rpc.UpdateImageAfterLabeledResponse
	14, // 14: auth.AuthService.RollbackLabeledImage:output_type -> rpc.RollbackLabeledImageResponse
	15, // 15: auth.AuthService.GetMyLabeledImage:output_type -> rpc.GetMyLabeledImageResponse
	16, // 16: auth.AuthService.CountMyLabeledImages:output_type -> rpc.CountMyLabeledImagesResponse
	17, // 17: auth.AuthService.GetLabeledImages:output_type -> rpc.GetLabeledImagesResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: backend/services.proto

package backend

import (
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

var File_backend_services_proto protoreflect.FileDescriptor

var file_backend_services_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x1a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x04, 0x0a, 0x0e, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x12, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x50, 0x6f,
	0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc8, 0x01, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09,
	0x50, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6e, 0x0a, 0x17, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69,
	0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x68, 0x6b, 0x69, 0x6c, 0x61, 0x62, 0x2f, 0x53, 0x55, 0x2d, 0x43, 0x53, 0x65, 0x78, 0x70, 0x41,
	0x2d, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_backend_services_proto_goTypes = []interface{}{
	(*GetRankingRequest)(nil),               // 0: backend.GetRankingRequest
	(*PostSubmitRequest)(nil),               // 1: backend.PostSubmitRequest
	(*GetSubmitRequest)(nil),                // 2: backend.GetSubmitRequest
	(*ListSubmitsRequest)(nil),              // 3: backend.ListSubmitsRequest
	(*PostLoginRequest)(nil),                // 4: backend.PostLoginRequest
	(*ListContestsRequest)(nil),             // 5: backend.ListContestsRequest
	(*VerifyTokenRequest)(nil),              // 6: backend.VerifyTokenRequest
	(*PingUnaryRequest)(nil),                // 7: backend.PingUnaryRequest
	(*PingServerSideStreamingRequest)(nil),  // 8: backend.PingServerSideStreamingRequest
	(*GetRankingResponse)(nil),              // 9: backend.GetRankingResponse
	(*PostSubmitResponse)(nil),              // 10: backend.PostSubmitResponse
	(*GetSubmitResponse)(nil),               // 11: backend.GetSubmitResponse
	(*ListSubmitsResponse)(nil),             // 12: backend.ListSubmitsResponse
	(*PostLoginResponse)(nil),               // 13: backend.PostLoginResponse
	(*ListContestsResponse)(nil),            // 14: backend.ListContestsResponse
	(*VerifyTokenResponse)(nil),             // 15: backend.VerifyTokenResponse
	(*PingUnaryResponse)(nil),               // 16: backend.PingUnaryResponse
	(*PingServerSideStreamingResponse)(nil), // 17: backend.PingServerSideStreamingResponse
}
var file_backend_services_proto_depIdxs = []int32{
	0,  // 0: backend.BackendService.GetRanking:input_type -> backend.GetRankingRequest
	1,  // 1: backend.BackendService.PostSubmit:input_type -> backend.PostSubmitRequest
	2,  // 2: backend.BackendService.GetSubmit:input_type -> backend.GetSubmitRequest
	3,  // 3: backend.BackendService.ListSubmits:input_type -> backend.ListSubmitsRequest
	4,  // 4: backend.BackendService.PostLogin:input_type -> backend.PostLoginRequest
	5,  // 5: backend.BackendService.ListContests:input_type -> backend.ListContestsRequest
	6,  // 6: backend.BackendService.VerifyToken:input_type -> backend.VerifyTokenRequest
	7,  // 7: backend.HealthcheckService.PingUnary:input_type -> backend.PingUnaryRequest
	8,  // 8: backend.HealthcheckService.PingServerSideStreaming:input_type -> backend.PingServerSideStreamingRequest
	9,  // 9: backend.BackendService.GetRanking:output_type -> backend.GetRankingResponse
	10, // 10: backend.BackendService.PostSubmit:output_type -> backend.PostSubmitResponse
	11, // 11: backend.BackendService.GetSubmit:output_type -> backend.GetSubmitResponse
	12, // 12: backend.BackendService.ListSubmits:output_type -> backend.ListSubmitsResponse
	13, // 13: backend.BackendService.PostLogin:output_type -> backend.PostLoginResponse
	14, // 14: backend.BackendService.ListContests:output_type -> backend.ListContestsResponse
	15, // 15: backend.BackendService.VerifyToken:output_type -> backend.VerifyTokenResponse
	16, // 16: backend.HealthcheckService.PingUnary:output_type -> backend.PingUnaryResponse
	17, // 17: backend.HealthcheckService.PingServerSideStreaming:output_type -> backend.PingServerSideStreamingResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_backend_services_proto_init() }
func file_backend_services_proto_init() {
	if File_backend_services_proto != nil {
		return
	}
	file_backend_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_backend_services_proto_goTypes,
		DependencyIndexes: file_backend_services_proto_depIdxs,
	}.Build()
	File_backend_services_proto = out.File
	file_backend_services_proto_rawDesc = nil
	file_backend_services_proto_goTypes = nil
	file_backend_services_proto_depIdxs = nil
}

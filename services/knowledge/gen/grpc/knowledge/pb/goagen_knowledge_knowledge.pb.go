// Code generated with goa v3.21.1, DO NOT EDIT.
//
// knowledge protocol buffer definition
//
// Command:
// $ goa gen
// github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/design/api
// -o ./services/knowledge/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: goagen_knowledge_knowledge.proto

package knowledgepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_goagen_knowledge_knowledge_proto protoreflect.FileDescriptor

const file_goagen_knowledge_knowledge_proto_rawDesc = "" +
	"\n" +
	" goagen_knowledge_knowledge.proto\x12\tknowledge2\v\n" +
	"\tKnowledgeB\x0eZ\f/knowledgepbb\x06proto3"

var file_goagen_knowledge_knowledge_proto_goTypes = []any{}
var file_goagen_knowledge_knowledge_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goagen_knowledge_knowledge_proto_init() }
func file_goagen_knowledge_knowledge_proto_init() {
	if File_goagen_knowledge_knowledge_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_goagen_knowledge_knowledge_proto_rawDesc), len(file_goagen_knowledge_knowledge_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goagen_knowledge_knowledge_proto_goTypes,
		DependencyIndexes: file_goagen_knowledge_knowledge_proto_depIdxs,
	}.Build()
	File_goagen_knowledge_knowledge_proto = out.File
	file_goagen_knowledge_knowledge_proto_goTypes = nil
	file_goagen_knowledge_knowledge_proto_depIdxs = nil
}

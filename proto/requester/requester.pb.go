// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/requester/requester.proto

package requester

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

type DispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_requester_requester_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_requester_requester_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_requester_requester_proto_rawDescGZIP(), []int{0}
}

func (x *DispatchRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DispatchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	LatencyMs uint32 `protobuf:"varint,3,opt,name=latency_ms,json=latencyMs,proto3" json:"latency_ms,omitempty"`
}

func (x *DispatchReply) Reset() {
	*x = DispatchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_requester_requester_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchReply) ProtoMessage() {}

func (x *DispatchReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_requester_requester_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchReply.ProtoReflect.Descriptor instead.
func (*DispatchReply) Descriptor() ([]byte, []int) {
	return file_proto_requester_requester_proto_rawDescGZIP(), []int{1}
}

func (x *DispatchReply) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DispatchReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DispatchReply) GetLatencyMs() uint32 {
	if x != nil {
		return x.LatencyMs
	}
	return 0
}

var File_proto_requester_requester_proto protoreflect.FileDescriptor

var file_proto_requester_requester_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x0d, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x73, 0x32, 0x4f, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_requester_requester_proto_rawDescOnce sync.Once
	file_proto_requester_requester_proto_rawDescData = file_proto_requester_requester_proto_rawDesc
)

func file_proto_requester_requester_proto_rawDescGZIP() []byte {
	file_proto_requester_requester_proto_rawDescOnce.Do(func() {
		file_proto_requester_requester_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_requester_requester_proto_rawDescData)
	})
	return file_proto_requester_requester_proto_rawDescData
}

var file_proto_requester_requester_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_requester_requester_proto_goTypes = []interface{}{
	(*DispatchRequest)(nil), // 0: requester.DispatchRequest
	(*DispatchReply)(nil),   // 1: requester.DispatchReply
}
var file_proto_requester_requester_proto_depIdxs = []int32{
	0, // 0: requester.Requester.Dispatch:input_type -> requester.DispatchRequest
	1, // 1: requester.Requester.Dispatch:output_type -> requester.DispatchReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_requester_requester_proto_init() }
func file_proto_requester_requester_proto_init() {
	if File_proto_requester_requester_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_requester_requester_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRequest); i {
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
		file_proto_requester_requester_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchReply); i {
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
			RawDescriptor: file_proto_requester_requester_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_requester_requester_proto_goTypes,
		DependencyIndexes: file_proto_requester_requester_proto_depIdxs,
		MessageInfos:      file_proto_requester_requester_proto_msgTypes,
	}.Build()
	File_proto_requester_requester_proto = out.File
	file_proto_requester_requester_proto_rawDesc = nil
	file_proto_requester_requester_proto_goTypes = nil
	file_proto_requester_requester_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: slowpizza.proto

package slowpizza

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

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item             string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	ConfirmCount     int32  `protobuf:"varint,2,opt,name=confirm_count,json=confirmCount,proto3" json:"confirm_count,omitempty"`
	ConfirmIntervalS int64  `protobuf:"varint,3,opt,name=confirm_interval_s,json=confirmIntervalS,proto3" json:"confirm_interval_s,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slowpizza_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slowpizza_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_slowpizza_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *OrderRequest) GetConfirmCount() int32 {
	if x != nil {
		return x.ConfirmCount
	}
	return 0
}

func (x *OrderRequest) GetConfirmIntervalS() int64 {
	if x != nil {
		return x.ConfirmIntervalS
	}
	return 0
}

type OrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OrderReply) Reset() {
	*x = OrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slowpizza_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReply) ProtoMessage() {}

func (x *OrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_slowpizza_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReply.ProtoReflect.Descriptor instead.
func (*OrderReply) Descriptor() ([]byte, []int) {
	return file_slowpizza_proto_rawDescGZIP(), []int{1}
}

func (x *OrderReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_slowpizza_proto protoreflect.FileDescriptor

var file_slowpizza_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x6c, 0x6f, 0x77, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x22,
	0x26, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x82, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x35, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2f, 0x73, 0x6c, 0x6f, 0x77, 0x70, 0x69, 0x7a, 0x7a, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_slowpizza_proto_rawDescOnce sync.Once
	file_slowpizza_proto_rawDescData = file_slowpizza_proto_rawDesc
)

func file_slowpizza_proto_rawDescGZIP() []byte {
	file_slowpizza_proto_rawDescOnce.Do(func() {
		file_slowpizza_proto_rawDescData = protoimpl.X.CompressGZIP(file_slowpizza_proto_rawDescData)
	})
	return file_slowpizza_proto_rawDescData
}

var file_slowpizza_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_slowpizza_proto_goTypes = []interface{}{
	(*OrderRequest)(nil), // 0: proto.OrderRequest
	(*OrderReply)(nil),   // 1: proto.OrderReply
}
var file_slowpizza_proto_depIdxs = []int32{
	0, // 0: proto.Agent.OrderItem:input_type -> proto.OrderRequest
	0, // 1: proto.Agent.OrderMultipleItems:input_type -> proto.OrderRequest
	1, // 2: proto.Agent.OrderItem:output_type -> proto.OrderReply
	1, // 3: proto.Agent.OrderMultipleItems:output_type -> proto.OrderReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_slowpizza_proto_init() }
func file_slowpizza_proto_init() {
	if File_slowpizza_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slowpizza_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_slowpizza_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReply); i {
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
			RawDescriptor: file_slowpizza_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slowpizza_proto_goTypes,
		DependencyIndexes: file_slowpizza_proto_depIdxs,
		MessageInfos:      file_slowpizza_proto_msgTypes,
	}.Build()
	File_slowpizza_proto = out.File
	file_slowpizza_proto_rawDesc = nil
	file_slowpizza_proto_goTypes = nil
	file_slowpizza_proto_depIdxs = nil
}

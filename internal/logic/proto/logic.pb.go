// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: internal/logic/proto/logic.proto

package logic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_logic_proto_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_internal_logic_proto_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_internal_logic_proto_logic_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type C2CMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId uint64     `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	To    uint64     `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Data  *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *C2CMsgReq) Reset() {
	*x = C2CMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_logic_proto_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CMsgReq) ProtoMessage() {}

func (x *C2CMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_logic_proto_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CMsgReq.ProtoReflect.Descriptor instead.
func (*C2CMsgReq) Descriptor() ([]byte, []int) {
	return file_internal_logic_proto_logic_proto_rawDescGZIP(), []int{1}
}

func (x *C2CMsgReq) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *C2CMsgReq) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *C2CMsgReq) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	MsgIds []uint64 `protobuf:"varint,2,rep,packed,name=msg_ids,json=msgIds,proto3" json:"msg_ids,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_logic_proto_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_internal_logic_proto_logic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_internal_logic_proto_logic_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Ack) GetMsgIds() []uint64 {
	if x != nil {
		return x.MsgIds
	}
	return nil
}

type C2CMsgRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgIds []uint64 `protobuf:"varint,1,rep,packed,name=msg_ids,json=msgIds,proto3" json:"msg_ids,omitempty"`
}

func (x *C2CMsgRsp) Reset() {
	*x = C2CMsgRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_logic_proto_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2CMsgRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2CMsgRsp) ProtoMessage() {}

func (x *C2CMsgRsp) ProtoReflect() protoreflect.Message {
	mi := &file_internal_logic_proto_logic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2CMsgRsp.ProtoReflect.Descriptor instead.
func (*C2CMsgRsp) Descriptor() ([]byte, []int) {
	return file_internal_logic_proto_logic_proto_rawDescGZIP(), []int{3}
}

func (x *C2CMsgRsp) GetMsgIds() []uint64 {
	if x != nil {
		return x.MsgIds
	}
	return nil
}

var File_internal_logic_proto_logic_proto protoreflect.FileDescriptor

var file_internal_logic_proto_logic_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x5c, 0x0a, 0x09, 0x43, 0x32, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x42, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x09, 0x43, 0x32, 0x43, 0x4d, 0x73, 0x67, 0x52,
	0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x73, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64, 0x78, 0x6a, 0x2f, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x2d, 0x77, 0x61, 0x67, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_logic_proto_logic_proto_rawDescOnce sync.Once
	file_internal_logic_proto_logic_proto_rawDescData = file_internal_logic_proto_logic_proto_rawDesc
)

func file_internal_logic_proto_logic_proto_rawDescGZIP() []byte {
	file_internal_logic_proto_logic_proto_rawDescOnce.Do(func() {
		file_internal_logic_proto_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_logic_proto_logic_proto_rawDescData)
	})
	return file_internal_logic_proto_logic_proto_rawDescData
}

var file_internal_logic_proto_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_logic_proto_logic_proto_goTypes = []interface{}{
	(*Error)(nil),     // 0: logic.Error
	(*C2CMsgReq)(nil), // 1: logic.C2CMsgReq
	(*Ack)(nil),       // 2: logic.Ack
	(*C2CMsgRsp)(nil), // 3: logic.C2CMsgRsp
	(*anypb.Any)(nil), // 4: google.protobuf.Any
}
var file_internal_logic_proto_logic_proto_depIdxs = []int32{
	4, // 0: logic.C2CMsgReq.data:type_name -> google.protobuf.Any
	0, // 1: logic.Ack.error:type_name -> logic.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_logic_proto_logic_proto_init() }
func file_internal_logic_proto_logic_proto_init() {
	if File_internal_logic_proto_logic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_logic_proto_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_internal_logic_proto_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CMsgReq); i {
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
		file_internal_logic_proto_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_internal_logic_proto_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2CMsgRsp); i {
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
			RawDescriptor: file_internal_logic_proto_logic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_logic_proto_logic_proto_goTypes,
		DependencyIndexes: file_internal_logic_proto_logic_proto_depIdxs,
		MessageInfos:      file_internal_logic_proto_logic_proto_msgTypes,
	}.Build()
	File_internal_logic_proto_logic_proto = out.File
	file_internal_logic_proto_logic_proto_rawDesc = nil
	file_internal_logic_proto_logic_proto_goTypes = nil
	file_internal_logic_proto_logic_proto_depIdxs = nil
}

package schedulerpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

var file_scheduler_proto_goTypes = []interface{}{(*RegisterRequest)(nil), (*RegisterResponse)(nil)}
var (
	file_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_proto_rawDescData = file_scheduler_proto_rawDesc
)
var file_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scheduler_proto_depIdxs = []int32{0, 1, 1, 0, 0, 0, 0}
var File_scheduler_proto protoreflect.FileDescriptor
var file_scheduler_proto_rawDesc = []byte{0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x22, 0x25, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x54, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x61, 0x70, 0x70, 0x70, 0x61, 0x74, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33}

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Status        string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Time          string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}
type x struct{}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (x *RegisterRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}
func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func init() {
	file_scheduler_proto_init()
}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{0}
}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{1}
}
func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}
func (*RegisterRequest) ProtoMessage() {
}
func (x *RegisterResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}
func (*RegisterResponse) ProtoMessage() {
}
func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func file_scheduler_proto_init() {
	if File_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
	out := protoimpl.TypeBuilder{File: protoimpl.DescBuilder{GoPackagePath: reflect.TypeOf(x{}).PkgPath(), RawDescriptor: file_scheduler_proto_rawDesc, NumEnums: 0, NumMessages: 2, NumExtensions: 0, NumServices: 1}, GoTypes: file_scheduler_proto_goTypes, DependencyIndexes: file_scheduler_proto_depIdxs, MessageInfos: file_scheduler_proto_msgTypes}.Build()
	File_scheduler_proto = out.File
	file_scheduler_proto_rawDesc = nil
	file_scheduler_proto_goTypes = nil
	file_scheduler_proto_depIdxs = nil
}
func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func file_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_proto_rawDescData)
	})
	return file_scheduler_proto_rawDescData
}
func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

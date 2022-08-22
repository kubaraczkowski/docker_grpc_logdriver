// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: logdriver.proto

package driver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogOptions) Reset() {
	*x = LogOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logdriver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogOptions) ProtoMessage() {}

func (x *LogOptions) ProtoReflect() protoreflect.Message {
	mi := &file_logdriver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogOptions.ProtoReflect.Descriptor instead.
func (*LogOptions) Descriptor() ([]byte, []int) {
	return file_logdriver_proto_rawDescGZIP(), []int{0}
}

type LogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string    `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Entry   *LogEntry `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *LogMessage) Reset() {
	*x = LogMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logdriver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMessage) ProtoMessage() {}

func (x *LogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_logdriver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMessage.ProtoReflect.Descriptor instead.
func (*LogMessage) Descriptor() ([]byte, []int) {
	return file_logdriver_proto_rawDescGZIP(), []int{1}
}

func (x *LogMessage) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *LogMessage) GetEntry() *LogEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ServicesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service []string `protobuf:"bytes,1,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *ServicesList) Reset() {
	*x = ServicesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logdriver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicesList) ProtoMessage() {}

func (x *ServicesList) ProtoReflect() protoreflect.Message {
	mi := &file_logdriver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicesList.ProtoReflect.Descriptor instead.
func (*ServicesList) Descriptor() ([]byte, []int) {
	return file_logdriver_proto_rawDescGZIP(), []int{2}
}

func (x *ServicesList) GetService() []string {
	if x != nil {
		return x.Service
	}
	return nil
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source             string                   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	TimeNano           int64                    `protobuf:"varint,2,opt,name=time_nano,json=timeNano,proto3" json:"time_nano,omitempty"`
	Line               []byte                   `protobuf:"bytes,3,opt,name=line,proto3" json:"line,omitempty"`
	Partial            bool                     `protobuf:"varint,4,opt,name=partial,proto3" json:"partial,omitempty"`
	PartialLogMetadata *PartialLogEntryMetadata `protobuf:"bytes,5,opt,name=partial_log_metadata,json=partialLogMetadata,proto3" json:"partial_log_metadata,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logdriver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_logdriver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_logdriver_proto_rawDescGZIP(), []int{3}
}

func (x *LogEntry) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *LogEntry) GetTimeNano() int64 {
	if x != nil {
		return x.TimeNano
	}
	return 0
}

func (x *LogEntry) GetLine() []byte {
	if x != nil {
		return x.Line
	}
	return nil
}

func (x *LogEntry) GetPartial() bool {
	if x != nil {
		return x.Partial
	}
	return false
}

func (x *LogEntry) GetPartialLogMetadata() *PartialLogEntryMetadata {
	if x != nil {
		return x.PartialLogMetadata
	}
	return nil
}

type PartialLogEntryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Last    bool   `protobuf:"varint,1,opt,name=last,proto3" json:"last,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ordinal int32  `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
}

func (x *PartialLogEntryMetadata) Reset() {
	*x = PartialLogEntryMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logdriver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialLogEntryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialLogEntryMetadata) ProtoMessage() {}

func (x *PartialLogEntryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_logdriver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialLogEntryMetadata.ProtoReflect.Descriptor instead.
func (*PartialLogEntryMetadata) Descriptor() ([]byte, []int) {
	return file_logdriver_proto_rawDescGZIP(), []int{4}
}

func (x *PartialLogEntryMetadata) GetLast() bool {
	if x != nil {
		return x.Last
	}
	return false
}

func (x *PartialLogEntryMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartialLogEntryMetadata) GetOrdinal() int32 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

var File_logdriver_proto protoreflect.FileDescriptor

var file_logdriver_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x4c, 0x6f, 0x67,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x51, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x29, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6c, 0x6f, 0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x12, 0x54, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x6f, 0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x12, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x4c,
	0x6f, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x17, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x6c, 0x32, 0x8e, 0x01, 0x0a, 0x10, 0x49, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x15, 0x2e, 0x6c, 0x6f, 0x67,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6c, 0x6f,
	0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x61, 0x72, 0x61, 0x63, 0x7a, 0x6b, 0x6f, 0x77, 0x73, 0x6b,
	0x69, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f,
	0x67, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logdriver_proto_rawDescOnce sync.Once
	file_logdriver_proto_rawDescData = file_logdriver_proto_rawDesc
)

func file_logdriver_proto_rawDescGZIP() []byte {
	file_logdriver_proto_rawDescOnce.Do(func() {
		file_logdriver_proto_rawDescData = protoimpl.X.CompressGZIP(file_logdriver_proto_rawDescData)
	})
	return file_logdriver_proto_rawDescData
}

var file_logdriver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_logdriver_proto_goTypes = []interface{}{
	(*LogOptions)(nil),              // 0: logdriver.LogOptions
	(*LogMessage)(nil),              // 1: logdriver.LogMessage
	(*ServicesList)(nil),            // 2: logdriver.ServicesList
	(*LogEntry)(nil),                // 3: logdriver.LogEntry
	(*PartialLogEntryMetadata)(nil), // 4: logdriver.PartialLogEntryMetadata
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_logdriver_proto_depIdxs = []int32{
	3, // 0: logdriver.LogMessage.entry:type_name -> logdriver.LogEntry
	4, // 1: logdriver.LogEntry.partial_log_metadata:type_name -> logdriver.PartialLogEntryMetadata
	0, // 2: logdriver.IDockerLogDriver.GetLogs:input_type -> logdriver.LogOptions
	5, // 3: logdriver.IDockerLogDriver.ListServices:input_type -> google.protobuf.Empty
	1, // 4: logdriver.IDockerLogDriver.GetLogs:output_type -> logdriver.LogMessage
	2, // 5: logdriver.IDockerLogDriver.ListServices:output_type -> logdriver.ServicesList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logdriver_proto_init() }
func file_logdriver_proto_init() {
	if File_logdriver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logdriver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogOptions); i {
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
		file_logdriver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMessage); i {
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
		file_logdriver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicesList); i {
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
		file_logdriver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_logdriver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialLogEntryMetadata); i {
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
			RawDescriptor: file_logdriver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logdriver_proto_goTypes,
		DependencyIndexes: file_logdriver_proto_depIdxs,
		MessageInfos:      file_logdriver_proto_msgTypes,
	}.Build()
	File_logdriver_proto = out.File
	file_logdriver_proto_rawDesc = nil
	file_logdriver_proto_goTypes = nil
	file_logdriver_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: agent.proto

package protobuf

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

type CollectCPUTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SqlDigest      string   `protobuf:"bytes,1,opt,name=sql_digest,json=sqlDigest,proto3" json:"sql_digest,omitempty"`
	PlanDigest     string   `protobuf:"bytes,2,opt,name=plan_digest,json=planDigest,proto3" json:"plan_digest,omitempty"`
	TimestampList  []uint64 `protobuf:"varint,10,rep,packed,name=timestamp_list,json=timestampList,proto3" json:"timestamp_list,omitempty"`     // timestamp in second
	CpuTimeMsList  []uint32 `protobuf:"varint,11,rep,packed,name=cpu_time_ms_list,json=cpuTimeMsList,proto3" json:"cpu_time_ms_list,omitempty"` // this can be greater than 1000 when counting concurrent running SQL queries
	NormalizedSql  string   `protobuf:"bytes,20,opt,name=normalized_sql,json=normalizedSql,proto3" json:"normalized_sql,omitempty"`             // SQL string with sensitive fields trimmed, potentially > 1M
	NormalizedPlan string   `protobuf:"bytes,21,opt,name=normalized_plan,json=normalizedPlan,proto3" json:"normalized_plan,omitempty"`          // potentially > 1M
	IsInternalSql  bool     `protobuf:"varint,22,opt,name=is_internal_sql,json=isInternalSql,proto3" json:"is_internal_sql,omitempty"`          // if true, then this sql and plan is internally generated by tidb itself, not user
}

func (x *CollectCPUTimeRequest) Reset() {
	*x = CollectCPUTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectCPUTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectCPUTimeRequest) ProtoMessage() {}

func (x *CollectCPUTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectCPUTimeRequest.ProtoReflect.Descriptor instead.
func (*CollectCPUTimeRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *CollectCPUTimeRequest) GetSqlDigest() string {
	if x != nil {
		return x.SqlDigest
	}
	return ""
}

func (x *CollectCPUTimeRequest) GetPlanDigest() string {
	if x != nil {
		return x.PlanDigest
	}
	return ""
}

func (x *CollectCPUTimeRequest) GetTimestampList() []uint64 {
	if x != nil {
		return x.TimestampList
	}
	return nil
}

func (x *CollectCPUTimeRequest) GetCpuTimeMsList() []uint32 {
	if x != nil {
		return x.CpuTimeMsList
	}
	return nil
}

func (x *CollectCPUTimeRequest) GetNormalizedSql() string {
	if x != nil {
		return x.NormalizedSql
	}
	return ""
}

func (x *CollectCPUTimeRequest) GetNormalizedPlan() string {
	if x != nil {
		return x.NormalizedPlan
	}
	return ""
}

func (x *CollectCPUTimeRequest) GetIsInternalSql() bool {
	if x != nil {
		return x.IsInternalSql
	}
	return false
}

type CollectCPUTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CollectCPUTimeResponse) Reset() {
	*x = CollectCPUTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectCPUTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectCPUTimeResponse) ProtoMessage() {}

func (x *CollectCPUTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectCPUTimeResponse.ProtoReflect.Descriptor instead.
func (*CollectCPUTimeResponse) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02,
	0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x50, 0x55, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x71, 0x6c, 0x5f, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x71, 0x6c,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x10, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x73, 0x71, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x71, 0x6c, 0x12, 0x27,
	0x0a, 0x0f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x71, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x71, 0x6c, 0x22,
	0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x50, 0x55, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x54, 0x0a, 0x0b, 0x54, 0x6f, 0x70,
	0x53, 0x51, 0x4c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x50, 0x55, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x43, 0x50, 0x55, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x50, 0x55, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agent_proto_goTypes = []interface{}{
	(*CollectCPUTimeRequest)(nil),  // 0: CollectCPUTimeRequest
	(*CollectCPUTimeResponse)(nil), // 1: CollectCPUTimeResponse
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: TopSQLAgent.CollectCPUTime:input_type -> CollectCPUTimeRequest
	1, // 1: TopSQLAgent.CollectCPUTime:output_type -> CollectCPUTimeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectCPUTimeRequest); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectCPUTimeResponse); i {
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
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}

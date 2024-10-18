// protoc --go_out=../grpc --go-grpc_out=../grpc node.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: node.proto

package node

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

type NewTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskIndex uint32 `protobuf:"varint,1,opt,name=task_index,json=taskIndex,proto3" json:"task_index,omitempty"`
	Task      *Task  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *NewTaskRequest) Reset() {
	*x = NewTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaskRequest) ProtoMessage() {}

func (x *NewTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaskRequest.ProtoReflect.Descriptor instead.
func (*NewTaskRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *NewTaskRequest) GetTaskIndex() uint32 {
	if x != nil {
		return x.TaskIndex
	}
	return 0
}

func (x *NewTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDetails               string `protobuf:"bytes,1,opt,name=task_details,json=taskDetails,proto3" json:"task_details,omitempty"`
	TaskCreatedBlock          uint32 `protobuf:"varint,2,opt,name=task_created_block,json=taskCreatedBlock,proto3" json:"task_created_block,omitempty"`
	QuorumNumbers             []byte `protobuf:"bytes,3,opt,name=quorum_numbers,json=quorumNumbers,proto3" json:"quorum_numbers,omitempty"`
	QuorumThresholdPercentage uint32 `protobuf:"varint,4,opt,name=quorum_threshold_percentage,json=quorumThresholdPercentage,proto3" json:"quorum_threshold_percentage,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetTaskDetails() string {
	if x != nil {
		return x.TaskDetails
	}
	return ""
}

func (x *Task) GetTaskCreatedBlock() uint32 {
	if x != nil {
		return x.TaskCreatedBlock
	}
	return 0
}

func (x *Task) GetQuorumNumbers() []byte {
	if x != nil {
		return x.QuorumNumbers
	}
	return nil
}

func (x *Task) GetQuorumThresholdPercentage() uint32 {
	if x != nil {
		return x.QuorumThresholdPercentage
	}
	return 0
}

type NewTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NewTaskResponse) Reset() {
	*x = NewTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaskResponse) ProtoMessage() {}

func (x *NewTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaskResponse.ProtoReflect.Descriptor instead.
func (*NewTaskResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *NewTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x22, 0xbe, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x74, 0x61, 0x73,
	0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x25, 0x0a,
	0x0e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x71, 0x75, 0x6f, 0x72, 0x75,
	0x6d, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x58, 0x0a, 0x15, 0x4d, 0x61, 0x6e, 0x75, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_node_proto_goTypes = []any{
	(*NewTaskRequest)(nil),  // 0: node.NewTaskRequest
	(*Task)(nil),            // 1: node.Task
	(*NewTaskResponse)(nil), // 2: node.NewTaskResponse
}
var file_node_proto_depIdxs = []int32{
	1, // 0: node.NewTaskRequest.task:type_name -> node.Task
	0, // 1: node.ManuscriptNodeService.ReceiveNewTask:input_type -> node.NewTaskRequest
	2, // 2: node.ManuscriptNodeService.ReceiveNewTask:output_type -> node.NewTaskResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NewTaskRequest); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Task); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NewTaskResponse); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}

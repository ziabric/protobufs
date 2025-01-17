// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.28.3
// source: tasks/tasks.proto

package __

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

type AddTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Desc          string                 `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	mi := &file_tasks_tasks_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_tasks_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *AddTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddTaskRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type AddTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Flag          bool                   `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	TaskId        string                 `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddTaskResponse) Reset() {
	*x = AddTaskResponse{}
	mi := &file_tasks_tasks_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskResponse) ProtoMessage() {}

func (x *AddTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_tasks_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskResponse.ProtoReflect.Descriptor instead.
func (*AddTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *AddTaskResponse) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

func (x *AddTaskResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type CompliteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompliteTaskRequest) Reset() {
	*x = CompliteTaskRequest{}
	mi := &file_tasks_tasks_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompliteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompliteTaskRequest) ProtoMessage() {}

func (x *CompliteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_tasks_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompliteTaskRequest.ProtoReflect.Descriptor instead.
func (*CompliteTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *CompliteTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type CompliteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompliteFlag  bool                   `protobuf:"varint,1,opt,name=complite_flag,json=compliteFlag,proto3" json:"complite_flag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompliteTaskResponse) Reset() {
	*x = CompliteTaskResponse{}
	mi := &file_tasks_tasks_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompliteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompliteTaskResponse) ProtoMessage() {}

func (x *CompliteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_tasks_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompliteTaskResponse.ProtoReflect.Descriptor instead.
func (*CompliteTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *CompliteTaskResponse) GetCompliteFlag() bool {
	if x != nil {
		return x.CompliteFlag
	}
	return false
}

var File_tasks_tasks_proto protoreflect.FileDescriptor

var file_tasks_tasks_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x53, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0x3e, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22,
	0x3b, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x32, 0x89, 0x01, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tasks_tasks_proto_rawDescOnce sync.Once
	file_tasks_tasks_proto_rawDescData = file_tasks_tasks_proto_rawDesc
)

func file_tasks_tasks_proto_rawDescGZIP() []byte {
	file_tasks_tasks_proto_rawDescOnce.Do(func() {
		file_tasks_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_tasks_tasks_proto_rawDescData)
	})
	return file_tasks_tasks_proto_rawDescData
}

var file_tasks_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tasks_tasks_proto_goTypes = []any{
	(*AddTaskRequest)(nil),       // 0: tasks.AddTaskRequest
	(*AddTaskResponse)(nil),      // 1: tasks.AddTaskResponse
	(*CompliteTaskRequest)(nil),  // 2: tasks.CompliteTaskRequest
	(*CompliteTaskResponse)(nil), // 3: tasks.CompliteTaskResponse
}
var file_tasks_tasks_proto_depIdxs = []int32{
	2, // 0: tasks.Task.CompliteTask:input_type -> tasks.CompliteTaskRequest
	0, // 1: tasks.Task.AddTask:input_type -> tasks.AddTaskRequest
	3, // 2: tasks.Task.CompliteTask:output_type -> tasks.CompliteTaskResponse
	1, // 3: tasks.Task.AddTask:output_type -> tasks.AddTaskResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tasks_tasks_proto_init() }
func file_tasks_tasks_proto_init() {
	if File_tasks_tasks_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tasks_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasks_tasks_proto_goTypes,
		DependencyIndexes: file_tasks_tasks_proto_depIdxs,
		MessageInfos:      file_tasks_tasks_proto_msgTypes,
	}.Build()
	File_tasks_tasks_proto = out.File
	file_tasks_tasks_proto_rawDesc = nil
	file_tasks_tasks_proto_goTypes = nil
	file_tasks_tasks_proto_depIdxs = nil
}

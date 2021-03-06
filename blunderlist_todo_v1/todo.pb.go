// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package blunderlist_todo_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TaskListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskListRequest) Reset()         { *m = TaskListRequest{} }
func (m *TaskListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskListRequest) ProtoMessage()    {}
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *TaskListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListRequest.Unmarshal(m, b)
}
func (m *TaskListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListRequest.Marshal(b, m, deterministic)
}
func (m *TaskListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListRequest.Merge(m, src)
}
func (m *TaskListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskListRequest.Size(m)
}
func (m *TaskListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListRequest proto.InternalMessageInfo

type TaskQuery struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskQuery) Reset()         { *m = TaskQuery{} }
func (m *TaskQuery) String() string { return proto.CompactTextString(m) }
func (*TaskQuery) ProtoMessage()    {}
func (*TaskQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *TaskQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskQuery.Unmarshal(m, b)
}
func (m *TaskQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskQuery.Marshal(b, m, deterministic)
}
func (m *TaskQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskQuery.Merge(m, src)
}
func (m *TaskQuery) XXX_Size() int {
	return xxx_messageInfo_TaskQuery.Size(m)
}
func (m *TaskQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskQuery.DiscardUnknown(m)
}

var xxx_messageInfo_TaskQuery proto.InternalMessageInfo

func (m *TaskQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TaskCreateRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Completed            bool     `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCreateRequest) Reset()         { *m = TaskCreateRequest{} }
func (m *TaskCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TaskCreateRequest) ProtoMessage()    {}
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *TaskCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateRequest.Unmarshal(m, b)
}
func (m *TaskCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateRequest.Marshal(b, m, deterministic)
}
func (m *TaskCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateRequest.Merge(m, src)
}
func (m *TaskCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TaskCreateRequest.Size(m)
}
func (m *TaskCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateRequest proto.InternalMessageInfo

func (m *TaskCreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskCreateRequest) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type TaskUpdateRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Completed            bool     `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskUpdateRequest) Reset()         { *m = TaskUpdateRequest{} }
func (m *TaskUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*TaskUpdateRequest) ProtoMessage()    {}
func (*TaskUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *TaskUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskUpdateRequest.Unmarshal(m, b)
}
func (m *TaskUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskUpdateRequest.Marshal(b, m, deterministic)
}
func (m *TaskUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskUpdateRequest.Merge(m, src)
}
func (m *TaskUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_TaskUpdateRequest.Size(m)
}
func (m *TaskUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskUpdateRequest proto.InternalMessageInfo

func (m *TaskUpdateRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskUpdateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskUpdateRequest) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type TaskResponse struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Completed            bool                 `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskResponse) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *TaskResponse) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskListRequest)(nil), "todo.blunderlist.v1.TaskListRequest")
	proto.RegisterType((*TaskQuery)(nil), "todo.blunderlist.v1.TaskQuery")
	proto.RegisterType((*TaskCreateRequest)(nil), "todo.blunderlist.v1.TaskCreateRequest")
	proto.RegisterType((*TaskUpdateRequest)(nil), "todo.blunderlist.v1.TaskUpdateRequest")
	proto.RegisterType((*TaskResponse)(nil), "todo.blunderlist.v1.TaskResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0xd3, 0x82, 0x3f, 0xbd, 0x18, 0x0d, 0xa3, 0x11, 0x52, 0x8c, 0x62, 0x63, 0x0c, 0xab,
	0x22, 0xb8, 0x74, 0x63, 0x10, 0xc3, 0x86, 0x05, 0x36, 0x28, 0x89, 0x1b, 0x02, 0xf4, 0x4a, 0x1a,
	0x5a, 0xa6, 0x76, 0xa6, 0x24, 0x3c, 0x85, 0x2f, 0xe1, 0x83, 0x9a, 0x99, 0x91, 0x00, 0xd5, 0xc6,
	0x2e, 0xdc, 0xb5, 0x77, 0xce, 0xf9, 0xda, 0x39, 0x67, 0x06, 0x80, 0x53, 0x97, 0xda, 0x61, 0x44,
	0x39, 0x25, 0xc7, 0xf2, 0x79, 0xec, 0xc7, 0x73, 0x17, 0x23, 0xdf, 0x63, 0xdc, 0x5e, 0x34, 0xcc,
	0xca, 0x94, 0xd2, 0xa9, 0x8f, 0x75, 0x29, 0x19, 0xc7, 0x6f, 0x75, 0x0c, 0x42, 0xbe, 0x54, 0x0e,
	0xf3, 0x22, 0xb9, 0xc8, 0xbd, 0x00, 0x19, 0x1f, 0x05, 0xa1, 0x12, 0x58, 0x45, 0x38, 0xea, 0x8f,
	0xd8, 0xac, 0xeb, 0x31, 0xee, 0xe0, 0x7b, 0x8c, 0x8c, 0x5b, 0x15, 0x30, 0xc4, 0xe8, 0x29, 0xc6,
	0x68, 0x49, 0x0e, 0x41, 0xf7, 0xdc, 0xb2, 0x56, 0xd5, 0x6a, 0x79, 0x47, 0xf7, 0x5c, 0xab, 0x03,
	0x45, 0xb1, 0xf8, 0x10, 0xe1, 0x88, 0xe3, 0xb7, 0x83, 0x9c, 0xc0, 0x0e, 0xf7, 0xb8, 0x8f, 0x52,
	0x67, 0x38, 0xea, 0x85, 0x9c, 0x81, 0x31, 0xa1, 0x41, 0xe8, 0x23, 0x47, 0xb7, 0xac, 0x57, 0xb5,
	0xda, 0xbe, 0xb3, 0x1e, 0x58, 0x03, 0x05, 0x7a, 0x0e, 0xdd, 0x0d, 0x50, 0xe2, 0x6b, 0x6b, 0xb0,
	0x9e, 0x0a, 0xce, 0x25, 0xc1, 0x1f, 0x1a, 0x1c, 0x08, 0xb2, 0x83, 0x2c, 0xa4, 0x73, 0x86, 0xff,
	0x01, 0x25, 0x77, 0x50, 0x98, 0xc8, 0x2d, 0x0f, 0x45, 0x80, 0xe5, 0x7c, 0x55, 0xab, 0x15, 0x9a,
	0xa6, 0xad, 0xd2, 0xb5, 0x57, 0xe9, 0xda, 0xfd, 0x55, 0xba, 0x0e, 0x28, 0xb9, 0x18, 0x34, 0x3f,
	0x73, 0x90, 0xef, 0x53, 0x97, 0x92, 0x17, 0x30, 0x44, 0xd0, 0xe2, 0xef, 0x18, 0xb9, 0xb2, 0x7f,
	0x69, 0xd3, 0x4e, 0x94, 0x61, 0x5e, 0xa6, 0xaa, 0x56, 0xfb, 0xbb, 0xd1, 0x48, 0x17, 0xf6, 0x3a,
	0x28, 0xb1, 0xe4, 0x3c, 0x55, 0x2f, 0xfb, 0xcc, 0xc0, 0x23, 0x03, 0x00, 0x55, 0xaf, 0x04, 0x5e,
	0xa7, 0x1a, 0xb6, 0xce, 0x40, 0x46, 0xb0, 0xaa, 0xfb, 0x0f, 0xf0, 0xd6, 0x99, 0xc8, 0x02, 0x6e,
	0x03, 0xb4, 0x51, 0x14, 0x95, 0x29, 0x82, 0xd3, 0x1f, 0xb5, 0x3d, 0x8a, 0x1b, 0xd3, 0xba, 0x87,
	0xd2, 0x84, 0x06, 0x5b, 0x3e, 0x09, 0x5a, 0x34, 0x5a, 0x20, 0xea, 0xeb, 0x09, 0x39, 0xeb, 0x69,
	0xaf, 0xa5, 0x70, 0x36, 0xad, 0x6f, 0xc8, 0x86, 0x42, 0x36, 0x5c, 0x34, 0xc6, 0xbb, 0x92, 0x78,
	0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x90, 0x36, 0xdd, 0xa0, 0xb4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	ListTasks(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (Todo_ListTasksClient, error)
	GetTask(ctx context.Context, in *TaskQuery, opts ...grpc.CallOption) (*TaskResponse, error)
	CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	UpdateTask(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	DeleteTask(ctx context.Context, in *TaskQuery, opts ...grpc.CallOption) (*empty.Empty, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) ListTasks(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (Todo_ListTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Todo_serviceDesc.Streams[0], "/todo.blunderlist.v1.Todo/ListTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoListTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todo_ListTasksClient interface {
	Recv() (*TaskResponse, error)
	grpc.ClientStream
}

type todoListTasksClient struct {
	grpc.ClientStream
}

func (x *todoListTasksClient) Recv() (*TaskResponse, error) {
	m := new(TaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoClient) GetTask(ctx context.Context, in *TaskQuery, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/todo.blunderlist.v1.Todo/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/todo.blunderlist.v1.Todo/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) UpdateTask(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/todo.blunderlist.v1.Todo/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) DeleteTask(ctx context.Context, in *TaskQuery, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/todo.blunderlist.v1.Todo/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	ListTasks(*TaskListRequest, Todo_ListTasksServer) error
	GetTask(context.Context, *TaskQuery) (*TaskResponse, error)
	CreateTask(context.Context, *TaskCreateRequest) (*TaskResponse, error)
	UpdateTask(context.Context, *TaskUpdateRequest) (*TaskResponse, error)
	DeleteTask(context.Context, *TaskQuery) (*empty.Empty, error)
}

// UnimplementedTodoServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (*UnimplementedTodoServer) ListTasks(req *TaskListRequest, srv Todo_ListTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (*UnimplementedTodoServer) GetTask(ctx context.Context, req *TaskQuery) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedTodoServer) CreateTask(ctx context.Context, req *TaskCreateRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedTodoServer) UpdateTask(ctx context.Context, req *TaskUpdateRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (*UnimplementedTodoServer) DeleteTask(ctx context.Context, req *TaskQuery) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServer).ListTasks(m, &todoListTasksServer{stream})
}

type Todo_ListTasksServer interface {
	Send(*TaskResponse) error
	grpc.ServerStream
}

type todoListTasksServer struct {
	grpc.ServerStream
}

func (x *todoListTasksServer) Send(m *TaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Todo_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.blunderlist.v1.Todo/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTask(ctx, req.(*TaskQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.blunderlist.v1.Todo/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).CreateTask(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.blunderlist.v1.Todo/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).UpdateTask(ctx, req.(*TaskUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.blunderlist.v1.Todo/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).DeleteTask(ctx, req.(*TaskQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.blunderlist.v1.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _Todo_GetTask_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Todo_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Todo_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Todo_DeleteTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTasks",
			Handler:       _Todo_ListTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}

package grpc

//go:generate protoc -I ./proto --go_out=plugins=grpc:. todo.proto

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"
	"google.golang.org/grpc"

	"github.com/tomasbasham/grpc-service-go/option"
	transport "github.com/tomasbasham/grpc-service-go/transport/grpc"
)

// Client is a wrapper type around a client connection to an RPC server and a
// concrete instance of a `TodoClient` type.
type Client struct {
	conn   *grpc.ClientConn
	client todopb.TodoClient
}

// NewClientWithTarget creates a new client connection to an RPC server at a
// specific endpoint.
func NewClientWithTarget(ctx context.Context, target string) (*Client, error) {
	return NewClient(ctx, option.WithEndpoint(target))
}

// NewClient creates a new client connectino to an RPC server accepting
// multiple connection options.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialInsecure(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		client: todopb.NewTodoClient(conn),
	}, nil
}

// Conn returns a client connection to an RPC server.
func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

// Close closes the client connection to an RPC server. If the connection has
// not been made then this method does nothing.
func (c *Client) Close() {
	if c.conn == nil {
		return
	}

	c.conn.Close()
}

// ListTasks delegates to the client, returning a type able to fetch tasks from
// a stream.
func (c *Client) ListTasks(ctx context.Context, in *todopb.TaskListRequest, opts ...grpc.CallOption) (todopb.Todo_ListTasksClient, error) {
	return c.client.ListTasks(ctx, in, opts...)
}

// GetTask delegates to the client, returning a task with the given id. If the
// task could not be found then an error is returned.
func (c *Client) GetTask(ctx context.Context, in *todopb.TaskQuery, opts ...grpc.CallOption) (*todopb.TaskResponse, error) {
	return c.client.GetTask(ctx, in, opts...)
}

// CreateTask delegates to the client, returning a persisted task with the
// given information. If the task could not be persisted then an error is
// returned.
func (c *Client) CreateTask(ctx context.Context, in *todopb.TaskCreateRequest, opts ...grpc.CallOption) (*todopb.TaskResponse, error) {
	return c.client.CreateTask(ctx, in, opts...)
}

// UpdateTask delegates to the client, returning a persisted task with the
// given information. If the task could not be found, or was unable to be
// persisted then an error is returned.
func (c *Client) UpdateTask(ctx context.Context, in *todopb.TaskUpdateRequest, opts ...grpc.CallOption) (*todopb.TaskResponse, error) {
	return c.client.UpdateTask(ctx, in, opts...)
}

// DeleteTask delegates to the client, removing the task with the given id. If
// the task could not be found, or it could not be deleted then as error is
// returned.
func (c *Client) DeleteTask(ctx context.Context, in *todopb.TaskQuery, opts ...grpc.CallOption) (*empty.Empty, error) {
	return c.client.DeleteTask(ctx, in, opts...)
}

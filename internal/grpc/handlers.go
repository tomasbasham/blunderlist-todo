package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	//details "google.golang.org/genproto/googleapis/rpc/errordetails"

	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// ListTasks streams all tasks from the store.
func (s *Server) ListTasks(_ *todopb.TaskListRequest, stream todopb.Todo_ListTasksServer) error {
	for _, task := range s.store.GetTasks() {
		if err := stream.Send(task); err != nil {
			return status.Error(codes.Unknown, "Unknown streaming error")
		}
	}

	return nil
}

// GetTask returns a single task from the store.
func (s *Server) GetTask(ctx context.Context, in *todopb.TaskQuery) (*todopb.TaskResponse, error) {
	task, err := s.store.GetTask(uint(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task not found with id %d", in.Id)
	}

	return task, nil
}

// CreateTask creates and returns a new task.
func (s *Server) CreateTask(ctx context.Context, in *todopb.TaskCreateRequest) (*todopb.TaskResponse, error) {
	task, err := s.store.CreateTask(in)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "failed to create task")
	}

	return task, nil
}

// UpdateTask updates and returns an existing task.
func (s *Server) UpdateTask(ctx context.Context, in *todopb.TaskUpdateRequest) (*todopb.TaskResponse, error) {
	task, err := s.store.UpdateTask(in)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task not found with id %d", in.Id)
	}

	return task, nil
}

// DeleteTask deletes a task from the store.
func (s *Server) DeleteTask(ctx context.Context, in *todopb.TaskQuery) (*empty.Empty, error) {
	if err := s.store.DeleteTask(uint(in.Id)); err != nil {
		return &empty.Empty{}, status.Errorf(codes.NotFound, "task not found with if %d", in.Id)
	}

	return &empty.Empty{}, nil
}

// Check is used by clients to know when the service is ready.
func (s *Server) Check(context.Context, *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	if err := s.store.IsAvailable(); err != nil {
		return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_NOT_SERVING}, nil
	}

	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

// Watch is used by clients to receive updates when the service status changes.
// Here it has no meaningful implementation just to satisfy the interface.
func (s *Server) Watch(*healthpb.HealthCheckRequest, healthpb.Health_WatchServer) error {
	return nil
}

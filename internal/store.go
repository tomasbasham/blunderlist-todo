package internal

import todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"

// Repository provides the storage requirements through which to access tasks.
type Repository interface {
	GetTasks() []*todopb.TaskResponse
	GetTask(uint) (*todopb.TaskResponse, error)
	CreateTask(*todopb.TaskCreateRequest) (*todopb.TaskResponse, error)
	UpdateTask(*todopb.TaskUpdateRequest) (*todopb.TaskResponse, error)
	DeleteTask(uint) error

	// Used to check of the storge media is ready to use. If not this method
	// should return an error.
	IsAvailable() error
}

// Store proxies requests to the underlying storage implementation.
type Store struct {
	repo Repository
}

// NewStore returns a Store containing a repository.
func NewStore(repo Repository) *Store {
	return &Store{repo: repo}
}

// GetTasks proxies to the repo.
func (s *Store) GetTasks() []*todopb.TaskResponse {
	return s.repo.GetTasks()
}

// GetTask proxies to the repo.
func (s *Store) GetTask(id uint) (*todopb.TaskResponse, error) {
	return s.repo.GetTask(id)
}

// CreateTask proxies to the repo.
func (s *Store) CreateTask(task *todopb.TaskCreateRequest) (*todopb.TaskResponse, error) {
	return s.repo.CreateTask(task)
}

// UpdateTask proxies to the repo.
func (s *Store) UpdateTask(task *todopb.TaskUpdateRequest) (*todopb.TaskResponse, error) {
	return s.repo.UpdateTask(task)
}

// DeleteTask proxies to the repo.
func (s *Store) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

// IsAvailable proxies to the repo.
func (s *Store) IsAvailable() error {
	return s.repo.IsAvailable()
}

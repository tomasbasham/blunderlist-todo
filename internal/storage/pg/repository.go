package pg

import (
	"database/sql"
	"time"

	// Initialise the pq package.
	_ "github.com/lib/pq"

	pb "github.com/golang/protobuf/ptypes"
	todopb "github.com/tomasbasham/blunderlist-todo/blunderlist_todo_v1"
)

// Storage implements the Repository interface, with an active connection to a
// PostgreSQL instance.
type Storage struct {
	db *sql.DB
}

// NewStorage returns a Storage that implements the Repository interface.
func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

// GetTasks makes a SQL query to retreive all tasks from the database.
func (s *Storage) GetTasks() []*todopb.TaskResponse {
	rows, err := s.db.Query("SELECT * FROM tasks ORDER BY created_at")
	if err != nil {
		return []*todopb.TaskResponse{}
	}

	var tasks []*todopb.TaskResponse

	for rows.Next() {
		task, err := scan(rows)
		if err != nil {
			continue
		}

		tasks = append(tasks, task)
	}

	return tasks
}

// GetTask makes a SQL query to retreive a single task from the database.
func (s *Storage) GetTask(id uint) (*todopb.TaskResponse, error) {
	row := s.db.QueryRow("SELECT * FROM tasks WHERE id=$1", id)

	task, err := scan(row)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// CreateTask makes a SQL query to insert a new task into the database.
func (s *Storage) CreateTask(task *todopb.TaskCreateRequest) (*todopb.TaskResponse, error) {
	var lastInsertID uint

	err := s.db.QueryRow("INSERT INTO tasks(title, completed, created_at) VALUES($1, $2, $3) RETURNING id", task.Title, task.Completed, time.Now()).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}

	return s.GetTask(lastInsertID)
}

// UpdateTask makes a SQL query to update a task already present in the
// database.
func (s *Storage) UpdateTask(task *todopb.TaskUpdateRequest) (*todopb.TaskResponse, error) {
	var lastUpdateID uint

	err := s.db.QueryRow("UPDATE tasks SET title=$1, completed=$2 WHERE id=$3 RETURNING id", task.Title, task.Completed, uint(task.Id)).Scan(&lastUpdateID)
	if err != nil {
		return nil, err
	}

	return s.GetTask(lastUpdateID)
}

// DeleteTask makes a SQL query to remove a task from the database.
func (s *Storage) DeleteTask(id uint) error {
	if _, err := s.db.Exec("DELETE FROM tasks WHERE id=$1", id); err != nil {
		return err
	}

	return nil
}

// IsAvailable enusues there is an active connection to the database,
// establishing a connection if necessary.
func (s *Storage) IsAvailable() error {
	if err := s.db.Ping(); err != nil {
		return err
	}

	// This is necessary because once a database connection is initiallly
	// established subsequent calls to the Ping method return success even if the
	// database goes down.
	//
	// https://github.com/lib/pq/issues/533
	if _, err := s.db.Exec("SELECT 1"); err != nil {
		return err
	}

	return nil
}

type resultScanner interface {
	Scan(dest ...interface{}) error
}

func scan(r resultScanner) (*todopb.TaskResponse, error) {
	var id uint
	var title string
	var completed bool
	var createdAt time.Time

	if err := r.Scan(&id, &title, &completed, &createdAt); err != nil {
		return nil, err
	}

	timestamp, err := pb.TimestampProto(createdAt)
	if err != nil {
		return nil, err
	}

	return &todopb.TaskResponse{
		Id:         uint64(id),
		Title:      title,
		Completed:  completed,
		CreateTime: timestamp,
	}, nil
}

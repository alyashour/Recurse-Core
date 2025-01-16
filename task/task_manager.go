package task

import (
	"fmt"
	"time"
)

type JSONTaskManager struct {
	fileManager FileManager
	Path        string
	Tasks       []Task
}

type FileManager interface {
}

func (tm JSONTaskManager) CreateTask(title, description, assignee string, dueDate time.Time) (*Task, error) {
	return nil, fmt.Errorf("not implemented error")
}

func (tm JSONTaskManager) GetTask(id int) (*Task, error) {
	return nil, fmt.Errorf("not implemented error")
}

func (tm JSONTaskManager) UpdateTask(id int, updates Task) error {
	return fmt.Errorf("not implemented error")
}

func (tm JSONTaskManager) DeleteTask(id int) error {
	return fmt.Errorf("not implemented error")
}

func (tm JSONTaskManager) ListTasks() ([]Task, error) {
	return nil, fmt.Errorf("not implemented error")
}

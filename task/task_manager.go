package task

import (
	"fmt"
	"time"
)

type TaskManager struct {
	Path  string
	Tasks []Task
}

func (tm TaskManager) CreateTask(title, description, assignee string, dueDate time.Time) (*Task, error) {
	return nil, fmt.Errorf("not implemented error")
}

func (tm TaskManager) GetTask(id int) (*Task, error) {
	return nil, fmt.Errorf("not implemented error")
}

func (tm TaskManager) UpdateTask(id int, updates Task) error {
	return fmt.Errorf("not implemented error")
}

func (tm TaskManager) DeleteTask(id int) error {
	return fmt.Errorf("not implemented error")
}

func (tm TaskManager) ListTasks() ([]Task, error) {
	return nil, fmt.Errorf("not implemented error")
}

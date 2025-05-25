package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks struct {
	TasksList []Task `json:"tasks_list"`
}

const (
	FilePath = "/home/logan/workspace/golang-projects/tasks-cli/assets/store.json"
)

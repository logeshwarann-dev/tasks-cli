package models

import "time"

type Task struct {
	Id          int
	Description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

type Tasks struct {
	TasksList []Task
}

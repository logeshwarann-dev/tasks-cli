package models

var (
	Command     string
	Status      string
	Id          int
	Description string
)

var CmdInputSlice []string

const (
	AddCommand       = "add"
	UpdateCommand    = "update"
	ListCommand      = "list"
	DeleteCommand    = "delete"
	MarkInProgress   = "mark-in-progress"
	MarkDone         = "mark-done"
	StatusDone       = "done"
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
)

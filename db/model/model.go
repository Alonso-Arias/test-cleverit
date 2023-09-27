package model

import "time"

type Task struct {
	Id          int32
	Title       string
	Description string
	DueDate     time.Time
	State       string
}

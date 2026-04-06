package simplesql

import "time"

type TaskModel struct {
	Id          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

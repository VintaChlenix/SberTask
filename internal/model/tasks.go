package model

import (
	"time"
)

type Task struct {
	Header      *string    `json:"header"`
	Description *string    `json:"description"`
	Date        *time.Time `json:"date"`
	IsDone      *bool      `json:"is_done"`
}

type Pagination struct {
	IsDone         *bool      `json:"is_done"`
	StartTimestamp *time.Time `json:"start_timestamp"`
	EndTimestamp   *time.Time `json:"end_timestamp"`
	Page           *int       `json:"page"`
	Rows           *int       `json:"rows"`
}

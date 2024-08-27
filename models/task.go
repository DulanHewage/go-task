package models

import (
	"database/sql"
	"time"
)
type Task struct {
    ID          int
    Title       string
    Description string
    Completed   sql.NullBool
    CreatedAt   time.Time
}
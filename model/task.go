package models

import "time"

//Task model structure...
type Task struct {
	ID      int
	name    string
	details string
	date    time.Time
	done    bool
}

package model

import (
	"time"
)

// User is Entity of user
type User struct {
	ID      int64
	Name    string `sql:"size:255"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
}

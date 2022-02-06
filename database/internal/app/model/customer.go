package model

import "time"

type Customer struct {
	ID         int64
	Email      string
	FirstName  string
	LastName   string
	DeleteTime *time.Time
}

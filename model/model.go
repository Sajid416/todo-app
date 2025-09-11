package model

import "time"

type User struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Created_At time.Time `json:"created_at"`
}

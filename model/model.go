package model

import "time"

type User struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Done       bool      `json:"done"`
	Created_At time.Time `json:"created_at"`
}

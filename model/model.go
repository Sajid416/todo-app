package model

import "time"

type User struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Created_At time.Time `json:"created_at"`
}

type UserInfo struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

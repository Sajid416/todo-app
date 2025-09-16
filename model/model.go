package model

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgUrl      string `json:"imgUrl"`
}

type UserInfo struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

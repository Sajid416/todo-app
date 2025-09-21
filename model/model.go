package model

type Product struct {
	Id          int    `json:"id" db:"id" `
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ImgUrl      string `json:"img_url" db:"img_url"`
}

type UserInfo struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

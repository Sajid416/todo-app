package user

import (
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	UserDB *sqlx.DB
}

func NewHandler(userDB *sqlx.DB) *Handler {
	return &Handler{UserDB: userDB}
}

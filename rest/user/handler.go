package user

type Handler struct {
	//UserDB *sqlx.DB
}

// userDB *sqlx.DB
func NewHandler() *Handler {
	return &Handler{}
}

//return &Handler{UserDB: userDB}

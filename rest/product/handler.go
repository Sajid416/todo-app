package product

type Handler struct {
}

// TodoDB *sqlx.DB
func NewHandler() *Handler {
	return &Handler{}
}

//return &Handler{TodoDB: TodoDB}

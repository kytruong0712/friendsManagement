package controller

// UserController: User Controller
type UserController struct {
	repo Repository
}

// NewUserController: Create new User Controller
func NewUserController(r Repository) *UserController {
	return &UserController{
		repo: r,
	}
}

package commonhelp

type Users struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Mobile   string `json:"mobile" validate:"required,min=10"`
	Password string `json:"password" validate:"required,min=6"`
}

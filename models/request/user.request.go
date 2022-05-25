package request

type UserCreateRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Nama    string `json:"nama"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

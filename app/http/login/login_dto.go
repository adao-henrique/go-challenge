package login

type LoginRequest struct {
	Cpf    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

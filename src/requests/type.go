package requests

type TypeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

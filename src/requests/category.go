package requests

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

package requests

type SourceRequest struct {
	Url string `json:"url" validate:"required,url"`
}

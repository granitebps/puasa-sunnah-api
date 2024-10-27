package requests

type TypeRequest struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description"`
	BackgroundColor string `json:"background_color" validate:"required,hexcolor"`
	TextColor       string `json:"text_color" validate:"required,hexcolor"`
}

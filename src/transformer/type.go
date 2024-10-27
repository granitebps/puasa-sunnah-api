package transformer

type TypeTransformer struct {
	ID              uint   `json:"id" example:"1"`
	Name            string `json:"name" example:"lorem ipsum"`
	Description     string `json:"description" example:"lorem ipsum"`
	BackgroundColor string `json:"background_color" example:"#FFFFFF"`
	TextColor       string `json:"text_color" example:"#FFFFFF"`
}

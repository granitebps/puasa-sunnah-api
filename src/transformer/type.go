package transformer

type TypeTransformer struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"lorem ipsum"`
	Description string `json:"description" example:"lorem ipsum"`
}

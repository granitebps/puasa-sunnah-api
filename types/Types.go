package types

type Type struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"lorem"`
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit."`
}

package transformer

type FastingTransformer struct {
	ID         uint   `json:"id" example:"1"`
	CategoryID uint   `json:"category_id" example:"1"`
	TypeID     uint   `json:"type_id" example:"1"`
	Date       string `json:"date" example:"2020-01-01"`
	Year       uint   `json:"year" example:"2020"`
	Month      uint   `json:"month" example:"1"`
	Day        uint   `json:"day" example:"1"`
	HumanDate  string `json:"human_date"`

	Category CategoryTransformer `json:"category"`
	Type     TypeTransformer     `json:"type"`
}

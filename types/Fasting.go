package types

type Fasting struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	TypeID     uint   `json:"type_id"`
	Date       string `json:"date"`
	Year       uint   `json:"year"`
	Month      uint   `json:"month"`
	Day        uint   `json:"day"`
}

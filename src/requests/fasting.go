package requests

type FastingRequest struct {
	TypeID string `query:"type_id"`
	Day    string `query:"day"`
	Month  string `query:"month"`
	Year   string `query:"year"`
}

type FastingCreateUpdateRequest struct {
	CategoryID uint   `json:"category_id" validate:"required"`
	TypeID     uint   `json:"type_id" validate:"required"`
	Date       string `json:"date" validate:"required,datetime=2006-01-02"`
	Year       uint   `json:"year" validate:"required,isYear"`
	Month      uint   `json:"month" validate:"required,min=1,max=12"`
	Day        uint   `json:"day" validate:"required,min=1,max=31"`
}

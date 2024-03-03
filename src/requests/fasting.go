package requests

type FastingRequest struct {
	TypeID string `query:"type_id"`
	Day    string `query:"day"`
	Month  string `query:"month"`
	Year   string `query:"year"`
}

package entities

type Event struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	LocationID    string `json:"location_id"`
	ImageUrl      string `json:"image_url"`
	Date          string `json:"date"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	OpenGatesTime string `json:"open_gates_time"`
	Description   string `json:"description"`
	Type          string `json:"type"`
}

package entities

type Event struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Artist        Artist   `json:"artist"`
	Location      Location `json:"location"`
	ImageUrl      string   `json:"image_url"`
	StartTime     string   `json:"start_time"`
	EndTime       string   `json:"end_time"`
	OpenGatesTime string   `json:"open_gates_time"`
	Description   string   `json:"description"`
	Type          string   `json:"type"`
}

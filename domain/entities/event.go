package entities

type Event struct {
	ID            string
	Name          string
	Artist        Artist
	Location      Location
	ImageUrl      string
	StartTime     string
	EndTime       string
	OpenGatesTime string
	Description   string
	Type          string
}

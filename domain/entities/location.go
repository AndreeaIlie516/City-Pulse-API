package entities

type Location struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	CityID  string `json:"city_id"`
	Address string `json:"address"`
}

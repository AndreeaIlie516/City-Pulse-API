package entities

type Location struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	City    City   `json:"city"`
	Address string `json:"address"`
}

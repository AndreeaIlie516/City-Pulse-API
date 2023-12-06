package entities

type Artist struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	IsBand      bool    `json:"is_band"`
	BandMembers string  `json:"band_members"`
	Description string  `json:"description"`
	Genres      []Genre `json:"genres"`
}

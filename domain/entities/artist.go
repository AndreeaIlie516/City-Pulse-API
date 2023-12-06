package entities

type Artist struct {
	ID          string
	Name        string
	IsBand      bool
	BandMembers string
	Description string
	Genres      []Genre
}

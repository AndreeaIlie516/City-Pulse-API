package entities

type ArtistGenre struct {
	ID       string `json:"id"`
	ArtistID string `json:"artist_id"`
	GenreID  string `json:"genre_id"`
	Period   string `json:"period"`
}

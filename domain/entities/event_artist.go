package entities

type EventArtist struct {
	ID        string `json:"id"`
	EventID   string `json:"event_id"`
	ArtistID  string `json:"artist_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type EventArtistDetail struct {
	Association EventArtist
	Event       EventDetails
	Artist      Artist
}

type EventWithArtists struct {
	Event   EventDetails
	Artists []Artist
}

type ArtistWithEvents struct {
	Artist Artist
	Events []EventDetails
}

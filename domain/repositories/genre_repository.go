package repositories

import "City-Pulse-API/domain/entities"

type GenreRepository interface {
	AllGenres() ([]entities.Genre, error)
	AllGenreIDs() []string
	GenreByID(id string) (*entities.Genre, error)
	CreateGenre(genre entities.Genre) (entities.Genre, error)
	UpdateGenre(id string, genre entities.Genre) (entities.Genre, error)
	DeleteGenre(id string) (entities.Genre, error)
}

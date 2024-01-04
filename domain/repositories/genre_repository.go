package repositories

import "City-Pulse-API/domain/entities"

type GenreRepository interface {
	AllGenres() ([]entities.Genre, error)
	AllGenreIDs() ([]uint, error)
	GenreByID(id uint) (*entities.Genre, error)
	CreateGenre(genre entities.Genre) (entities.Genre, error)
	UpdateGenre(id uint, updatedGenre entities.Genre) (entities.Genre, error)
	DeleteGenre(id uint) (entities.Genre, error)
}

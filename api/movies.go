package api

import "gorm.io/gorm"

type MovieGateway interface {
	Add(m *Movie) error
	FindByTitle(title string) ([]Movie, error)
	FindByID(id uint) (Movie, error)
	FindByIMDB(s string) (Movie, error)
}

type Movie struct {
	gorm.Model `json:"gorm_model"`
	Title      string `json:"title" json:"title,omitempty"`
	Year       string `json:"year" json:"year,omitempty"`
	Genre      string `json:"genre" json:"genre,omitempty"`
	Director   string `json:"director" json:"director,omitempty"`
	Writer     string `json:"writer" json:"writer,omitempty"`
	ImdbID     string `json:"imdb_id" json:"imdb_id,omitempty" gorm:"uniqueIndex"`
}

type MovieStorage struct{}

func (m *MovieStorage) Add(movie *Movie) error {
	panic("implement me")
}

func (m *MovieStorage) FindByTitle(title string) ([]Movie, error) {
	panic("implement me")
}

func (m *MovieStorage) FindByID(id uint) (Movie, error) {
	panic("implement me")
}

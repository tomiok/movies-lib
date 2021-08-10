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
	Title      string `json:"title"`
	Year       string `json:"year"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Writer     string `json:"writer"`
	ImdbID     string `json:"imdb_id" gorm:"uniqueIndex"`
	Reviews []Review  `json:"reviews"`
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

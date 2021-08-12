package api

import "gorm.io/gorm"

type MovieGateway interface {
	Add(m *Movie) error
	FindByTitle(title string) ([]Movie, error)
	FindByID(id uint) (Movie, error)
	FindByIMDB(s string) (Movie, error)
	GetReviews(id uint) []Review
}

type Movie struct {
	gorm.Model
	Title    string   `json:"title"`
	Year     string   `json:"year"`
	Genre    string   `json:"genre"`
	Director string   `json:"director"`
	Writer   string   `json:"writer"`
	ImdbID   string   `json:"imdb_id" gorm:"uniqueIndex"`
	Reviews  []Review `json:"reviews"`
}

type MovieStorage struct {
	db *gorm.DB
}

func (m *MovieStorage) Add(movie *Movie) error {
	return m.db.Create(movie).Error
}

func (m *MovieStorage) FindByTitle(title string) ([]Movie, error) {
	panic("implement me")
}

func (m *MovieStorage) FindByID(id uint) (Movie, error) {
	var movie Movie
	err := m.db.First(&movie, "id = ?", id).Error

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}

func (m *MovieStorage) GetReviews(id uint) []Review {
	return nil	
}
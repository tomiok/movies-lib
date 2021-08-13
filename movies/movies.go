package movies

import (
	"errors"
	"github.com/tomiok/movies-lib/reviews"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MovieGateway interface {
	Add(m *Movie) error
	FindByTitle(title string) ([]Movie, error)
	FindByID(id uint) (Movie, error)
	FindByIMDB(s string) (Movie, error)
	GetReviews(id uint) []reviews.Review
}

type Movie struct {
	gorm.Model
	Title    string   `json:"title"`
	Year     string   `json:"year"`
	Genre    string   `json:"genre"`
	Director string   `json:"director"`
	Writer   string   `json:"writer"`
	ImdbID   string           `json:"imdb_id" gorm:"uniqueIndex"`
	Reviews  []reviews.Review `json:"reviews"`
}

type MovieStorage struct {
	db *gorm.DB
}

func (m *MovieStorage) Add(movie *Movie) error {
	return m.db.Create(movie).Error
}

func (m *MovieStorage) FindByTitle(title string) ([]Movie, error) {
	panic("implement me with some full text")
}

func (m *MovieStorage) FindByID(id uint) (Movie, error) {
	var movie Movie
	err := m.db.First(&movie, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.S().Warnf("cannot find movie with id %d", id)
		return Movie{}, nil
	}

	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}

func (m *MovieStorage) GetReviews(id uint) []reviews.Review {
	return nil
}

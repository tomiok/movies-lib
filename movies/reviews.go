package movies

import (
	"gorm.io/gorm"
)

type ReviewGateway interface {
	Add(comment, movie string) error
	ReadFromMovie(movie string) ([]Review, error)
}

type Review struct {
	gorm.Model
	Comment string `json:"comment" gorm:"column:comment"`
	MovieID uint
}

type ReviewStorage struct {
	MovieGtw MovieGateway
	db       *gorm.DB
}

func (m *ReviewStorage) Add(comment string, imdbID string) error {
	movie, err := m.MovieGtw.FindByIMDB(imdbID)

	if err != nil {
		return err
	}

	return m.db.Create(Review{
		Comment: comment,
		MovieID: movie.ID,
	}).Error
}

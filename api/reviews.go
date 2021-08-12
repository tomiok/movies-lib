package api

import "gorm.io/gorm"

type ReviewGateway interface {
	Add(comment, movie string) error
	ReadFromMovie(movie string) ([]Review, error)
}

type Review struct {
	gorm.Model
	Comment string `json:"comment" gorm:"column:comment"`
	MovieID uint
}

type MoviesReview struct {
	MovieGtw MovieGateway
	db       *gorm.DB
}

func (m *MoviesReview) Add(comment string, imdbID string) error {
	movie, err := m.MovieGtw.FindByIMDB(imdbID)

	if err != nil {
		return err
	}

	return m.db.Create(Review{
		Comment: comment,
		MovieID: movie.ID,
	}).Error
}

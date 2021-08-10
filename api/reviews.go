package api

import "gorm.io/gorm"

type ReviewGateway interface {
	Add(comment, movie string) error
	ReadFromMovie(movie string) ([]Review, error)
}

type Review struct {
	gorm.Model
	Comment string `gorm:"column:comment"`
	MovieID uint
}

type MoviesReview struct {
	MovieGtw MovieGateway
}

func (m *MoviesReview) Add(comment string, imdbID string) error {
	db := Get()

	movie, err := m.MovieGtw.FindByIMDB(imdbID)

	if err != nil {
		return err
	}

	return db.Create(Review{
		Comment: comment,
		MovieID: movie.ID,
	}).Error
}

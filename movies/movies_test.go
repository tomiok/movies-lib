package movies

import (
	"github.com/tomiok/movies-lib/storage"
	"testing"
)

func Droppable() []interface{} {
	return []interface{}{&Movie{}, &Review{}}
}
func TestGORMV2(t *testing.T) {
	DB := storage.Get(true)

	storage.Migrate(true, Droppable()...)
	movie := &Movie{
		Title:    "Blade Runner",
		Year:     "1982",
		Genre:    "sci fi",
		Director: "Ridley Scott",
		Writer:   "Phillip Dick",
	}
	DB.Create(movie)

	dest := movie
	DB.First(&dest, "id = ?", movie.ID)

}

package movies_lib

type Search interface {
	ByTitle(title string) ([]*MovieResponse, error)
}

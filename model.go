package movies_lib

// OpenAPIResponse is the response mapped from the movie database API. When search with 's' param.
type OpenAPIResponse struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	IMDB   string `json:"imdbID"`
	Poster string `json:"poster"`
}

// OpenAPIMovie is the response mapped from the movie database with only a title. Param 't.'
type OpenAPIMovie struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
}

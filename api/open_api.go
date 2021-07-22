package api

import (
	"bytes"
	"github.com/buger/jsonparser"
	"net/http"
	"net/url"
	"time"
)

var _ Search = (*openAPIDB)(nil)

type Search interface {
	ByQueryTitle(title string) ([]OpenAPIResponse, error)
	ByTitle(title string) (*OpenAPIMovie, error)
}

type openAPIDB struct {
	client http.Client
	key    string
}

func newOA() *openAPIDB {
	client := http.Client{
		Timeout: time.Millisecond * 3000,
	}

	return &openAPIDB{
		client: client,
	}
}

func (o *openAPIDB) ByQueryTitle(search string) ([]OpenAPIResponse, error) {
	req := buildRequest(search, "s")

	res, err := o.client.Do(&req)

	if err != nil {
		return nil, err
	}

	body := res.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)

	var result []OpenAPIResponse

	v, _, _, err := jsonparser.Get(buf.Bytes(), "Search")

	_, err = jsonparser.ArrayEach(v, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		var movie OpenAPIResponse
		movie.Title, err = jsonparser.GetString(value, "Title")
		movie.Year, err = jsonparser.GetString(value, "Year")
		movie.IMDB, err = jsonparser.GetString(value, "imdbID")
		movie.Poster, err = jsonparser.GetString(value, "Poster")
		result = append(result, movie)
	})

	return result, nil
}

func (o *openAPIDB) ByTitle(title string) (*OpenAPIMovie, error) {
	req := buildRequest(title, "t")

	res, err := o.client.Do(&req)

	if err != nil {
		return nil, err
	}

	body := res.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)
	value := buf.Bytes()
	var movie OpenAPIMovie
	movie.Title, _ = jsonparser.GetString(value, "Title")
	movie.Year, _ = jsonparser.GetString(value, "Year")
	movie.Genre, _ = jsonparser.GetString(value, "Genre")
	movie.Director, _ = jsonparser.GetString(value, "Director")
	movie.Writer, _ = jsonparser.GetString(value, "Writer")

	return &movie, nil
}

const baseURL = "https://www.omdbapi.com/"

func buildRequest(title, query string) http.Request {
	u, _ := url.Parse(baseURL)
	req := http.Request{
		URL:    u,
		Method: http.MethodGet,
	}

	q := req.URL.Query()

	q.Add("apikey", "4ecb0111")
	q.Add(query, title)
	req.URL.RawQuery = q.Encode()
	return req
}

package movies_lib

import (
	"bytes"
	"fmt"
	"github.com/buger/jsonparser"
	"net/http"
	"os"
	"time"
)

type Search interface {
	ByTitle(title string) (*MovieResponse, error)
}

type openAPIDB struct {
	client http.Client
	key    string
}

func newOA() *openAPIDB {
	client := http.Client{
		Timeout: time.Millisecond * 900,
	}

	return &openAPIDB{
		client: client,
		key:    os.Getenv("API_KEY"),
	}
}

func (o *openAPIDB) Search(title string) (*MovieResponse, error) {
	req := buildRequest(o.key, title)

	res, err := o.client.Do(&req)

	if err != nil {
		return nil, err
	}

	body := res.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)

	var result = new(MovieResponse)

	result.Title, err = jsonparser.GetString(buf.Bytes(), "Title")
	result.Year, err = jsonparser.GetString(buf.Bytes(), "Year")
	result.Genre, err = jsonparser.GetString(buf.Bytes(), "Genre")
	result.Director, err = jsonparser.GetString(buf.Bytes(), "Director")
	result.Writer, err = jsonparser.GetString(buf.Bytes(), "Writer")

	return result, nil
}

const baseURL = "http://www.omdbapi.com/?apikey=%s&t=%s"

func buildRequest(key, title string) http.Request {
	host := fmt.Sprintf(baseURL, key, title)
	return http.Request{Host: host}
}

package movies_lib

import (
	"bytes"
	"fmt"
	"github.com/buger/jsonparser"
	"net/http"
	"net/url"
	"os"
	"time"
)

var _ Search = (*openAPIDB)(nil)

type Search interface {
	ByTitle(title string) ([]OpenAPIResponse, error)
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
		key:    os.Getenv("API_KEY"),
	}
}

func (o *openAPIDB)ByTitle(title string) ([]OpenAPIResponse, error) {
	req := buildRequest(o.key, title)

	res, err := o.client.Do(&req)

	if err != nil {
		return nil, err
	}

	body := res.Body
	defer body.Close()
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)

	var result []OpenAPIResponse
	fmt.Println(string(buf.Bytes()))

	v, _, _, err := jsonparser.Get(buf.Bytes(), "Search")

	_, err = jsonparser.ArrayEach(v, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		var movie OpenAPIResponse
		movie.Title, err = jsonparser.GetString(value, "Title")
		movie.Year, err = jsonparser.GetString(value, "Year")
		movie.IMDB, err = jsonparser.GetString(value, "imdbID")
		movie.Poster, err = jsonparser.GetString(value, "Poster")

	})

	return result, nil
}

const baseURL = "https://www.omdbapi.com/"

func buildRequest(key, title string) http.Request {
	u, _ := url.Parse(baseURL)
	req := http.Request{
		URL:    u,
		Method: http.MethodGet,
	}

	q := req.URL.Query()

	q.Add("apikey", key)
	q.Add("s", title)
	req.URL.RawQuery = q.Encode()
	return req
}

package utils

import (
	"io"
	"net/http"
)

type HttpClient struct {
	Method  string
	BaseUrl string
	Header  map[string]string
	Body    io.Reader
}

func (c HttpClient) MakeGet() (*http.Response, error) {

	request, err := http.NewRequest(c.Method, c.BaseUrl, c.Body)

	if err != nil {
		return nil, err
	}

	for key, value := range c.Header {
		request.Header.Add(key, value)
	}

	return http.DefaultClient.Do(request)
}

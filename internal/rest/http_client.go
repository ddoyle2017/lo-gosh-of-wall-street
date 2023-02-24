package rest

import "net/http"

type HttpClient interface {
	Get(url string) (*http.Response, error)
	Do(request *http.Request) (*http.Response, error)
}

func NewHttpClient() *http.Client {
	return &http.Client{}
}

package rest

import "net/http"

type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

func NewHttpClient() *http.Client {
	return &http.Client{}
}

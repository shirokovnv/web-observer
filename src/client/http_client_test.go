package client_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/shirokovnv/web-observer/src/client"
)

var (
	httpMock client.HTTPClientMock
	site     client.Site
)

func setUpHttp() {
	httpMock = client.HTTPClientMock{}
	client.Client = &httpMock
	site = client.Site{URL: "https://foo.com", Interval: "10m", Code: 200}
}

func TestSuccessfulResponseCode(t *testing.T) {
	setUpHttp()
	httpMock.DoFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Body:       io.NopCloser(strings.NewReader("Successful response")),
			StatusCode: 200,
		}, nil
	}

	code, _ := site.CheckUrlStatus()

	if code != int(site.Code) {
		t.Fatalf("want %v, got %v", int(site.Code), code)
	}
}

func TestFailedResponseCode(t *testing.T) {
	setUpHttp()
	httpMock.DoFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			Body:       io.NopCloser(strings.NewReader("Not found")),
			StatusCode: 404,
		}, nil
	}

	code, _ := site.CheckUrlStatus()

	if code == int(site.Code) {
		t.Fatal("Status codes must not be equal")
	}
}

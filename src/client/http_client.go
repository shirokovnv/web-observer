package client

import (
	"context"
	"net/http"
	"time"
)

const DefaultHttpReqTimeout = 30 * time.Second

type Site struct {
	URL      string `yaml:"url"`
	Interval string `yaml:"interval"`
	Code     uint16 `yaml:"code"`
}

type SiteConfig struct {
	Sites []Site `yaml:"sites"`
}

func (site Site) CheckUrlStatus() (int, error) {
	ctx, cncl := context.WithTimeout(context.Background(), DefaultHttpReqTimeout)
	defer cncl()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, site.URL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := Client.Do(req)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

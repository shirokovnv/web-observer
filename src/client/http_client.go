package client

import (
	"net/http"
)

type Site struct {
	URL      string `yaml:"url"`
	Interval string `yaml:"interval"`
	Code     uint16 `yaml:"code"`
}

type SiteConfig struct {
	Sites []Site `yaml:"sites"`
}

func (site Site) CheckUrlStatus() (int, error) {
	req, err := http.NewRequest(http.MethodGet, site.URL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := Client.Do(req)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

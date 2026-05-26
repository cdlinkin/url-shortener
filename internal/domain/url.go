package domain

import (
	"fmt"
	"strings"
	"time"
)

type URL struct {
	ID          int
	OriginalUrl string
	ShortCode   string
	Clicks      int

	CreatedAt time.Time
}

type CreateUrlRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
	Code     string `json:"code"`
}

type StatsResponse struct {
	URL       string    `json:"url"`
	ShortCode string    `json:"short_code"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *URL) ValidateURL() error {
	if u.OriginalUrl == "" {
		return fmt.Errorf("url is empty")
	}

	if !strings.HasPrefix(u.OriginalUrl, "http://") &&
		!strings.HasPrefix(u.OriginalUrl, "https://") {
		return fmt.Errorf("url start with 'http://' or 'https://'")
	}
	return nil
}

func (u *CreateUrlRequest) ValidateURLReq() error {
	if u.URL == "" {
		return fmt.Errorf("url is empty")
	}

	if !strings.HasPrefix(u.URL, "http://") &&
		!strings.HasPrefix(u.URL, "https://") {
		return fmt.Errorf("url start with 'http://' or 'https://'")
	}
	return nil
}

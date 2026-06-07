package domain

import (
	"fmt"
	"strings"
	"time"
)

type URL struct {
	ID          int    `db:"id"`
	OriginalUrl string `db:"original_url"`
	ShortCode   string `db:"short_code"`
	Clicks      int    `db:"clicks"`

	CreatedAt time.Time `db:"created_at"`
}

type CreateUrlRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
	Code     string `json:"code"`
}

type StatsResponse struct {
	URL       string    `json:"url" db:"original_url"`
	ShortCode string    `json:"short_code" db:"short_code"`
	Clicks    int       `json:"clicks" db:"clicks"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
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

package domain

import "time"

type UrlOrigin struct {
	ID        int
	URL       string
	ShortCode string
	Clicks    int

	CreatedAt time.Time
}

type CreateUrlRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
	Code     string `json:"code"`
}

// type StatsResponse struct {} // для статистики

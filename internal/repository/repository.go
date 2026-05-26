package repository

import "github.com/cdlinkin/url-shortener/internal/domain"

type Repository interface {
	CreateURLShort(url domain.URL) (domain.URL, error)
	GetByCode(code string) (domain.URL, error)
	GetCodeStats(code string) (domain.StatsResponse, error)
	Delete(code string) error
}

package service

import (
	"fmt"
	"math/rand"

	"github.com/cdlinkin/url-shortener/internal/domain"
	"github.com/cdlinkin/url-shortener/internal/repository"
)

type UrlService struct {
	repo repository.Repository
}

func NewUrlService(repo repository.Repository) *UrlService {
	return &UrlService{
		repo: repo,
	}
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *UrlService) CreateURLShort(req domain.CreateUrlRequest) (domain.URL, error) {
	if err := req.ValidateURLReq(); err != nil {
		return domain.URL{}, fmt.Errorf("service error: %w", err)
	}
}

func (s *UrlService) GetByCode(code string) (domain.Url, error) {
	return s.repo.GetByCode(code)
}

func (s *UrlService) GetCodeStats(code string) (domain.StatsResponse, error) {
	return s.repo.GetCodeStats(code)
}

func (s *UrlService) Delete(code string) error {
	return s.repo.Delete(code)
}

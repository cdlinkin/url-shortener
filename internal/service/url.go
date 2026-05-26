package service

import (
	"fmt"
	"math/rand"
	"time"

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
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func (s *UrlService) CreateURLShort(req domain.CreateUrlRequest) (domain.URL, error) {
	if err := req.ValidateURLReq(); err != nil {
		return domain.URL{}, fmt.Errorf("service error: %w", err)
	}
	shortCode := generateShortCode()
	urlModel := domain.URL{
		OriginalUrl: req.URL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
	}
	urlDomain, err := s.repo.CreateURLShort(urlModel)
	if err != nil {
		return domain.URL{}, fmt.Errorf("service error: %w", err)
	}
	return urlDomain, nil
}

func (s *UrlService) GetByCode(code string) (domain.URL, error) {
	urlDomain, err := s.repo.GetByCode(code)
	if err != nil {
		return domain.URL{}, fmt.Errorf("service error: %w", err)
	}
	return urlDomain, nil
}

func (s *UrlService) GetCodeStats(code string) (domain.StatsResponse, error) {
	urlStatsDomain, err := s.repo.GetCodeStats(code)
	if err != nil {
		return domain.StatsResponse{}, fmt.Errorf("service error: %w", err)
	}
	return urlStatsDomain, nil
}

func (s *UrlService) Delete(code string) error {
	return s.repo.Delete(code)
}

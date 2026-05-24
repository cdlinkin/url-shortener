package repository

import (
	"fmt"

	"github.com/cdlinkin/url-shortener/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p *PostgresRepository) CreateURLShort(url domain.URL) (domain.URL, error) {
	query := `
	INSERT INTO urls (original_url, short_code, clicks, created_at)
	VALUES ($1, $2, $3, $4) 
	RETURNING id, original_url, short_code, clicks, created_at
	`
	if err := p.db.QueryRowx(query,
		url.OriginalUrl,
		url.ShortCode,
		url.Clicks,
		url.CreatedAt,
	).StructScan(&url); err != nil {
		return domain.URL{}, fmt.Errorf("failed to create url_short: %w", err)
	}

	return url, nil
}

func (p *PostgresRepository) GetByCode(code string) (domain.URL, error) {
	query := `
	UPDATE urls SET clicks = clicks + 1
	WHERE short_code = $1
	RETURNING original_url
	`

	urlDomain := domain.URL{}
	err := p.db.QueryRowx(query, code).StructScan(&urlDomain)
	if err != nil {
		return domain.URL{}, fmt.Errorf("failed to get clicks: %w", err)
	}

	return urlDomain, nil
}

func (p *PostgresRepository) GetCodeStats(code string) (domain.StatsResponse, error) {
	query := `
	SELECT id, original_url, short_code, clicks, created_at
	FROM urls
	WHERE short_code = $1
	`
	stats := domain.StatsResponse{}
	err := p.db.QueryRowx(query, code).StructScan(&stats)
	if err != nil {
		return domain.StatsResponse{}, fmt.Errorf("failed to get stats_url: %w", err)
	}

	return stats, nil
}

func (p *PostgresRepository) Delete(code string) error {
	query := `DELETE FROM urls WHERE short_code = $1`

	res, err := p.db.Exec(query, code)
	if err != nil {
		return fmt.Errorf("failed to delete url in database")
	}
	check, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to rows affected in database")
	}
	if check == 0 {
		return fmt.Errorf("url not found")
	}
	return nil
}

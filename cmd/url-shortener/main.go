package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cdlinkin/url-shortener/internal/cache"
	"github.com/cdlinkin/url-shortener/internal/handler"
	core_middleware "github.com/cdlinkin/url-shortener/internal/middleware"
	"github.com/cdlinkin/url-shortener/internal/repository"
	"github.com/cdlinkin/url-shortener/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}

	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	db, err := sqlx.Connect("postgres", source)
	if err != nil {
		log.Fatal("error in connect to postgres")
	}

	redis := cache.NewRedisCache(os.Getenv("REDIS_ADDR"))
	urlRepo := repository.NewPostgresRepository(db)
	urlService := service.NewUrlService(urlRepo, redis)
	urlHandler := handler.NewUrlHandler(urlService)
	authHandler := handler.NewAuthHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("GET /{code}", urlHandler.GetByCode)

	mux.HandleFunc("POST /shorten", withAuth(urlHandler.CreateURLShort))
	mux.HandleFunc("GET /{code}/stats", withAuth(urlHandler.GetCodeStats))
	mux.HandleFunc("DELETE /{code}", withAuth(urlHandler.Delete))

	if err := http.ListenAndServe(":8080", core_middleware.Logger(mux)); err != nil {
		log.Fatal(err)
	}
}

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		core_middleware.Authorization(next).ServeHTTP(w, r)
	}
}

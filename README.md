# URL Shortener API

REST API for shortening URLs. Built in Go to practice clean architecture and work with PostgreSQL, Redis, and JWT auth without third-party routers just standard `net/http`.

## Stack

- **Go** — standard `net/http` with native path parameters
- **PostgreSQL** — stores URLs and click counts
- **Redis** — caches short code → original URL for fast redirects
- **JWT** — protects write endpoints
- **Docker** — PostgreSQL and Redis via docker-compose

## Structure

```
internal/
├── domain/      # types and validation
├── handler/     # HTTP layer
├── service/     # business logic, short code generation
├── repository/  # PostgreSQL queries
├── middleware/  # JWT auth, logger
├── cache/       # Redis wrapper
└── auth/        # token generation and validation
```

## Run

```bash
git clone https://github.com/YOUR_USERNAME/url-shortener
cd url-shortener
```

Create `.env`

Start database and cache:

```bash
docker-compose up -d
go run cmd/url-shortener/main.go
```

## API

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | /auth/login | — | get JWT token |
| POST | /shorten | yes | create short URL |
| GET | /{code} | — | redirect to original URL |
| GET | /{code}/stats | yes | click count and info |
| DELETE | /{code} | yes | delete short URL |

### Example

```bash
# login
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'

# shorten
curl -X POST http://localhost:3000/shorten \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'

# redirect
curl http://localhost:3000/aB3xK9

# stats
curl http://localhost:3000/aB3xK9/stats \
  -H "Authorization: Bearer TOKEN"
```
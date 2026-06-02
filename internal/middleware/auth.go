package core_middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	authorization "github.com/cdlinkin/url-shortener/internal/auth"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		autoHeader := r.Header.Get("Authorization")
		if autoHeader == "" {
			writeError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenStr := strings.TrimPrefix(autoHeader, "Bearer ")

		_, err := authorization.ValidateToken(tokenStr)
		if err != nil {
			writeError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

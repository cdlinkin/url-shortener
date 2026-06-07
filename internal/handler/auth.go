package handler

import (
	"encoding/json"
	"net/http"

	authorization "github.com/cdlinkin/url-shortener/internal/auth"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Username != "admin" || req.Password != "admin" {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := authorization.GenerateToken(1)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"token": token})
}

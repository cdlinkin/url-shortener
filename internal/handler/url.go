package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cdlinkin/url-shortener/internal/domain"
	"github.com/cdlinkin/url-shortener/internal/service"
)

type UrlHandler struct {
	service *service.UrlService
}

func NewUrlHandler(service *service.UrlService) *UrlHandler {
	return &UrlHandler{service: service}
}

func (h *UrlHandler) CreateURLShort(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateUrlRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "error decoder")
		return
	}

	urlDomain, err := h.service.CreateURLShort(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid create url short in handler")
		return
	}

	response := domain.URLResponse{
		ShortURL: "http://localhost:8080/" + urlDomain.ShortCode,
		Code:     urlDomain.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	writeJSON(w, http.StatusCreated, response)
}

func (h *UrlHandler) GetByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	ctx := context.Background()
	urlDomain, err := h.service.GetByCode(ctx, code)
	if err != nil {
		writeError(w, http.StatusNotFound, "url not found")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	http.Redirect(w, r, urlDomain.OriginalUrl, http.StatusFound)
}

func (h *UrlHandler) GetCodeStats(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	statsDomain, err := h.service.GetCodeStats(code)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid get stats in handler")
		return

	}

	response := domain.StatsResponse{
		URL:       statsDomain.URL,
		ShortCode: statsDomain.ShortCode,
		Clicks:    statsDomain.Clicks,
	}

	w.Header().Set("Content-Type", "application/json")
	writeJSON(w, http.StatusOK, response)
}

func (h *UrlHandler) Delete(w http.ResponseWriter, r *http.Request) {
	reqCode := strings.TrimPrefix(r.URL.Path, "/code/")

	if err := h.service.Delete(reqCode); err != nil {
		writeError(w, http.StatusBadRequest, "invalid delete in handler")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

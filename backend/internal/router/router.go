package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/samuelt37/GoStarter/internal/handler"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", handler.Health)
	
	return r
}

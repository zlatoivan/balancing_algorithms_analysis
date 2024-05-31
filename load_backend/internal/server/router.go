package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) createRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
		middleware.RequestID,
		//middleware.Logger,
	)

	r.Get("/", s.Duration)

	return r
}

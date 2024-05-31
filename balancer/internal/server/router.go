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

	r.Get("/*", LoadStatic)
	r.Get("/", MainPage)

	r.Get("/balancer", s.Balancer)
	r.Get("/reload", s.Reload)
	r.Get("/metrics", s.Metrics)

	return r
}

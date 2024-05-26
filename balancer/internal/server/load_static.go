package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

// LoadStatic https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
func LoadStatic(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	rctx := chi.RouteContext(r.Context())
	pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir(workDir)))
	fs.ServeHTTP(w, r)
}

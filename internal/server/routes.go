package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.MainGame)

	// Serve static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(filesDir))))

	return r
}

func (s *Server) MainGame(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("assets/html", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

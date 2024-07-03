package server

import (
	"cookie-clicker/entity"
	"encoding/json"
	"io"
	"log"
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

	r.Post("/register", s.Register)
	r.Post("/login", s.Login)
	r.Post("/clicked", s.Clicked)

	// Serve static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(filesDir))))

	return r
}

func (s *Server) MainGame(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("assets/html", "index.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("usernameInput")
	password := r.FormValue("passwordInput")

	filePath := "data/player-score.json"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	var players []entity.Player
	if err := json.Unmarshal(byteValue, &players); err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}

	for _, account := range players {
		if account.Username == username && account.Password == password {
			data := entity.Player{
				Username: account.Username,
				Score:    account.Score,
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}

func (s *Server) Clicked(w http.ResponseWriter, r *http.Request) {

}

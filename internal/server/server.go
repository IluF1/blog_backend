package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"guthub.com/server/internal/middleware"
	"guthub.com/server/internal/service"
)

type server struct {
	addr   string
	router *mux.Router
}

func New(addr string) *server {
	s := &server{
		addr:   addr,
		router: mux.NewRouter(),
	}
	s.SetupRouter()
	return s
}

func (s *server) Run() error {
	return http.ListenAndServe(s.addr, s.router)
}

func (s *server) SetupRouter() {
	s.router.HandleFunc("/api/posts", service.GetPosts).Methods("GET")
	// s.router.HandleFunc("/api/post/{id}", service.GetPostById).Methods("POST")
	s.router.Handle("/api/user", middleware.Headers(http.HandlerFunc(service.GetUser))).Methods("POST")
	
}

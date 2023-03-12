package http

import (
	"net/http"
	"os"

	tangent "github.com/dfsantos-source/Tangent"
	"github.com/go-chi/chi/v5"
)

// server struct to hold instances needed
// wraps all HTTP functionality
type Server struct {
	server *http.Server
	router *chi.Mux

	MapboxUtil tangent.Util
	YelpUtil   tangent.Util

	UserService     tangent.UserService
	LocationService tangent.LocationService
}

// creates an instance of a server
// using Go http library and go-chi library router
func CreateServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: chi.NewRouter(),
	}

	s.MapboxUtil = *tangent.CreateUtil(os.Getenv("MAP_BOXTOKEN"))
	s.YelpUtil = *tangent.CreateUtil(os.Getenv("MAP_BOXTOKEN"))

	s.registerUserRoutes(s.router)

	return s
}

func (s *Server) RunServer() error {
	err := http.ListenAndServe(":3000", s.router)
	if err != nil {
		return err
	}
	return nil
}

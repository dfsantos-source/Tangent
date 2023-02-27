package http

import (
	"fmt"
	"net/http"

	tangent "github.com/dfsantos-source/Tangent"
	"github.com/gin-gonic/gin/render"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerUserRoutes(r *chi.Mux) {
	r.Get("/users/:id", s.getUsers)
	r.Get("/users", s.getUsers)
	r.Post("/users", s.createUser)
	r.Delete("/users/:id", s.deleteUser)
}

func (s *Server) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.UserService.Users()
	if err != nil {
		return
	}
	if users == nil {
		fmt.Printf("UserService.Users() returns nil")
		return
	}
	render.WriteJSON(w, users)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	user := &tangent.User{
		Name:     "Bob Marley",
		Email:    "bobmarley@umass.edu",
		Password: "123",
	}
	err := s.UserService.CreateUser(user)
	if err != nil {
		panic("error")
	}
	w.Write([]byte("Hello World!"))
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	panic(err)
}

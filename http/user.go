package http

import (
	"net/http"
	"strconv"

	tangent "github.com/dfsantos-source/Tangent"
	"github.com/gin-gonic/gin/render"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerUserRoutes(r *chi.Mux) {
	r.Get("/users/{id}", s.getUser)
	r.Get("/users", s.getUsers)
	r.Post("/users", s.createUser)
	r.Delete("/users/{id}", s.deleteUser)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	user, err := s.UserService.User(intId)
	if err != nil {
		render.WriteJSON(w, err)
		return
	}
	if user == nil {
		render.WriteJSON(w, []byte("User was nil"))
		return
	}
	render.WriteJSON(w, user)
}

func (s *Server) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.UserService.Users()
	if err != nil {
		return
	}
	if users == nil {
		w.Write([]byte("Users were nil"))
		return
	}
	render.WriteJSON(w, users)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	// TODO: endpoint parsing instead of hard-coding here
	user := &tangent.User{
		Name:     "Bob Marley",
		Email:    "bobmarley@umass.edu",
		Password: "123",
	}
	id, err := s.UserService.CreateUser(user)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	user.ID = id
	user.Password = ""
	render.WriteJSON(w, user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	err := s.UserService.DeleteUser(1)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	render.WriteJSON(w, nil)
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	panic(err)
}

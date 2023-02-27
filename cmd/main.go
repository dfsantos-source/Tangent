package main

import (
	"github.com/dfsantos-source/Tangent/http"
	"github.com/dfsantos-source/Tangent/postgres"
)

type Main struct {
	HTTPServer *http.Server
	DB         *postgres.DB
}

func CreateMain() *Main {
	return &Main{
		HTTPServer: http.CreateServer(),
		DB:         postgres.CreateDB(),
	}
}

func (m *Main) Run() {
	m.DB.OpenDB()
	m.DB.InitDB()

	m.HTTPServer.UserService = postgres.CreateUserServiceDB(m.DB)
	m.HTTPServer.RunServer()
}

func main() {
	m := CreateMain()
	m.Run()
}

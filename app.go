package main

import (
	"database/sql"
	
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Initializer(user, password, database string) {}

func (a *App) Runner(addr string) {}
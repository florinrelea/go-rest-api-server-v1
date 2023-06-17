package main

import (
	"database/sql"
	"log"

	"example.com/server"
	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	DB, err := sql.Open("sqlite3", "./practice.db")

	if err != nil {
		log.Fatal(err.Error())
	}

	server := server.ServerApp{
		Port:   ":9003",
		Router: rtr,
		DB:     DB,
	}

	server.InitRoutes()
	server.Run()
}

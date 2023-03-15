package main

import (
	"github.com/Tassil0/bazos-watcher/internal/database"
	"github.com/Tassil0/bazos-watcher/internal/routes"
	"log"
	"net/http"
)

//go:generate sqlboiler --wipe psql

func main() {

	database.Init()
	log.Fatal(http.ListenAndServe("localhost:8080", routes.Routes()))
}

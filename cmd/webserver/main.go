package main

import (
	"log"
	"net/http"
	"short_url/config"
	"short_url/internal/app"
)

func main() {
	// Load Env from .env file
	config.LoadEnv()

	// Init PSQL Connection
	db := app.InitDatabase()

	// Init components
	repository := app.InitRepository(db)
	service := app.InitService(repository)
	handler := app.InitHandler(service)

	// Routers
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/create", handler.Create)
	mux.HandleFunc("/{shortKey}", handler.Redirect)

	port := config.GetEnv(config.EnvServerPort, ":8080")
	log.Println("Server running at port" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}

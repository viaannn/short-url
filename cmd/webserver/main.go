package main

import (
	"log"
	"net/http"
	"short_url/config"
	"short_url/internal/app"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.Handle("/amour/v1/ping", app.LoggingMiddleware(http.HandlerFunc(handler.Ping)))
	router.Handle("/amour/v1/create", app.LoggingMiddleware(http.HandlerFunc(handler.Create)))
	router.Handle("/{key}", app.LoggingMiddleware(http.HandlerFunc(handler.Redirect)))

	// Start the server
	port := config.GetEnv(config.EnvServerPort, ":8080")
	log.Println("Server running at port" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

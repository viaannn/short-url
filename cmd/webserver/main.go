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
	app.InitDatabase()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HomeHandler)

	port := config.GetEnv(config.EnvServerPort, ":8080")
	log.Println("Server running at port" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}

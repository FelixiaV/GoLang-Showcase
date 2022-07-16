package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s",webPort)

	// Define the server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}

	// start the server

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
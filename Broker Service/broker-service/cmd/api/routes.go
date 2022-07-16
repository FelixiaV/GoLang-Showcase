package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	
//Ön kontrol isteği tarafından belirtilen CORS isteğine yetki verilmişse,
//sunucu ön kontrol isteğine izin verilen kaynağı, yöntemleri ve başlıkları 
//belirten bir mesajla yanıt verecektir. 

		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	// Make sure service is still alive
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/",app.Broker)
	return mux
}
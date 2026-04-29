package main

import (
	"log"
	"main/app/internal/handlers"
	"main/app/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(handlers.HandlerDevices)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST","DELETE","OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
    	MaxAge: 300,
	}))
	
	r.Get("/ip", handlers.HandleGetIP)
	r.Post("/upload", handlers.HandleUploadedFile)
	r.Get("/files", handlers.HandleShowFiles)
	r.Get("/files/{id}", handlers.HandleGetFile)
	r.Get("/download/{id}", handlers.HandleDownloadFile)
	r.Delete("/delete/{id}", handlers.HandleDeleteFile)
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("web/static"))))
	log.Printf("Server starting on %s:8080", utils.CheckIP())

    http.ListenAndServe(":8080", r)
}
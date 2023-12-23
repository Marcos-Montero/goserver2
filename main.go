package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("port is not found in the environment")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	v1Rouer := chi.NewRouter()

	v1Rouer.Get("/healthz", handlerReadiness)
	v1Rouer.Get("/err", handlerError)
	router.Mount("/v1", v1Rouer)

	srv := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(":%s", portString),
	}

	log.Printf("Server is running on port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	

}
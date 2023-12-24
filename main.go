package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/MarcosMRod/goserver2/server"
	"github.com/MarcosMRod/goserver2/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main () {
	feed, err := utils.UrlToFeed("https://wagslane.dev/index.xml");
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("dbUrl is not found in the environment")
	}

	connection, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db := database.New(connection)
	apiCfg := server.NewApiConfig(db)

	go utils.StartScraping(db, 10, time.Minute )

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", server.HandlerReadiness)
	v1Router.Get("/err", server.HandlerError)

	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))

	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))

	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(":%s", portString),
	}

	log.Printf("Server is running on port %s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	

}
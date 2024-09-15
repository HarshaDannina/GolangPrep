package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HarshaDannina/GolangPrep/RSSaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("PORT not found in environment")
	}
	if db_url == "" {
		log.Fatal("DB_URL not found in environment")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	apiConfig := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/user", apiConfig.handlerCreateUser)
	v1Router.Get("/user", apiConfig.middlewareAuth(apiConfig.handlerGetUser))
	v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerCreateFeed))
	v1Router.Get("/feeds", apiConfig.handlerGetFeeds)
	v1Router.Delete("/feeds/{feedId}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeed))
	v1Router.Post("/feed_follow", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1Router.Get("/feed_follow", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1Router.Delete("/feed_follow/{feedFollowId}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server starting on port: %s\n", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

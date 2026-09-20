package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dev-sandip/hamrobazar-api/internal/config"
	"github.com/dev-sandip/hamrobazar-api/internal/db"
	"github.com/dev-sandip/hamrobazar-api/internal/handlers"
	"github.com/dev-sandip/hamrobazar-api/internal/middleware"
)

// @title HamroBazar API
// @version 1.0
// @description API for selling items .
// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Default().Println("Database Connection Established!")
	listingHandler := handlers.NewListingHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("POST /listings", listingHandler.Create)
	mux.HandleFunc("GET /listings", listingHandler.List)
	handler := middleware.RequestID(mux)
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is running on port http://localhost:%s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server Failed : %v", err)
	}
}

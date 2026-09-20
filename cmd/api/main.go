package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dev-sandip/hamrobazar-api/cmd/internal/config"
	"github.com/dev-sandip/hamrobazar-api/cmd/internal/db"
	"github.com/dev-sandip/hamrobazar-api/cmd/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Default().Println("Database Connection Established!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is running on port http://localhost:%s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server Failed : %v", err)
	}
}

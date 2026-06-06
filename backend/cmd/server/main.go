package main

import (
	"log"
	"net/http"
	"os"

	"drakemaye/backend/internal/api"
	"drakemaye/backend/internal/storage"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./drakemaye.db"
	}

	db, err := storage.NewDB(dbPath)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	if err := storage.RunMigrations(db); err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}

	if err := storage.SeedData(db); err != nil {
		log.Fatalf("failed to seed db: %v", err)
	}

	r := api.NewRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("DrakeMaye backend listening on :%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

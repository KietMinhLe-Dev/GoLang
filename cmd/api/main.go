package main

import (
	"log"

	"GO/internal/config"
	"GO/internal/db"
	"GO/internal/routes"
)

func main() {
	// 1. Load config
	connect := config.LoadConfig()

	// 2. Initialize DB
	db.ConnnectDB(connect)

	// 3. Setup Router
	r := routes.SetupRouter()

	// 4. Chạy server
	log.Printf("Starting server on port %s", connect.PORT)
	if err := r.Run(":" + connect.PORT); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

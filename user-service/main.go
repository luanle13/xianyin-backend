package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/luanle13/xianyin-backend/user-service/models"
	"github.com/luanle13/xianyin-backend/user-service/repositories"
)

func main() {
	// Load database configuration from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Ping to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// Initialize UserRepository
	repo := repositories.NewUserRepository(db)

	// Example usage (remove or replace with your actual logic)
	ctx := context.Background()
	user := &models.User{Username: "testuser", Email: "test@example.com", Password: "secret"}
	err = repo.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	fmt.Printf("Created user: %+v\n", user)
}

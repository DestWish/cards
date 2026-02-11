package main

import (
	"fmt"
	"log"

	"github.com/DestWish/cards/internal/handlers"
	"github.com/DestWish/cards/internal/models"
	"github.com/DestWish/cards/internal/repository"
	"github.com/DestWish/cards/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Database init failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	cardRepo := repository.NewCardRepository(db)

	userService := service.NewUserService(userRepo)
	cardService := service.NewCardService(cardRepo, userRepo)

	userHandler := handlers.NewUserHandler(userService)
	cardHandler := handlers.NewCardHandler(cardService)

	r := gin.Default()
	userHandler.RegisterRoutes(r)
	cardHandler.RegisterRoutes(r)

	// r.StaticFile("/", "./web/index.html")

	log.Printf("Server starting on: :3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 port=9920 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Card{}); err != nil {
		return nil, fmt.Errorf("Failed to migrate: %w", err)
	}

	return db, nil
}

package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/AFelipeTrujillo/homysync-go-api/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		dbPath = "homysync.db"
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, _ := DB.DB()
	// Enable foreign key constraints
	sqlDB.Exec("PRAGMA foreign_keys = ON;")
	fmt.Println("Database connection established")
	runMigrations()

}

func runMigrations() {

	err := DB.AutoMigrate(
		&domain.User{},
		&domain.Household{},
		&domain.Membership{},
		&domain.CatalogItem{},
		&domain.Recipe{},
		&domain.RecipeIngredient{},
	)

	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	fmt.Println("Database migrations completed")
}

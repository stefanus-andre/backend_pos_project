package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"project_backend/Config"
	"project_backend/Models"
)

var GormDB *gorm.DB

func Migrate() {
	cfg, err := Config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = db.AutoMigrate(&Models.User{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	GormDB = db
	log.Println("Database migrated")

}

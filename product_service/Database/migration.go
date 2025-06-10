package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"product_service/Config"
	"product_service/Models"
)

var GormDB *gorm.DB

func Migrate() {
	cfg, err := Config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database : ", err.Error())
	}

	err = db.AutoMigrate(&Models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database : ", err.Error())
	}

	GormDB = db
	log.Println("Database migrated")
}

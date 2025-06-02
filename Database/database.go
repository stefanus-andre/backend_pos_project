package Database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"project_backend/Config"
)

var DB *pgxpool.Pool

func InitDB() {
	cfg, err := Config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL :", err)
	}

	DB = dbpool
	log.Println("Connected to PostgreSQL successfully")
}

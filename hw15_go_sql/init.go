package main

import (
	"fmt"
	"log"

	"github.com/ddpakhomov/home-work-golang-otus/hw15_go_sql/config"
	"github.com/ddpakhomov/home-work-golang-otus/hw15_go_sql/db"
)

func InitDB() (*db.DB, error) {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	database, err := db.NewDB(connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Println("Connected to the database successfully")
	return database, nil
}

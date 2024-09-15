package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ddpakhomov/home-work-golang-otus/crud-3t/config"
	"github.com/ddpakhomov/home-work-golang-otus/crud-3t/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Загрузка конфигурации
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Создание обработчиков
	termHandler := &handlers.TermHandler{
		Client:     client,
		Database:   config.MongoDB.Database,
		Collection: config.MongoDB.Collection,
	}

	// Создание маршрутизатора
	router := mux.NewRouter()

	// Определение маршрутов
	router.HandleFunc("/terms", termHandler.CreateTerm).Methods("POST")
	router.HandleFunc("/terms/theme/{theme}", termHandler.CreateTermTheme).Methods("POST")
	router.HandleFunc("/terms", termHandler.GetTerms).Methods("GET")
	router.HandleFunc("/terms/theme/{theme}", termHandler.GetThemeTerms).Methods("GET")
	router.HandleFunc("/terms/{term}", termHandler.GetTerm).Methods("GET")
	router.HandleFunc("/terms/theme/{theme}/term/{term}", termHandler.GetThemeTerm).Methods("GET")
	router.HandleFunc("/terms/theme/{theme}/term/{term}", termHandler.UpdateTerm).Methods("PUT")
	router.HandleFunc("/terms/theme/{theme}/term/{term}", termHandler.DeleteTerm).Methods("DELETE")
	router.HandleFunc("/terms/theme/{theme}", termHandler.DeleteTheme).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Starting the server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

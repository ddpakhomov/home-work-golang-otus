package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ddpakhomov/home-work-golang-otus/hw15_go_sql/handlers"
)

func main() {
	// Подключение к базе данных и загрузка конфигурации
	database, err := InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			log.Printf("Error closing the database: %v", cerr)
		}
	}()

	// Инициализация HTTP-хендлеров
	handler := handlers.NewHandler(database)

	http.HandleFunc("/users", handler.UsersHandler)
	http.HandleFunc("/products", handler.ProductsHandler)
	http.HandleFunc("/orders", handler.OrdersHandler)

	// Настройка HTTP-сервера с таймаутами
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Канал для передачи ошибки
	errCh := make(chan error, 1)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	// Проверка ошибки
	srvErr := <-errCh
	if srvErr != nil && !errors.Is(srvErr, http.ErrServerClosed) {
		log.Printf("Error starting server: %v", srvErr)
	}

	// Завершение программы после выполнения всех defer
	if srvErr != nil {
		log.Printf("Server error: %v", srvErr)
	}
}

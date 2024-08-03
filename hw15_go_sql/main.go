package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ddpakhomov/home-work-golang-otus/hw15_go_sql/handlers"
)

func main() {
	database, err := InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			log.Printf("Error closing the database: %v", cerr)
		}
	}()

	handler := handlers.NewHandler(database)

	http.HandleFunc("/users", handler.UsersHandler)
	http.HandleFunc("/products", handler.ProductsHandler)
	http.HandleFunc("/orders", handler.OrdersHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	errCh := make(chan error, 1)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	srvErr := <-errCh
	if srvErr != nil && !errors.Is(srvErr, http.ErrServerClosed) {
		log.Printf("Error starting server: %v", srvErr)
	}

	if srvErr != nil {
		log.Printf("Server error: %v", srvErr)
	}
}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ddpakhomov/home-work-golang-otus/tech-term-trainer/config"
	"github.com/ddpakhomov/home-work-golang-otus/tech-term-trainer/handlers"
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

	// Первая заставка
	fmt.Println(`
  _____                 _       _____                             _____                  _
 |_   _|   ___    ___  | |__   |_   _|   ___   _ __   _ __ ___   |_   _|  _ __    __ _  (_)  _ __     ___   _ __
   | |    / _ \  / __| | '_ \    | |    / _ \ | '__| | '_ ' _ \    | |   | '__|  / _' | | | | '_ \   / _ \ | '__|
   | |   |  __/ | (__  | | | |   | |   |  __/ | |    | | | | | |   | |   | |    | (_| | | | | | | | |  __/ | |
   |_|    \___|  \___| |_| |_|   |_|    \___| |_|    |_| |_| |_|   |_|   |_|     \__,_| |_| |_| |_|  \___| |_|

  это консольное приложение для помощи пользователю повторять технические термины на основе указанной темы.

===================================================================================================================
`)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\tВведи тему: ")
	theme, _ := reader.ReadString('\n')
	theme = strings.TrimSpace(theme) // Удаление начальных и конечных пробелов и символов новой строки
	fmt.Println(`

`)

	// Создание обработчика команд
	termHandler := &handlers.TermHandler{
		Client:     client,
		Database:   config.MongoDB.Database,
		Collection: config.MongoDB.Collection,
		NumDocs:    config.MongoDB.NumDocuments,
		Theme:      theme,
	}

	// Выполнение обработки команд
	termHandler.FetchAndProcessTerms()

	// Ожидание нажатия клавиши Enter для завершения программы
	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

package handlers

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ddpakhomov/home-work-golang-otus/tech-term-trainer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TermHandler struct {
	Client     *mongo.Client
	Database   string
	Collection string
	NumDocs    int
	Theme      string
}

func (h *TermHandler) FetchAndProcessTerms() {
	// Получение случайных документов из MongoDB
	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Использование количества документов из конфигурации
	numDocs := h.NumDocs
	//log.Printf("Number of documents to fetch: %d\n", numDocs)

	// Создание конвейера агрегации с фильтром по теме и выборкой случайных документов
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"theme", h.Theme}}}},
		{{"$sample", bson.D{{"size", numDocs}}}},
	}
	//log.Printf("Aggregation pipeline: %v\n", pipeline)
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatalf("Failed to fetch documents: %v", err)
	}
	defer cursor.Close(ctx)

	var terms []models.Term
	for cursor.Next(ctx) {
		var term models.Term
		if err := cursor.Decode(&term); err != nil {
			log.Fatalf("Failed to decode document: %v", err)
		}
		terms = append(terms, term)
	}

	if len(terms) == 0 {
		log.Println("No terms found for the specified theme.")
		return
	}

	// Отладочное сообщение для проверки извлеченных данных
	//log.Printf("Fetched %d terms from the database.\n", len(terms))

	// Переменные для отслеживания статистики
	var correctCount, incorrectCount int

	// Чтение ввода пользователя и проверка правильности ввода
	reader := bufio.NewReader(os.Stdin)
	for _, term := range terms {
		attempts := 0
		for {
			fmt.Println("*************************************************")
			fmt.Printf("описание термина: %s\n", term.Description)
			fmt.Println("*************************************************")
			fmt.Print("введи термин соответствующий описанию: ")
			userInput, _ := reader.ReadString('\n')
			userInput = strings.TrimSpace(userInput) // Удаление начальных и конечных пробелов и символов новой строки
			if userInput == term.Term {
				fmt.Println(`
  /$$$$$$                                                            
 /$$__  $$                                                           
| $$  \__/ /$$   /$$  /$$$$$$$  /$$$$$$$  /$$$$$$   /$$$$$$$ /$$$$$$$
|  $$$$$$ | $$  | $$ /$$_____/ /$$_____/ /$$__  $$ /$$_____//$$_____/
 \____  $$| $$  | $$| $$      | $$      | $$$$$$$$|  $$$$$$|  $$$$$$ 
 /$$  \ $$| $$  | $$| $$      | $$      | $$_____/ \____  $$\____  $$
|  $$$$$$/|  $$$$$$/|  $$$$$$$|  $$$$$$$|  $$$$$$$ /$$$$$$$//$$$$$$$/
 \______/  \______/  \_______/ \_______/ \_______/|_______/|_______/ 

`)
				fmt.Println("\t\tВерно, ты капитальный молодец!\n\n")
				correctCount++
				break
			} else {
				attempts++
				if attempts >= 3 {
					fmt.Println(`
 /$$$$$$$$                                         /$$
| $$_____/                                        | $$
| $$        /$$$$$$   /$$$$$$   /$$$$$$   /$$$$$$ | $$
| $$$$$    /$$__  $$ /$$__  $$ /$$__  $$ /$$__  $$| $$
| $$__/   | $$  \__/| $$  \__/| $$  \ $$| $$  \__/|__/
| $$      | $$      | $$      | $$  | $$| $$          
| $$$$$$$$| $$      | $$      |  $$$$$$/| $$       /$$
|________/|__/      |__/       \______/ |__/      |__/
`)
					fmt.Printf("Правильным был ответ: %s\n", term.Term)
					fmt.Println("---------------------------------------------------------------------------------")
					fmt.Printf("Термин: %s,\n Description: %s,\n", term.Term, term.Description)
					fmt.Println("---------------------------------------------------------------------------------")

					incorrectCount++
					break
				} else {
					fmt.Println(fmt.Sprintf(`
    __________________|      |________________________________________________________________
         ,--.    ,--.          ,--.   ,--.
        |oo  | _  \  '.       | oo | |  oo|
    o  o|~~  |(_) /   ;       | ~~ | |  ~~|o твой ответ -->  %s  <-- на этот раз пропал  o
        |/\/\|   '._,'        |/\/\| |/\/\|
    __________________        ________________________________________________________________
                      |      |
`, userInput))
					fmt.Println("Не верно, соберись и попробуй снова\n\n")
				}
			}
		}
	}

	// Вывод статистики
	fmt.Printf("\nСтатистика:\n")
	fmt.Printf("Правильных ответов: %d\n", correctCount)
	fmt.Printf("Не ответил: %d раз\n", incorrectCount)

	fmt.Println(`
  ________                                          __________ 
 /  _____/_____    _____   ____     _______  __ ____\______   \
/   \  ___\__  \  /     \_/ __ \   /  _ \  \/ // __ \|       _/
\    \_\  \/ __ \|  Y Y  \  ___/  (  <_> )   /\  ___/|    |   \
 \______  (____  /__|_|  /\___  >  \____/ \_/  \___  >____|_  /
        \/     \/      \/     \/                   \/       \/ 
`)

	// Завершение программы с сообщением
	fmt.Println("Приходи еще раз и помни залог успеха в формировании хороших привычек! Будь здоров!")
}

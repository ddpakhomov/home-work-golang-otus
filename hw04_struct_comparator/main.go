package main

import "fmt"

type ComparisonMode struct {
	ComparisonType ComparisonType
}

type ComparisonType int

const (
	ByYear ComparisonType = iota
	BySize
	ByRate
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewComparisonMode(compareBy ComparisonType) *ComparisonMode {
	return &ComparisonMode{
		ComparisonType: compareBy,
	}
}

func (c ComparisonMode) CompareBooks(book1, book2 *Book) bool {
	switch c.ComparisonType {
	case ByYear:
		return book1.year > book2.year
	case BySize:
		return book1.size > book2.size
	case ByRate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

func NewBook(title string, author string, year int, size int, rate float64) *Book {
	return &Book{
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func main() {
	myBook := Book{
		id:     1,
		title:  "Go Programming",
		author: "John Doe",
		year:   2021,
		size:   300,
		rate:   4.5,
	}
	fmt.Println("\n====================================================")
	fmt.Println("Вывод информации о книге ДО использования методов")
	fmt.Println("Book ID:", myBook.id)
	fmt.Println("Title:", myBook.title)
	fmt.Println("Author:", myBook.author)
	fmt.Println("Year:", myBook.year)
	fmt.Println("Size:", myBook.size)
	fmt.Println("Rate:", myBook.rate)
	fmt.Println("====================================================")

	fmt.Println("\nИспользуя метод GetID, получили значение id:", myBook.GetID())
	myBook.SetID(2)
	fmt.Println("Используя метод SetID, обновили значание id на:", myBook.GetID())

	fmt.Println("\nИспользуя метод GetTitle, получили значение title:", myBook.GetTitle())
	myBook.SetTitle("Golang Programming")
	fmt.Println("Используя метод SetTitle, обновили значание title на:", myBook.GetTitle())

	fmt.Println("\nИспользуя метод GetAuthor, получили значение author:", myBook.GetAuthor())
	myBook.SetAuthor("Женек Дой")
	fmt.Println("Используя метод SetAuthor, обновили значание author на:", myBook.GetAuthor())

	fmt.Println("\nИспользуя метод GetYear, получили значение year:", myBook.GetYear())
	myBook.SetYear(2024)
	fmt.Println("Используя метод SetYear, обновили значание year на:", myBook.GetYear())

	fmt.Println("\nИспользуя метод GetSize, получили значение size:", myBook.GetSize())
	myBook.SetSize(426)
	fmt.Println("Используя метод SetSize, обновили значание size на:", myBook.GetSize())

	fmt.Println("\nИспользуя метод GetRate, получили значение rate:", myBook.GetRate())
	myBook.SetRate(4.3)
	fmt.Println("Используя метод SetRate, обновили значание rate на:", myBook.GetRate())

	fmt.Println("\n====================================================")
	fmt.Println("Вывод информации о книге ПОСЛЕ использования методов")
	fmt.Println("Book ID:", myBook.id)
	fmt.Println("Title:", myBook.title)
	fmt.Println("Author:", myBook.author)
	fmt.Println("Year:", myBook.year)
	fmt.Println("Size:", myBook.size)
	fmt.Println("Rate:", myBook.rate)
	fmt.Println("====================================================")

	book1 := NewBook("Book 1", "Author 1", 2020, 350, 4.2)
	book2 := NewBook("Book 2", "Author 2", 2021, 360, 4.0)

	mode := NewComparisonMode(ByRate)

	if mode.CompareBooks(book1, book2) {
		fmt.Println("Book 1 is greater than Book 2")
	} else {
		fmt.Println("Book 2 is greater than Book 1")
	}
}

func (b *Book) SetID(newID int) {
	b.id = newID
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) SetTitle(newTitle string) {
	b.title = newTitle
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(newAuthor string) {
	b.author = newAuthor
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetYear(newYear int) {
	b.year = newYear
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) SetSize(newSize int) {
	b.size = newSize
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) SetRate(newRate float64) {
	b.rate = newRate
}

func (b *Book) GetRate() float64 {
	return b.rate
}

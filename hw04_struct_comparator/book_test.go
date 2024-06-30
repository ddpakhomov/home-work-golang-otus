package main

import "testing"

func TestNewBook(t *testing.T) {
	title, author, year, size, rate := "Test Book", "Test Author", 2000, 250, 4.0
	book := NewBook(title, author, year, size, rate)

	if book.title != "Test Book" {
		t.Errorf("Ожидается, что заголовок книги будет %s, получили %s", title, book.title)
	}

	if book.author != "Test Author" {
		t.Errorf("Ожидается, что автор книги будет %s, получили %s", author, book.author)
	}

	if book.year != 2000 {
		t.Errorf("Ожидается, что год книги будет %d, получили %d", year, book.year)
	}

	if book.size != 250 {
		t.Errorf("Ожидается, что размер книги будет %d, получили %d", size, book.size)
	}

	if book.rate != 4.0 {
		t.Errorf("Ожидается, что рейтинг книги будет %f, получили %f", rate, book.rate)
	}
}

func TestCompareBooksByYear(t *testing.T) {
	book1 := NewBook("Book 1", "Author 1", 2022, 350, 4.2)
	book2 := NewBook("Book 2", "Author 2", 2021, 360, 4.0)

	mode := NewComparisonMode(ByYear)

	if !mode.CompareBooks(book1, book2) {
		t.Error("Ожидается, что book1 вышла позже чем книга book2")
	}
}

func TestCompareBooksBySize(t *testing.T) {
	book1 := NewBook("Book 1", "Author 1", 2020, 370, 4.2)
	book2 := NewBook("Book 2", "Author 2", 2021, 360, 4.0)

	mode := NewComparisonMode(BySize)

	if !mode.CompareBooks(book1, book2) {
		t.Error("Ожидается, что book1 больше книг book2")
	}
}

func TestCompareBooksByRate(t *testing.T) {
	book1 := NewBook("Book 1", "Author 1", 2020, 350, 4.2)
	book2 := NewBook("Book 2", "Author 2", 2021, 360, 4.5)

	mode := NewComparisonMode(ByRate)

	if mode.CompareBooks(book1, book2) {
		t.Error("Ожидается, что рейтинг book2 выше чем book1")
	}
}

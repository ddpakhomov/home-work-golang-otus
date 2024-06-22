package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Book struct {
	ID     int32   `json:"id" protobuf:"varint,1,opt,name=id,proto3"`
	Title  string  `json:"title" protobuf:"bytes,2,opt,name=title,proto3"`
	Author string  `json:"author" protobuf:"bytes,3,opt,name=author,proto3"`
	Year   int32   `json:"year" protobuf:"varint,4,opt,name=year,proto3"`
	Size   int32   `json:"size" protobuf:"varint,5,opt,name=size,proto3"`
	Rate   float64 `json:"rate" protobuf:"fixed64,6,opt,name=rate,proto3"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(b),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	return json.Unmarshal(data, aux)
}

func (b *Book) Reset() {
	*b = Book{}
}

func (b *Book) String() string {
	return proto.CompactTextString(b)
}

func (b *Book) ProtoMessage() {}

func SerializeBooksToJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeBooksFromJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}

type Books struct {
	Books []Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books"`
}

func main() {
	book := &Book{
		ID:     1,
		Title:  "Hello Proto",
		Author: "Ivan Ivanov",
		Year:   2025,
		Size:   218,
		Rate:   4.5,
	}

	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	var unmarshaledBook Book
	err = json.Unmarshal(jsonData, &unmarshaledBook)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Unmarshaled Book: %+v\n", unmarshaledBook)

	data, err := proto.Marshal(book)
	if err != nil {
		fmt.Println("Marshaling error:", err)
	}

	newBook := &Book{}
	err = proto.Unmarshal(data, newBook)
	if err != nil {
		fmt.Println("Unmarshaling error:", err)
	}

	fmt.Println(newBook)
	fmt.Println(data)

	books := []Book{
		{ID: 1, Title: "Hello Proto", Author: "Ivan Ivanov", Year: 2025, Size: 218, Rate: 4.5},
		{ID: 2, Title: "Go in Action", Author: "John Doe", Year: 2021, Size: 300, Rate: 4.7},
	}

	jsonDatas, err := SerializeBooksToJSON(books)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonDatas))

	unmarshaledBooks, err := DeserializeBooksFromJSON(jsonDatas)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Unmarshaled Books: %+v\n", unmarshaledBooks)
}

package main

import (
	"encoding/json"
	"testing"

	"github.com/ddpakhomov/home-work-golang-otus/hw09_serialize/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestSerializeBooksToJSON(t *testing.T) {
	books := []Book{
		{ID: 1, Title: "Hello Proto", Author: "Ivan Ivanov", Year: 2025, Size: 218, Rate: 4.5},
		{ID: 2, Title: "Go in Action", Author: "John Doe", Year: 2021, Size: 300, Rate: 4.7},
	}

	jsonData, err := SerializeBooksToJSON(books)
	assert.NoError(t, err, "Expected no error during JSON serialization")
	assert.NotNil(t, jsonData, "Expected non-nil JSON data")

	var deserializedBooks []Book
	err = json.Unmarshal(jsonData, &deserializedBooks)
	assert.NoError(t, err, "Expected no error during JSON deserialization")
	assert.Equal(t, books, deserializedBooks, "Expected deserialized books to match original")
}

func TestDeserializeBooksFromJSON(t *testing.T) {
	jsonData := `[
                {"id":1,"title":"Hello Proto","author":"Ivan Ivanov","year":2025,"size":218,"rate":4.5},
                {"id":2,"title":"Go in Action","author":"John Doe","year":2021,"size":300,"rate":4.7}
        ]`

	books, err := DeserializeBooksFromJSON([]byte(jsonData))
	assert.NoError(t, err, "Expected no error during JSON deserialization")
	assert.Len(t, books, 2, "Expected two books to be deserialized")

	expectedBooks := []Book{
		{ID: 1, Title: "Hello Proto", Author: "Ivan Ivanov", Year: 2025, Size: 218, Rate: 4.5},
		{ID: 2, Title: "Go in Action", Author: "John Doe", Year: 2021, Size: 300, Rate: 4.7},
	}
	assert.Equal(t, expectedBooks, books, "Expected deserialized books to match expected")
}

func TestProtoSerialization(t *testing.T) {
	protoBook := &pb.Book{
		Id:     1,
		Title:  "Hello Proto",
		Author: "Ivan Ivanov",
		Year:   2025,
		Size:   218,
		Rate:   4.5,
	}

	data, err := proto.Marshal(protoBook)
	assert.NoError(t, err, "Expected no error during Proto serialization")
	assert.NotNil(t, data, "Expected non-nil Proto data")

	newBook := new(pb.Book)
	err = proto.Unmarshal(data, newBook)
	assert.NoError(t, err, "Expected no error during Proto deserialization")
	assertProtoBooksEqual(t, protoBook, newBook)
}

func TestJSONSerialization(t *testing.T) {
	book := &Book{
		ID:     1,
		Title:  "Hello Proto",
		Author: "Ivan Ivanov",
		Year:   2025,
		Size:   218,
		Rate:   4.5,
	}

	jsonData, err := json.Marshal(book)
	assert.NoError(t, err, "Expected no error during JSON serialization")
	assert.NotNil(t, jsonData, "Expected non-nil JSON data")

	var newBook Book
	err = json.Unmarshal(jsonData, &newBook)
	assert.NoError(t, err, "Expected no error during JSON deserialization")
	assert.Equal(t, *book, newBook, "Expected deserialized book to match original")
}

func assertProtoBooksEqual(t *testing.T, expected, actual *pb.Book) {
	t.Helper() // This will mark this function as a test helper
	assert.Equal(t, expected.Id, actual.Id, "Expected ID to match")
	assert.Equal(t, expected.Title, actual.Title, "Expected Title to match")
	assert.Equal(t, expected.Author, actual.Author, "Expected Author to match")
	assert.Equal(t, expected.Year, actual.Year, "Expected Year to match")
	assert.Equal(t, expected.Size, actual.Size, "Expected Size to match")
	assert.Equal(t, expected.Rate, actual.Rate, "Expected Rate to match")
}

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/ddpakhomov/home-work-golang-otus/crud-3t/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TermHandler struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

func (h *TermHandler) CreateTerm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var term models.Term
	err := json.NewDecoder(r.Body).Decode(&term)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Проверка наличия всех необходимых полей
	if term.Term == "" || term.Name == "" || term.Description == "" || term.Theme == "" {
		http.Error(w, "All fields (term, name, description, theme) are required", http.StatusBadRequest)
		return
	}

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверка на существование термина
	var existingTerm models.Term
	err = collection.FindOne(ctx, bson.M{"term": term.Term}).Decode(&existingTerm)
	if err == nil {
		http.Error(w, "Term already exists", http.StatusConflict)
		return
	} else if err != mongo.ErrNoDocuments {
		http.Error(w, "Failed to check term existence", http.StatusInternalServerError)
		return
	}

	// Вставка нового термина
	result, err := collection.InsertOne(ctx, term)
	if err != nil {
		http.Error(w, "Failed to insert term", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (h *TermHandler) CreateTermTheme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	theme := vars["theme"]

	var term models.Term
	err := json.NewDecoder(r.Body).Decode(&term)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Проверка наличия всех необходимых полей
	if term.Term == "" || term.Name == "" || term.Description == "" {
		http.Error(w, "All fields (term, name, description) are required", http.StatusBadRequest)
		return
	}

	// Установка темы из URL
	term.Theme = theme

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверка на существование термина в данной теме
	var existingTerm models.Term
	err = collection.FindOne(ctx, bson.M{"term": term.Term, "theme": term.Theme}).Decode(&existingTerm)
	if err == nil {
		http.Error(w, "Term already exists in this theme", http.StatusConflict)
		return
	} else if err != mongo.ErrNoDocuments {
		http.Error(w, "Failed to check term existence", http.StatusInternalServerError)
		return
	}

	// Вставка нового термина
	result, err := collection.InsertOne(ctx, term)
	if err != nil {
		http.Error(w, "Failed to insert term", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (h *TermHandler) GetTerms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var terms []models.Term
	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var term models.Term
		cursor.Decode(&term)
		terms = append(terms, term)
	}
	json.NewEncoder(w).Encode(terms)
}

func (h *TermHandler) GetThemeTerms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	themeValue := params["theme"]

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"theme": themeValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var terms []models.Term
	for cursor.Next(ctx) {
		var term models.Term
		if err := cursor.Decode(&term); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		terms = append(terms, term)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(terms) == 0 {
		http.Error(w, "Terms not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(terms)
}

func (h *TermHandler) GetTerm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	termValue := params["term"]

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var term models.Term
	err := collection.FindOne(ctx, bson.M{"term": termValue}).Decode(&term)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Term not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(term)
}

func (h *TermHandler) GetThemeTerm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	termValue := params["term"]
	themeValue := params["theme"]

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var term models.Term
	err := collection.FindOne(ctx, bson.M{"theme": themeValue, "term": termValue}).Decode(&term)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Term not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(term)
}

func (h *TermHandler) UpdateTerm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	themeValue := params["theme"]
	termValue := params["term"]

	var term models.Term
	err := json.NewDecoder(r.Body).Decode(&term)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Проверка наличия всех необходимых полей
	if term.Term == "" || term.Name == "" || term.Description == "" || term.Theme == "" {
		http.Error(w, "All fields (term, name, description, theme) are required", http.StatusBadRequest)
		return
	}

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверка существования термина
	var existingTerm models.Term
	err = collection.FindOne(ctx, bson.M{"theme": themeValue, "term": termValue}).Decode(&existingTerm)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Term not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	update := bson.M{
		"$set": term,
	}
	_, err = collection.UpdateOne(ctx, bson.M{"theme": themeValue, "term": termValue}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(term)
}

func (h *TermHandler) DeleteTerm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	themeValue := params["theme"]
	termValue := params["term"]

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверка существования термина в указанной теме
	count, err := collection.CountDocuments(ctx, bson.M{"theme": themeValue, "term": termValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		http.Error(w, "Term not found in the specified theme", http.StatusNotFound)
		return
	}

	// Удаление найденного термина
	_, err = collection.DeleteMany(ctx, bson.M{"theme": themeValue, "term": termValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Формирование успешного ответа
	response := map[string]string{
		"message": fmt.Sprintf("Term '%s' in theme '%s' successfully deleted", termValue, themeValue),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *TermHandler) DeleteTheme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	themeValue := params["theme"]

	collection := h.Client.Database(h.Database).Collection(h.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверка существования терминов в указанной теме
	count, err := collection.CountDocuments(ctx, bson.M{"theme": themeValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		http.Error(w, "No terms found for the specified theme", http.StatusNotFound)
		return
	}

	// Удаление всех терминов в указанной теме
	_, err = collection.DeleteMany(ctx, bson.M{"theme": themeValue})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Формирование успешного ответа
	response := map[string]string{
		"message": fmt.Sprintf("All terms in theme '%s' successfully deleted", themeValue),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

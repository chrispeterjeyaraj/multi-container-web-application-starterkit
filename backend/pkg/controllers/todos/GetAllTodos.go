package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/configs"
	models "github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/pkg/models"
)

func GetAllTodos(response http.ResponseWriter, request *http.Request) {
	// Check if the request method is GET
	if request.Method != http.MethodGet {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a context with a timeout for database operations
	ctx, cancel := context.WithTimeout(request.Context(), 10*time.Second)
	defer cancel()

	// Get the MongoDB collection for todos
	collection := configs.GetCollection(configs.DB, "todos")

	// Find all todos in the collection
	allTodos, err := collection.Find(ctx, bson.M{})
	if err != nil {
		// Handle internal server error if the database query fails
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer allTodos.Close(ctx)

	// Create a slice to store the retrieved todos
	var results []models.Todo

	// Iterate over the todos and decode them into the results slice
	for allTodos.Next(ctx) {
		var singleTodo models.Todo
		if err := allTodos.Decode(&singleTodo); err != nil {
			// Handle internal server error if decoding fails
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		results = append(results, singleTodo)
	}

	// Check if any todos were found
	if len(results) == 0 {
		// Return a not found error if no todos were found
		http.Error(response, "No todos found", http.StatusNotFound)
		return
	}

	// Set the response content type to JSON
	response.Header().Set("Content-Type", "application/json")

	// Encode and send the results as a JSON response
	if err := json.NewEncoder(response).Encode(&results); err != nil {
		// Handle internal server error if encoding fails
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
}

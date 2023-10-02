package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/configs"
	models "github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/pkg/models"
)

func GetAllTodos(CollectionName string) (bool, []models.Todo) {
	var results []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	allTodos, errFind := collection.Find(ctx, bson.M{})

	if errFind != nil {
		return false, nil
	}

	for allTodos.Next(ctx) {
		var singleTodo models.Todo
		if errFind = allTodos.Decode(&singleTodo); errFind != nil {
			return false, nil
		}

		results = append(results, singleTodo)
	}

	defer cancel()
	return true, results
}

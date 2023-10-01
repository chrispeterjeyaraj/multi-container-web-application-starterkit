package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/chrispeterjeyaraj/server-ui-container-boilerplate/backend/database"
)

func GetAllTodos(response http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	getErr, results := database.GetAllTodos("todos")

	if getErr {
		response.WriteHeader(http.StatusOK)
		if results == nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{\"message\": \"No todos Found\"}"))
			return
		} else {
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(&results)
			return
		}

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}

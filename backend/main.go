package main

import (
	"log"
	"net/http"

	"github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/configs"
	"github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.TodosRoute(router)

	log.Fatal(http.ListenAndServe(":4000", router))
}

package routes

import (
	controllers "github.com/chrispeterjeyaraj/multi-container-web-application-starterkit/backend/pkg/controllers/todos"

	"github.com/gorilla/mux"
)

func TodosRoute(router *mux.Router) {
	router.HandleFunc("/gettodos", controllers.GetAllTodos)

}

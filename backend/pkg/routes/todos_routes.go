package routes

import (
	controllers "github.com/chrispeterjeyaraj/server-ui-container-boilerplate/backend/pkg/controllers/todos"

	"github.com/gorilla/mux"
)

func TodosRoute(router *mux.Router) {
	router.HandleFunc("/gettodos", controllers.GetAllTodos)

}

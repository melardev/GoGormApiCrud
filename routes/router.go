package routes

import (
	"github.com/melardev/GoGormApiCrud/controllers"
	"github.com/melardev/GoGormApiCrud/dtos"
	"net/http"
	"strconv"
	"strings"
)

func RegisterRoutes() {

	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method == http.MethodGet {
			controllers.GetAllTodos(w, r)
			return
		} else if method == http.MethodPost {
			controllers.CreateTodo(w, r)
			return
		} else if method == http.MethodDelete {
			controllers.DeleteAllTodos(w, r)
			return
		}

		controllers.SendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage("Unsupported HTTP Method"))
	})

	http.HandleFunc("/api/todos/completed", controllers.GetAllCompletedTodos)
	http.HandleFunc("/api/todos/pending", controllers.GetAllPendingTodos)

	// Very important to finish the URL in / otherwise we will not receive /api/todos/id
	http.HandleFunc("/api/todos/", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		// remaining := strings.TrimPrefix(r.URL.Path, "/api/todos")
		parts := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			controllers.SendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("Invalid Id format"))
			return
		}
		if method == http.MethodGet {
			controllers.GetTodoById(id, w, r)
			return
		} else if method == http.MethodPut {
			controllers.UpdateTodo(id, w, r)
			return
		} else if method == http.MethodDelete {
			controllers.DeleteTodo(id, w, r)
			return
		}

		controllers.SendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage("Unsopported HTTP Method"))
	})
}

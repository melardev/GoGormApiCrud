package controllers

import (
	"encoding/json"
	"github.com/melardev/GoGormApiCrud/dtos"
	"github.com/melardev/GoGormApiCrud/models"
	"github.com/melardev/GoGormApiCrud/services"
	"net/http"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.FetchTodos()
	SendAsJson(w, http.StatusOK, dtos.GetTodoListDto(todos))
}

func GetAllPendingTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.FetchPendingTodos()
	SendAsJson(w, http.StatusOK, dtos.GetTodoListDto(todos))
}
func GetAllCompletedTodos(w http.ResponseWriter, r *http.Request) {
	todos := services.FetchCompletedTodos()
	SendAsJson(w, http.StatusOK, dtos.GetTodoListDto(todos))
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		SendAsJson(w, http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	defer r.Body.Close()

	todo, err := services.CreateTodo(todo.Title, todo.Description, todo.Completed)
	if err != nil {
		SendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	SendAsJson(w, http.StatusCreated, dtos.GetTodoDetaislDto(&todo))
}

func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	services.DeleteAllTodos()
	SendAsJson(w, http.StatusNoContent, nil)
}

func GetTodoById(id int, w http.ResponseWriter, r *http.Request) {
	/*id, ok := r.URL.Query()["id"]

	if !ok || len(id[0]) < 1 {
		SendAsJson(w, http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide an Id"))
	}
	*/

	// id64, _ := strconv.ParseUint(id[0], 10, 32)
	// todo, err := services.FetchById(uint(id64))
	todo, err := services.FetchById(uint(id))
	if err != nil {
		SendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not find Todo"))
		return
	}

	sendAsJson2(w, http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}
func UpdateTodo(id int, w http.ResponseWriter, r *http.Request) {

	var todoInput models.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todoInput); err != nil {
		SendAsJson(w, http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	defer r.Body.Close()

	todo, err := services.UpdateTodo(uint(id), todoInput.Title, todoInput.Description, todoInput.Completed)
	if err != nil {
		SendAsJson(w, http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	SendAsJson(w, http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}

func DeleteTodo(id int, w http.ResponseWriter, r *http.Request) {
	todo, err := services.FetchById(uint(id))
	if err != nil {
		SendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("todo not found"))
		return
	}

	err = services.DeleteTodo(&todo)

	if err != nil {
		SendAsJson(w, http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not delete Todo"))
		return
	}

	SendAsJson(w, http.StatusNoContent, nil)
}

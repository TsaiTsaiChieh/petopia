package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"petopia-server/services"
	"strconv"

	"github.com/gorilla/mux"
)

var TodoList []Todo

type Todo struct {
	Id   int64
	Item string
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	if err != nil {
		fmt.Println(err)
	}
	var addtodo Todo
	_ = json.Unmarshal(body, &addtodo) // to json
	TodoList = append(TodoList, addtodo)
	defer r.Body.Close()
	response := ApiResponse{"200", TodoList}

	services.ResponseWithJson(w, http.StatusOK, response)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"] // get the id
	var targetTodo Todo
	for _, Todo := range TodoList {
		intQueryId, _ := strconv.ParseInt(queryId, 10, 64) // string to int64
		if Todo.Id == intQueryId {
			targetTodo = Todo
		}
	}
	response := ApiResponse{"200", targetTodo}
	services.ResponseWithJson(w, http.StatusOK, response)

}

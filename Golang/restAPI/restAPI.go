package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	todos = append(todos, newTodo)
	context.JSON(http.StatusCreated, newTodo)
}

func getTodoById(context *gin.Context) {
	id := context.Param("id")
	todo, err := findTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, todo)
}

func updateTodoByID(context *gin.Context) {
	id := context.Param("id")
	var updatedTodo todo
	if err := context.BindJSON(&updatedTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	index, err := findTodoIndexByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todos[index] = updatedTodo
	context.JSON(http.StatusOK, updatedTodo)
}

func deleteTodoByID(context *gin.Context) {
	id := context.Param("id")
	index, err := findTodoIndexByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func patchTodoByID(context *gin.Context) {
	id := context.Param("id")
	var updatedFields map[string]interface{}
	if err := context.BindJSON(&updatedFields); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	index, err := findTodoIndexByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Update the fields that were sent in the PATCH request
	if val, ok := updatedFields["item"]; ok {
		todos[index].Item = val.(string)
	}
	if val, ok := updatedFields["completed"]; ok {
		todos[index].Completed = val.(bool)
	}

	context.JSON(http.StatusOK, todos[index])
}

func findTodoByID(id string) (*todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func findTodoIndexByID(id string) (int, error) {
	for i, t := range todos {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.POST("/todos", addTodo)
	router.PUT("/todos/:id", updateTodoByID)
	router.DELETE("/todos/:id", deleteTodoByID)
	router.PATCH("/todos/:id", patchTodoByID)
	router.Run("localhost:8040")
}

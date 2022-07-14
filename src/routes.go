package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodoList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoByID(id) 

	if err != nil { 
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"}) 
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil { return }

	todos = append(todos, newTodo)
	
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func completeTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id) 

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not fonund"})
	}

	todo.Completed = true
	
	context.IndentedJSON(http.StatusOK, todo)
}
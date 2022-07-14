package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        string `json:"id"`
	Objective string `json:"objective"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Objective: "clean room", Completed: false},
	{ID: "2", Objective: "do hw", Completed: false},
	{ID: "3", Objective: "stuff", Completed: false},
	{ID: "4", Objective: "read book", Completed: false},
}

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	if (os.Getenv("mode") == "release") {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/todos", getTodoList)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", completeTodo)
	router.POST("/todos", addTodo)




	router.Run("localhost:3000")
	fmt.Println("server started on port 3000")
}
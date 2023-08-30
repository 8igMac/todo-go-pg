package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []todo{
	{ID: "1", Title: "Buy surfboard", Done: false},
	{ID: "2", Title: "Solve 1 leetcode", Done: false},
	{ID: "3", Title: "Read 1 chapter of book.", Done: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func getTodosByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found."})
}

func postTodos(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodosByID)
	router.POST("/todos", postTodos)
	router.Run("localhost:8080")
}

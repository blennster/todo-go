package main

import (
	// "fmt"
	"fmt"
	"net/http"
	// "net/http"

	// "github.com/blennster/todo-go/routes"
	"github.com/blennster/todo-go/internal"
	"github.com/blennster/todo-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/**/*.html")
	r.Static("/static", "web/static")

	store := internal.NewTodoStore()
	store.Add(internal.Todo{Name: "Todo 1", Completed: false})

	r.Use(internal.TodoMiddleware(&store))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/todos", routes.GetTodos)
	r.PATCH("/todos/:name", routes.ToggleTodo)
	r.POST("/todos", routes.AddTodo)
	r.DELETE("/todos/:name", routes.DeleteTodo)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

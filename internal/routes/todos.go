package routes

import (
	"net/http"

	"github.com/blennster/todo-go/internal"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	store := internal.GetStore(c)
	store.DoRead(func(todos internal.Todos) {
		c.HTML(http.StatusOK, "todos.html", todos)
	})
}

func AddTodo(c *gin.Context) {
	store := internal.GetStore(c)
	name := c.PostForm("name")
	if store.Add(internal.Todo{Name: name, Completed: false}) {
		GetTodos(c)
	} else {
		Error(c, "todo already exists", http.StatusConflict)
	}
}

func ToggleTodo(c *gin.Context) {
	store := internal.GetStore(c)
	name := c.Param("name")
	if store.Toggle(name) {
		GetTodos(c)
	} else {
		Error(c, "todo not found", http.StatusConflict)
	}
}

func DeleteTodo(c *gin.Context) {
	store := internal.GetStore(c)
	name := c.Param("name")
	if store.Delete(name) {
		GetTodos(c)
	} else {
		Error(c, "todo not found", http.StatusConflict)
	}
}

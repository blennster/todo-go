package internal

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Todos map[string]*Todo

type TodoStore struct {
	sync.RWMutex
	todos Todos
}

func NewTodoStore() TodoStore {
	return TodoStore{
		todos: make(Todos),
	}
}

func (t *TodoStore) DoRead(f func(Todos)) {
	t.RLock()
	defer t.RUnlock()
	f(t.todos)
}

// Returns false if item already exists
func (t *TodoStore) Add(v Todo) bool {
	t.Lock()
	defer t.Unlock()
	_, ok := t.todos[v.Name]
	if ok {
		return false
	}
	t.todos[v.Name] = &v
	return true
}

func (t *TodoStore) Update(v Todo) {
	t.Lock()
	defer t.Unlock()
	t.todos[v.Name] = &v
}

func (t *TodoStore) Toggle(name string) bool {
	t.Lock()
	defer t.Unlock()
	todo, ok := t.todos[name]
	if !ok {
		return false
	}

	todo.Completed = !todo.Completed
	return true
}

func (t *TodoStore) Delete(name string) bool {
	t.Lock()
	defer t.Unlock()
	_, ok := t.todos[name]
	if !ok {
		return false
	}
	delete(t.todos, name)
	return true
}

func TodoMiddleware(store *TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	}
}

func GetStore(c *gin.Context) *TodoStore {
	store, ok := c.MustGet("store").(*TodoStore)
	if !ok {
		panic("store is not of correct type")
	}
	return store
}

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Todo struct represents a task in the Todo list
// @Description Todo represents a task in the list
type Todo struct {
	ID     int    `json:"id"`     // The ID of the todo task
	Task   string `json:"task"`   // The name or description of the task
	Status string `json:"status"` // Status of the task: "pending", "completed"
}

// ErrorResponse is a structure for error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// DeleteResponse represents the response for a successful deletion
type DeleteResponse struct {
	Message string `json:"message"`
}

var todos = []Todo{
	{ID: 1, Task: "Learn Go", Status: "pending"},
	{ID: 2, Task: "Learn Gin", Status: "completed"},
}

// @title Todo API
// @description This is a simple API to manage a Todo list
// @contact.name Your Name
// @contact.url https://yourwebsite.com
// @contact.email your-email@example.com
// @version 1.0
// @BasePath /
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Serve Swagger UI HTML file at /docs
	// @Summary Serve Swagger UI HTML
	// @Router /docs [get]
	router.GET("/docs", func(c *gin.Context) {
		c.File("./swagger-ui.html") // Serve the Swagger UI HTML file
	})

	// Serve Swagger JSON file at /swagger.json
	// @Summary Serve Swagger JSON documentation
	// @Router /swagger.json [get]
	router.GET("/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json") // Serve the Swagger JSON file
	})

	// Define API routes with Swagger annotations

	// @Summary Get all todos
	// @Description Get all todos from the list
	// @Tags Todo
	// @Accept json
	// @Produce json
	// @Success 200 {array} Todo
	// @Router /todos [get]
	router.GET("/todos", getTodos)

	// @Summary Get a todo by ID
	// @Description Get a todo by its ID
	// @Tags Todo
	// @Accept json
	// @Produce json
	// @Param id path int true "Todo ID"
	// @Success 200 {object} Todo
	// @Failure 404 {object} ErrorResponse
	// @Router /todos/{id} [get]
	router.GET("/todos/:id", getTodoByID)

	// @Summary Create a new todo
	// @Description Create a new todo task
	// @Tags Todo
	// @Accept json
	// @Produce json
	// @Param todo body Todo true "Create new todo"
	// @Success 201 {object} Todo
	// @Failure 400 {object} ErrorResponse
	// @Router /todos [post]
	router.POST("/todos", createTodo)

	// @Summary Update the status of a todo
	// @Description Update the status of a todo task
	// @Tags Todo
	// @Accept json
	// @Produce json
	// @Param id path int true "Todo ID"
	// @Param todo body Todo true "Updated Todo"
	// @Success 200 {object} Todo
	// @Failure 404 {object} ErrorResponse
	// @Router /todos/{id} [put]
	router.PUT("/todos/:id", updateTodoStatus)

	// @Summary Delete a todo by its ID
	// @Description Delete a todo task by its ID
	// @Tags Todo
	// @Accept json
	// @Produce json
	// @Param id path int true "Todo ID"
	// @Success 200 {object} DeleteResponse
	// @Failure 404 {object} ErrorResponse
	// @Router /todos/{id} [delete]
	router.DELETE("/todos/:id", deleteTodoByID)

	router.Run(":8080")
}

// getTodos returns the list of all todos
// @Summary Get all todos
// @Description Get all todos from the list
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// getTodoByID returns a todo by ID
// @Summary Get a todo by ID
// @Description Get a todo by its ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [get]
func getTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Todo not found"})
}

// createTodo creates a new todo
// @Summary Create a new todo
// @Description Create a new todo task
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body Todo true "Create new todo"
// @Success 201 {object} Todo
// @Failure 400 {object} ErrorResponse
// @Router /todos [post]
func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// updateTodoStatus updates the status of a todo
// @Summary Update the status of a todo
// @Description Update the status of a todo task
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body Todo true "Updated Todo"
// @Success 200 {object} Todo
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [put]
func updateTodoStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Status = updatedTodo.Status
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Todo not found"})
}

// deleteTodoByID deletes a todo by its ID
// @Summary Delete a todo by its ID
// @Description Delete a todo task by its ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} DeleteResponse
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [delete]
func deleteTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, DeleteResponse{Message: "Todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, ErrorResponse{Message: "Todo not found"})
}

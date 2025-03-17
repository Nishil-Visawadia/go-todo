package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "your_module_name/docs"  // Ensure the correct import path for your docs
    "net/http"
    "strconv"
)

// Todo struct represents a task in the Todo list
// @Description Todo represents a task in the list
type Todo struct {
    ID     int    `json:"id"`     // The ID of the todo task
    Task   string `json:"task"`   // The name or description of the task
    Status string `json:"status"` // Status of the task: "pending", "completed"
}

var todos = []Todo{
    {ID: 1, Task: "Learn Go", Status: "pending"},
    {ID: 2, Task: "Learn Gin", Status: "completed"},
}

// @Title Todo API
// @Description This is a simple API to manage a Todo list
// @Contact name="Your Name" url="https://yourwebsite.com" email="your-email@example.com"
// @Version 1.0
// @BasePath /
func main() {
    router := gin.Default()
    router.Use(cors.Default())

    // Set up Swagger documentation
    docs.SwaggerInfo.BasePath = "/"
    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
    // @Failure 404 {object} gin.H{"message": "Todo not found"}
    // @Router /todos/{id} [get]
    router.GET("/todos/:id", getTodoByID)

    // @Summary Create a new todo
    // @Description Create a new todo task
    // @Tags Todo
    // @Accept json
    // @Produce json
    // @Param todo body Todo true "Create new todo"
    // @Success 201 {object} Todo
    // @Failure 400 {object} gin.H{"error": "Invalid input"}
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
    // @Failure 404 {object} gin.H{"message": "Todo not found"}
    // @Router /todos/{id} [put]
    router.PUT("/todos/:id", updateTodoStatus)

    // @Summary Delete a todo by its ID
    // @Description Delete a todo task by its ID
    // @Tags Todo
    // @Accept json
    // @Produce json
    // @Param id path int true "Todo ID"
    // @Success 200 {object} gin.H{"message": "Todo deleted"}
    // @Failure 404 {object} gin.H{"message": "Todo not found"}
    // @Router /todos/{id} [delete]
    router.DELETE("/todos/:id", deleteTodoByID)

    // Start the server
    router.Run(":8080")
}

func getTodos(c *gin.Context) {
    c.JSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, todo := range todos {
        if todo.ID == id {
            c.JSON(http.StatusOK, todo)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func createTodo(c *gin.Context) {
    var newTodo Todo
    if err := c.ShouldBindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTodo.ID = len(todos) + 1
    todos = append(todos, newTodo)
    c.JSON(http.StatusCreated, newTodo)
}

func updateTodoStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedTodo Todo
    if err := c.ShouldBindJSON(&updatedTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for i, todo := range todos {
        if todo.ID == id {
            todos[i].Status = updatedTodo.Status
            c.JSON(http.StatusOK, todos[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func deleteTodoByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, todo := range todos {
        if todo.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

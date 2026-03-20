package controllers

import ( 
    "net/http"
    "github.com/gin-gonic/gin"
)

// Task represents a task object.
type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

var tasks = []Task{}

// CreateTask creates a new task.
func CreateTask(c *gin.Context) {
    var newTask Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task data"})
        return
    }
    newTask.ID = generateID()  // Assume generateID() is a function to generate unique IDs
    tasks = append(tasks, newTask)
    c.JSON(http.StatusCreated, newTask)
}

// GetAllTasks retrieves all tasks.
func GetAllTasks(c *gin.Context) {
    c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID.
func GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    for _, task := range tasks {
        if task.ID == id {
            c.JSON(http.StatusOK, task)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// UpdateTask updates an existing task.
func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    for i, task := range tasks {
        if task.ID == id {
            var updatedTask Task
            if err := c.ShouldBindJSON(&updatedTask); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task data"})
                return
            }
            updatedTask.ID = id
            tasks[i] = updatedTask
            c.JSON(http.StatusOK, updatedTask)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// DeleteTask deletes a task by its ID.
func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// GetTaskStatistics retrieves statistics about tasks.
func GetTaskStatistics(c *gin.Context) {
    totalTasks := len(tasks)
    completedTasks := 0
    for _, task := range tasks {
        if task.Completed {
            completedTasks++
        }
    }
    c.JSON(http.StatusOK, gin.H{"total": totalTasks, "completed": completedTasks})
}

// CompleteTask marks a task as completed.
func CompleteTask(c *gin.Context) {
    id := c.Param("id")
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Completed = true
            c.JSON(http.StatusOK, tasks[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Assume generateID function implementation goes here.

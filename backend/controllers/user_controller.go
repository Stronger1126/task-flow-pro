package controllers

import (
\t"net/http"
\t"github.com/gin-gonic/gin"
\t"stronger1126/task-flow-pro/models"
) 

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
\tvar user models.User
\tif err := c.ShouldBindJSON(&user); err != nil {
\t\tc.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
\t\treturn
\t}
\t// Logic to save user in database
\tc.JSON(http.StatusCreated, user)
}

// GetAllUsers retrieves all users
func GetAllUsers(c *gin.Context) {
\tvar users []models.User
\t// Logic to retrieve users from database
\tc.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
\tid := c.Param("id")
\tvar user models.User
\t// Logic to retrieve user by id from database
\tc.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
\tid := c.Param("id")
\tvar user models.User
\tif err := c.ShouldBindJSON(&user); err != nil {
\t\tc.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
\t\treturn
\t}
\t// Logic to update user in database
\tc.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
\tid := c.Param("id")
\t// Logic to delete user from database
\tc.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
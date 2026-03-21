package controllers

import (
	"golang_rest_orm_unit_test/database"
	"golang_rest_orm_unit_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseError represents an error response
type ResponseError struct {
	Error string `json:"error" example:"User not found"`
}

// ResponseMessage represents a simple message response
type ResponseMessage struct {
	Message string `json:"message" example:"User deleted"`
}

const UserNotFoundMessage = "User not found"

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user in the database. The ID field is auto-generated and should not be provided in the request body.
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User payload" example({"name":"John Doe","email":"john@example.com"})
// @Success 201 {object} models.User "User created successfully"
// @Failure 400 {object} ResponseError "Invalid request body"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Returns a single user matching the given ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID" minimum(1)
// @Success 200 {object} models.User "User found"
// @Failure 404 {object} ResponseError "User not found"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: UserNotFoundMessage})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary List all users
// @Description Returns a list of all registered users. Returns an empty array if no users exist.
// @Tags users
// @Produce json
// @Success 200 {array} models.User "List of users"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Updates the name and/or email of an existing user. The user must exist or a 404 is returned.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" minimum(1)
// @Param user body models.User true "Updated user payload" example({"name":"Jane Smith","email":"jane.smith@example.com"})
// @Success 200 {object} models.User "User updated successfully"
// @Failure 400 {object} ResponseError "Invalid request body"
// @Failure 404 {object} ResponseError "User not found"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: UserNotFoundMessage})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Permanently removes a user from the database. The user must exist or a 404 is returned.
// @Tags users
// @Produce json
// @Param id path int true "User ID" minimum(1)
// @Success 200 {object} ResponseMessage "User deleted successfully"
// @Failure 404 {object} ResponseError "User not found"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: UserNotFoundMessage})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{Message: "User deleted"})
}

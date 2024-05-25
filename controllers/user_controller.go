package controllers

import (
	"golang_rest_orm_unit_test/database"
	"golang_rest_orm_unit_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseError represents an error response
type ResponseError struct {
	Error string `json:"error"`
}

// ResponseMessage represents a simple message response
type ResponseMessage struct {
	Message string `json:"message"`
}

const UserNotFoundMessage = "User not found"

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User to create"
// @Success 200 {object} models.User
// @Failure 400 {object} ResponseError
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} ResponseError
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
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update a user by ID with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User to update"
// @Success 200 {object} models.User
// @Failure 400 {object} ResponseError
// @Failure 404 {object} ResponseError
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

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} ResponseMessage
// @Failure 404 {object} ResponseError
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Delete(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{Message: "User deleted"})
}

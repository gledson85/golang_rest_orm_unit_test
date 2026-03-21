package test

import (
	"encoding/json"
	"fmt"
	"golang_rest_orm_unit_test/controllers"
	"golang_rest_orm_unit_test/database"
	"golang_rest_orm_unit_test/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	database.SetupTestDB()
	m.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/users", controllers.GetAllUsers)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	return router
}

func cleanDB() {
	database.DB.Exec("DELETE FROM users")
}

func createTestUser(name, email string) models.User {
	user := models.User{Name: name, Email: email}
	database.DB.Create(&user)
	return user
}

// --- CreateUser ---

func TestCreateUser(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(`{"name":"John Doe","email":"john@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var user models.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(`{invalid json}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

// --- GetUser ---

func TestGetUser(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	user := createTestUser("Jane Doe", "jane@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, response.ID)
	assert.Equal(t, "Jane Doe", response.Name)
	assert.Equal(t, "jane@example.com", response.Email)
}

func TestGetUser_NotFound(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/99999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), controllers.UserNotFoundMessage)
}

// --- GetAllUsers ---

func TestGetAllUsers(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	createTestUser("Jane Doe", "jane@example.com")
	createTestUser("John Doe", "john@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGetAllUsers_Empty(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Empty(t, users)
}

// --- UpdateUser ---

func TestUpdateUser(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	user := createTestUser("Jane Doe", "jane@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", user.ID), strings.NewReader(`{"name":"Jane Smith","email":"jane.smith@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Smith", response.Name)
	assert.Equal(t, "jane.smith@example.com", response.Email)
}

func TestUpdateUser_NotFound(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/99999", strings.NewReader(`{"name":"Ghost","email":"ghost@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), controllers.UserNotFoundMessage)
}

func TestUpdateUser_InvalidJSON(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	user := createTestUser("Jane Doe", "jane@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", user.ID), strings.NewReader(`{invalid json}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

// --- DeleteUser ---

func TestDeleteUser(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	user := createTestUser("Jane Doe", "jane@example.com")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User deleted")

	// Verifica que o usuário foi realmente removido
	var count int64
	database.DB.Model(&models.User{}).Where("id = ?", user.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestDeleteUser_NotFound(t *testing.T) {
	cleanDB()
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/99999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), controllers.UserNotFoundMessage)
}

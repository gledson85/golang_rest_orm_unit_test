package test

import (
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

func TestCreateUser(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(`{"name":"John Doe", "email":"john@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"John Doe"`)
	assert.Contains(t, w.Body.String(), `"email":"john@example.com"`)
}

func TestGetUser(t *testing.T) {
	router := SetupRouter()

	// Criar o usuário no banco de dados
	user := models.User{Name: "Jane Doe", Email: "jane@example.com"}
	result := database.DB.Create(&user)
	assert.Nil(t, result.Error) // Verifica se não houve erro na criação
	assert.NotZero(t, user.ID)  // Verifica se o ID foi preenchido

	// Realizar a requisição GET para buscar o usuário criado
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/"+strings.TrimSpace(fmt.Sprintf("%d", user.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Jane Doe"`)
	assert.Contains(t, w.Body.String(), `"email":"jane@example.com"`)
}

func TestGetAllUsers(t *testing.T) {
	router := SetupRouter()
	database.DB.Create(&models.User{Name: "Jane Doe", Email: "jane@example.com"})
	database.DB.Create(&models.User{Name: "John Doe", Email: "john@example.com"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Jane Doe"`)
	assert.Contains(t, w.Body.String(), `"name":"John Doe"`)
}

func TestUpdateUser(t *testing.T) {
	router := SetupRouter()
	database.DB.Create(&models.User{Name: "Jane Doe", Email: "jane@example.com"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(`{"name":"Jane Smith", "email":"jane.smith@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Jane Smith"`)
	assert.Contains(t, w.Body.String(), `"email":"jane.smith@example.com"`)
}

func TestDeleteUser(t *testing.T) {
	router := SetupRouter()
	database.DB.Create(&models.User{Name: "Jane Doe", Email: "jane@example.com"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `User deleted`)
}

package main

import (
	"fmt"
	"golang_rest_orm_unit_test/controllers"
	"golang_rest_orm_unit_test/database"
	_ "golang_rest_orm_unit_test/docs"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title User API
// @version 1.0
// @description REST API for user management with CRUD operations.
// @contact.name API Support
// @contact.url https://github.com/gleds/golang_rest_orm_unit_test
// @license.name MIT
// @host localhost:8081
// @BasePath /
// @schemes http
func main() {
	// Initialize the database
	database.InitDatabase()

	// Load server configuration
	var config struct {
		Server struct {
			Port int
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Swagger endpoint
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	r.GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/" {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			return
		}
		swaggerHandler(c)
	})

	serverPort := fmt.Sprintf(":%d", config.Server.Port)
	if err := r.Run(serverPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

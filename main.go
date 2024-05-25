package main

import (
	"fmt"
	"golang_rest_orm_unit_test/controllers"
	"golang_rest_orm_unit_test/database"
	_ "golang_rest_orm_unit_test/docs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title User API
// @version 1.0
// @description This is a sample server for a user management API.
// @host localhost:8081
// @BasePath /
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
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}

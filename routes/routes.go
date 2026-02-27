package routes

import (
	"backend1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/register", handlers.Register)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.POST("/login", handlers.Login)
}
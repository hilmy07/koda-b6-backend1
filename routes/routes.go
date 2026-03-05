package routes

import (
	"backend1/handlers"

	docs "backend1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	// user
	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/register", handlers.Register)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.POST("/login", handlers.Login)

	// product
	r.GET("/products", handlers.GetAllProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	docs.SwaggerInfo.BasePath = "/"

	docPath := r.Group("/docs")
	{
		docPath.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
package main

import (
	"backend1/routes"

	"github.com/gin-gonic/gin"
)

// @title             Backend Apps
// @version           1.0.0
// @description       This is basic backend apps
// @host              localhost:8000
// @BasePath          /

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run("localhost:8000")
}
package main

import (
	"github.com/FabricioCes/cabys-api-go/internal/cabys"
	"github.com/FabricioCes/cabys-api-go/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://tu-dominio-frontend.com"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	data := utils.LoadCabysData("data/cabys.json")
	handler := cabys.NewHandler(data)

	r.GET("/api/cabys", handler.Search)
	r.GET("/api/cabys/:id", handler.GetByID)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}

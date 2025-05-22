package main

import (
	"cabys-api-go/internal/cabys"
	"cabys-api-go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	data := utils.LoadCabysData("data/cabys.json")
	handler := cabys.NewHandler(data)

	r.GET("/api/cabys", handler.Search)
	r.GET("/api/cabys/:id", handler.GetByID)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}

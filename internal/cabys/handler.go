package cabys

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Data []CabysCode
}

func NewHandler(data []CabysCode) *Handler {
	return &Handler{Data: data}
}

func (h *Handler) Search(c *gin.Context) {
	query := strings.ToLower(c.Query("query"))

	// Si no hay query, devolver todos los datos
	if query == "" {
		c.JSON(http.StatusOK, h.Data)
		return
	}

	var results []CabysCode
	for _, code := range h.Data {
		if strings.Contains(strings.ToLower(code.Description), query) {
			results = append(results, code)
		}
	}

	c.JSON(http.StatusOK, results)
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	for _, code := range h.Data {
		if code.ID == id {
			c.JSON(http.StatusOK, code)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Código CABYS no encontrado"})
}

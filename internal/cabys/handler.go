package cabys

import (
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Data []CabysCode
}

func NewHandler(data []CabysCode) *Handler {
	return &Handler{Data: data}
}

// normalize remueve tildes y pasa a minúscula
func normalize(input string) string {
	// Forma NFD separa caracteres base y tildes
	t := norm.NFD.String(input)
	var result []rune
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue // ignora marcas diacríticas (tildes)
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func (h *Handler) Search(c *gin.Context) {
	query := normalize(c.Query("query"))

	if query == "" {
		c.JSON(http.StatusOK, h.Data)
		return
	}

	var results []CabysCode
	queryWords := strings.Fields(query)

	// Buscar por descripción (todas las palabras deben coincidir)
	for _, code := range h.Data {
		desc := normalize(code.Description)

		match := true
		for _, word := range queryWords {
			if !strings.Contains(desc, word) {
				match = false
				break
			}
		}

		if match {
			results = append(results, code)
		}
	}

	// Si no encontró nada, buscar por ID
	if len(results) == 0 {
		for _, code := range h.Data {
			if strings.Contains(strings.ToLower(code.ID), query) {
				results = append(results, code)
			}
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

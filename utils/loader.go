package utils

import (
	"encoding/json"
	"log"
	"os"

	"cabys-api-go/internal/cabys"
)

func LoadCabysData(path string) []cabys.CabysCode {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("❌ Error al leer el archivo JSON: %v", err)
	}

	var data []cabys.CabysCode
	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatalf("❌ Error al parsear el archivo JSON: %v", err)
	}

	log.Printf("✅ Se cargaron %d códigos CABYS en memoria", len(data))
	return data
}

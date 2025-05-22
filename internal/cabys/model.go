package cabys

type CabysCode struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	TaxPercent  float64 `json:"tax_percent"`
}

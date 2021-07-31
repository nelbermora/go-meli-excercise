package domain

import "time"

type ProdcutBatch struct {
	ID                int       `json:"id"`
	BatchNumber       string    `json:"batch_number"`
	CurrentQuantity   int       `json:"current_quantity"`
	CurrentTemp       float32   `json:"current_temperature"`
	DueDate           time.Time `json:"due_date"`
	InitialQuantity   int       `json:"initial_quantity"`
	ManufacturingDate time.Time `json:"manufacturing_date"`
	MinTemperature    float32   `json:"minimum_temperature"`
	ProductId         int       `json:"product_id"`
	SectionId         int       `json:"section_id"`
}

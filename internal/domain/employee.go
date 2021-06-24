package domain

type Employee struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WarehouseID int    `json:"warehouse_id"`
}

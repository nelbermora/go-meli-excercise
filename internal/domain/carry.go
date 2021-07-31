package domain

type Carry struct {
	ID        int    `json:"id"`
	Batch     int    `json:"batch_number"`
	Cid       string `json:"cid"`
	Company   string `json:"company_name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Locality  int    `json:"locality_id"`
}

package domain

type Locality struct {
	ID       int    `json:"locality_id"`
	Name     string `json:"locality_name"`
	Province string `json:"province_name"`
	Country  string `json:"country_name"`
	Sellers  *int   `json:"sellers_count,omitempty"`
	Carries  *int   `json:"carries_count,omitempty"`
}

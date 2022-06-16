package products_repo

type ProductsRetailers struct {
	PName    string  `json:"p_name"`
	RName    string  `json:"r_name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

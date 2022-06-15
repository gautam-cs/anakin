package service

type ProductsRetailers struct {
	Pname    string  `json:"pname"`
	Rname    string  `json:"rname"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

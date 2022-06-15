package query

type SignUpRequest struct {
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	UserName  string  `json:"user_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type ProductRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type PromotionsRequest struct {
	ProductID  int64   `json:"product_id"`
	RetailerID int64   `json:"retailer_id"`
	Discount   float64 `json:"discount"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
}

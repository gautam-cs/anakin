package models

type ProductRetailer struct {
	ID           int64 `gorm:"primaryKey"`
	UUID         string
	ProductID    int64
	RetailerID   int64
	Price        float64
	Quantity     int64
	CreatedDate  DateTime
	ModifiedDate DateTime

	Retailers *Retailers `gorm:"foreignKey:RetailerID;references:ID"`
	Products  *Products  `gorm:"foreignKey:ProductID;references:ID"`
}

func (ProductRetailer) TableName() string {
	return "products_retailers"
}

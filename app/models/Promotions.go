package models

type Promotions struct {
	ID           int64 `gorm:"primaryKey"`
	UUID         string
	ProductID    int64
	RetailerID   int64
	Discount     float64
	StartTime    *DateTime
	EndTime      *DateTime
	IsActive     int
	CreatedDate  DateTime
	ModifiedDate DateTime

	Retailers *Retailers `gorm:"foreignKey:RetailerID;references:ID"`
	Products  *Products  `gorm:"foreignKey:ProductID;references:ID"`
}

func (Promotions) TableName() string {
	return "promotions"
}

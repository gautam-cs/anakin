package models

type Products struct {
	ID           int64 `gorm:"primaryKey"`
	UUID         string
	Name         string
	Brand        string
	CreatedDate  DateTime
	ModifiedDate DateTime
}

func (Products) TableName() string {
	return "products"
}

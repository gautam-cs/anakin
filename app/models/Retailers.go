package models

type Retailers struct {
	ID           int64 `gorm:"primaryKey"`
	UUID         string
	Email        string
	Name         string
	CreatedDate  DateTime
	ModifiedDate DateTime
}

func (Retailers) TableName() string {
	return "retailers"
}

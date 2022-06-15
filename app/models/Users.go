package models

type Users struct {
	ID           int64 `gorm:"primaryKey"`
	UUID         string
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	CreatedDate  DateTime
	ModifiedDate DateTime
}

func (Users) TableName() string {
	return "users"
}

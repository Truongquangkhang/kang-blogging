package model

type Role struct {
	ID   string `gorm:"primaryKey;column:id;default:UUID()"`
	Name string `gorm:"column:name"`
}

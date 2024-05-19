package model

type Category struct {
	ID   string `gorm:"primaryKey;column:id;default:UUID()"`
	Name string `gorm:"column:name"`
}

func (Category) TableName() string {
	return "categories"
}

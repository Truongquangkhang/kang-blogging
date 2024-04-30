package model

type BlogCategories struct {
	ID         string `gorm:"primaryKey;column:id;default:UUID()"`
	BlogID     string `gorm:"column:blog_id"`
	CategoryID string `gorm:"column:category_id"`
}

func (BlogCategories) TableName() string {
	return "blog_categories"
}

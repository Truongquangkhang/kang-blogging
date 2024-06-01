package model

import "time"

type Category struct {
	ID          string    `gorm:"primaryKey;column:id;default:UUID()"`
	Name        string    `gorm:"column:name"`
	Description *string   `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

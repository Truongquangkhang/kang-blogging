package model

import (
	"time"
)

// User represents the MySQL table structure
type User struct {
	ID          string     `gorm:"type:varchar(40);primaryKey;not null;default:uuid();column:id"`
	AccountID   string     `gorm:"type:varchar(40);index;not null;column:account_id"`
	RoleID      string     `gorm:"type:varchar(40);index;not null;column:role_id"`
	Name        string     `gorm:"type:varchar(255);not null;column:name"`
	Email       string     `gorm:"type:varchar(255);not null;column:email"`
	PhoneNumber *string    `gorm:"type:varchar(255);column:phone_number"`
	CreatedAt   time.Time  `gorm:"not null;default:now();column:created_at"`
	UpdatedAt   time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}

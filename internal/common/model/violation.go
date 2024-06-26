package model

import "time"

type Violation struct {
	ID                string    `gorm:"primaryKey;column:id;default:UUID()"`
	UserID            string    `gorm:"not null;column:user_id"`
	ViolationType     string    `gorm:"not null;column:violation_type"`
	ViolationTargetID string    `gorm:"not null;column:violation_target_id"`
	Description       *string   `gorm:"column:description"`
	CreatedAt         time.Time `gorm:"not null;default:now();column:created_at"`
	User              *User     `gorm:"foreignkey:UserID"`
}

func (Violation) TableName() string {
	return "violations"
}

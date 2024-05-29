package model

import "time"

type Comment struct {
	ID             string     `gorm:"primaryKey;column:id;default:UUID()"`
	UserID         string     `gorm:"column:user_id"`
	BlogID         string     `gorm:"column:blog_id"`
	Content        string     `gorm:"column:content"`
	IsToxicity     bool       `gorm:"column:is_toxicity"`
	Level          int        `gorm:"column:level"`
	ReplyCommentID *string    `gorm:"column:reply_comment_id"`
	CreatedAt      time.Time  `gorm:"not null;default:now();column:created_at"`
	UpdatedAt      time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"`
	User           User       `gorm:"foreignKey:UserID;references:ID"`
	Blog           Blog       `gorm:"foreignKey:BlogID;references:ID"`
}

func (Comment) TableName() string {
	return "comments"
}

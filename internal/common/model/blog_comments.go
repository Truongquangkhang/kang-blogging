package model

type BlogComment struct {
	ID        string `gorm:"primaryKey;column:id;default:UUID()"`
	BlogID    string `gorm:"column:blog_id"`
	CommentID string `gorm:"column:comment_id"`
}

func (BlogComment) TableName() string {
	return "blog_comments"
}

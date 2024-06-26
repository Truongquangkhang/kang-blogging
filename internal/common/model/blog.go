package model

import (
	"time"
)

type Blog struct {
	ID                string     `gorm:"type:varchar(40);primaryKey;not null;default:uuid();column:id"`
	AuthorID          string     `gorm:"type:varchar(40);index;not null;column:author_id"`
	Title             string     `gorm:"type:varchar(255);not null;column:title"`
	Summary           *string    `gorm:"type:varchar(255);column:summary"`
	Thumbnail         *string    `gorm:"type:varchar(255);column:thumbnail"`
	Content           *string    `gorm:"type:text;column:content"`
	Published         bool       `gorm:"type:bool;column:published;default:false"`
	IsDeprecated      bool       `gorm:"type:bool;column:is_deprecated;default:false"`
	CreatedAt         time.Time  `gorm:"not null;default:now();column:created_at"`
	UpdatedAt         time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at"`
	User              *User      `gorm:"foreignkey:AuthorID"`
	Categories        []Category `gorm:"many2many:blog_categories;"`
	TotalBlogComments *int32
}

func (Blog) TableName() string {
	return "blogs"
}

package model

import (
	"time"
)

// User represents the MySQL table structure
type User struct {
	ID                string     `gorm:"type:varchar(40);primaryKey;not null;default:uuid();column:id"`
	RoleID            string     `gorm:"type:varchar(40);index;not null;column:role_id"`
	Name              string     `gorm:"type:varchar(255);not null;column:name"`
	Email             string     `gorm:"type:varchar(255);not null;column:email"`
	Description       *string    `gorm:"type:text;column:description"`
	PhoneNumber       *string    `gorm:"type:varchar(255);column:phone_number"`
	DisplayName       string     `gorm:"type:varchar(255);column:display_name"`
	Gender            *bool      `gorm:"type:bool;column:gender;default:false"`
	Avatar            *string    `gorm:"type:varchar(255);column:avatar"`
	BirthOfDay        *int64     `gorm:"type:bigint;column:birth_of_day"`
	IsActive          bool       `gorm:"type:bool;column:is_active;default:true"`
	TotalViolation    int32      `gorm:"type:int;column:total_violation;default:0"`
	ExpireWarningTime *int64     `gorm:"type:bigint;column:expire_warning_time"`
	CreatedAt         time.Time  `gorm:"not null;default:now();column:created_at"`
	UpdatedAt         time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at"`
	Blogs             []Blog     `gorm:"foreignKey:AuthorID;references:ID"`
	Comments          []Comment  `gorm:"foreignKey:UserID;references:ID"`
	TotalBlogs        int32
	TotalComments     int32
	TotalFollowers    int32
	TotalFolloweds    int32
	IsFollower        bool
	IsFollowed        bool
}

func (User) TableName() string {
	return "users"
}

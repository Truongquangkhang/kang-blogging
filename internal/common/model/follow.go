package model

import "time"

type Follow struct {
	FollowerID string    `gorm:"primaryKey;column:follower_id"`
	FollowedID string    `gorm:"primaryKey;column:followed_id"`
	CreatedAt  time.Time `gorm:"not null;default:now();column:created_at"`
}

func (Follow) TableName() string {
	return "follows"
}

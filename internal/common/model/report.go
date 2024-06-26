package model

import "time"

type Report struct {
	ID             string    `gorm:"primaryKey;column:id;default:UUID()"`
	ReporterID     string    `gorm:"not null;column:reporter_id"`
	ReportType     string    `gorm:"not null;column:report_type"`
	ReportTargetID string    `gorm:"not null;column:report_target_id"`
	Reason         string    `gorm:"not null;column:reason"`
	Description    *string   `gorm:"column:description"`
	IsClosed       bool      `gorm:"column:is_closed;default:false"`
	CreatedAt      time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at"`
	User           *User     `gorm:"foreignkey:ReporterID"`
}

func (Report) TableName() string {
	return "reports"
}

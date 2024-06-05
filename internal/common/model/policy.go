package model

type Policy struct {
	ID    string  `gorm:"primaryKey;column:id;default:UUID()"`
	Type  *string `gorm:"column:type"`
	Value *int64  `gorm:"column:value"`
}

func (Policy) TableName() string {
	return "policies"
}

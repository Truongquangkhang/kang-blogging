package model

type Account struct {
	ID       string `gorm:"primaryKey;column:id;default:UUID()"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (Account) TableName() string {
	return "accounts"
}

package model

type BillPO struct {
	ID         string `gorm:"column:id"`
	FromUserID string `gorm:"column:from_user_id"`
	ToUserID   string `gorm:"column:to_user_id"`
	Amount     string `gorm:"column:amount"`
	Currency   string `gorm:"column:currency"`
}

func (BillPO) TableName() string {
	return "bill"
}

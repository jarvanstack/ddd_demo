package model

import (
	"time"

	user_model "ddd_demo/internal/user/model"
)

const (
	DefaultBillIDValue = "0"
)

type BillID struct {
	value string
}

func NewBillID(billID string) (*BillID, error) {
	// 省略参数检查
	return &BillID{
		value: billID,
	}, nil
}

func (b *BillID) Value() string {
	if b == nil {
		return DefaultBillIDValue
	}

	return b.value
}

type Bill struct {
	ID        *BillID
	FromUser  *user_model.User
	ToUser    *user_model.User
	Amount    *user_model.Amount
	Rate      *user_model.Rate
	CreatedAt time.Time
}

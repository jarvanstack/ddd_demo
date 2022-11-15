package model

import (
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
	ID         *BillID
	FromUserID *user_model.UserID
	ToUserID   *user_model.UserID
	Amount     *user_model.Amount
	Currency   *user_model.Currency
}

func (b *Bill) ToPO() *BillPO {
	return &BillPO{
		ID:         b.ID.Value(),
		FromUserID: b.FromUserID.Value(),
		ToUserID:   b.ToUserID.Value(),
		Amount:     b.Amount.Value().String(),
		Currency:   b.Currency.Value(),
	}
}

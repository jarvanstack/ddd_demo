package model

import (
	"errors"
	"strconv"

	"github.com/shopspring/decimal"
)

var (
	ErrUserIDIsEmpty = errors.New("user id is empty")
)

// po (presentation object) 持久化对象

type UserPO struct {
	ID       int64
	Username string
	Password string
	Currency string
	Amount   decimal.Decimal `gorm:"type:decimal(20,2);"`
}

func (UserPO) TableName() string {
	return "user"
}

// ToDomain converts a UserRepo to a domain.User
func (u *UserPO) ToDomain() (*User, error) {
	user := &User{}
	if u.ID == 0 {
		return nil, ErrUserIDIsEmpty
	}

	user.ID, _ = NewUserID(strconv.FormatInt(u.ID, 10))

	if u.Username != "" {
		user.Username, _ = NewUsername(u.Username)
	}

	if u.Password != "" {
		user.Password, _ = NewPassword(u.Password)
	}

	if u.Currency != "" {
		user.Currency, _ = NewCurrency(u.Currency)
	}

	if !u.Amount.IsZero() {
		user.Amount, _ = NewAmount(u.Amount)
	}

	return user, nil
}

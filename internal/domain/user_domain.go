package domain

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// domain 领域对象

type UserID struct {
	value string
}

func NewUserID(userID string) (*UserID, error) {
	// 省略参数检查
	return &UserID{
		value: userID,
	}, nil
}

func (u *UserID) Value() string {
	if u == nil {
		return ""
	}

	return u.value
}

type Username struct {
	value string
}

func NewUsername(username string) (*Username, error) {
	// 省略参数检查
	return &Username{
		value: username,
	}, nil
}

func (u *Username) Value() string {
	if u == nil {
		return ""
	}

	return u.value
}

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	// 省略参数检查
	return &Password{
		value: password,
	}, nil
}

func (u *Password) Value() string {
	if u == nil {
		return ""
	}

	return u.value
}

type Currency struct {
	value string
}

func NewCurrency(currency string) (*Currency, error) {
	// 省略参数检查
	return &Currency{
		value: currency,
	}, nil
}

func (u *Currency) Value() string {
	if u == nil {
		return ""
	}

	return u.value
}

type Balance struct {
	value decimal.Decimal
}

func NewBalance(balance decimal.Decimal) (*Balance, error) {
	// 省略参数检查
	return &Balance{
		value: balance,
	}, nil
}

func (u *Balance) Value() decimal.Decimal {
	if u == nil {
		return decimal.NewFromFloat(0)
	}

	return u.value
}

type User struct {
	ID       *UserID
	Username *Username
	Password *Password
	Currency *Currency
	Balance  *Balance
}

func (u *User) ToLoginResp(token string) *S2C_Login {
	return &S2C_Login{
		UserID:   u.ID.Value(),
		Username: u.Username.Value(),
		Token:    token,
	}
}

func (u *User) ToUserInfo() *S2C_UserInfo {
	return &S2C_UserInfo{
		UserID:   u.ID.Value(),
		Username: u.Username.Value(),
		Balance:  u.Balance.Value().String(),
		Currency: u.Currency.Value(),
	}
}

func (u *User) ToPO() *UserPO {
	id, _ := strconv.ParseInt(u.ID.Value(), 10, 64)
	return &UserPO{
		ID:       id,
		Username: u.Username.Value(),
		Password: u.Password.Value(),
		Currency: u.Currency.Value(),
		Balance:  u.Balance.Value(),
	}
}

type LoginParams struct {
	Username *Username
	Password *Password
}

type RegisterParams struct {
	Username *Username
	Password *Password
}

func (c *RegisterParams) ToDomain() *User {
	return &User{
		Username: c.Username,
		Password: c.Password,
	}
}

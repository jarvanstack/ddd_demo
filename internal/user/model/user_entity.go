package model

import (
	"errors"
	"strconv"

	"github.com/shopspring/decimal"
)

// domain 领域对象

var (
	DefaultUserIDValue   = "0"
	DefaultUsernameValue = ""
	DefaultPasswordValue = ""
	DefaultCurrencyValue = "CNY"
	DefaultAmountValue   = decimal.NewFromFloat(0)
	DefaultFeeValue, _   = NewAmount(decimal.NewFromFloat(0))
)

var (
	ErrAmountNotEnough = errors.New("余额不足")
)

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
		return DefaultUserIDValue
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
		return DefaultUsernameValue
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
		return DefaultPasswordValue
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
		return DefaultCurrencyValue
	}

	return u.value
}

type Amount struct {
	value decimal.Decimal
}

func NewAmount(amount decimal.Decimal) (*Amount, error) {
	// 省略参数检查
	return &Amount{
		value: amount,
	}, nil
}

func (m *Amount) Value() decimal.Decimal {
	if m == nil {
		return DefaultAmountValue
	}

	return m.value
}

func (m *Amount) Add(amount *Amount) *Amount {
	return &Amount{
		value: m.value.Add(amount.value),
	}
}

type User struct {
	ID       *UserID
	Username *Username
	Password *Password
	Currency *Currency
	Amount   *Amount
}

func (u *User) CalcFee(fromAmount *Amount) (*Amount, error) {
	return NewAmount(fromAmount.Value().Mul(DefaultFeeValue.Value()))
}

// 付款
func (u *User) Pay(amount *Amount) error {
	// 省略参数检查
	if u.Amount.Value().LessThan(amount.Value()) {
		return ErrAmountNotEnough
	}

	u.Amount.value = u.Amount.Value().Sub(amount.Value())

	return nil
}

// 收款
func (u *User) Receive(amount *Amount) error {
	// 省略参数检查

	u.Amount.value = u.Amount.value.Add(amount.value)

	return nil
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
		Amount:   u.Amount.Value().String(),
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
		Amount:   u.Amount.Value(),
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

type Rate struct {
	rate decimal.Decimal
}

func NewRate(rate decimal.Decimal) (*Rate, error) {
	// 省略参数检查
	return &Rate{
		rate: rate,
	}, nil
}

func (r *Rate) Exchange(amount *Amount) (*Amount, error) {
	return NewAmount(amount.Value().Mul(r.rate))
}

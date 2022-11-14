package domain

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
	return u.value
}

type User struct {
	ID       *UserID
	Username *Username
	Password *Password
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
	}
}

func (u *User) ToPO() *UserPO {
	return &UserPO{
		ID:       u.ID.Value(),
		Username: u.Username.Value(),
		Password: u.Password.Value(),
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

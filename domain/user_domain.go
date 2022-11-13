package domain

// domain 领域对象

type User struct {
	ID       string
	Username string
	Password string
}

func (u *User) ToLoginResp(token string) *S2C_Login {
	return &S2C_Login{
		UserID:   u.ID,
		Username: u.Username,
		Token:    token,
	}
}

func (u *User) ToUserInfo() *S2C_UserInfo {
	return &S2C_UserInfo{
		UserID:   u.ID,
		Username: u.Username,
	}
}

func (u *User) ToPO() *UserPO {
	return &UserPO{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}

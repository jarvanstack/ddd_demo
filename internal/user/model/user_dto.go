package model

// dto (data transfer object) 数据传输对象

// C2S_Login Web登录请求
type C2S_Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *C2S_Login) ToDomain() (*LoginParams, error) {
	username, err := NewUsername(c.Username)
	if err != nil {
		return nil, err
	}
	password, err := NewPassword(c.Password)
	if err != nil {
		return nil, err
	}

	return &LoginParams{
		Username: username,
		Password: password,
	}, nil
}

// S2C_Login Web登录响应
type S2C_Login struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type S2C_UserInfo struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type C2S_Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *C2S_Register) ToDomain() (*RegisterParams, error) {
	username, err := NewUsername(c.Username)
	if err != nil {
		return nil, err
	}
	password, err := NewPassword(c.Password)
	if err != nil {
		return nil, err
	}

	return &RegisterParams{
		Username: username,
		Password: password,
	}, nil
}

type C2S_Transfer struct {
	ToUserID string `json:"to_user_id"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

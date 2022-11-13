package domain

// dto (data transfer object) 数据传输对象

// C2S_Login Web登录请求
type C2S_Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *C2S_Login) Validate() error {
	// 省略参数检查
	return nil
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
}

type C2S_Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *C2S_Register) Validate() error {
	// 省略参数检查
	return nil
}

func (c *C2S_Register) ToDomain() *User {
	return &User{
		Username: c.Username,
		Password: c.Password,
	}
}

func ValidateUserID(userID string) error {
	// 省略参数检查
	return nil
}

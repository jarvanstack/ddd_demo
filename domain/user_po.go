package domain

// po (presentation object) 持久化对象

type UserPO struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (UserPO) TableName() string {
	return "user"
}

// ToDomain converts a UserRepo to a domain.User
func (u *UserPO) ToDomain() *User {
	return &User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}

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
func (u *UserPO) ToDomain() (*User, error) {
	username, err := NewUsername(u.Username)
	if err != nil {
		return nil, err
	}
	password, err := NewPassword(u.Password)
	if err != nil {
		return nil, err
	}
	userID, err := NewUserID(u.ID)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       userID,
		Username: username,
		Password: password,
	}, nil
}

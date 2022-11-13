package domain

import "encoding/json"

type AuthInfo struct {
	UserID string `json:"user_id"`
}

func (s *AuthInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *AuthInfo) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, s)
}

func NewAuth(userID string) (*AuthInfo, error) {
	// 省略参数检查
	return &AuthInfo{
		UserID: userID,
	}, nil
}

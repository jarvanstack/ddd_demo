package test

import "testing"

type User struct {
	value string
}

func (u *User) Value() string {
	if u == nil {
		return "默认值"
	}
	return u.value
}

func Test_Nil(t *testing.T) {
	var u *User
	t.Log(u.Value())
}

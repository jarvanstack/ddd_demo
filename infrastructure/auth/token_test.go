package auth

import (
	"ddd_demo/domain"
	"fmt"
	"testing"
	"time"
)

func TestTokenAuth_SetAndGet(t *testing.T) {
	auth := NewJwtAuth("123456", time.Hour*2)
	authInfo := &domain.AuthInfo{
		UserID: "123",
	}
	s, err := auth.Set(authInfo)
	fmt.Printf("err: %v\n", err)
	ai, err := auth.Get(s)
	fmt.Printf("ai: %v\n", ai)
	fmt.Printf("err: %v\n", err)
}

func TestTokenAuth_Get(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjgzMzE4MzMsInVzZXJuYW1lIjoiMSJ9.Gl5AOSmItstaTJ3blROH1lTtU1moLZqBkwJL9MvHUZI"
	auth := NewJwtAuth("123456", time.Hour*2)
	ai, err := auth.Get(token)
	fmt.Printf("ai: %v\n", ai)
	fmt.Printf("err: %v\n", err)
}

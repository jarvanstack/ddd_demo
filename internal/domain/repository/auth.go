package repository

import "ddd_demo/internal/domain"

// redis : cookie + session 认证
// redis 储存 + 返回 web cookie
// redis 获取 + 通过 cookie 返回 session

// token + refreshToken 认证
// struct 加密 string
// string 解密 struct

/*
Auth 使用的最佳方法是, 无感知的切换 Token 和 Cookie + Session
	app.Login() 返回 S2C_Login { token }
	handler 有 gin, 调用
*/

type AuthInterface interface {
	Set(*domain.AuthInfo) (string, error)
	Get(string) (*domain.AuthInfo, error)
	Del(string) error
}

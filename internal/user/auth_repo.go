package user

import (
	"context"
	"ddd_demo/internal/user/model"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

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
	Set(*model.AuthInfo) (string, error)
	Get(string) (*model.AuthInfo, error)
	Del(string) error
}

const (
	TokenKeyUserID     = "username"
	TokenKeyExpireTime = "exp"
)

var _ AuthInterface = &TokenAuth{}

type TokenAuth struct {
	priKey     string
	expireTime time.Duration
}

func NewJwtAuth(priKey string, expireTime time.Duration) AuthInterface {
	return &TokenAuth{
		priKey:     priKey,
		expireTime: expireTime,
	}
}

// Set 生成 token
func (r *TokenAuth) Set(auth *model.AuthInfo) (string, error) {
	// 对称加密 auth
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		TokenKeyUserID:     auth.UserID,
		TokenKeyExpireTime: time.Now().Add(r.expireTime).Unix(),
	})

	return token.SignedString([]byte(r.priKey))
}

func (r *TokenAuth) Get(token string) (*model.AuthInfo, error) {
	// 解密 token
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.priKey), nil
	})
	if err != nil {
		return nil, err
	}

	// 获取 token 中的数据
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// 判断 token 是否过期
	if time.Now().Unix() > int64(claims[TokenKeyExpireTime].(float64)) {
		return nil, err
	}

	return &model.AuthInfo{
		UserID: claims[TokenKeyUserID].(string),
	}, nil
}

func (r *TokenAuth) Del(token string) error {
	return nil
}

const (
	encryptKeyPrefix = "auth_"
)

// redis: cookie + session 认证
var _ AuthInterface = &RedisAuth{}

type RedisAuth struct {
	c          *redis.Client
	expireTime time.Duration
}

func NewRedisAuthRepo(c *redis.Client, expireTime time.Duration) AuthInterface {
	return &RedisAuth{
		c:          c,
		expireTime: expireTime,
	}
}

func (r *RedisAuth) Set(auth *model.AuthInfo) (string, error) {
	key := encryptKeyPrefix + auth.UserID
	status := r.c.Set(context.Background(), key, auth, r.expireTime)
	return key, status.Err()
}

func (r *RedisAuth) Get(token string) (*model.AuthInfo, error) {
	auth := &model.AuthInfo{}
	err := r.c.Get(context.Background(), token).Scan(auth)
	return auth, err
}

func (r *RedisAuth) Del(token string) error {
	return r.c.Del(context.Background(), token).Err()
}

package auth

import (
	"ddd_demo/domain"
	"ddd_demo/domain/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TokenKeyUserID     = "username"
	TokenKeyExpireTime = "exp"
)

var _ repository.AuthInterface = &TokenAuth{}

type TokenAuth struct {
	priKey     string
	expireTime time.Duration
}

func NewJwtAuth(priKey string, expireTime time.Duration) repository.AuthInterface {
	return &TokenAuth{
		priKey:     priKey,
		expireTime: expireTime,
	}
}

// Set 生成 token
func (r *TokenAuth) Set(auth *domain.AuthInfo) (string, error) {
	// 对称加密 auth
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		TokenKeyUserID:     auth.UserID,
		TokenKeyExpireTime: time.Now().Add(r.expireTime).Unix(),
	})

	return token.SignedString([]byte(r.priKey))
}

func (r *TokenAuth) Get(token string) (*domain.AuthInfo, error) {
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

	return &domain.AuthInfo{
		UserID: claims[TokenKeyUserID].(string),
	}, nil
}

func (r *TokenAuth) Del(token string) error {
	return nil
}

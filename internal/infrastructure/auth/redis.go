package auth

import (
	"context"
	"ddd_demo/config"
	"ddd_demo/internal/domain"
	"ddd_demo/internal/domain/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	encryptKeyPrefix = "auth_"
)

// redis: cookie + session 认证
var _ repository.AuthInterface = &RedisAuth{}

type RedisAuth struct {
	c          *redis.Client
	expireTime time.Duration
}

func NewRedisAuthRepo(redisConf config.Redis, expireTime time.Duration) repository.AuthInterface {
	c := redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + redisConf.Port,
		Password: redisConf.Password,
		DB:       0,
	})
	return &RedisAuth{
		c: c,
	}
}

func (r *RedisAuth) Set(auth *domain.AuthInfo) (string, error) {
	key := encryptKeyPrefix + auth.UserID
	status := r.c.Set(context.Background(), key, auth, r.expireTime)
	return key, status.Err()
}

func (r *RedisAuth) Get(token string) (*domain.AuthInfo, error) {
	auth := &domain.AuthInfo{}
	err := r.c.Get(context.Background(), token).Scan(auth)
	return auth, err
}
func (r *RedisAuth) Del(token string) error {
	return r.c.Del(context.Background(), token).Err()
}

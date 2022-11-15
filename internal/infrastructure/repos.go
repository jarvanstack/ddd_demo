package infrastructure

import (
	"ddd_demo/config"
	"ddd_demo/internal/domain/repository"
	"ddd_demo/internal/infrastructure/auth"
	"ddd_demo/internal/infrastructure/persistence"
)

type Repos struct {
	UserRepo repository.UserRepo
	AuthRepo repository.AuthInterface
}

func NewRepos(cfg *config.SugaredConfig) *Repos {
	// 持久化类型的 repo
	persistenceRepos := persistence.NewMysqlRepos(cfg)

	// auth 策略
	var authRepo repository.AuthInterface
	if cfg.Auth.Active == "redis" {
		authRepo = auth.NewRedisAuthRepo(cfg.Redis, cfg.AuthExpireTime)
	} else {
		authRepo = auth.NewJwtAuth(cfg.Auth.PrivateKey, cfg.AuthExpireTime)
	}

	return &Repos{
		UserRepo: persistenceRepos.UserRepo,
		AuthRepo: authRepo,
	}
}

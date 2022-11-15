package servers

import (
	"ddd_demo/config"
	"fmt"

	"ddd_demo/internal/bill"
	"ddd_demo/internal/user"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"

	//  mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbDriver = "mysql"
	dbURLFmt = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

type Repos struct {
	UserRepo user.UserRepo
	AuthRepo user.AuthInterface
	BillRepo bill.BillRepo
}

func NewDB(cfg *config.SugaredConfig) *gorm.DB {
	dbURL := fmt.Sprintf(dbURLFmt, cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)

	db, err := gorm.Open(dbDriver, dbURL)
	if err != nil {
		panic(err)
	}

	db.LogMode(cfg.Log.Env == "dev")

	return db
}

func NewCache(cfg *config.SugaredConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       0,
	})
}

func NewRepos(cfg *config.SugaredConfig) *Repos {
	// 持久化类型的 repo
	db := NewDB(cfg)
	userRepo := user.NewMysqlUserRepo(db)
	billRepo := bill.NewMysqlBillRepo(db)

	// auth 策略
	var authRepo user.AuthInterface
	if cfg.Auth.Active == "redis" {
		authRepo = user.NewRedisAuthRepo(NewCache(cfg), cfg.AuthExpireTime)
	} else {
		authRepo = user.NewJwtAuth(cfg.Auth.PrivateKey, cfg.AuthExpireTime)
	}

	return &Repos{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		BillRepo: billRepo,
	}
}

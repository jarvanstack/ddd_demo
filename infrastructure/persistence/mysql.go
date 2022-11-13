package persistence

import (
	"ddd_demo/config"
	"ddd_demo/domain/repository"
	"fmt"

	"github.com/jinzhu/gorm"

	//  mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbDriver = "mysql"
	dbURLFmt = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

type MysqlRepos struct {
	UserRepo repository.UserInterface
}

func NewMysqlRepos(cfg *config.SugaredConfig) *MysqlRepos {
	dbURL := fmt.Sprintf(dbURLFmt, cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)
	db, err := gorm.Open(dbDriver, dbURL)
	db.LogMode(cfg.Log.Env == "dev")
	if err != nil {
		panic(err)
	}

	repos := &MysqlRepos{
		UserRepo: NewMysqlUserRepo(db),
	}

	return repos
}

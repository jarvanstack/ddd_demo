package main

// 创建数据库初始化表

import (
	"ddd_demo/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	sqlPath    = "./ddd_demo.sql"
	configPath = "../../config.yaml"
)

func main() {
	sc := config.NewConfig(configPath)

	// 设置默认的打开的数据库解决数据库不存在
	sc.Mysql.Database = "mysql"

	_, err := os.Stat(sqlPath)
	if os.IsNotExist(err) {
		log.Println("数据库SQL文件不存在:", err)
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", sc.Mysql.User, sc.Mysql.Password, sc.Mysql.Host, sc.Mysql.Port, sc.Mysql.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("数据库连接失败:", err)
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(59 * time.Second)

	sqls, _ := ioutil.ReadFile(sqlPath)
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := db.Exec(sql).Error
		if err != nil {
			log.Println("数据库导入失败:" + err.Error())
			panic(err)
		} else {
			log.Println(sql, "\t success!")
		}
	}
	return
}

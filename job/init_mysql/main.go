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
	sqlPath         = "../../sql/ddd_demo.sql"
	configPath      = "../../config.yaml"
	defaultDatabase = "mysql"
)

func main() {
	// 读取配置文件
	sc := config.NewConfig(configPath)

	// sql文件检查
	_, err := os.Stat(sqlPath)
	if os.IsNotExist(err) {
		log.Println("数据库SQL文件不存在:", err)
		panic(err)
	}

	// 创建连接
	db := newDB(sc)

	// 创建数据库
	createDatabase(db, sc.Mysql.Database)
	if err != nil {
		log.Println("创建数据库失败:", err)
		panic(err)
	}

	// 执行sql文件
	execSqlFile(db, sqlPath)
	if err != nil {
		log.Println("执行sql文件失败:", err)
		panic(err)
	}

	return
}

func newDB(sc *config.SugaredConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", sc.Mysql.User, sc.Mysql.Password, sc.Mysql.Host, sc.Mysql.Port, defaultDatabase)

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

	return db
}

func createDatabase(db *gorm.DB, dbName string) {
	// 创建数据库
	err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;", dbName)).Error
	if err != nil {
		log.Println("创建数据库失败:", err)
		panic(err)
	}

	// 使用数据库
	err = db.Exec(fmt.Sprintf("USE %s;", dbName)).Error
	if err != nil {
		log.Println("使用数据库失败:", err)
		panic(err)
	}
}

func execSqlFile(db *gorm.DB, sqlPath string) {
	sqls, err := ioutil.ReadFile(sqlPath)
	if err != nil {
		panic(err)
	}

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
}

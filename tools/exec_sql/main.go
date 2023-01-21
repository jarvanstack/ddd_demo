package main

// 创建数据库初始化表

import (
	"database/sql"
	"ddd_demo/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
)

const (
	sqlPath    = "../../sql/ddd_demo.sql"
	configPath = "../../config.yaml"

	// 默认用来创建连接的数据库
	defaultDBname = "mysql"
)

var (
	isForce = pflag.BoolP("force", "f", false, "config file")
)

func main() {

	// 读取配置文件
	sc := config.NewConfig(configPath)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", sc.Mysql.User, sc.Mysql.Password, sc.Mysql.Host, sc.Mysql.Port, defaultDBname)
	dbName := sc.Mysql.Database

	// sql文件检查
	_, err := os.Stat(sqlPath)
	if os.IsNotExist(err) {
		log.Println("[数据库SQL文件不存在]:", err)
		panic(err)
	}

	// 创建连接
	db := newDB(dns)

	if *isForce {
		// 删除数据库
		deleteDatabase(db, dbName)
	}

	// 创建数据库
	createDatabase(db, dbName)
	if err != nil {
		log.Println("[创建数据库失败]:", err)
		panic(err)
	}

	// 执行sql文件
	execSqlFile(db, sqlPath)
	if err != nil {
		log.Println("[执行sql文件失败]:", err)
		panic(err)
	}

	return
}

func newDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("[数据库连接失败]:", err)
		panic(err)
	}

	return db
}

func deleteDatabase(db *sql.DB, dbName string) {
	// 删除数据库
	sqlStr := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)
	_, err := db.Exec(sqlStr)
	if err != nil {
		log.Println(sqlStr, "\t [删除数据库失败]")
		panic(err)
	}
	log.Println(sqlStr, "\t [删除数据库成功]")
}

func createDatabase(db *sql.DB, dbName string) {
	// 创建数据库
	sqlStr := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;", dbName)
	_, err := db.Exec(sqlStr)
	if err != nil {
		log.Println(sqlStr, "\t [创建数据库失败]")
		panic(err)
	}
	log.Println(sqlStr, "\t [创建数据库成功]")

	// 使用数据库
	sqlStr = fmt.Sprintf("USE %s;", dbName)
	_, err = db.Exec(sqlStr)
	if err != nil {
		log.Println(sqlStr, "\t [使用数据库成功]")
		panic(err)
	}
	log.Println(sqlStr, "\t [使用数据库成功]")
}

func execSqlFile(db *sql.DB, sqlPath string) {
	sqls, err := ioutil.ReadFile(sqlPath)
	if err != nil {
		panic(err)
	}

	sqlArr := strings.Split(string(sqls), ";")
	for _, sqlStr := range sqlArr {
		sqlStr = strings.TrimSpace(sqlStr)
		if sqlStr == "" {
			continue
		}
		_, err := db.Exec(sqlStr)
		if err != nil {
			log.Println(sqlStr, "\t [数据库导入失败]:"+err.Error())
			panic(err)
		} else {
			log.Println(sqlStr, "\t [执行成功]!")
		}
	}
}

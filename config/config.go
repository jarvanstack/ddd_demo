package config

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Web   Web   `yaml:"web"`
	RPC   RPC   `yaml:"rpc"`
	Mysql Mysql `yaml:"mysql"`
	Auth  Auth  `yaml:"auth"`
	Redis Redis `yaml:"redis"`
	Log   Log   `yaml:"log"`
}
type Web struct {
	Port string `yaml:"port"`
}
type RPC struct {
	Port string `yaml:"port"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Auth struct {
	Active     string `yaml:"active"`
	ExpireTime string `yaml:"expireTime"`
	PrivateKey string `yaml:"privateKey"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
type Log struct {
	Env        string `yaml:"env"`
	Path       string `yaml:"path"`
	Encoding   string `yaml:"encoding"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

// SugaredConfig 将配置文件的参数解析,比如解析时间为 time.Ticker
type SugaredConfig struct {
	*Config
	AuthExpireTime time.Duration
}

func NewConfig(filePath string) *SugaredConfig {
	// 初始化配置文件
	pflag.StringP("config", "c", filePath, "config file")
	pflag.Parse()
	viper.SetConfigType("yaml")
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	conf := viper.GetString("config")
	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("load config %s fail: %v", conf, err))
	}

	// 解析初始配置
	baseConf := &Config{}
	if err := viper.Unmarshal(baseConf); err != nil {
		if err != nil {
			panic(err)
		}
	}

	// AuthExpireTime 解析为 time.Duration
	authExpireTime, err := time.ParseDuration(baseConf.Auth.ExpireTime)
	if err != nil {
		panic(err)
	}

	// 构造 SugaredConfig
	sugaredConfig := &SugaredConfig{
		Config:         baseConf,
		AuthExpireTime: authExpireTime,
	}

	return sugaredConfig
}

package utils

import (
	"bufio"
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppModel string         `json:"app_model"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:"database"`
	Sms      SmsConfig      `json:"sms"`
	Redis    RedisConfig    `json:"redis"`
}

// mysql数据库配置
type DatabaseConfig struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	DbName    string `json:"db_name"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parse_time"`
	Loc       string `json:"loc"`
}

// redis
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// sms短信
type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

// 解析配置文件
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg)
	if err != nil {
		panic(err)
	}
	return _cfg, err
}

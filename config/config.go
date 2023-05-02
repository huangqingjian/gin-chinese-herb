package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// 配置
type Config struct {
	App     `ini:"app"`
	Redis   `ini:"redis"`
	Mysql   `ini:"mysql"`
	Session `ini:"session"`
	Zap     `ini:"zap"`
}

// app
type App struct {
	Name      string `ini:"name"`
	HttpPort  string `ini:"httpPort"`
	RunMode   string `ini:"runMode"`
	ImagePath string `ini:"imagePath"`
}

// redis
type Redis struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
	Db       int    `ini:"db"`
}

// mysql
type Mysql struct {
	Dsn string `ini:"dsn"`
}

// session
type Session struct {
	Id     string `ini:"id"`
	Secret string `ini:"secret"`
	MaxAge int    `ini:"maxAge"`
}

// zap
type Zap struct {
	FileName string `ini:"fileName"`
	MaxSize  int    `ini:"maxSize"`
	MaxAge   int    `ini:"maxAge"`
	Compress bool   `ini:"compress"`
	Time     string `ini:"compress"`
	Format   string `ini:"compress"`
}

var Global *Config

// 初始化配置
func InitConfig() {
	Global = new(Config)
	err := ini.MapTo(Global, "./conf/app.ini")
	if err != nil {
		fmt.Printf("Config Init Error:%v", err)
		panic(any(err))
	}
}

// 初始化配置（指定路径）
func InitConfigV2(filePath string) {
	Global = new(Config)
	err := ini.MapTo(Global, filePath)
	if err != nil {
		fmt.Printf("Config Init Error:%v", err)
		panic(any(err))
	}
}

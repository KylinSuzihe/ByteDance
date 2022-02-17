package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Port        int          `ini:"port"`
	Release     bool         `ini:"release"`
	MysqlConf   *MySQLConfig `ini:"mysql"`
	RedisConfig *RedisConfig `ini:"redis"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	DB       string `ini:"db"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Addr     string `ini:"addr"`
	Password string `ini:"password"`
	DB       int    `ini:"DB"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}

package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	viper.SetConfigFile("../conf/config.yaml") // 指定配置文件路径
	viper.SetDefault("server.ip", "0.0.0.0")
	viper.SetDefault("server.port", 8008)
	viper.SetDefault("clickhouse.ip", "127.0.0.1")
	viper.SetDefault("clickhouse.port", 9000)

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

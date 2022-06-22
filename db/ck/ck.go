package ck

import (
	"context"
	"dbproxy/utils/logger"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/spf13/viper"
	"log"
	"net"
	"time"
)

var DB clickhouse.Conn

func Connect() {
	var err error
	DB, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{net.JoinHostPort(viper.GetString("clickhouse.ip"), viper.GetString("clickhouse.port"))},
		Auth: clickhouse.Auth{
			Database: viper.GetString("clickhouse.database"),
			Username: viper.GetString("clickhouse.username"),
			Password: viper.GetString("clickhouse.password"),
		},
		//Debug:           true,
		DialTimeout:     time.Duration(viper.GetInt("clickhouse.dial_timeout")) * time.Second,
		MaxOpenConns:    viper.GetInt("clickhouse.max_open_conns"),
		MaxIdleConns:    viper.GetInt("clickhouse.max_idle_conns"),
		ConnMaxLifetime: time.Duration(viper.GetInt("clickhouse.conn_max_lifetime")) * time.Second,
	})

	if err = DB.Ping(context.Background()); err != nil {
		logger.Global.Errorf(err.Error())
		log.Fatalf(err.Error())
		return
	}
}

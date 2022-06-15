/*
* @Author: wangqilong
* @Description:
* @File: clickhouse
* @Date: 2021/11/30 11:04 上午
 */

package ck

import (
	"database/sql"
	"dbproxy/utils/config"
	"dbproxy/utils/log"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
)

var DB *sql.DB

func Connect() {
	var err error
	sdn := fmt.Sprintf("tcp://%v:%v?read_timeout=%v&write_timeout=%v",
		config.Cfg.Clickhouse.Ip,
		config.Cfg.Clickhouse.Port,
		config.Cfg.Clickhouse.ReadTimeout,
		config.Cfg.Clickhouse.WriteTimeout)
	db, err := sql.Open("clickhouse", sdn)
	if err != nil {
		log.L.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.L.Errorf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.L.Error(err)
		}
		return
	}
	log.L.Infof("clickhouse connet success %v:%v", config.Cfg.Clickhouse.Ip, config.Cfg.Clickhouse.Port)
	DB = db
}

/*
* @Author: wangqilong
* @Description:
* @File: clickhouse
* @Date: 2021/11/30 4:00 下午
 */

package model

import (
	"dbproxy/db/ck"
	"dbproxy/utils/log"
	"fmt"
	"strings"
)

type InsertData struct {
	Database string          `json:"database"`
	Table    string          `json:"table"`
	Columns  []string        `json:"columns"`
	Values   [][]interface{} `json:"values"`
}

func (i *InsertData) Insert() error {
	tx, _ := ck.DB.Begin()
	var values []string
	for n := 0; n < len(i.Columns); n++ {
		values = append(values, "?")
	}

	sql := fmt.Sprintf("INSERT INTO %v.%v(%v) VALUES(%v)", i.Database, i.Table, strings.Join(i.Columns, ","), strings.Join(values, ","))
	stmt, _ := tx.Prepare(sql)
	defer stmt.Close()
	for _, v := range i.Values {
		if _, err := stmt.Exec(v...); err != nil {
			log.L.Errorf("SQL: [%v], stmt.Exec err: %v", sql, err.Error())
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		log.L.Errorf("SQL: [%v], commit err: %v", sql, err.Error())
		return err
	}

	return nil
}

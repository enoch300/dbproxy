/*
* @Author: wangqilong
* @Description:
* @File: config
* @Date: 2021/11/30 3:13 下午
 */

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var cfgFile = "../conf/dbproxy.yaml"
var Cfg *GlobalCfg
var once sync.Once

type Log struct {
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}

type Server struct {
	Ip   string `yaml:"ip"`
	Port int64  `yaml:"port"`
}

type Args struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	Ip             string `yaml:"ip"`
	Port           int    `yaml:"port"`
	Database       string `yaml:"database"`
	MaxIdleConns   int    `yaml:"max_idle_cons"`
	MaxOpenConns   int    `yaml:"max_open_cons"`
	ConnectTimeout int    `yaml:"connect_timeout"`
}

type Mysql struct {
	IpaasServer Args `yaml:"ipaas_server"`
	DataCenter  Args `yaml:"data_center"`
}

type Clickhouse struct {
	Ip           string `yaml:"ip"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type GlobalCfg struct {
	Log        Log        `yaml:"log"`
	Mysql      Mysql      `yaml:"mysql"`
	Server     Server     `yaml:"server"`
	Clickhouse Clickhouse `yaml:"clickhouse"`
}

func LoadConfig() {
	once.Do(func() {
		content, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(content, &Cfg)
		if err != nil {
			log.Fatal(err)
		}
	})
}

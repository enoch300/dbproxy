package app

import (
	"dbproxy/db/ck"
	"dbproxy/route"
	"dbproxy/utils/config"
	"dbproxy/utils/logger"
)

func Run() {
	config.LoadConfig()
	logger.InitLog()
	ck.Connect()
	route.Listen()
}

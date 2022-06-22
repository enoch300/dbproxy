/*
* @Author: wangqilong
* @Description:
* @File: root
* @Date: 2021/11/30 3:30 下午
 */

package cmd

import (
	"dbproxy/app"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:     "dbproxy",
	Short:   "数据库API代理",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

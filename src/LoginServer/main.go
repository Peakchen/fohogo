/*
* CopyRight(C) Stefan e-mail:2572915286@qq.com
 */

package main

import (
	"LoginServer/LogicMsg"
	"LoginServer/dbo"
	"flag"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
}

func main() {
	akLog.FmtPrintln("start login server.")

	logincfg := serverConfig.GLoginconfigConfig.Get(0)
	server := logincfg.Zone + logincfg.No
	dbo.StartDBSerice(server)
	client := Kcpnet.NewKcpClient(logincfg.Listenaddr,
		logincfg.Pprofaddr,
		logincfg.Name,
		define.ERouteId_ER_Login,
		nil)

	client.Run()
	dbStatistics.DBStatisticsStop()
	return
}

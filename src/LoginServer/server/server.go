package server

import (
	"LoginServer/LogicMsg"
	"LoginServer/dbo"
	"flag"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
}

func StartServer() {
	akLog.FmtPrintf("start Login server.")
	logincfg := serverConfig.GLoginconfigConfig.Get()
	server := logincfg.Zone + logincfg.No
	dbo.StartDBSerice(server)
	client := Kcpnet.NewKcpClient(logincfg.Name,
		logincfg.Listenaddr,
		logincfg.Pprofaddr)

	client.Run()
	dbStatistics.DBStatisticsStop()
}

package server

import (
	"LoginServer/LogicMsg"
	"LoginServer/dbo"
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"flag"
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
	gameSvr := Kcpnet.NewClient(logincfg.Listenaddr,
		logincfg.Pprofaddr,
		define.ERouteId_ER_Login,
		nil,
		logincfg.Name)

	gameSvr.Run()
	dbStatistics.DBStatisticsStop()
}

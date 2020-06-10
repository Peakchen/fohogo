package server

// add by stefan

import (
	"GameServer/LogicMsg"
	"GameServer/dbo"
	"GameServer/rpc"
	"flag"
	"syscall"

	"github.com/Peakchen/xgameCommon/Config/LogicConfig"
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/HotUpdate"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	akLog.InitLogBroker([]string{"192.168.126.128:9092"})
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
	rpc.Init()
}

func reloadConfig() {
	LogicConfig.LoadLogicAll()
}

func StartServer() {
	Gamecfg := serverConfig.GGameconfigConfig.Get(0)
	server := Gamecfg.Zone + Gamecfg.No
	dbo.StartDBSerice(server)
	// for kill pid to emit signal to do action...
	HotUpdate.RunHotUpdateCheck(&HotUpdate.TServerHotUpdateInfo{
		Recvsignal: syscall.SIGTERM,
		HUCallback: reloadConfig,
	})
	gameSvr := Kcpnet.NewKcpClient(
		Gamecfg.Listenaddr,
		Gamecfg.Pprofaddr,
		Gamecfg.Name,
		define.ERouteId_ER_Game,
		nil)

	gameSvr.Run()
	dbStatistics.DBStatisticsStop()
}

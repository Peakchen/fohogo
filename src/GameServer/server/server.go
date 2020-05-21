package server

// add by stefan

import (
	"GameServer/LogicMsg"
	"GameServer/dbo"
	"GameServer/rpc"
	"github.com/Peakchen/xgameCommon/Config/LogicConfig"
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/HotUpdate"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"flag"
	"syscall"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
	rpc.Init()
}

func reloadConfig() {
	LogicConfig.LoadLogicAll()
}

func StartServer() {
	Gamecfg := serverConfig.GGameconfigConfig.Get()
	server := Gamecfg.Zone + Gamecfg.No
	dbo.StartDBSerice(server)
	// for kill pid to emit signal to do action...
	HotUpdate.RunHotUpdateCheck(&HotUpdate.TServerHotUpdateInfo{
		Recvsignal: syscall.SIGTERM,
		HUCallback: reloadConfig,
	})
	gameSvr := Kcpnet.NewClient(Gamecfg.Listenaddr,
		Gamecfg.Pprofaddr,
		define.ERouteId_ER_Game,
		nil,
		Gamecfg.Name)

	gameSvr.Run()
	dbStatistics.DBStatisticsStop()
}

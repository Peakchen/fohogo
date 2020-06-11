package main

import (
	"GameServer/LogicMsg"
	"flag"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
)

func init() {
	var cfgPath string
	flag.StringVar(&cfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	akLog.InitLogBroker([]string{"192.168.126.128:9092"})
	serverConfig.LoadSvrAllConfig(cfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
}

func runKcpServer() {
	centercfg := serverConfig.GCenterconfigConfig.Get(0)
	newkcpServer := Kcpnet.NewKcpServer(centercfg.Name,
		centercfg.Listenaddr,
		centercfg.Pprofaddr,
		define.ERouteId_ER_MMS,
		nil)

	newkcpServer.Run()
}

func main() {
	akLog.FmtPrintln("start MMGServer...")
	runKcpServer()
}

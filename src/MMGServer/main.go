package main

import (
	"flag"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	var cfgPath string
	flag.StringVar(&cfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	serverConfig.LoadSvrAllConfig(cfgPath)
}

func runKcpServer() {
	centercfg := serverConfig.GCenterconfigConfig.Get(0)
	newkcpServer := Kcpnet.NewKcpServer(centercfg.Name,
		centercfg.Listenaddr,
		centercfg.Pprofaddr,
		nil)

	newkcpServer.Run()
}

func main() {
	akLog.FmtPrintln("start MMGServer...")
	runKcpServer()
}

/*
* CopyRight(C) StefanChen e-mail:2572915286@qq.com
 */

package main

import (
	//"log"

	"ExternalWebGateway/kcpclient"
	"ExternalWebGateway/kcpserver"
	"ExternalWebGateway/webserver"
	"GameServer/LogicMsg"
	"flag"
	"sync"

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

func main() {
	akLog.FmtPrintln("start ExternalWebGateway.")
	var (
		sw    sync.WaitGroup
		excol = Kcpnet.NewExternalCollection()
	)
	sw.Add(2)
	go webserver.StartServer()
	go kcpserver.StartServer(excol)
	go kcpclient.StartKcpClient(&sw, excol)
	sw.Wait()
}

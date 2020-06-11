// add by stefan

package main

import (
	"ExternalGateway/LogicMsg"
	"ExternalGateway/ado"
	"ExternalGateway/server"
	"flag"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	akLog.InitLogBroker([]string{"192.168.126.128:9092"})
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
}

func main() {
	akLog.FmtPrintf("start ExternalGateWay.")
	ado.StartDBSerice("ExternalGateWay")
	server.Start()
	dbStatistics.DBStatisticsStop()
}

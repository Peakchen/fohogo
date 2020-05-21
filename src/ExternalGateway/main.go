// add by stefan

package main

import (
	"ExternalGateway/LogicMsg"
	"ExternalGateway/ado"
	"flag"

	"github.com/Peakchen/xgameCommon/Kcpnet"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
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
	akLog.FmtPrintf("start ExternalGateWay.")
	ado.StartDBSerice("ExternalGateWay")
	externalgw := serverConfig.GExternalgwconfigConfig.Get()
	newExternalServer := Kcpnet.NewKcpServer(externalgw.Name,
		externalgw.Listenaddr,
		externalgw.Pprofaddr)

	newExternalServer.Run()
	dbStatistics.DBStatisticsStop()
}

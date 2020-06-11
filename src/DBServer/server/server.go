package server

// add by stefan

import (
	"flag"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/ado/service"
	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	akLog.InitLogBroker([]string{"192.168.126.128:9092"})
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
}

/*
	run db server.
*/
func StartDBServer() {
	server := "sever1"
	service.StartMultiDBProvider(server)
	dbStatistics.DBStatisticsStop()
}

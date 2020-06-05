package kcpserver

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/Peakchen/xgameCommon/define"
)

func StartServer(excol *Kcpnet.ExternalCollection) {
	akLog.FmtPrintf("start kcp server.")
	externalgw := serverConfig.GExternalgwconfigConfig.Get(1)
	newtcpServer := akNet.NewTcpServer(
		externalgw.Listenaddr,
		externalgw.Pprofaddr,
		define.ERouteId_ER_ESG,
		externalgw.Name)

	newtcpServer.Run()
}

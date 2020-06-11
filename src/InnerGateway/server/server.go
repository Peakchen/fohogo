package server

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
)

func StartServer() {
	akLog.FmtPrintf("start InnerGateway server.")
	Innergw := serverConfig.GInnergwconfigConfig.Get(0)
	newInnerServer := Kcpnet.NewKcpServer(Innergw.Name,
		Innergw.Listenaddr,
		Innergw.Pprofaddr,
		define.ERouteId_ER_ISG,
		nil)

	newInnerServer.Run()
}

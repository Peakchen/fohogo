package server

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/Kcpnet"
)

func StartServer() {
	akLog.FmtPrintf("start InnerGateway server.")
	Innergw := serverConfig.GInnergwconfigConfig.Get()
	newInnerServer := Kcpnet.NewTcpServer(Innergw.Listenaddr,
		Innergw.Pprofaddr,
		define.ERouteId_ER_ISG,
		Innergw.Name)

	newInnerServer.Run()
}

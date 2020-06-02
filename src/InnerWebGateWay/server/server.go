package server

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
)

func StartServer() {
	akLog.FmtPrintf("start InnerGateway server.")
	Innergw := serverConfig.GInnergwconfigConfig.Get(0)
	newInnerServer := Kcpnet.NewKcpServer(Innergw.Name,
		Innergw.Listenaddr,
		Innergw.Pprofaddr)

	newInnerServer.Run()
}

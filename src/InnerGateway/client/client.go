package client

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
)

func StartClient() {
	akLog.FmtPrintf("start InnerGateway client.")
	Innergw := serverConfig.GInnergwconfigConfig.Get(0)
	gameSvr := Kcpnet.NewKcpClient(
		Innergw.Connectaddr,
		Innergw.Pprofaddr,
		Innergw.Name,
		define.ERouteId_ER_ISG,
		Innergw.Id,
		nil)

	gameSvr.Run()
}

package client

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/Kcpnet"
)

func StartClient() {
	akLog.FmtPrintf("start InnerGateway client.")
	Innergw := serverConfig.GInnergwconfigConfig.Get()
	gameSvr := Kcpnet.NewClient(Innergw.Connectaddr,
		Innergw.Pprofaddr,
		define.ERouteId_ER_ISG,
		nil,
		Innergw.Name)

	gameSvr.Run()
}

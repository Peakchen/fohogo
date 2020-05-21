package client

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
)

func StartClient() {
	akLog.FmtPrintf("start InnerGateway client.")
	Innergw := serverConfig.GInnergwconfigConfig.Get()
	gameSvr := Kcpnet.NewKcpClient(Innergw.Name,
		Innergw.Connectaddr,
		Innergw.Pprofaddr)

	gameSvr.Run()
}

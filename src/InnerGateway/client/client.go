package client

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
)

func StartClient() {
	akLog.FmtPrintf("start InnerGateway client.")
	Innergw := serverConfig.GInnergwconfigConfig.Get()
	gameSvr := Kcpnet.NewKcpClient(
		Innergw.Connectaddr,
		Innergw.Pprofaddr,
		Innergw.Name)

	gameSvr.Run()
}

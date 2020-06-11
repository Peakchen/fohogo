package kcpclient

import (
	"strconv"
	"sync"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/define"
)

func StartKcpClient(sw *sync.WaitGroup, excol *Kcpnet.ExternalCollection) {
	gitem := serverConfig.GServerglobalconfig.Get(0)
	if gitem == nil {
		return
	}
	isOpen, err := strconv.Atoi(gitem.Value)
	if err != nil {
		return
	}
	if isOpen != 1 {
		return
	}
	sw.Add(1)
	externalgw := serverConfig.GExternalgwconfigConfig.Get(2)
	newtcpClient := Kcpnet.NewKcpClient(
		externalgw.Listenaddr,
		externalgw.Pprofaddr,
		externalgw.Name,
		define.ERouteId_ER_ESG,
		externalgw.Id,
		excol)

	newtcpClient.Run()
}

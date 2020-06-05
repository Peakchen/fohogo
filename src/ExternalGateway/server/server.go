package server

import (
	"strconv"
	"sync"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/Peakchen/xgameCommon/define"
)

func runKcpServer(sw *sync.WaitGroup, excol *Kcpnet.ExternalCollection) {
	externalgw := serverConfig.GExternalgwconfigConfig.Get(0)
	newkcpServer := Kcpnet.NewKcpServer(externalgw.Name,
		externalgw.Listenaddr,
		externalgw.Pprofaddr,
		excol)

	newkcpServer.Run()
}

func runTcpServer(sw *sync.WaitGroup) {
	externalgw := serverConfig.GExternalgwconfigConfig.Get(1)
	newtcpServer := akNet.NewTcpServer(
		externalgw.Listenaddr,
		externalgw.Pprofaddr,
		define.ERouteId_ER_ESG,
		externalgw.Name)

	newtcpServer.Run()
}

func runKcpClient(sw *sync.WaitGroup, excol *Kcpnet.ExternalCollection) {
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
		excol)

	newtcpClient.Run()
}

func Start() {
	var (
		sw    sync.WaitGroup
		excol = Kcpnet.NewExternalCollection()
	)
	sw.Add(2)
	go runTcpServer(&sw)
	go runKcpServer(&sw, excol)
	go runKcpClient(&sw, excol)
	sw.Wait()
}

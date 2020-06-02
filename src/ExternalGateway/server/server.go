package server

import (
	"sync"

	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/Peakchen/xgameCommon/define"
)

func runKcpServer(sw *sync.WaitGroup) {
	externalgw := serverConfig.GExternalgwconfigConfig.Get()
	newkcpServer := Kcpnet.NewKcpServer(externalgw.Name,
		externalgw.Listenaddr,
		externalgw.Pprofaddr)

	newkcpServer.Run()
}

func runTcpServer(sw *sync.WaitGroup) {
	externalgw := serverConfig.GExternalgwconfigConfig.Get()
	newtcpServer := akNet.NewTcpServer(
		externalgw.Listenaddr,
		externalgw.Pprofaddr,
		define.ERouteId_ER_ESG,
		externalgw.Name)

	newtcpServer.Run()
}

func Start() {
	var sw sync.WaitGroup
	sw.Add(2)
	go runTcpServer(&sw)
	go runKcpServer(&sw)
	sw.Wait()
}
package webserver

import (
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/akWebNet"
)

func StartServer() {
	// start websock server.
	externalgw := serverConfig.GExternalgwconfigConfig.Get(0)
	websvr := akWebNet.NewWebsocketSvr(externalgw.Listenaddr, externalgw.Pprofaddr)
	websvr.Run()
}

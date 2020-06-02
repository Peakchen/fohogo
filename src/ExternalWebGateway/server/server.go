package server

import "github.com/Peakchen/xgameCommon/akWebNet"

func StartServer() {
	// start websock server.
	websvr := akWebNet.NewWebsocketSvr("172.0.0.1:8080")
	websvr.Run()
}

/*
* CopyRight(C) StefanChen e-mail:2572915286@qq.com
 */

package main

import (
	//"log"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/akWebNet"
)

func init() {

}

func main() {
	akLog.FmtPrintln("start ExternalWebGateway.")
	// start websock server.
	websvr := akWebNet.NewWebsocketSvr("172.0.0.1:8080")
	websvr.Run()
}

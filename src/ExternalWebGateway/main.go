/*
* CopyRight(C) StefanChen e-mail:2572915286@qq.com
 */

package main

import (
	//"log"
	"ExternalWebGateway/client"
	"ExternalWebGateway/server"
	"sync"

	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {

}

func run() {
	var sw sync.WaitGroup
	sw.Add(2)
	server.StartServer()
	client.StartClient()
	sw.Wait()
}

func main() {
	akLog.FmtPrintln("start ExternalWebGateway.")
	run()
}

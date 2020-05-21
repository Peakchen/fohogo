/*
* CopyRight(C) Stefan e-mail:2572915286@qq.com
 */

package main

import (
	"LoginServer/server"
	"github.com/Peakchen/xgameCommon/akLog"
)

func main() {
	akLog.FmtPrintln("start login server.")

	server.StartServer()
	return
}

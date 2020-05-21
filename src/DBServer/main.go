// add by stefan
package main

import (
	"DBServer/server"
	"github.com/Peakchen/xgameCommon/akLog"
)

func main() {
	akLog.FmtPrintln("run db server.")
	server.StartDBServer()
}

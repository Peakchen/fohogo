// add by stefan

package main

import (
	"GameServer/server"
	"github.com/Peakchen/xgameCommon/akLog"
	//"log"
)

func init() {

}

func main() {
	akLog.FmtPrintf("start gameServer.")

	server.StartServer()
	return
}

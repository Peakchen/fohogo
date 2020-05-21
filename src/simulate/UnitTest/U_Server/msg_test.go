package U_Server

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Server"
	"simulate/TestCommon"
	"testing"
)

var serverhost string = "127.0.0.1:51001"

func TestServer(t *testing.T) {
	akLog.FmtPrintf("server msg test.")
	req := &MSG_Server.CS_EnterServer_Req{}
	req.Enter = 2
	serverM := TestCommon.NewModule(serverhost, "server")
	serverM.PushMsg(uint16(define.ERouteId_ER_Login),
		uint16(MSG_MainModule.MAINMSG_SERVER),
		uint16(MSG_Server.SUBMSG_CS_ServerRegister),
		req)
	serverM.Run()
}

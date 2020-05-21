package Player

import (
	"GameServer/logic"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/Kcpnet"
)

type TPlayerReady struct {
}

func (this *TPlayerReady) EnterReady(session Kcpnet.TcpSession) {
	akLog.FmtPrintln("enter ready.")
	player := GetPlayer(session.GetIdentify())
	if player == nil {
		akLog.Error("can not find ")
		return
	}

	//for test
	//RunModuleRpc4GetPlayerInfoTest(session, cstRpcModule_GetPlayerInfo, cstRpcFunc_GetPlayerInfo)
	//RunRpc4GetPlayerInfoTest(session, cstRpcFunc_GetPlayerInfo)
}

func (this *TPlayerReady) LeaveReady(session Kcpnet.TcpSession) {

}

func (this *TPlayerReady) ReconnectReady(session Kcpnet.TcpSession) {

}

func init() {
	logic.RegisterEnterReadyModule(cstPlayerSubModule, &TPlayerReady{})
	logic.RegisterReconnReadyModule(cstPlayerSubModule, &TPlayerReady{})
	logic.RegisterLeaveReadyModule(cstPlayerSubModule, &TPlayerReady{})
}

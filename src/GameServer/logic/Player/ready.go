package Player

import (
	"GameServer/logic"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
)

type TPlayerReady struct {
}

func (this *TPlayerReady) EnterReady(session Kcpnet.TSession) {
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

func (this *TPlayerReady) LeaveReady(session Kcpnet.TSession) {

}

func (this *TPlayerReady) ReconnectReady(session Kcpnet.TSession) {

}

func init() {
	logic.RegisterEnterReadyModule(cstPlayerSubModule, &TPlayerReady{})
	logic.RegisterReconnReadyModule(cstPlayerSubModule, &TPlayerReady{})
	logic.RegisterLeaveReadyModule(cstPlayerSubModule, &TPlayerReady{})
}

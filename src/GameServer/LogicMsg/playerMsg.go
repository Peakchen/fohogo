package LogicMsg

import (
	"GameServer/logic/Player"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Player"
)

func onGetPlayerInfo(session Kcpnet.TSession, req *MSG_Player.CS_PlayerInfo_Req) (succ bool, err error) {
	akLog.FmtPrintf("[onGetPlayerInfo] SessionID: %v.", session.GetIdentify())

	rsp := &MSG_Player.SC_PlayerInfo_Rsp{}
	rsp.Ret = MSG_Player.ErrorCode_Success
	data := Player.GetPlayer(session.GetIdentify())
	if data == nil {
		return
	}
	akLog.FmtPrintf("get player info: %v.", data.BaseInfo[MSG_Player.EmBaseInfo_Name])
	return session.SendInnerClientMsg(uint16(MSG_MainModule.MAINMSG_PLAYER),
		uint16(MSG_Player.SUBMSG_SC_PlayerInfo),
		rsp)
}

func init() {
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_PLAYER), uint16(MSG_Player.SUBMSG_CS_PlayerInfo), onGetPlayerInfo)
}

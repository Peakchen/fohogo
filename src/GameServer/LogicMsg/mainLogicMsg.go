package LogicMsg

import (
	"GameServer/logic"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Player"
)

func Init() {

}

func onEnterServer(session Kcpnet.TcpSession, req *MSG_Player.CS_EnterServer_Req) (succ bool, err error) {
	akLog.FmtPrintf("enter Server player(%v) enter game server.", session.GetIdentify())
	logic.EnterGameReady(session)
	rsp := &MSG_Player.SC_EnterServer_Rsp{}
	rsp.Ret = MSG_Player.ErrorCode_Success
	return session.SendInnerMsg(uint16(MSG_MainModule.MAINMSG_PLAYER),
		uint16(MSG_Player.SUBMSG_SC_EnterServer),
		rsp)
}

func onLeaveServer(session Kcpnet.TcpSession, req *MSG_Player.CS_LeaveServer_Req) (succ bool, err error) {
	akLog.FmtPrintf("leave Server player(%v).", session.GetIdentify())
	return true, nil
}

func init() {
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_PLAYER), uint16(MSG_Player.SUBMSG_CS_EnterServer), onEnterServer)
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_PLAYER), uint16(MSG_Player.SUBMSG_CS_LeaveServer), onLeaveServer)
}

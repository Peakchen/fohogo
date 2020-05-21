package gameMsg

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Player"
	"simulate/AutoTest/msgImp"
)

func Init() {
	akLog.FmtPrintln("run game player test.")

	msgImp.RegisterMsg(uint16(MSG_MainModule.MAINMSG_PLAYER),
		uint16(MSG_Player.SUBMSG_CS_EnterServer),
		"CS_EnterServer_Req",
		&MSG_Player.CS_EnterServer_Req{})

	msgImp.RegisterMsg(uint16(MSG_MainModule.MAINMSG_PLAYER),
		uint16(MSG_Player.SUBMSG_SC_EnterServer),
		"SC_EnterServer_Rsp",
		&MSG_Player.SC_EnterServer_Rsp{})
}

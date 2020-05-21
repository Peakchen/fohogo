package LogicMsg

// add by stefan

import (
	"fmt"
	"net"

	"github.com/Peakchen/xgameCommon/Kcpnet"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_HeartBeat"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Server"
	"github.com/golang/protobuf/proto"
)

func ExternalGatewayMessageCallBack(c net.Conn, mainID uint16, subID uint16, msg proto.Message) {
	akLog.FmtPrintf("exec external gateway server message call back: %v, %v.", c.RemoteAddr(), c.LocalAddr())
}

func onSvrRegister(session Kcpnet.TcpSession, req *MSG_Server.CS_ServerRegister_Req) (succ bool, err error) {
	akLog.FmtPrintf("onSvrRegister, StrIdentify: %v, recv: %v.", session.GetIdentify(), req.ServerType)
	var (
		msgfmt string
	)

	session.Push(define.ERouteId(req.ServerType))
	for _, id := range req.Msgs {
		mainid, subid := akNet.DecodeCmd(uint32(id))
		msgfmt += fmt.Sprintf("[mainid: %v, subid: %v]\t", mainid, subid)
	}

	msgfmt += "\n"
	akLog.FmtPrintln("message context: ", msgfmt)
	return Kcpnet.RegisterMessageRet(session)
}

func onHeartBeat(session Kcpnet.TcpSession, req *MSG_HeartBeat.CS_HeartBeat_Req) (succ bool, err error) {
	return Kcpnet.ResponseHeartBeat(session)
}

func init() {
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_CS_ServerRegister), onSvrRegister)
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_HEARTBEAT), uint16(MSG_HeartBeat.SUBMSG_CS_HeartBeat), onHeartBeat)
}

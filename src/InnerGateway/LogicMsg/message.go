// add by stefan

package LogicMsg

import (
	"common/Define"
	"common/Log"
	"common/akNet"
	"common/msgProto/MSG_HeartBeat"
	"common/msgProto/MSG_MainModule"
	"common/msgProto/MSG_Server"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
)

func InnerGatewayMessageCallBack(c net.Conn, mainID uint16, subID uint16, msg proto.Message) {
	Log.FmtPrintf("exec [innter gateway] server message call back.", c.RemoteAddr(), c.LocalAddr())
}

func onSvrRegister(session akNet.TcpSession, req *MSG_Server.CS_ServerRegister_Req) (succ bool, err error) {
	Log.FmtPrintf("onSvrRegister: StrIdentify: %v, recv: %v.", session.GetIdentify(), req.ServerType)
	var (
		msgfmt string
	)

	session.Push(Define.ERouteId(req.ServerType))
	for _, id := range req.Msgs {
		mainid, subid := akNet.DecodeCmd(uint32(id))
		msgfmt += fmt.Sprintf("[mainid: %v, subid: %v]\t", mainid, subid)
	}

	msgfmt += "\n"
	Log.FmtPrintln("message context: ", msgfmt)

	rsp := &MSG_Server.SC_ServerRegister_Rsp{}
	rsp.Ret = MSG_Server.ErrorCode_Success
	rsp.Identify = session.GetModuleName()
	return session.SendInnerMsg(uint16(MSG_MainModule.MAINMSG_SERVER),
		uint16(MSG_Server.SUBMSG_SC_ServerRegister),
		rsp)
}

func onHeartBeat(session akNet.TcpSession, req *MSG_HeartBeat.CS_HeartBeat_Req) (succ bool, err error) {
	return akNet.ResponseHeartBeat(session)
}

func init() {
	akNet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_CS_ServerRegister), onSvrRegister)
	akNet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_HEARTBEAT), uint16(MSG_HeartBeat.SUBMSG_CS_HeartBeat), onHeartBeat)
}

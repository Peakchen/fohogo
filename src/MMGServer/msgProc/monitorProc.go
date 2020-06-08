package msgProc

import (
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_CenterGate"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
)

func onEnter(session Kcpnet.TcpSession, req *MSG_CenterGate.CS_PlayerOnline_Req) (succ bool, err error) {
	session.GetCenterSession().AddPlayerSession(req.PlayerIdentify, session)
	return true, nil
}

func onLeave(session Kcpnet.TcpSession, req *MSG_CenterGate.CS_PlayerOffline_Req) (succ bool, err error) {
	session.GetCenterSession().ClearPlayerSession(req.PlayerIdentify)
	return true, nil
}

func onGetSessions(session Kcpnet.TcpSession, req *MSG_CenterGate.CS_GetBroadCastSessions_Req) (succ bool, err error) {
	sessionPlayers := map[string][]string{}
	for _, key := range req.PlayerIdentifys {
		pSess := session.GetCenterSession().GetPlayerSession(key)
		if pSess == nil {
			continue
		}
		vkeys, exist := sessionPlayers[pSess.GetRemoteAddr()]
		if !exist {
			vkeys = []string{}
		}
		vkeys = append(vkeys, key)
		sessionPlayers[pSess.GetRemoteAddr()] = vkeys
	}
	for sk, vPkeys := range sessionPlayers {
		svrSess := session.GetCenterSession().GetSvrSession(sk)
		rsp := &SC_GetBroadCastSessions_Rsp{
			PlayerIdentifys: vPkeys,
			Data:            req.Data,
		}
		svrSess.SendInnerClientMsg(uint16(MSG_MainModule.MAINMSG_CENTERGATE),
			uint16(MSG_CenterGate.SUBMSG_SC_GetBroadCastSessions),
			rsp)
	}
	return true, nil
}

func init() {
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_CENTERGATE), uint16(MSG_CenterGate.SUBMSG_CS_PlayerOnline), onEnter)
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_CENTERGATE), uint16(MSG_CenterGate.SUBMSG_CS_PlayerOffline), onLeave)
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_CENTERGATE), uint16(MSG_CenterGate.SUBMSG_CS_GetBroadCastSessions), onGetSessions)
}

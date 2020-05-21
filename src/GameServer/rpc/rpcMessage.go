package rpc

/*
	rpc process message module
	date: 20191203
	author: stefan
	version: 1.0
*/

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Rpc"
)

func Init() {

}

/*
	@func: onRpcProcess 接收处理rpc消息
	@param1: session obj
	@param2: req content (module, func, data)
*/
func onRpcProcess(session Kcpnet.TcpSession, req *MSG_Rpc.CS_Rpc_Req) (succ bool, err error) {
	akLog.FmtPrintf("rpc process, rpc module: %v, func: %v.", req.Rpcmodule, req.Rpcfunc)
	if len(req.Rpcmodule) == 0 {
		succ, err = onSingleRpc(req.Rpcfunc, req.Data)
	} else {
		succ, err = onModuleRpcProcess(req.Rpcmodule, req.Rpcfunc, req.Data)
	}
	return
}

func init() {
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_RPC), uint16(MSG_Rpc.SUBMSG_CS_Rpc), onRpcProcess)
}

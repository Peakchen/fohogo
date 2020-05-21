package logic

import "github.com/Peakchen/xgameCommon/Kcpnet"

/*
	by stefan
	date: 20191111 16:08
*/
type ILogicReady interface {
	EnterReady(session Kcpnet.TcpSession)
	LeaveReady(session Kcpnet.TcpSession)
	ReconnectReady(session Kcpnet.TcpSession)
}

var (
	GEnterReadyModule map[string]ILogicReady = map[string]ILogicReady{}
)

var (
	GLeaveReadyModule map[string]ILogicReady = map[string]ILogicReady{}
)

var (
	GReconnReadyModule map[string]ILogicReady = map[string]ILogicReady{}
)

func RegisterEnterReadyModule(module string, data ILogicReady) {
	GEnterReadyModule[module] = data
}

func RegisterReconnReadyModule(module string, data ILogicReady) {
	GReconnReadyModule[module] = data
}

func RegisterLeaveReadyModule(module string, data ILogicReady) {
	GLeaveReadyModule[module] = data
}

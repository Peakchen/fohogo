package logic

import (
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
	"reflect"
)

// after player login, need getting ready.
func EnterGameReady(session Kcpnet.TcpSession) {
	params := []reflect.Value{reflect.ValueOf(session)}
	for module, obj := range GEnterReadyModule {
		enter := reflect.ValueOf(obj).MethodByName("EnterReady")
		if enter.IsNil() || !enter.IsValid() {
			akLog.ErrorIDCard(session.GetIdentify(), "can not find EnterReady method, module: ", module)
			return
		}

		enter.Call(params)
	}
}

//before leave, get ready.
func LeaveGameReady(session Kcpnet.TcpSession) {
	params := []reflect.Value{reflect.ValueOf(session)}
	for module, obj := range GLeaveReadyModule {
		enter := reflect.ValueOf(obj).MethodByName("LeaveReady")
		if enter.IsNil() || !enter.IsValid() {
			akLog.ErrorIDCard(session.GetIdentify(), "can not find LeaveReady method, module: ", module)
			return
		}

		enter.Call(params)
	}
}

//broken link reconnct enter game ready.
func ReconnectEnterReady(session Kcpnet.TcpSession) {
	params := []reflect.Value{reflect.ValueOf(session)}
	for module, obj := range GReconnReadyModule {
		enter := reflect.ValueOf(obj).MethodByName("ReconnectReady")
		if enter.IsNil() || !enter.IsValid() {
			akLog.ErrorIDCard(session.GetIdentify(), "can not find ReconnectReady method, module: ", module)
			return
		}

		enter.Call(params)
	}
}

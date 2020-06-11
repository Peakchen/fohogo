package LogicMsg

import (
	"LoginServer/Logic/UserAccount"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Login"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
)

func onUserBind(key string, req *MSG_Login.CS_UserBind_Req) (succ bool, err error) {
	akLog.FmtPrintf("onUserBind recv: %v, %v.", key, req.Account, req.Passwd)

	return
}

func onUserRegister(session Kcpnet.TSession, req *MSG_Login.CS_UserRegister_Req) (succ bool, err error) {
	akLog.FmtPrintf("[onUserRegister] name: %v, identify: %v, Account: %v, Passwd: %v, DeviceSerial: %v, DeviceName: %v.", session.GetModuleName(), session.GetIdentify(), req.Account, req.Passwd, req.DeviceSerial, req.DeviceName)
	rsp := &MSG_Login.SC_UserRegister_Rsp{}
	rsp.Ret = MSG_Login.ErrorCode_Success

	acc := &UserAccount.TUserAcc{
		UserName:   req.Account,
		Passwd:     req.Passwd,
		DeviceNo:   req.DeviceSerial,
		DeviceType: req.DeviceName,
	}

	if err, exist := UserAccount.RegisterUseAcc(acc); err != nil || !exist {
		rsp.Ret = MSG_Login.ErrorCode_Fail
	}

	session.SetIdentify(acc.Identify())

	return session.SendInnerClientMsg(uint16(MSG_MainModule.MAINMSG_LOGIN),
		uint16(MSG_Login.SUBMSG_SC_UserRegister),
		rsp)
}

func onUserLogin(session Kcpnet.TSession, req *MSG_Login.CS_Login_Req) (succ bool, err error) {
	akLog.FmtPrintf("[onUserLogin] identify: %v, Account: %v, Passwd: %v, DeviceSerial: %v, DeviceName: %v.", session.GetIdentify(), req.Account, req.Passwd, req.DeviceSerial, req.DeviceName)

	rsp := &MSG_Login.SC_Login_Rsp{}
	rsp.Ret = MSG_Login.ErrorCode_Success

	acc := &UserAccount.TUserAcc{
		UserName:   req.Account,
		Passwd:     req.Passwd,
		DeviceNo:   req.DeviceSerial,
		DeviceType: req.DeviceName,
	}

	if _, exist := UserAccount.GetUserAcc(acc); !exist {
		rsp.Ret = MSG_Login.ErrorCode_UserNotExistOrPasswdErr
	}
	session.SetIdentify(acc.Identify())
	return session.SendInnerClientMsg(uint16(MSG_MainModule.MAINMSG_LOGIN),
		uint16(MSG_Login.SUBMSG_SC_Login),
		rsp)
}

func init() {
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_LOGIN), uint16(MSG_Login.SUBMSG_CS_UserBind), onUserBind)
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_LOGIN), uint16(MSG_Login.SUBMSG_CS_UserRegister), onUserRegister)
	Kcpnet.RegisterMessage(uint16(MSG_MainModule.MAINMSG_LOGIN), uint16(MSG_Login.SUBMSG_CS_Login), onUserLogin)
}

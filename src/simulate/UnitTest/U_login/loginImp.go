package U_login

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Login"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"simulate/TestCommon"
	"simulate/UnitTest/U_config"
	"time"
)

var (
	cstSendInterval                           = 200
	loginM          *TestCommon.TModuleCommon = nil
)

func Init(host string, module string) {
	loginM = TestCommon.NewModule(host, module)
}

func Run(idx int) {
	UserRegister(idx)
}

func UserRegister(idx int) {
	akLog.FmtPrintf("user register.")
	simulateItems := U_config.GloginConfig.Get()
	if len(simulateItems) == idx {
		return
	}

	item := simulateItems[idx]
	if item.Register != U_config.CstRegister_No {
		req := &MSG_Login.CS_UserRegister_Req{}
		req.Account = item.Username
		req.Passwd = item.Passwd
		req.DeviceSerial = "123"
		req.DeviceName = "androd"
		akLog.FmtPrintln("UserRegister: ", item.Username, item.Passwd)
		loginM.PushMsg(uint16(MSG_MainModule.MAINMSG_LOGIN),
			uint16(MSG_Login.SUBMSG_CS_UserRegister),
			req)
		go loginM.Run()
		time.Sleep(time.Duration(cstSendInterval) * time.Millisecond)
	}

	UserLogin(item)
}

func UserLogin(item *U_config.TSimulateLoginBase) {
	akLog.FmtPrintf("user login.")
	if item.Login == U_config.CstLogin_No {
		return
	}

	req := &MSG_Login.CS_Login_Req{}
	req.Account = item.Username
	req.Passwd = item.Passwd
	req.DeviceSerial = "456"
	req.DeviceName = "iso"
	akLog.FmtPrintln("UserLogin: ", item.Username, item.Passwd)
	loginM.PushMsg(uint16(MSG_MainModule.MAINMSG_LOGIN),
		uint16(MSG_Login.SUBMSG_CS_Login),
		req)
	go loginM.Run()
	time.Sleep(time.Duration(cstSendInterval) * time.Millisecond)
}

package udxNet

import "common"

type UdxNetServer struct {
	udxDll *common.TDLLMgr
	udxObj interface{}
}

func NewUdxNetServer(dll *common.TDLLMgr) *UdxNetServer {
	return &UdxNetServer{
		udxDll: dll,
	}
}

func (this *UdxNetServer) Run() {
	this.udxObj = this.udxDll.DoProc("CreateFastUdx")
}

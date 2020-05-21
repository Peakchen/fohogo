// +build windows

package common

import "syscall"

type TDLLMgr struct {
	dll syscall.LazyDLL
}

func NewDLL(dllpath string) *TDLLMgr {
	return &TDLLMgr{
		dll: syscall.NewLazyDLL(dllpath),
	}
}

func (this *TDLLMgr) DoProc(fn string, params ...interface{}) (ret uintptr) {
	f := this.dll.NewProc(fn)
	var (
		uptrs = []*uintptr{}
	)
	for _, p := range params {
		uptrs = append(uptrs, uintptr(p))
	}
	ret, _, _ = f.Call(uptrs...)
	return
}

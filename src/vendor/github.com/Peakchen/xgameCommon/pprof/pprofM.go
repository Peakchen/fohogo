package pprof

// add by stefan 20190606 16:12
import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/aktime"
	"fmt"
	"os"
	"path"
	"runtime/pprof"
	"strings"
	"time"
	//"log"
	"github.com/Peakchen/xgameCommon/utls"
	"context"
	"sync"
)

const (
	const_PProfWriteInterval = int32(60 * 1)
)

type TPProfMgr struct {
	ctx context.Context
	wg  sync.WaitGroup
	cpu *os.File
	mem *os.File
}

var (
	_pprofobj *TPProfMgr
)

func init() {
	_pprofobj = &TPProfMgr{}
}

func Run(ctx context.Context) {
	_pprofobj.StartPProf(ctx)
}

func Exit() {
	_pprofobj.Exit()
}

func (this *TPProfMgr) StartPProf(ctx context.Context) {
	this.ctx = ctx
	this.wg.Add(1)
	checkcreateTempDir()
	this.cpu = createCpu()
	this.mem = createMem()
	go this.loop()
}

func (this *TPProfMgr) Exit() {
	akLog.FmtPrintln("pprof exist.")
	this.flush()
	if this.cpu != nil {
		this.cpu.Close()
	}
	if this.mem != nil {
		pprof.WriteHeapProfile(this.mem)
	}
}

func (this *TPProfMgr) flush() {
	if this.cpu != nil {
		pprof.StopCPUProfile()
	}
	if this.mem != nil {
		pprof.WriteHeapProfile(this.mem)
	}
}

func (this *TPProfMgr) loop() {
	defer this.wg.Done()
	t := time.NewTicker(time.Duration(const_PProfWriteInterval) * time.Second)
	for {
		select {
		case <-this.ctx.Done():
			this.Exit()
			return
		case <-t.C:
			// do nothing...
			this.flush()
		}
	}
}

func Newpprof(file string) (retfile string) {
	timeformat := aktime.Now().Format("2006-01-02")
	retfile = timeformat + "_" + file
	execpath, err := os.Executable()
	if err != nil {
		return
	}
	execpath = strings.Replace(execpath, "\\", "/", -1)
	_, sfile := path.Split(execpath)
	arrfile := strings.Split(sfile, ".")
	retfile = fmt.Sprintf("./pprof/%s_%v.prof", arrfile[0], retfile)
	return
}

func checkcreateTempDir() {
	exepath := utls.GetExeFilePath()
	filepath := exepath + "/pprof"
	exist, err := utls.IsPathExisted(filepath)
	if err != nil {
		panic("check path exist err: " + err.Error())
		return
	}

	if false == exist {
		err = os.Mkdir(filepath, os.ModePerm)
		if err != nil {
			panic("log mkdir fail, err: " + err.Error())
			return
		}
	}
}

func createCpu() (file *os.File) {
	cpuf := Newpprof("cpu")
	f, err := os.OpenFile(cpuf, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		akLog.FmtPrintln("cpu pprof open fail, err: ", err)
		return
	}
	pprof.StartCPUProfile(f)
	return f
}

func createMem() (file *os.File) {
	cpuf := Newpprof("mem")
	f, err := os.OpenFile(cpuf, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		akLog.FmtPrintln("mem pprof open fail, err: ", err)
		return
	}
	return f
}
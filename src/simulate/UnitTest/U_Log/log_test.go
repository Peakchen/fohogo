package U_Log

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"sync"
	"testing"
)

func TestLogNormal(t *testing.T) {
	akLog.FmtPrintln("test log: ", "yes")
	akLog.Error("test error.")
}

func TestLogLoopWrite(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			akLog.Info("info, idx: %v.", i)
		}
	}()

	//time.Sleep(time.Duration(30 * time.Second))
	wg.Wait()
}

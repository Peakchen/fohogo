package U_Channel

import (
	"sync"
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
)

func TestNormalChannel(t *testing.T) {
	akLog.FmtPrintln("test normal channel...")
	var (
		maxLoop = 1000
		nch     = make(chan int, maxLoop)
		count   int
	)
	var sw sync.WaitGroup
	sw.Add(2)
	go func() {
		for {
			akLog.FmtPrintln("write ch: ", count)
			nch <- count
			if count >= maxLoop {
				break
			}
			count++
		}
		sw.Done()
		akLog.FmtPrintln("===write end===")
	}()
	go func() {
		for {
			select {
			case v := <-nch:
				akLog.FmtPrintln("recv ch: ", v)
			default:
				//akLog.FmtPrintln("nothing...")
				break
			}
		}
		sw.Done()
		akLog.FmtPrintln("===read end===")
	}()
	sw.Wait()
}

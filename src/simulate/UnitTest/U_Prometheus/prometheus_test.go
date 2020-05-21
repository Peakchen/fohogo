package U_Prometheus

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"testing"
	"time"
)

func TestPrometheus(t *testing.T) {
	//初始一个http handler
	http.Handle("/metrics", promhttp.Handler())

	//初始化一个容器
	diskPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memeory_percent",
		Help: "memeory use percent",
	},
		[]string{"percent"},
	)
	prometheus.MustRegister(diskPercent)

	// 启动web服务，监听9090端口
	go func() {
		akLog.FmtPrintln("ListenAndServe at:localhost:9090")
		err := http.ListenAndServe("localhost:9090", nil)
		if err != nil {
			akLog.Error("ListenAndServe: ", err)
		}
	}()

	//收集内存使用的百分比
	for {
		akLog.FmtPrintln("start collect memory used percent!")
		v, err := mem.VirtualMemory()
		if err != nil {
			akLog.FmtPrintln("get memeory use percent error:%s", err)
			continue
		}

		akLog.FmtPrintln("get memeory use percent:", v.UsedPercent)
		diskPercent.WithLabelValues("usedMemory").Set(v.UsedPercent)
		time.Sleep(time.Second * 5)
	}
}

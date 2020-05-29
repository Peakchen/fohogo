module github.com/Peakchen/fohogo

go 1.12

// github link latest
// for example: github.com/pkg/sftp latest
// go clean -modcache 清除缓存
// go mod vendor 自动创建vendor 目录
// 更新到某个分支最新的代码 go get github.com/xx/xx@master

require (
	github.com/Peakchen/xgameCommon v0.0.6-0.20200527031411-9e96268ad551
	github.com/Shopify/sarama v1.26.4 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/bsm/sarama-cluster v2.1.15+incompatible // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gonutz/ide v0.0.0-20180502124734-e9fc8c14ed56 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/klauspost/reedsolomon v1.9.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/prometheus/client_golang v1.6.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shirou/gopsutil v2.20.4+incompatible
	github.com/sony/sonyflake v1.0.0 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b // indirect
	github.com/tjfoc/gmsm v1.3.0 // indirect
	github.com/urfave/cli v1.22.4 // indirect
	github.com/xtaci/kcp-go v5.4.20+incompatible // indirect
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
)

replace (
	golang.org/x/arch => github.com/golang/arch v0.0.0-20200312215426-ff8b605520f4
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200429183012-4b2356b1ed79
	golang.org/x/net => github.com/golang/net v0.0.0-20200506145744-7e3656a0809f
	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200509044756-6aff5f38e54f
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200509030707-2212a7e161a5
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
)

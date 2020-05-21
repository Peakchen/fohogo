package AutoTest

// add by stefan

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"simulate/AutoTest/gameMsg"
	"simulate/AutoTest/msgImp"
	"simulate/AutoTest/testconfig"
	"simulate/TestCommon"
	"simulate/UnitTest/U_config"
	"simulate/UnitTest/U_login"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
)

type TAokoTest struct {
	testConf []testconfig.TArrConfig4Test
	ConnConf *testconfig.TGateWay
}

func (this *TAokoTest) init() {
	if this.testConf == nil || len(this.testConf) == 0 {
		this.testConf = []testconfig.TArrConfig4Test{}
	}

	if this.ConnConf == nil {
		this.ConnConf = &testconfig.TGateWay{}
	}
}

var (
	_testobj     *TAokoTest = &TAokoTest{}
	_testCfgPath string
)

func init() {
	_testCfgPath = testconfig.GetTestCfgPath()
	gameMsg.Init()
}

func Start() {
	_testobj.loadAndRun()
}

func (this *TAokoTest) loadAndRun() {
	this.loadTestCheck()
	U_login.Init(this.ConnConf.ConnAddr, "login")
	var (
		sw sync.WaitGroup
	)
	sw.Add(1)
	go this.loopRun()
	sw.Wait()
}

func (this *TAokoTest) loopRun() {
	for idx, _ := range U_config.GloginConfig.Get() {
		this.runItem(idx)
	}
}

func (this *TAokoTest) runItem(idx int) {
	U_login.Run(idx)
	this.gameEnter()
}

func (this *TAokoTest) loadConnCheck(dir, fileName string) {
	fileobj, err := ioutil.ReadFile(_testCfgPath + dir + fileName)
	if err != nil {
		akLog.Error("read config fail, info: ", fileName, err)
		return
	}

	data := &testconfig.TArrGateWay{}
	err = json.Unmarshal(fileobj, &data)
	if err != nil {
		akLog.Error("unmarshal config fail, info: ", fileName, err)
		return
	}

	this.ConnConf = (*data)[0]
}

func (this *TAokoTest) loadTestCheck() {
	this.init()

	rd, err := ioutil.ReadDir(_testCfgPath)
	if err != nil {
		akLog.Error("read test config path fail, err: ", err)
		return
	}

	for _, file := range rd {
		if !file.IsDir() {
			continue
		}

		testrd, err := ioutil.ReadDir(_testCfgPath + file.Name())
		if err != nil {
			akLog.Error("read test config path fail, err: ", err)
			return
		}

		if strings.Contains(file.Name(), "T_Test") {
			continue
		}

		fileName := testrd[0].Name()
		if strings.Contains(fileName, ".json") &&
			strings.Contains(file.Name(), "gateway") {
			this.loadConnCheck(file.Name()+"/", fileName)
			continue
		}

		if !strings.Contains(fileName, ".json") &&
			!strings.Contains(fileName, "T_") {
			continue
		}

		fileobj, err := ioutil.ReadFile(_testCfgPath + file.Name() + "/" + fileName)
		if err != nil {
			akLog.Error("read config fail, info: ", fileName, err)
			return
		}

		// load module test msg
		data := &testconfig.TArrConfig4Test{}
		err = json.Unmarshal(fileobj, &data)
		if err != nil {
			akLog.Error("unmarshal config fail, info: ", fileName, err)
			return
		}

		this.testConf = append(this.testConf, *data)
	}
}

func (this *TAokoTest) gameEnter() {
	for _, data := range this.testConf {
		this.Run(data)
	}
}

func (this *TAokoTest) Run(data testconfig.TArrConfig4Test) {
	if len(data) == 0 {
		return
	}

	ConnAddr := data[0].ConnAddr
	module := data[0].Module
	mainId := data[0].MainId
	src := data[0].Msg
	pack := TestCommon.NewModule(ConnAddr, module)
	for _, msgitem := range src.Msg {
		for _, Item := range msgitem.Data {
			SubId := Item.Request.SubId
			Req := Item.Request.Name
			Params := Item.Request.Params
			akLog.FmtPrintln("data: ", mainId, SubId, Req, Params)
			byparams, err := json.Marshal(Params)
			if err != nil {
				akLog.Error("json marshal fail, err: ", err)
				continue
			}

			_cmd := Kcpnet.EncodeCmd(uint16(mainId), uint16(SubId))
			PbItem, exist := msgImp.GetMsgPb(_cmd)
			if !exist {
				continue
			}

			dst := reflect.New(PbItem.MT.Elem()).Interface()
			if dst == nil {
				return
			}

			err = json.Unmarshal(byparams, dst)
			if err != nil {
				akLog.Error("proto Unmarshal fail, err: ", err)
				continue
			}

			dstpb := dst.(proto.Message)
			pack.PushMsg(uint16(mainId),
				uint16(SubId),
				dstpb)

			akLog.FmtPrintln("pb: ", dstpb)
			go pack.Run()
		}
	}
}

package U_Interface

import (
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
)

/*
	map[key]value 特殊存在
	如下例子：arrC 在for 循环外定义时，赋值并加入arrB中，下次再次赋值时，则会覆盖之前的值
	另外，arrD验证，当它是int，string，struct，结构体指针类型时，赋值加入到arrE不会有问题
*/

func TestInterface(t *testing.T) {
	arrA := []int{1, 2, 3}
	arrB := []interface{}{}
	var arrD interface{}
	arrE := []interface{}{}
	for _, a := range arrA {
		arrC := map[string]interface{}{}
		arrC["item"] = a
		arrB = append(arrB, arrC)

		arrD = &struct {
			A int
		}{a}
		arrE = append(arrE, arrD)
	}
	akLog.FmtPrintln("comfire: ", arrB, arrE)
	for _, e := range arrE {
		akLog.FmtPrintln("e val: ", e)
	}
}

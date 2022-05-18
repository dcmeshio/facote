package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		println(fmt.Sprintf("最后"))
		// 判断是否宕机
		if info := recover(); info != nil {
			println(fmt.Sprintf("触发了宕机，并恢复"))
		} else {
			println(fmt.Sprintf("未触发宕机"))
		}
	}()
	println(fmt.Sprintf("执行"))
	//panic("宕机")
	println(fmt.Sprintf("继续执行"))
}

func panicExecute(param int32) (err error) {
	defer func() {
		// 判断是否宕机
		if info := recover(); info != nil {
			println(fmt.Sprintf("触发了宕机，并恢复: %s", info))
			err = errors.New(fmt.Sprintf("panic: %s", info))
		} else {
			println(fmt.Sprintf("未触发宕机"))
		}
	}()
	println(fmt.Sprintf("宕机前判断：%d", param))
	if param == 0 {
		panic("执行宕机")
	} else {
		err = nil
		return
	}
}

func TestExecute(t *testing.T) {
	err := panicExecute(0)
	if err != nil {
		println(fmt.Sprintf("Error %s", err))
	} else {
		println(fmt.Sprintf("Successful"))
	}
}

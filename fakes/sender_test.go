package fakes

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	err := Send(nil)
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}

}

package fakec

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {

	err := Send(nil, "idimesh.helmsnets.com:5855", 10000, 20000)
	if err != nil {
		println(fmt.Sprintf("%s", err))
	}

}

func TestCreateRequest(t *testing.T) {
	bytes, err := CreateRequest("idimesh.helmsnets.com:5855", 10000, 20000)
	if err != nil {
		println(err)
		return
	}
	println(fmt.Sprintf("%d", bytes))
	println(fmt.Sprintf("%s", string(bytes)))
}

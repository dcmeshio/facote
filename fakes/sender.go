package fakes

import (
	"fmt"
	"github.com/dcmeshio/facote"
	"github.com/dcmeshio/facote/common"
	"io"
	"net/http"
	"time"
)

func Send(writer io.Writer) error {
	option := facote.GetOption()
	builder := CreateBuilder()
	// 请求头 Date
	Date := fmt.Sprintf("Date: %s", time.Now().UTC().Format(http.TimeFormat))
	builder.AddHeader(Date)
	// 请求头 Other
	for _, v := range option.ResponseHeaders {
		Header := ""
		if v.Single {
			Header = fmt.Sprintf("%s: %s", v.Name, v.Value)
		} else {
			index := common.Rint(len(v.Values))
			Header = fmt.Sprintf("%s: %s", v.Name, v.Values[index])
		}
		builder.AddHeader(Header)
	}
	bytes := builder.Build()
	_, err := writer.Write(bytes)
	return err
}

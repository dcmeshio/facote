package fakec

import (
	"errors"
	"github.com/dcmeshio/facote"
	"strings"
)

func Receive(bc *facote.BufferConn) error {
	option := facote.GetOption()
	headers := make(map[string]*facote.Header)
	for _, v := range option.ResponseHeaders {
		headers[v.Name] = v
	}
	ch := &checker{
		bc:      bc,
		headers: headers,
	}
	return ch.check()
}

type checker struct {
	bc      *facote.BufferConn
	headers map[string]*facote.Header // 待校验的 header
}

func (c *checker) check() error {
	// 首行
	firstLine, err := c.bc.ReadBytes(byte(10))
	if err != nil {
		return err
	}
	max := len(firstLine)
	if max < 2 {
		return errors.New("[FakeError] firstLine not enough")
	}
	Line := string(firstLine[:max-2])
	if Line != "HTTP/1.1 200 OK" {
		return errors.New("[FakeError] not http firstLine format")
	}
	// 请求头
	for true {
		header, err := c.bc.ReadBytes(byte(10))
		if err != nil {
			return err
		}
		max = len(header)
		// 结束
		if max == 2 && header[0] == byte(13) {
			if len(c.headers) != 0 {
				return errors.New("[FakeError] incomplete http protocol")
			} else {
				break
			}
		}
		Line = string(header[:max-2])
		// 单独处理 Date
		if strings.HasPrefix(Line, "Date") {
			continue
		}
		// 统一处理请求头
		ok := c.checkHeader(Line)
		if !ok {
			return errors.New("[FakeError] header format error")
		}
	}
	return nil
}

// 校验除 Date 外的请求头
func (c *checker) checkHeader(header string) bool {
	Lines := strings.Split(header, ": ")
	if len(Lines) != 2 {
		return false
	}
	h := c.headers[Lines[0]]
	if h == nil {
		return false
	}
	if h.Single {
		if Lines[1] != h.Value {
			return false
		}
	} else {
		if !facote.In(Lines[1], h.Values) {
			return false
		}
	}
	delete(c.headers, h.Name)
	return true
}

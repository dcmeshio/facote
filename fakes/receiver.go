package fakes

import (
	"errors"
	"github.com/dcmeshio/facote"
	"strings"
)

/*
 * Return
 * []byte forerunner   	到出错时已读的字节内容
 * error err
 */
func Receive(bc *facote.BufferConn) (int, int, []byte, error) {
	option := facote.GetOption()
	headers := make(map[string]*facote.Header)
	for _, v := range option.RequestHeaders {
		headers[v.Name] = v
	}
	ch := &checker{
		bc:      bc,
		headers: headers,
		checked: false,
	}
	return ch.check()
}

type checker struct {
	bc         *facote.BufferConn
	headers    map[string]*facote.Header // 待校验的 header
	forerunner []byte                    // 已读的字节数据
	usercode   int
	quotecode  int
	checked    bool // 是否完成口令校验
}

func (c *checker) check() (int, int, []byte, error) {
	// 首行方法
	buf := make([]byte, 1)
	for i := 0; i < 4; i++ {
		_, err := c.bc.Read(buf)
		if err != nil {
			return 0, 0, c.forerunner, err
		}
		c.forerunner = append(c.forerunner, buf...)
		ch := string(buf)
		if i == 0 && ch != "G" {
			return 0, 0, c.forerunner, errors.New("[FakeError] not GET method")
		}
		if i == 1 && ch != "E" {
			return 0, 0, c.forerunner, errors.New("[FakeError] not GET method")
		}
		if i == 2 && ch != "T" {
			return 0, 0, c.forerunner, errors.New("[FakeError] not GET method")
		}
		if i == 3 && ch != " " {
			return 0, 0, c.forerunner, errors.New("[FakeError] not GET method")
		}
	}
	// 首行其他
	firstLine, err := c.bc.ReadBytes(byte(10))
	c.forerunner = append(c.forerunner, firstLine...)
	if err != nil {
		return 0, 0, c.forerunner, err
	}
	max := len(firstLine)
	if max < 2 {
		return 0, 0, c.forerunner, errors.New("[FakeError] firstLine not enough")
	}
	Line := string(firstLine[:max-2])
	if !strings.HasSuffix(Line, "HTTP/1.1") {
		return 0, 0, c.forerunner, errors.New("[FakeError] not HTTP/1.1")
	}
	Lines := strings.Split(Line, " ")
	if len(Lines) != 2 {
		return 0, 0, c.forerunner, errors.New("[FakeError] not http firstLine format")
	}
	ok := c.checkFirstLine(Lines[0])
	if !ok {
		return 0, 0, c.forerunner, errors.New("[FakeError] firstLine path or param error")
	}
	// 请求头
	for true {
		header, err := c.bc.ReadBytes(byte(10))
		if err != nil {
			return 0, 0, c.forerunner, err
		}
		c.forerunner = append(c.forerunner, header...)
		max = len(header)
		// 结束
		if max == 2 && header[0] == byte(13) {
			if !c.checked {
				return 0, 0, c.forerunner, errors.New("[FakeError] token not checked")
			}
			if len(c.headers) != 0 {
				return 0, 0, c.forerunner, errors.New("[FakeError] incomplete http protocol")
			} else {
				return c.usercode, c.quotecode, nil, nil
			}
		}
		Line = string(header[:max-2])
		// 单独处理 Host
		if strings.HasPrefix(Line, "Host") {
			continue
		}
		// 单独处理 鉴权
		if strings.HasPrefix(Line, "X-token") || strings.HasPrefix(Line, "Ps") {
			ok = c.checkToken(Line)
			if !ok {
				return 0, 0, c.forerunner, errors.New("[FakeError] token error")
			}
			continue
		}
		// 统一处理请求头
		ok = c.checkHeader(Line)
		if !ok {
			return 0, 0, c.forerunner, errors.New("[FakeError] header format error")
		}
	}
	return 0, 0, c.forerunner, errors.New("[FakeError] Unknown")
}

// 传入的是去除 GET 头和 HTTP/1.1 尾的字符串数据
func (c *checker) checkFirstLine(value string) bool {
	option := facote.GetOption()
	values := strings.Split(value, "?")
	if len(values) != 2 {
		return false
	}
	ok := false
	for _, v := range option.Lines {
		if values[0] == v.PathName && strings.HasPrefix(values[1], v.ParamName) {
			ok = true
			break
		}
	}
	return ok
}

// 校验除 Host、Token 外的请求头
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

// 鉴权
func (c *checker) checkToken(header string) bool {
	Lines := strings.Split(header, ": ")
	if len(Lines) != 2 {
		return false
	}
	uc, qc, err := facote.CheckToken(Lines[1])
	if err != nil {
		return false
	}
	c.usercode = uc
	c.quotecode = qc
	c.checked = true
	return true
}

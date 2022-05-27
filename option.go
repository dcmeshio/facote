package facote

import (
	"sync"
)

var option *Option

type FirstLine struct {
	PathName  string
	ParamName string
}

type Header struct {
	Name   string
	Single bool // 是单个值或者多个值
	Value  string
	Values []string
}

// 根据 option 创建和检验
type Option struct {
	mu                  sync.Mutex
	Lines               []*FirstLine // 请求首行
	RequestHeaders      []*Header    // 请求头
	ResponseHeaders     []*Header    // 应答头
	TimestampDifference int64        // 客户端与服务器的时差
	Key                 []byte       // 乱序秘钥 256 个乱序字节码 21 .... 91, 93 ... 125
	Unkey               []byte
	Ps                  string // 密码，传输内容中包含的字段值，在解密数据后的校验
}

func SetOption(opt *Option) {
	option = opt
}

func GetOption() *Option {
	return option
}

func CreateOption(key []byte, ps string) *Option {
	opt := &Option{
		Lines:               make([]*FirstLine, 0),
		RequestHeaders:      make([]*Header, 0),
		ResponseHeaders:     make([]*Header, 0),
		TimestampDifference: 0,
		Key:                 key,
		Ps:                  ps,
	}
	bytes := make([]byte, 256)
	for i := 0; i < len(key); i++ {
		index := key[i]
		bytes[index] = byte(i)
	}
	opt.Unkey = bytes
	return opt
}

func (o *Option) AddFirstLine(line *FirstLine) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.Lines = append(o.Lines, line)
}

func (o *Option) AddRequestHeader(header *Header) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.RequestHeaders = append(o.RequestHeaders, header)
}

func (o *Option) AddResponseHeader(header *Header) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.ResponseHeaders = append(o.ResponseHeaders, header)
}

func (o *Option) SetTimestampDifference(td int64) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.TimestampDifference = td
}

func In(str string, strs []string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}

func init() {
	opt := defaultOption()
	SetOption(opt)
}

func defaultOption() *Option {

	// Lines
	Lines := make([]*FirstLine, 0)
	FirstLineA := &FirstLine{
		PathName:  "/queryA",
		ParamName: "aParam",
	}
	Lines = append(Lines, FirstLineA)
	FirstLineB := &FirstLine{
		PathName:  "/queryB",
		ParamName: "bParam",
	}
	Lines = append(Lines, FirstLineB)
	FirstLineC := &FirstLine{
		PathName:  "/queryC",
		ParamName: "cParam",
	}
	Lines = append(Lines, FirstLineC)

	// RequestHeaders
	RequestHeaders := make([]*Header, 0)
	UserAgent := &Header{
		Name:   "User-Agent",
		Single: false,
		Values: []string{"Dart/2.14 (dart:io)", "Go/1.16"},
	}
	RequestHeaders = append(RequestHeaders, UserAgent)
	AcceptEncoding := &Header{
		Name:   "Accept-Encoding",
		Single: true,
		Value:  "gzip, deflate",
	}
	RequestHeaders = append(RequestHeaders, AcceptEncoding)
	ConnectionA := &Header{
		Name:   "Connection",
		Single: true,
		Value:  "keep-alive",
	}
	RequestHeaders = append(RequestHeaders, ConnectionA)

	// ResponseHeaders
	ResponseHeaders := make([]*Header, 0)
	ContentType := &Header{
		Name:   "Content-Type",
		Single: false,
		Values: []string{"application/octet-stream", "video/mpeg", "video/mpeg4", "audio/wav"},
	}
	ResponseHeaders = append(ResponseHeaders, ContentType)
	TransferEncoding := &Header{
		Name:   "Transfer-Encoding",
		Single: true,
		Value:  "chunked",
	}
	ResponseHeaders = append(ResponseHeaders, TransferEncoding)
	ConnectionB := &Header{
		Name:   "Connection",
		Single: true,
		Value:  "keep-alive",
	}
	ResponseHeaders = append(ResponseHeaders, ConnectionB)
	Server := &Header{
		Name:   "Server",
		Single: true,
		Value:  "nginx/1.21.2",
	}
	ResponseHeaders = append(ResponseHeaders, Server)

	opt := &Option{
		Lines:               Lines,
		RequestHeaders:      RequestHeaders,
		ResponseHeaders:     ResponseHeaders,
		TimestampDifference: 0,
		Key:                 []byte{15, 60, 88, 72, 34, 58, 22, 70, 49, 73, 117, 122, 37, 7, 71, 102, 66, 20, 123, 47, 69, 0, 107, 96, 16, 21, 83, 124, 126, 8, 92, 52, 79, 97, 74, 85, 76, 29, 82, 1, 30, 111, 84, 56, 18, 19, 94, 91, 53, 90, 112, 100, 87, 65, 23, 14, 5, 89, 3, 86, 42, 35, 13, 98, 44, 101, 67, 103, 31, 61, 75, 4, 26, 106, 41, 24, 77, 57, 36, 78, 33, 95, 127, 99, 109, 48, 62, 54, 105, 46, 81, 59, 104, 63, 108, 9, 27, 80, 40, 43, 114, 110, 25, 118, 11, 120, 64, 38, 113, 116, 55, 50, 115, 68, 32, 6, 121, 39, 125, 45, 10, 17, 51, 2, 93, 28, 119, 12},
		Ps:                  "aaa",
	}

	bytes := make([]byte, len(opt.Key))
	for i := 0; i < len(opt.Key); i++ {
		index := opt.Key[i]
		bytes[index] = byte(i)
	}
	opt.Unkey = bytes
	return opt

}

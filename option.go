package facote

import "sync"

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
	Type                int          // 区分项目
	Key                 string       // AES 秘钥，16位、32位
	KeyBytes            []byte       // 乱序秘钥
}

func SetOption(opt *Option) {
	option = opt
}

func GetOption() *Option {
	return option
}

func CreateOption(codeType int, key string) *Option {
	return &Option{
		Lines:               make([]*FirstLine, 0),
		RequestHeaders:      make([]*Header, 0),
		ResponseHeaders:     make([]*Header, 0),
		TimestampDifference: 0,
		Type:                codeType,
		Key:                 key,
	}
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

	return &Option{
		Lines:               Lines,
		RequestHeaders:      RequestHeaders,
		ResponseHeaders:     ResponseHeaders,
		Type:                1,
		TimestampDifference: 0,
		Key:                 "where1is23macwin",
	}

}

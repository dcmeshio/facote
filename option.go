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
	Key                 []byte       // 乱序秘钥 256 个乱序字节码
	Unkey               []byte
	Ps                  string // 密码，传输内容中包含的字段值，在解密数据后的校验
}

func SetOption(opt *Option) {
	option = opt
}

func GetOption() *Option {
	return option
}

func CreateOption(codeType int, key []byte) *Option {
	opt := &Option{
		Lines:               make([]*FirstLine, 0),
		RequestHeaders:      make([]*Header, 0),
		ResponseHeaders:     make([]*Header, 0),
		TimestampDifference: 0,
		Type:                codeType,
		Key:                 key,
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
		Type:                1,
		TimestampDifference: 0,
		Key:                 []byte{235, 204, 188, 68, 203, 3, 121, 162, 5, 70, 139, 216, 236, 191, 102, 220, 79, 93, 34, 234, 153, 167, 184, 91, 232, 174, 119, 254, 84, 108, 182, 113, 39, 179, 189, 82, 137, 1, 106, 28, 141, 155, 149, 109, 152, 98, 252, 104, 53, 249, 90, 205, 138, 187, 125, 63, 23, 118, 66, 25, 0, 31, 100, 46, 123, 74, 227, 75, 169, 211, 197, 157, 165, 247, 17, 180, 166, 241, 16, 177, 242, 136, 251, 185, 114, 172, 37, 218, 30, 186, 36, 22, 9, 47, 50, 55, 107, 120, 51, 111, 44, 54, 210, 32, 4, 105, 229, 77, 76, 112, 206, 248, 150, 78, 147, 58, 101, 170, 134, 221, 183, 94, 224, 117, 228, 26, 126, 156, 159, 253, 176, 29, 52, 81, 196, 225, 8, 144, 42, 20, 226, 62, 223, 217, 64, 140, 142, 18, 238, 243, 231, 161, 88, 163, 71, 59, 250, 87, 145, 21, 246, 10, 173, 116, 190, 14, 6, 96, 175, 193, 131, 85, 178, 146, 171, 245, 83, 45, 194, 160, 164, 240, 130, 86, 7, 89, 230, 122, 57, 192, 12, 237, 244, 11, 201, 132, 48, 202, 127, 33, 67, 27, 40, 214, 60, 80, 13, 15, 19, 198, 199, 208, 222, 38, 72, 56, 148, 95, 128, 215, 209, 129, 97, 181, 103, 239, 61, 65, 41, 99, 110, 143, 43, 233, 151, 133, 2, 255, 135, 219, 213, 212, 69, 154, 158, 124, 115, 35, 92, 195, 168, 207, 73, 200, 49, 24},
	}

	bytes := make([]byte, 256)
	for i := 0; i < len(opt.Key); i++ {
		index := opt.Key[i]
		bytes[index] = byte(i)
	}
	opt.Unkey = bytes
	return opt

}

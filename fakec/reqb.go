package fakec

type RequestBuilder struct {
	firstLine string
	headers   []string
}

func CreateBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (rb *RequestBuilder) AddFirstLine(firstLine string) *RequestBuilder {
	rb.firstLine = firstLine
	return rb
}

func (rb *RequestBuilder) AddHeader(header string) *RequestBuilder {
	rb.headers = append(rb.headers, header)
	return rb
}

func (rb *RequestBuilder) Build() []byte {
	last := []byte{13, 10}
	bytes := make([]byte, 0)
	bytes = append(bytes, []byte(rb.firstLine)...)
	bytes = append(bytes, last...)
	for _, v := range rb.headers {
		bytes = append(bytes, []byte(v)...)
		bytes = append(bytes, last...)
	}
	bytes = append(bytes, last...)
	return bytes
}

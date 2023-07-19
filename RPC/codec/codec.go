package codec

import "io"

type Header struct {
	ServiceMethod string //Service.Method
	Seq           uint64
	Error         string
}

// 抽象出对消息体进行编解码的接口 Code
type Codec interface {
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
	io.Closer // response + close
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}

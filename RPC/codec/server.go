package codec

import (
	"encoding/json"
	"io"
	"log"
	"net"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	ContentType Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	ContentType: GobType,
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (s *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln("rpc server:accept error:", err)
		}
		go s.ServerConn(conn)
	}
}
func Accept(lis net.Listener) { DefaultServer.Accept(lis) }
func (s *Server) ServerConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server:options error ", err)
	}
	if opt.MagicNumber != MagicNumber {
		return
	}
	f := NewCodecFuncMap[opt.ContentType]
	if f == nil {
		return
	}
	Server.ServeCodec(f(conn))
}

func (server *Server) ServeCodec(cc Codec) {

}

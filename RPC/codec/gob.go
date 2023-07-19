package codec

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
)

type GobCodec struct {
	conn io.ReadWriteCloser //由构造函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例
	buf  *bufio.Writer      //带缓冲的Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(header *Header) error {
	fmt.Print("ReadHeader")
	return c.dec.Decode(header)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	fmt.Print("ReadBody")
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(header *Header, body interface{}) (err error) {
	fmt.Print("Write")
	return
}
func (c *GobCodec) Close() error {
	return c.conn.Close()
}

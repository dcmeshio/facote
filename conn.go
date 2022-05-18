package facote

import (
	"bufio"
	"net"
	"time"
)

type BufferConn struct {
	conn net.Conn
	br   *bufio.Reader
}

func NewBufferConn(conn net.Conn) *BufferConn {
	return &BufferConn{
		conn: conn,
		br:   bufio.NewReader(conn),
	}
}

func (bc *BufferConn) ReadBytes(delim byte) ([]byte, error) {
	return bc.br.ReadBytes(delim)
}

func (bc *BufferConn) Read(b []byte) (n int, err error) {
	return bc.br.Read(b)
}

func (bc *BufferConn) Write(b []byte) (n int, err error) {
	return bc.conn.Write(b)
}

func (bc *BufferConn) Close() error {
	return bc.conn.Close()
}

func (bc *BufferConn) LocalAddr() net.Addr {
	return bc.conn.LocalAddr()
}

func (bc *BufferConn) RemoteAddr() net.Addr {
	return bc.conn.RemoteAddr()
}

func (bc *BufferConn) SetDeadline(t time.Time) error {
	return bc.conn.SetDeadline(t)
}

func (bc *BufferConn) SetReadDeadline(t time.Time) error {
	return bc.conn.SetReadDeadline(t)
}

func (bc *BufferConn) SetWriteDeadline(t time.Time) error {
	return bc.conn.SetWriteDeadline(t)
}

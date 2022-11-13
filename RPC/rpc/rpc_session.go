package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

// Session 是 RPC 通信的一个会话连接
type Session struct {
	conn net.Conn
}

// NewSession 从网络连接新建一个 Session
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// Write 向 Session 中写数据
// 写入数据格式为：
//    4字节头(Header) + 具体数据(Data)
// 其中， Header 为 uint32 (大端法表示)，表示 Data 的长度
// Data 即传入的参数 data，为任意的 []byte，
func (s *Session) Write(data []byte) error {
	buf := make([]byte, 4+len(data))
	// Header
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// Data
	copy(buf[4:], data)

	_, err := s.conn.Write(buf)

	return err
}

// Read 从 Session 中读数据
// 读取数据格式为：
//    4字节头(Header) + 具体数据(Data)
// 该函数返回读取出的 Data
func (s *Session) Read() ([]byte, error) {
	// 读取 Header，获取 Data 长度信息
	header := make([]byte, 4)
	if _, err := io.ReadFull(s.conn, header); err != nil {
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(header)

	// 按照 dataLen 读取 Data
	data := make([]byte, dataLen)
	if _, err := io.ReadFull(s.conn, data); err != nil {
		return nil, err
	}
	return data, nil
}

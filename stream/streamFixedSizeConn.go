package stream

import "net"

type StreamFixedSizeConn struct {
	size int
	net.Conn
}

func NewStreamFixedSizeConn(size int, conn net.Conn) net.Conn {
	return &StreamFixedSizeConn{size: size, Conn: conn}
}

func UnPaddingStream(data []byte) (n int, realData []byte, err error) {
	return 0, nil, nil
}

func PaddingStream(data []byte) (n int, paddingData []byte, err error) {
	return 0, nil, nil
}

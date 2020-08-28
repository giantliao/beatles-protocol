package stream

import "net"

func NewCipherSteamConn(conn net.Conn, aesKey [32]byte) net.Conn {
	sconn := &StreamConn{Conn: conn}
	c, err := NewCipherConn(sconn, aesKey)
	if err != nil {
		return nil
	}
	return c
}

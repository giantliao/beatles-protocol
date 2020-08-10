package stream

import "net"

func NewCipherSteamConn(conn net.Conn, aesKey [32]byte) net.Conn {
	sconn := &StreamConn{Conn: conn}
	return NewCipherConn(sconn, aesKey)
}

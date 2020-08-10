package stream

import (
	"github.com/kprc/libeth/util"
	"github.com/pkg/errors"
	"io"
	"net"
)



type CipherConn struct {
	aeskey [32]byte
	net.Conn
}

func NewCipherConn(conn net.Conn,key [32]byte) (net.Conn)  {
	return &CipherConn{aeskey: key,Conn:conn}
}

func (cc *CipherConn)Read(b []byte)(n int, err error)  {
	n,err = cc.Conn.Read(b)
	if (err!=nil && err!=io.EOF) || n == 0{
		return 0,errors.New("Read error")
	}

	if err == io.EOF && n == 0{
		return 0,io.EOF
	}

	_,err = util.Decrypt(cc.aeskey[:],b[:n])
	if err!=nil{
		return 0, err
	}

	return n,nil
}

func (cc *CipherConn)Write(b []byte) (n int,err error)  {
	cipherTxt,err:=util.Encrypt(cc.aeskey[:],b)
	if err!=nil{
		return 0,err
	}
	lc:=len(cipherTxt)

	n,err = cc.Conn.Write(cipherTxt)
	if (err!=nil && err!=io.EOF) || lc != n{
		return 0,errors.New("write cipherTxt error")
	}

	return
}


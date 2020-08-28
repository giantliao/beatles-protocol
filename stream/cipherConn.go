package stream

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/kprc/libeth/util"
	"github.com/pkg/errors"
	"io"
	"net"
)

type CipherConn struct {
	aeskey [32]byte
	net.Conn
	encStream cipher.Stream
	decStream cipher.Stream
	iv        [aes.BlockSize]byte
}

func NewCipherConn(conn net.Conn, key [32]byte) (net.Conn, error) {
	cc := &CipherConn{aeskey: key, Conn: conn}
	var err error
	cc.encStream, cc.iv, err = util.NewEncStream(key[:])
	if err != nil {
		return nil, err
	}

	cc.decStream, err = util.NewDecStreamWithIv(key[:], cc.iv)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (cc *CipherConn) GetIV() [aes.BlockSize]byte {
	return cc.iv
}

func NewCipherConnWithIv(conn net.Conn, key [32]byte, iv [aes.BlockSize]byte) (net.Conn, error) {
	cc := &CipherConn{aeskey: key, Conn: conn}
	var err error
	cc.encStream, err = util.NewEncStreamWithIv(key[:], iv)
	if err != nil {
		return nil, err
	}

	if cc.decStream, err = util.NewDecStreamWithIv(key[:], iv); err != nil {
		return nil, err
	}

	cc.iv = iv

	return cc, nil
}

func (cc *CipherConn) Read(b []byte) (n int, err error) {
	n, err = cc.Conn.Read(b)
	if (err != nil && err != io.EOF) || n == 0 {
		return 0, errors.New("Read error")
	}

	if err == io.EOF && n == 0 {
		return 0, io.EOF
	}
	//fmt.Println("read cipher:",hex.EncodeToString(b[:n]))
	var plaintxt []byte
	plaintxt = util.Decrypt2(cc.decStream, b[:n])

	//fmt.Println("read plain:",hex.EncodeToString(plaintxt))

	return len(plaintxt), nil
}

func (cc *CipherConn) Write(b []byte) (n int, err error) {

	//fmt.Println("wirte plain:",hex.EncodeToString(b))

	cipherTxt := util.Encrypt2(cc.encStream, b)

	lc := len(cipherTxt)
	//fmt.Println("write cipher:",hex.EncodeToString(cipherTxt))
	n, err = cc.Conn.Write(cipherTxt)
	if (err != nil && err != io.EOF) || lc != n {
		return 0, errors.New("write cipherTxt error")
	}

	return
}

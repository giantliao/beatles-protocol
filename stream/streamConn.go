package stream

import (
	"encoding/binary"
	"errors"
	"github.com/kprc/nbsnetwork/tools"
	"io"
	"math/rand"
	"net"
	"strconv"
)

type StreamConn struct {
	net.Conn
}

func NewStreamBuf() []byte {
	buf := make([]byte, bufferSize)

	return buf
}

func ShadowLen(l int32) (int, error) {
	if l > bufferSize {
		return 0, errors.New("length must less than :" + strconv.Itoa(int(bufferSize)))
	}

	rand.Seed(tools.GetNowMsTime())
	randn := rand.Intn(int(lengthMagicMax))

	l = l ^ int32(randn)

	shadowlen := (randn << int(bufferBitsSize)) | int(l)

	return shadowlen, nil
}

func RealLen(shadowLen int32) int32 {
	l := shadowLen & ((1 << bufferBitsSize) - 1)

	randn := (shadowLen - l) >> bufferBitsSize

	l = l ^ randn

	return l
}

func (sc *StreamConn) Read(b []byte) (n int, err error) {
	headLenBuff := make([]byte, int32Size)

	var hl int

	hl, err = tools.SafeRead(sc.Conn, headLenBuff)
	if (err != nil && err != io.EOF) || int32(hl) != int32Size {
		return 0, errors.New("read buffer len error")
	}
	shal := binary.BigEndian.Uint32(headLenBuff)
	l := int(RealLen(int32(shal)))

	if len(b) < l {
		panic("buffer size too small...")
	}

	n, err = tools.SafeRead(sc.Conn, b[:l])
	if (err != nil && err != io.EOF) || n != l {
		return 0, errors.New("read buffer error")
	}

	return

}

func (sc *StreamConn) Write(b []byte) (n int, err error) {
	l := len(b)
	var shal int
	shal, err = ShadowLen(int32(l))

	headLenbuf := make([]byte, int32Size)
	binary.BigEndian.PutUint32(headLenbuf, uint32(shal))

	n, err = sc.Conn.Write(headLenbuf)
	if err != nil && int32(n) != int32Size {
		return 0, errors.New("write buffer len error")
	}

	n, err = sc.Conn.Write(b)
	if err != nil && n != l {
		return 0, errors.New("write buffer error")
	}

	return
}

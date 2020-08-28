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

func ShadowLen(l int) (uint32, error) {
	if int32(l) > bufferSize {
		return 0, errors.New("length must less than :" + strconv.Itoa(int(bufferSize)))
	}

	rand.Seed(tools.GetNowMsTime())
	randn := rand.Intn(int(lengthMagicMax))

	ll := uint32(l) ^ uint32(randn)

	shadowlen := (uint32(randn) << int(bufferBitsSize)) | ll

	return shadowlen, nil
}

func RealLen(shadowLen uint32) int {
	l := shadowLen & ((1 << bufferBitsSize) - 1)

	randn := (shadowLen - l) >> bufferBitsSize

	l = l ^ randn

	return int(l)
}

func (sc *StreamConn) Read(b []byte) (n int, err error) {
	headLenBuff := make([]byte, int32Size)

	var hl int

	hl, err = tools.SafeRead(sc.Conn, headLenBuff)
	if (err != nil && err != io.EOF) || int32(hl) != int32Size {
		return 0, errors.New("read buffer len error")
	}

	shal := binary.BigEndian.Uint32(headLenBuff)
	l := RealLen(shal)
	//log.Println("---->",hex.EncodeToString(headLenBuff),l)
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
	var shal uint32
	shal, _ = ShadowLen(l)

	headLenbuf := make([]byte, int32Size)
	binary.BigEndian.PutUint32(headLenbuf, uint32(shal))
	//log.Println("=====>",hex.EncodeToString(headLenbuf),l)
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

package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/giantliao/beatles-protocol/stream"
)

func main() {
	sl, _ := stream.ShadowLen(221233)

	fmt.Println(sl)

	l := stream.RealLen(sl)

	fmt.Println(l)

	buf := make([]byte, 4)

	binary.BigEndian.PutUint32(buf, uint32(sl))

	fmt.Println(hex.EncodeToString(buf))

}

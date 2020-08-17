package token

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
)

func TokenCovert(tk string) (string,error) {
	tkbytes,err:=hex.DecodeString(tk)
	if err!=nil{
		return "", err
	}

	return "at"+base58.Encode(tkbytes),nil
}

func TokenRevert(tk string) string {

	tkbytes := base58.Decode(tk[2:])

	return hex.EncodeToString(tkbytes)
}






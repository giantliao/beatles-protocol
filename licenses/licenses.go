package licenses

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kprc/libeth/account"
	"github.com/kprc/libeth/util"
)

type LicenseContent struct {
	Provider   account.BeatleAddress `json:"provider"`
	Receiver   account.BeatleAddress `json:"receiver"`
	Name       string                `json:"name"`
	Email      string                `json:"email"`
	Cell       string                `json:"cell"`
	ExpireTime int64                 `json:"expire_time"`
}

type License struct {
	Signature string         `json:"signature"`
	Content   LicenseContent `json:"content"`
}

type LicenseRenew struct {
	Receiver       account.BeatleAddress `json:"receiver"`
	EthAddress     common.Address        `json:"eth_address"`
	EthTransaction common.Hash           `json:"eth_transaction"`
	Name           string                `json:"name"`
	Email          string                `json:"email"`
	Cell           string                `json:"cell"`
	CurrentPrice   float64               `json:"current_price"`
	Month          int64                 `json:"month"`
}

func (lr *LicenseRenew) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*lr)
	return util.Encrypt(key, j)
}

func (lr *LicenseRenew) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(plainTxt, lr)
	if err != nil {
		return err
	}

	return nil
}

func (l *License) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*l)
	return util.Encrypt(key, j)
}

func (l *License) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(plainTxt, l)
	if err != nil {
		return err
	}

	return nil
}

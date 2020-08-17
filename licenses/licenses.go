package licenses

import (
	"crypto/ed25519"
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
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

type NoncePrice struct {
	Nonce    uint64                `json:"nonce"`
	Receiver account.BeatleAddress `json:"receiver"`
	EthAddr  common.Address        `json:"eth_addr"`
	Month    int64                 `json:"month"`
}

func (np *NoncePrice) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*np)
	return util.Encrypt(key, j)
}

func (np *NoncePrice) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(plainTxt, np)
	if err != nil {
		return err
	}

	return nil
}

//Total = PricePerMonth * Month
//TotalEth = Total/EthPrice
type NoncePriceContent struct {
	Nonce         uint64                `json:"nonce"`
	Receiver      account.BeatleAddress `json:"receiver"`
	EthAddr       common.Address        `json:"eth_addr"`
	PricePerMonth float64               `json:"price_per_month"`
	Month         int64                 `json:"month"`
	Total         float64               `json:"total"`
	TotalEth      float64               `json:"total_eth"`
	EthPrice      float64               `json:"eth_price"`
}

type NoncePriceSig struct {
	Sig     string            `json:"sig"`
	Content NoncePriceContent `json:"content"`
}

func (nps *NoncePriceSig) Sign(sig func([]byte) []byte) error {
	j, err := json.Marshal(nps.Content)
	if err != nil {
		return err
	}
	sigbyte := sig(j)

	nps.Sig = base58.Encode(sigbyte)

	return nil
}

func (nps *NoncePriceSig) ValidSig(pk ed25519.PublicKey) bool {
	j, _ := json.Marshal(nps.Content)

	return ed25519.Verify(pk, j, base58.Decode(nps.Sig))
}

func (nps *NoncePriceSig) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*nps)
	return util.Encrypt(key, j)
}

func (nps *NoncePriceSig) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(plainTxt, nps)
	if err != nil {
		return err
	}

	return nil
}

type LicenseRenew struct {
	TXSig          NoncePriceSig `json:"tx_sig"`
	EthTransaction common.Hash   `json:"eth_transaction"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Cell           string        `json:"cell"`
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

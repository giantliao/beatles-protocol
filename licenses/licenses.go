package licenses

import (
	"crypto/ed25519"
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kprc/libeth/account"
	"github.com/kprc/libeth/util"
	"github.com/kprc/nbsnetwork/tools"
	"strconv"
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

func (l *License) String() string {
	msg := ""
	msg += "sig: " + l.Signature

	msg += "\r\n" + "Receiver: " + l.Content.Receiver.String()
	msg += "\r\n" + "Provider: " + l.Content.Provider.String()
	msg += "\r\n" + "Name: " + l.Content.Name
	msg += "\r\n" + "Email: " + l.Content.Email
	msg += "\r\n" + "Cell: " + l.Content.Cell
	msg += "\r\n" + "Expire: " + tools.Int64Time2String(l.Content.ExpireTime)

	return msg
}

type NoncePrice struct {
	Nonce    uint64                `json:"nonce"`
	Receiver account.BeatleAddress `json:"receiver"`
	Payer    common.Address        `json:"payer"`
	PayTyp   int 			   		`json:"pay_typ"`
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


const(
	PayTypETH  int = iota
	PayTypBTLC
	PayTypTRX
	PayTypTrxUsd
)

var gPayType []string = []string{"eth","btlc","trx","trx-usdt"}


//Total = PricePerMonth * Month
//TotalEth = Total/EthPrice
type NoncePriceContent struct {
	Nonce         	uint64                `json:"nonce"`
	Receiver      	account.BeatleAddress `json:"receiver"`
	Payer         	common.Address        `json:"payer"`
	PricePerMonth 	float64               `json:"price_per_month"`
	Month         	int64                 `json:"month"`
	Total         	float64               `json:"total"`
	TotalPrice      float64               `json:"total_price"`
	MarketPrice      float64              `json:"market_price"`
	PayTyp           int               	   `json:"pay_typ"`
}

type NoncePriceSig struct {
	Sig     string            `json:"sig"`
	Content NoncePriceContent `json:"content"`
}

func float64toString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func (nps *NoncePriceSig) String() string {
	msg := ""
	msg += "sig: " + nps.Sig
	msg += "\r\n" + "nonce: " + strconv.FormatUint(nps.Content.Nonce, 10)
	msg += "\r\n" + "Receiver: " + nps.Content.Receiver.String()
	msg += "\r\n" + "Payer: " + nps.Content.Payer.String()
	msg += "\r\n" + "PricePerMonth: " + float64toString(nps.Content.PricePerMonth)
	msg += "\r\n" + "Month: " + strconv.FormatInt(nps.Content.Month, 10)
	msg += "\r\n" + "Total: " + float64toString(nps.Content.Total)
	msg += "\r\n" + "TotalPrice: " + float64toString(nps.Content.TotalPrice)
	msg += "\r\n" + "MarketPrice: " + float64toString(nps.Content.MarketPrice)
	msg += "\r\n" + "PayType:" + gPayType[nps.Content.PayTyp]

	return msg
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
	TxStr          string        `json:"tx_str"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Cell           string        `json:"cell"`
}


func (lr *LicenseRenew) String() string {
	msg := ""
	msg += "\r\n" + "EthTransaction: " + lr.TxStr
	msg += "\r\n" + "Name: " + lr.Name
	msg += "\r\n" + "Email: " + lr.Email
	msg += "\r\n" + "Cell: " + lr.Cell
	msg += "\r\n" + "TxSig: \r\n" + (&lr.TXSig).String()

	return msg
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

type FreshLicenseReq struct {
	Receiver account.BeatleAddress
}

func (fl *FreshLicenseReq)Marshal(key []byte) ([]byte, error)  {
	j,_ := json.Marshal(*fl)
	return util.Encrypt(key,j)
}

func (fl *FreshLicenseReq)UnMarshal(key ,data []byte) error  {
	plainTxt, err := util.Decrypt(key,data)
	if err!=nil {
		return err
	}

	err = json.Unmarshal(plainTxt,fl)
	if err!=nil{
		return err
	}

	return nil
}

type FreshLicensResult struct {
	TxStr string `json:"tx_str"`
	License
}


func (flr *FreshLicensResult)Marshal(key []byte) ([]byte, error)  {
	j,_ := json.Marshal(*flr)
	return util.Encrypt(key,j)
}

func (flr *FreshLicensResult)UnMarshal(key ,data []byte) error  {
	plainTxt, err := util.Decrypt(key,data)
	if err!=nil {
		return err
	}

	err = json.Unmarshal(plainTxt,flr)
	if err!=nil{
		return err
	}

	return nil
}



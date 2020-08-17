package miners

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/kprc/libeth/account"
	"github.com/kprc/libeth/util"
)

type GithubDownLoadPoint struct {
	Owner      string `json:"ownew"`
	Repository string `json:"repository"`
	Path       string `json:"path"`
	ReadToken  string `json:"access_token"`
}

type BootsTrapMiners struct {
	Boots             []*Miner               `json:"boots"`
	BeatlesMasterAddr account.BeatleAddress  `json:"beatles_master_addr"`
	BeatlesEthAddr    string                 `json:"beatles_eth_addr"`
	BeatlesTrxAddr    string                 `json:"beatles_trx_addr"`
	EthAccPoint       string                 `json:"eth_acc_point"`
	TrxAccPoint       string                 `json:"trx_acc_point"`
	Price             float64                `json:"price"`
	NextDownloadPoint []*GithubDownLoadPoint `json:"next_download_point"`
}

var seckey = `We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, 
				provide for the common defense, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, 
				do ordain and establish this Constitution for the United States of America.`

func (gd *GithubDownLoadPoint) String() string {
	return "owner: " + gd.Owner + "  repo: " + gd.Repository + "  path: " + gd.Path + "  token: " + gd.ReadToken
}

func (btms *BootsTrapMiners) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*btms)
	return util.Encrypt(key, j)
}

func (btms *BootsTrapMiners) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(plainTxt, btms)
	if err != nil {
		return err
	}
	return nil
}

func SecKey() []byte {

	key := sha256.Sum256([]byte(seckey))

	return key[:]
}

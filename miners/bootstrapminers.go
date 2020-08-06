package miners

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/kprc/libeth/account"
	"github.com/kprc/libeth/util"
)

type BootsTrapMiners struct {
	Boots             []*Miner              `json:"boots"`
	BeatlesMasterAddr account.BeatleAddress `json:"beatles_master_addr"`
	EthAccPoint       string                `json:"eth_acc_point"`
	TrxAccPoint       string                `json:"trx_acc_point"`
	NextDownloadPoint []string              `json:"next_download_point"`
}

var seckey = `We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, 
				provide for the common defense, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, 
				do ordain and establish this Constitution for the United States of America.`

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

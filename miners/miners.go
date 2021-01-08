package miners

import (
	"encoding/json"
	"fmt"
	"github.com/kprc/libeth/account"
	"github.com/kprc/libeth/util"
)

type Miner struct {
	Ipv4Addr string                `json:"ipv_4_addr"`
	Port     int                   `json:"port"`
	Location string                `json:"location"`
	MinerId  account.BeatleAddress `json:"miner_id"`
}

func (m *Miner)String() string  {
	msg := ""
	msg += fmt.Sprintf("%-50s", m.MinerId.String())
	msg += fmt.Sprintf("%-18s", m.Ipv4Addr)
	msg += fmt.Sprintf("%-7d", m.Port)
	msg += fmt.Sprintf("%-12s", m.Location)

	return msg
}

type BestMiners struct {
	Ms []Miner `json:"ms"`
}

func (bm *BestMiners) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*bm)
	return util.Encrypt(key, j)
}

func (bm *BestMiners) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(plainTxt, bm)
	if err != nil {
		return err
	}
	return nil
}

func (m *Miner) Marshal(key []byte) ([]byte, error) {
	j, _ := json.Marshal(*m)
	return util.Encrypt(key, j)
}

func (m *Miner) UnMarshal(key []byte, data []byte) error {
	plainTxt, err := util.Decrypt(key, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(plainTxt, m)
	if err != nil {
		return err
	}
	return nil
}

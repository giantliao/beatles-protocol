package meta

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

type Meta struct {
	Content  []byte
	ContentS string
}

func (m *Meta) Marshal(sender string, cipherTxt []byte) {
	var randbytes [128]byte

	rand.Read(randbytes[:])

	m.Content = randbytes[:]

	var lSender byte = byte(len(sender))
	m.Content = append(m.Content, lSender)
	m.Content = append(m.Content, []byte(sender)...)
	m.Content = append(m.Content, cipherTxt...)

	m.ContentS = base64.StdEncoding.EncodeToString(m.Content)

	return

}

func (m *Meta) UnMarshal() (sender string, cipherTxt []byte, err error) {

	m.Content, err = base64.StdEncoding.DecodeString(m.ContentS)
	if err != nil {
		return "", nil, err
	}

	if len(m.Content) <= 129 {
		return "", nil, errors.New("meta data error")
	}

	lSender := int(m.Content[128])

	if len(m.Content) < 129+lSender {
		return "", nil, errors.New("meta sender data error")
	}

	sender = string(m.Content[129 : 129+lSender])

	cipherTxt = m.Content[129+lSender:]

	return
}

package meta

import (
	"crypto/rand"
	"errors"
)

type Meta struct {
	Content []byte
}

func (m *Meta) Marshal(sender string, cipherTxt []byte) {
	var randbytes [128]byte
	rand.Read(randbytes[:])

	m.Content = randbytes[:]

	var lSender byte = byte(len(sender))
	m.Content = append(m.Content, lSender)
	m.Content = append(m.Content, ([]byte(sender))...)
	m.Content = append(m.Content, cipherTxt...)

	return

}

func (m *Meta) UnMarshal() (sender string, cipherTxt []byte, err error) {
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

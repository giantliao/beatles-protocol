package meta

import "github.com/giantliao/beatles-protocol/licenses"

type SessionBuild struct {
	ShadowLen  int32
	CAddress   string
	ShadownLen int32
	CLicense   licenses.License
}

type SessionFastBuild struct {
	ShadowLen int32
	SessionId [32]byte
}

package model

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
)

func newMD5Filter(w io.Writer) *md5Filter {
	return &md5Filter{
		hash: md5.New(),
		w:    w,
	}
}

type md5Filter struct {
	hash hash.Hash
	w    io.Writer
}

func (m *md5Filter) Write(p []byte) (int, error) {
	n, err := m.hash.Write(p)
	if err != nil {
		return n, err
	}
	return m.w.Write(p)
}

func (m *md5Filter) Sum() string {
	return hex.EncodeToString(m.hash.Sum(nil))
}

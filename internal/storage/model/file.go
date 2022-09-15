package model

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"io/fs"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type File interface {
	io.Seeker
	fs.File
}

type FileInfo struct {
	ID        uint64
	CreatedAt time.Time
	Filename  string
	UserID    uint64
	MD5       string

	path string
}

func (fi *FileInfo) Open() (File, error) {
	return os.Open(fi.path)
}

func tempName() string {
	buf := make([]byte, 10)
	n, err := rand.Read(buf)
	if n == 0 {
		logrus.Errorf("temp name: %s", err)
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return base32.StdEncoding.EncodeToString(buf)
}

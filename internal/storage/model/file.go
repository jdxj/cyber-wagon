package model

import (
	"crypto/rand"
	"encoding/base32"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type File struct {
	ID        uint64
	CreatedAt time.Time
	Filename  string
	UserID    uint64
	MD5       string
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

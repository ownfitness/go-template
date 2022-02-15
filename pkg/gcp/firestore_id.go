package gcp

import (
	"crypto/md5"
	"encoding/hex"
)

func encodeID(id string) string {
	hash := md5.Sum([]byte(id))
	return hex.EncodeToString(hash[:])
}

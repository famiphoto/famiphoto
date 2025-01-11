package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(val string) string {
	h := md5.Sum([]byte(val))
	return hex.EncodeToString(h[:])
}

func MD5Bytes(val []byte) string {
	h := md5.Sum(val)
	return hex.EncodeToString(h[:])
}

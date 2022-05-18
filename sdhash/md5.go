package sdhash

import (
	"bytes"
	"crypto/md5"

	"github.com/gaorx/stardust4/sdbytes"
)

func MD5(data []byte) sdbytes.Slice {
	sum := md5.Sum(data)
	return sum[:]
}

func ValidMD5(data, expected []byte) bool {
	sum := md5.Sum(data)
	return bytes.Equal(sum[:], expected)
}

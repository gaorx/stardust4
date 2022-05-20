package sdhash

import (
	"bytes"
	"crypto/md5"

	"github.com/gaorx/stardust4/sdbytes"
)

func Md5(data []byte) sdbytes.Slice {
	sum := md5.Sum(data)
	return sum[:]
}

func ValidMd5(data, expected []byte) bool {
	sum := md5.Sum(data)
	return bytes.Equal(sum[:], expected)
}

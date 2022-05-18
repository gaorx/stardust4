package sdhash

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/gaorx/stardust4/sdbytes"
)

func SHA1(data []byte) sdbytes.Slice {
	sum := sha1.Sum(data)
	return sum[:]
}

func SHA256(data []byte) sdbytes.Slice {
	sum := sha256.Sum256(data)
	return sum[:]
}

func SHA512(data []byte) sdbytes.Slice {
	sum := sha512.Sum512(data)
	return sum[:]
}

func HMACSHA1(data, key []byte) sdbytes.Slice {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func HMACSHA256(data, key []byte) sdbytes.Slice {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func HMACSHA512(data, key []byte) sdbytes.Slice {
	mac := hmac.New(sha512.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func ValidHMACSHA1(data, key, expected []byte) bool {
	actual := HMACSHA1(data, key)
	return hmac.Equal(actual, expected)
}

func ValidHMACSHA256(data, key, expected []byte) bool {
	actual := HMACSHA256(data, key)
	return hmac.Equal(actual, expected)
}

func ValidHMACSHA512(data, key, expected []byte) bool {
	actual := HMACSHA512(data, key)
	return hmac.Equal(actual, expected)
}

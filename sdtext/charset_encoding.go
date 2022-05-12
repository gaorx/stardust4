package sdtext

import (
	"github.com/gaorx/stardust4/sderr"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type CharsetEncoding interface {
	Encode(s string) ([]byte, error)
	Decode(encoded []byte) (string, error)
	MustDecode(encoded []byte) string
}

var (
	GBK     CharsetEncoding = simplifiedChineseEncoding{name: "GBK", encoding: simplifiedchinese.GBK}
	GB2312  CharsetEncoding = simplifiedChineseEncoding{name: "GB2312", encoding: simplifiedchinese.HZGB2312}
	GB18030 CharsetEncoding = simplifiedChineseEncoding{name: "GB18030", encoding: simplifiedchinese.GB18030}
)

type simplifiedChineseEncoding struct {
	name     string
	encoding encoding.Encoding
}

func (e simplifiedChineseEncoding) Encode(s string) ([]byte, error) {
	b, err := e.encoding.NewEncoder().Bytes([]byte(s))
	if err != nil {
		return nil, sderr.Wrapf(err, "string to %s error", e.name)
	}
	return b, nil
}

func (e simplifiedChineseEncoding) Decode(encoded []byte) (string, error) {
	if len(encoded) <= 0 {
		return "", nil
	}
	b, err := e.encoding.NewDecoder().Bytes(encoded)
	if err != nil {
		return "", sderr.Wrapf(err, "%s to string error", e.name)
	}
	return string(b), nil
}

func (e simplifiedChineseEncoding) MustDecode(encoded []byte) string {
	r, err := e.Decode(encoded)
	if err != nil {
		panic(err)
	}
	return r
}

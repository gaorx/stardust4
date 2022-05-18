package sdcompress

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/gaorx/stardust4/sderr"
)

type GzipLevel int

const (
	GzipNoCompression      = GzipLevel(gzip.NoCompression)
	GzipBestSpeed          = GzipLevel(gzip.BestSpeed)
	GzipBestCompression    = GzipLevel(gzip.BestCompression)
	GzipDefaultCompression = GzipLevel(gzip.DefaultCompression)
	GzipHuffmanOnly        = GzipLevel(gzip.HuffmanOnly)
)

var (
	GzipAllLevels = []GzipLevel{
		GzipNoCompression,
		GzipBestSpeed,
		GzipBestCompression,
		GzipDefaultCompression,
		GzipHuffmanOnly,
	}
)

func Gzip(data []byte, level GzipLevel) ([]byte, error) {
	if data == nil {
		return nil, sderr.New("sdcompress_gzip nil data")
	}
	buff := new(bytes.Buffer)
	w, err := gzip.NewWriterLevel(buff, int(level))
	if err != nil {
		return nil, sderr.Wrapf(err, "sdcompress gzip make writer error(%d)", level)
	}
	_, err = w.Write(data)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress gzip write error")
	}
	err = w.Close()
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress gzip close error")
	}
	return buff.Bytes(), nil
}

func Ungzip(data []byte) ([]byte, error) {
	if data == nil {
		return nil, sderr.New("sdcompress ungzip nil data")
	}
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress ungzip make reader error")
	}
	defer func() { _ = r.Close() }()

	to, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress ungzip read error")
	}
	return to, nil
}

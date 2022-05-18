package sdcompress

import (
	"bytes"
	"io/ioutil"

	"github.com/gaorx/stardust4/sderr"
	"github.com/pierrec/lz4/v4"
)

type Lz4Level lz4.CompressionLevel

const (
	Lz4Fast   = Lz4Level(lz4.Fast)
	Lz4Level1 = Lz4Level(lz4.Level1)
	Lz4Level2 = Lz4Level(lz4.Level2)
	Lz4Level3 = Lz4Level(lz4.Level3)
	Lz4Level4 = Lz4Level(lz4.Level4)
	Lz4Level5 = Lz4Level(lz4.Level5)
	Lz4Level6 = Lz4Level(lz4.Level6)
	Lz4Level7 = Lz4Level(lz4.Level7)
	Lz4Level8 = Lz4Level(lz4.Level8)
	Lz4Level9 = Lz4Level(lz4.Level9)
)

var (
	Lz4AllLevels = []Lz4Level{
		Lz4Fast,
		Lz4Level1,
		Lz4Level2,
		Lz4Level3,
		Lz4Level4,
		Lz4Level5,
		Lz4Level6,
		Lz4Level7,
		Lz4Level8,
		Lz4Level9,
	}
)

func Lz4(data []byte, level Lz4Level) ([]byte, error) {
	if data == nil {
		return nil, sderr.New("sdcompress lz4 nil data")
	}
	buff := new(bytes.Buffer)
	w := lz4.NewWriter(buff)
	_ = w.Apply(lz4.CompressionLevelOption(lz4.CompressionLevel(level)))
	_, err := w.Write(data)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress lz4 write error")
	}
	err = w.Close()
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress lz4 close error")
	}
	return buff.Bytes(), nil
}

func Unlz4(data []byte) ([]byte, error) {
	if data == nil {
		return nil, sderr.New("sdcompress unlz4 nil data")
	}
	r := lz4.NewReader(bytes.NewReader(data))
	to, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, sderr.Wrap(err, "sdcompress unlz4 read error")
	}
	return to, nil
}

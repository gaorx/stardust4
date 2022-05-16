package sdfile

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gaorx/stardust4/sderr"
)

func ReadBytes(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, sderr.Wrap(err, "sdfile read bytes error")
	}
	return data, nil
}

func ReadBytesDef(filename string, def []byte) []byte {
	data, err := ReadBytes(filename)
	if err != nil {
		return def
	}
	return data
}

func WriteBytes(filename string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(filename, data, perm)
	return sderr.Wrap(err, "sdfile write bytes error")
}

func AppendBytes(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return sderr.Wrap(err, "sdfile append bytes: open file error")
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	_ = f.Close()
	return sderr.Wrap(err, "sdfile append bytes: write error")
}

func ReadText(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", sderr.Wrap(err, "sdfile read text: read file erro")
	}
	return string(data), nil
}

func ReadTextDef(filename, def string) string {
	data, err := ReadText(filename)
	if err != nil {
		return def
	}
	return data
}

func WriteText(filename string, text string, perm os.FileMode) error {
	return WriteBytes(filename, []byte(text), perm)
}

func AppendText(filename string, text string, perm os.FileMode) error {
	return AppendBytes(filename, []byte(text), perm)
}

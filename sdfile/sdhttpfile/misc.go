package sdhttpfile

import (
	"io/ioutil"
	"net/http"

	"github.com/gaorx/stardust4/sderr"
)

func HttpReadBytes(hfs http.FileSystem, name string) ([]byte, error) {
	if hfs == nil {
		return nil, sderr.New("sdhttpfile nil hfs")
	}
	f, err := hfs.Open(name)
	if err != nil {
		return nil, sderr.Wrap(err, "sdhttpfile open error")
	}
	defer func() { _ = f.Close() }()
	r, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, sderr.Wrap(err, "sdhttpfile read error")
	}
	return r, nil
}

func HttpReadText(hfs http.FileSystem, name string) (string, error) {
	b, err := HttpReadBytes(hfs, name)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func HttpReadTextDef(hfs http.FileSystem, name, def string) string {
	s, err := HttpReadText(hfs, name)
	if err != nil {
		return def
	}
	return s
}

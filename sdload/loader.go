package sdload

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gaorx/stardust4/sderr"
)

// Loader

type Loader interface {
	LoadBytes(loc string) ([]byte, error)
}

// LoaderFunc

type LoaderFunc func(loc string) ([]byte, error)

func (f LoaderFunc) LoadBytes(loc string) ([]byte, error) {
	return f(loc)
}

// Loaders

var (
	loaders = map[string]Loader{
		"":      LoaderFunc(fileLoader),
		"file":  LoaderFunc(fileLoader),
		"http":  LoaderFunc(httpLoader),
		"https": LoaderFunc(httpLoader),
	}
)

func RegisterLoader(scheme string, loader Loader) {
	if scheme == "" {
		panic(sderr.New("sdload no scheme"))
	}
	if loader == nil {
		panic(sderr.New("sdload nil loader"))
	}
	loaders[scheme] = loader
}

// default loader

func fileLoader(loc string) ([]byte, error) {
	loc = strings.TrimPrefix(loc, "file://")
	data, err := ioutil.ReadFile(loc)
	if err != nil {
		return nil, sderr.Wrap(err, "sdload read file error")
	}
	return data, nil
}

func httpLoader(loc string) ([]byte, error) {
	resp, err := (&http.Client{Timeout: 7 * time.Second}).Get(loc)
	if err != nil {
		return nil, sderr.Wrap(err, "sdload http get error")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sderr.Newf("sdload response HTTP status error(%d, '%s')", resp.StatusCode, loc)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, sderr.Wrap(err, "sdload read http response body error")
	}
	return data, nil
}

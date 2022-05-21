package sdecho

import (
	"io/ioutil"

	"github.com/gaorx/stardust4/sderr"
)

func (c Context) RequestBodyBytes() ([]byte, error) {
	reader := c.Request().Body
	r, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, sderr.Wrap(err, "sdecho read request body error")
	}
	return r, nil
}

func (c Context) RequestBodyString() (string, error) {
	b, err := c.RequestBodyBytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

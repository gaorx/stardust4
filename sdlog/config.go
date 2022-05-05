package sdlog

import (
	"io"
	"os"
	"strings"

	"github.com/gaorx/stardust4/sderr"
	"github.com/samber/lo"
)

type Config struct {
	Tag        string   `json:"tag" toml:"tag"`
	Level      string   `json:"level" toml:"level"`
	TimeFormat string   `json:"time_format" toml:"time_format"`
	Outputs    []string `json:"outputs" toml:"outputs"`
	Others     []Config `json:"others" toml:"others"`
}

func (c Config) flatten() []Config {
	return nil
}

func (c Config) parseOutputs() (io.Writer, error) {
	openOne := func(output string) (io.Writer, error) {
		lout := strings.ToLower(output)
		switch lout {
		case "", "stdout":
			return os.Stdout, nil
		case "stderr":
			return os.Stderr, nil
		case "null", "discard":
			return io.Discard, nil
		default:
			f, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				return nil, sderr.Wrap(err, "create log file error")
			}
			return f, nil
		}
	}
	outputs := lo.Filter(c.Outputs, func(v string, _ int) bool {
		return v != ""
	})
	if len(outputs) <= 0 {
		return openOne("STDOUT")
	}
	var writers []io.Writer
	for _, output := range outputs {
		w, err := openOne(output)
		if err != nil {
			return nil, err
		}
		writers = append(writers, w)
	}
	if len(writers) == 1 {
		return writers[0], nil
	} else {
		return io.MultiWriter(writers...), nil
	}
}

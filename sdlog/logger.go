package sdlog

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gaorx/stardust4/sderr"
	"github.com/gaorx/stardust4/sdstr"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Level   string   `json:"level" toml:"level"`
	Format  string   `json:"format" toml:"format"`
	File    string   `json:"file" toml:"file"`
	Outputs []Output `json:"outputs" toml:"outputs"`
}

type Output struct {
	Level string `json:"level" toml:"level"`
	Tag   string `json:"tag" toml:"tag"`
	File  string `json:"file" toml:"file"`
}

func New(config Config) (*Logger, error) {
	l := logrus.New()
	if err := setup(l, &config); err != nil {
		return nil, sderr.Wrap(err, "setup log error")
	}
	return l, nil
}

func setup(l *Logger, config *Config) error {
	level, err := ParseLevel(ifEmptyAs(config.Level, "debug"))
	if err != nil {
		return sderr.Wrap(err, "parse log level error")
	}
	output, err := parseFile(config.File)
	if err != nil {
		return err
	}
	l.SetLevel(level)
	switch strings.ToLower(config.Format) {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	case "pretty":
		l.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	default:
		l.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}
	l.SetOutput(output)
	return nil
}

func parseFile(outFn string) (io.Writer, error) {
	parseOptions := func(s string) map[string]string {
		if s == "" {
			return map[string]string{}
		}
		segs := strings.Split(s, ",")
		m := map[string]string{}
		for _, seg := range segs {
			k, v := sdstr.Split2s(seg, "=")
			k, v = strings.TrimSpace(k), strings.TrimSpace(v)
			if k != "" {
				m[k] = ifEmptyAs(v, "true")
			}
		}
		return m
	}

	getInt := func(options map[string]string, k string, def int) int {
		v, ok := options[k]
		if !ok || v == "" {
			return def
		}
		iv, err := strconv.Atoi(v)
		if err != nil {
			return def
		}
		return iv
	}
	lout := strings.ToLower(strings.TrimSpace(outFn))
	switch lout {
	case "", "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "null":
		return io.Discard, nil
	default:
		fn, options := sdstr.Split2s(outFn, "|")
		fn, options = strings.TrimSpace(fn), strings.TrimSpace(options)
		absFn, err := absPath(fn)
		if err != nil {
			return nil, err
		}
		options1 := parseOptions(options)
		return &lumberjack.Logger{
			Filename:   absFn,
			MaxSize:    getInt(options1, "block", 200),
			MaxBackups: getInt(options1, "backups", 100),
			MaxAge:     getInt(options1, "days", 30),
			Compress:   false,
		}, nil
	}
}

type outputHook struct {
	logger *Logger
}

func newOutputHook(output Output) (*outputHook, error) {
	// TODO
	return nil, nil
}

func (h *outputHook) Levels() []Level {
	return AllLevels
}

func (h *outputHook) Fire(entry *Entry) error {
	return nil
}

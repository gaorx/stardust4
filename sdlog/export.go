package sdlog

import (
	"github.com/sirupsen/logrus"
)

const (
	TAG               = "tag"
	DefaultTimeLayout = "2006-01-02 15:04:05.000"
	PanicLevel        = logrus.PanicLevel
	FatalLevel        = logrus.FatalLevel
	ErrorLevel        = logrus.ErrorLevel
	WarnLevel         = logrus.WarnLevel
	InfoLevel         = logrus.InfoLevel
	DebugLevel        = logrus.DebugLevel
	TraceLevel        = logrus.TraceLevel
)

type (
	Logger = logrus.Logger
	Level  = logrus.Level
	Fields = logrus.Fields
	Entry  = logrus.Entry
)

var (
	// fields
	WithFields = logrus.WithFields

	// Level
	ParseLevel = logrus.ParseLevel
	SetLevel   = logrus.SetLevel
	GetLevel   = logrus.GetLevel
	AllLevels  = logrus.AllLevels

	// With info
	WithError = logrus.WithError
	WithField = logrus.WithField

	// Log
	Debug   = logrus.Debug
	Print   = logrus.Print
	Info    = logrus.Info
	Warn    = logrus.Warn
	Warning = logrus.Warning
	Error   = logrus.Error
	Panic   = logrus.Panic
	Fatal   = logrus.Fatal

	// Logf
	Debugf   = logrus.Debugf
	Printf   = logrus.Printf
	Infof    = logrus.Infof
	Warnf    = logrus.Warnf
	Warningf = logrus.Warningf
	Errorf   = logrus.Errorf
	Panicf   = logrus.Panicf
	Fatalf   = logrus.Fatalf

	// LogFn
	DebugFn   = logrus.DebugFn
	PrintFn   = logrus.PrintFn
	InfoFn    = logrus.InfoFn
	WarnFn    = logrus.WarnFn
	WarningFn = logrus.WarningFn
	ErrorFn   = logrus.ErrorFn
	PanicFn   = logrus.PanicFn
	FatalFn   = logrus.FatalFn
)

func init() {
	Init(Config{
		Format: "pretty",
		Level:  "trace",
		File:   "stdout",
	})
}

func Init(config Config) error {
	return setup(logrus.StandardLogger(), &config)
}

func MustInit(config Config) {
	err := Init(config)
	if err != nil {
		panic(err)
	}
}

func WithTag(tag string) *logrus.Entry {
	return logrus.WithField(TAG, tag)
}

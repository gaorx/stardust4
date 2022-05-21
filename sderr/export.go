package sderr

import (
	"runtime"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
)

// export types
type (
	StackFrame    = eris.StackFrame
	MultipleError = multierror.Error
)

// export errors

var (
	New   = eris.New
	Newf  = eris.Errorf
	Wrap  = eris.Wrap
	Wrapf = eris.Wrapf
)

// errors

func WithStack(err error) error {
	return Wrap(err, "")
}

func Cause(err error) error {
	return eris.Cause(err)
}

func Unwrap(err error) error {
	return eris.Unwrap(err)
}

func Is(err, target error) bool {
	return eris.Is(err, target)
}

func As[E error](err error) (E, bool) {
	var e E
	if eris.As(err, &e) {
		return e, true
	} else {
		return e, false
	}
}

func Append(err error, errs ...error) error {
	return multierror.Append(err, errs...)
}

// Utils

func StackOf(err error) []StackFrame {
	if err == nil {
		return []StackFrame{}
	}
	rawFrames := eris.StackFrames(err)
	if len(rawFrames) <= 0 {
		return []StackFrame{}
	}
	var frames []StackFrame
	callersFrames := runtime.CallersFrames(rawFrames)
	for {
		callerFrames, more := callersFrames.Next()
		i := strings.LastIndex(callerFrames.Function, "/")
		name := callerFrames.Function[i+1:]
		frames = append(frames, StackFrame{
			Name: name,
			File: callerFrames.File,
			Line: callerFrames.Line,
		})
		if !more {
			break
		}
	}
	return frames
}

func ToErr(v any) error {
	switch err := v.(type) {
	case nil:
		return nil
	case error:
		return err
	case string:
		return New(err)
	default:
		return Newf("%v", err)
	}
}

func Multiple(errs []error) error {
	var merged error = nil
	for _, err := range errs {
		if err != nil {
			merged = Append(merged, err)
		}
	}
	return merged
}

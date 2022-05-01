package sderr

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// export types
type (
	Frame         = errors.Frame
	StackTrace    = errors.StackTrace
	MultipleError = multierror.Error
)

// export functions
var (
	New          = errors.New
	Newf         = errors.Errorf
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	WithStack    = errors.WithStack
	Cause        = errors.Cause
	Unwrap       = errors.Unwrap
	As           = errors.As
	Is           = errors.Is
	Append       = multierror.Append
	Flatten      = multierror.Flatten
)

// Utils

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
	var merr error = nil
	for _, err := range errs {
		if err != nil {
			merr = Append(merr, err)
		}
	}
	return merr
}

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

// create errors
var (
	New          = errors.New
	Newf         = errors.Errorf
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	WithStack    = errors.WithStack
)

// errors
func Cause(err error) error {
	return errors.Cause(err)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As[E error](err error) (E, bool) {
	var e E
	if errors.As(err, &e) {
		return e, true
	} else {
		return e, false
	}
}

func Append(err error, errs ...error) error {
	return multierror.Append(err, errs...)
}

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

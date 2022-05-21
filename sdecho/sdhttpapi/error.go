package sdhttpapi

import (
	"fmt"

	"github.com/gaorx/stardust4/sderr"
	"github.com/samber/lo"
)

type Error struct {
	Code  any
	Msg   string
	Stack []string
}

type ErrorExtractor func(err, codeOk, codeUnknownError any) Error

func defaultErrorExtractor(err, codeOk, codeUnknownError any) Error {
	if err == nil {
		panic("sdhttpapi err is nil")
	}
	if msg, ok := err.(string); ok {
		return Error{
			Code:  codeUnknownError,
			Msg:   msg,
			Stack: nil,
		}
	} else if err1, ok := err.(error); ok {
		var code any
		err1Root := sderr.Cause(err1)
		if err1RootE, ok := err1Root.(interface{ ErrorCode() int }); ok {
			code = err1RootE.ErrorCode()
		} else if err1RootE, ok := err1Root.(interface{ ErrorCode() string }); ok {
			code = err1RootE.ErrorCode()
		} else if err1RootE, ok := err1Root.(interface{ ErrorCode() any }); ok {
			code = err1RootE.ErrorCode()
		} else {
			code = codeUnknownError
		}
		frames := lo.Map(sderr.StackOf(err1), func(frame sderr.StackFrame, _ int) string {
			return fmt.Sprintf("%s (%s:%d)", frame.Name, frame.File, frame.Line)
		})
		return Error{
			Code:  code,
			Msg:   err1.Error(),
			Stack: frames,
		}
	} else {
		return Error{
			Code:  codeUnknownError,
			Msg:   fmt.Sprintf("%v", err),
			Stack: nil,
		}
	}
}

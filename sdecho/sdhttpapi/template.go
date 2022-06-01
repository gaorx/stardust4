package sdhttpapi

type ResultTemplate struct {
	CodeOk           any
	CodeUnknownError any
	ErrorExtractor   ErrorExtractor
	Renderer         Renderer
}

func (rt *ResultTemplate) Ok(data any) *Result {
	return &Result{
		// template
		template: rt,

		// data & error
		Code:         rt.CodeOk,
		Data:         data,
		ErrorMessage: "",
		ErrorStack:   nil,
	}
}

func (rt *ResultTemplate) Err(err any) *Result {
	if err == nil {
		return Ok(nil)
	}

	extractor := rt.ErrorExtractor
	if extractor == nil {
		extractor = defaultErrorExtractor
	}
	info := extractor(err, rt.CodeOk, rt.CodeUnknownError)
	return &Result{
		// template
		template: rt,

		// data & error
		Code:         info.Code,
		Data:         nil,
		ErrorMessage: info.Msg,
		ErrorStack:   info.Stack,
	}
}

func (rt *ResultTemplate) Of(data any, err any) *Result {
	if err != nil {
		return rt.Err(err)
	} else {
		return rt.Ok(data)
	}
}

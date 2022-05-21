package sdhttpapi

var resultTemplate = &ResultTemplate{
	CodeOk:           0,
	CodeUnknownError: -1,
	ErrorExtractor:   nil,
	Renderer:         nil,
}

func SetCodeOk(code any) {
	resultTemplate.CodeOk = code
}

func SetCodeUnknownError(code any) {
	resultTemplate.CodeUnknownError = code
}

func SetErrorExtractor(extractor ErrorExtractor) {
	resultTemplate.ErrorExtractor = extractor
}

func SetRenderer(renderer Renderer) {
	resultTemplate.Renderer = renderer
}

func Ok(data any) *Result {
	return resultTemplate.Ok(data)
}

func Err(err any) *Result {
	return resultTemplate.Err(err)
}

func Of(data, err any) *Result {
	return resultTemplate.Of(data, err)
}

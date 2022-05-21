package sdhttpapi

import (
	"github.com/gaorx/stardust4/sderr"
	"github.com/labstack/echo/v4"
)

type Result struct {
	template       *ResultTemplate
	HttpStatusCode int
	Code           any
	Data           any
	ErrorMessage   string
	ErrorStack     []string
	Facade         any
	Fields         map[string]any
	Renderer       Renderer
}

func (r *Result) IsOk() bool {
	return r.Code == r.template.CodeOk
}

func (r *Result) Template() *ResultTemplate {
	return r.template
}

func (r *Result) WithHttpStatusCode(statusCode int) *Result {
	r.HttpStatusCode = statusCode
	return r
}

func (r *Result) WithFacade(facade any) *Result {
	r.Facade = facade
	return r
}

func (r *Result) WithField(k string, v any) *Result {
	r.ensureFields()
	r.Fields[k] = v
	return r
}

func (r *Result) WithFields(fields map[string]any) *Result {
	r.ensureFields()
	for k, v := range fields {
		r.Fields[k] = v
	}
	return r
}

func (r *Result) WithRenderer(renderer Renderer) *Result {
	r.Renderer = renderer
	return r
}

func (r *Result) Render(c echo.Context) error {
	renderer := r.Renderer
	if renderer == nil {
		renderer = r.template.Renderer
	}
	if renderer == nil {
		renderer = RenderJSON
	}
	err := renderer(c, r)
	return sderr.WithStack(err)
}

func (r *Result) ensureFields() {
	if r.Fields == nil {
		r.Fields = make(map[string]any)
	}
}

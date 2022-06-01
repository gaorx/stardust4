package sdhttpapi

import (
	"github.com/labstack/echo/v4"
)

type HandlerOptions struct {
	ResultTemplate *ResultTemplate
	Facade         any
	Fields         map[string]any
}

func WithResultTemplateOption(t *ResultTemplate) func(*HandlerOptions) {
	return func(opts *HandlerOptions) {
		opts.ResultTemplate = t
	}
}

func WithFacadeOption(facade any) func(option *HandlerOptions) {
	return func(opts *HandlerOptions) {
		opts.Facade = facade
	}
}

func WithFieldsOption(fields map[string]any) func(option *HandlerOptions) {
	return func(opts *HandlerOptions) {
		if opts.Fields == nil {
			opts.Fields = map[string]any{}
		}
		for k, v := range fields {
			opts.Fields[k] = v
		}
	}
}

func WithFieldOption(k string, v any) func(option *HandlerOptions) {
	return func(opts *HandlerOptions) {
		if opts.Fields == nil {
			opts.Fields = map[string]any{}
		}
		opts.Fields[k] = v
	}
}

func H[S any, I any, O any](state S, h func(S, I) (O, error), opts ...func(*HandlerOptions)) echo.HandlerFunc {
	var opts1 HandlerOptions
	for _, opt := range opts {
		if opt != nil {
			opt(&opts1)
		}
	}
	if opts1.ResultTemplate == nil {
		opts1.ResultTemplate = DefaultResultTemplate()
	}
	rt := opts1.ResultTemplate

	return func(c echo.Context) error {
		var in I
		err := c.Bind(&in)
		if err != nil {
			return rt.Err(err).Render(c)
		}
		out, err := h(state, in)
		if err != nil {
			return rt.Err(err).Render(c)
		}
		return rt.Ok(out).WithFacade(opts1.Facade).WithFields(opts1.Fields).Render(c)
	}
}

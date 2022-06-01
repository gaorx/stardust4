package sdmansys

import "github.com/gaorx/stardust4/sdecho/sdhttpapi"

type System struct {
	ResultTemplate *sdhttpapi.ResultTemplate
	PathPrefix     string
	Endpoints      []Endpoint
	Methods        []Method
	Services       []Service
}

func New() *System {
	return &System{}
}

func (sys *System) SetPathPrefix(pathPrefix string) {
	sys.PathPrefix = pathPrefix
}

func (sys *System) SetResultTemplate(rt *sdhttpapi.ResultTemplate) {
	sys.ResultTemplate = rt
}

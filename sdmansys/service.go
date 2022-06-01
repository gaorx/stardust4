package sdmansys

type Service struct {
	PathPrefix string
	Facade     any
	Service    any
}

func (sys *System) AddService(category string, facade any, service any) {
	sys.Services = append(sys.Services, Service{
		PathPrefix: sys.PathPrefix,
		Facade:     facade,
		Service:    service,
	})
}

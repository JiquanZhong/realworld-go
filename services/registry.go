package services

type Registry struct {
	User UserService
}

var registry *Registry

func initServices() {
	registry = &Registry{
		User: NewUserService(),
	}
}

func Services() *Registry {
	if registry == nil {
		initServices()
	}
	return registry
}

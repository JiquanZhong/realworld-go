package services

type Registry struct {
	User   UserService
	Mcp    McpService
	McpTag McpTagService
}

var registry *Registry

func initServices() {
	registry = &Registry{
		User:   NewUserService(),
		Mcp:    NewMcpService(),
		McpTag: NewMcpTagService(),
	}
}

func Services() *Registry {
	if registry == nil {
		initServices()
	}
	return registry
}

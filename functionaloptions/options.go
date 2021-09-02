package functionaloptions

type ServerOptions struct {
	host string
	port int
}

type ServerOption interface {
	apply(*ServerOptions)
}

type funcServerOption struct {
	f func(*ServerOptions)
}

func (f funcServerOption) apply(opt *ServerOptions) {
	f.f(opt)
}

func newFuncServerOption(f func(*ServerOptions)) *funcServerOption {
	return &funcServerOption{
		f: f,
	}
}

func WithHost(host string) ServerOption {
	return newFuncServerOption(func(so *ServerOptions) {
		so.host = host
	})
}

func WithPort(port int) ServerOption {
	return newFuncServerOption(func(so *ServerOptions) {
		so.port = port
	})
}

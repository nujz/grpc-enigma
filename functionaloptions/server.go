package functionaloptions

type Server struct {
	options *ServerOptions
}

func (s Server) Serve() {
}

func NewServer(opt ...ServerOption) *Server {
	opts := &ServerOptions{
		host: "127.0.0.1", // default
		port: 8080,        // default
	}
	for _, o := range opt {
		o.apply(opts)
	}
	return &Server{
		options: opts,
	}
}

package floz

type Scope struct {
	server *Server
	prefix string
}

func (s *Scope) Scope(prefix string) *Scope {
	return &Scope{
		prefix: s.prefix + prefix,
		server: s.server,
	}
}

func (s *Scope) addRoute(method, path string, handler ReqHandler) *Scope {
	p := s.prefix + path
	s.server.addRoute(method, p, handler)
	return s
}

func (s *Scope) Get(path string, handler ReqHandler) *Scope {
	return s.addRoute("GET", path, handler)
}

func (s *Scope) Post(path string, handler ReqHandler) *Scope {
	return s.addRoute("POST", path, handler)
}

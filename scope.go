package floz

type Scope struct {
	server *Server
	prefix string
	mw     *MiddleWare
}

func (s *Scope) Scope(prefix string, middlewares ...ReqHandler) *Scope {
	if prefix == "/" {
		return s
	}
	sp := &Scope{
		prefix: s.prefix + prefix,
		server: s.server,
		mw:     NewMW(middlewares...),
	}
	s.server.scopes = append(s.server.scopes, sp)
	return sp
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

func (s *Scope) Wrap(middlewares ...ReqHandler) *Scope {
	s.mw.addMW(middlewares...)
	return s
}

package floz

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

type Server struct {
	router map[string]*trie
}

func (s *Server) Handle(ctx *fasthttp.RequestCtx) {
	node, _ := s.getRoute(string(ctx.Method()), string(ctx.Path()))
	if node != nil {
		node.handler(ctx)
	} else {
		fmt.Fprintf(ctx, "404 NOT FOUND:%p", ctx.URI())
	}
}

func newServer() *Server {
	return &Server{
		router: map[string]*trie{},
	}
}

func (s *Server) addRoute(method, path string, handler ReqHandler) *Server {
	parts := parsePath(path)
	if _, ok := s.router[method]; !ok {
		s.router[method] = newTrie()
	}
	s.router[method].insert(path, parts, handler)
	return s
}

func (s *Server) getRoute(method, path string) (*node, map[string]string) {
	parts := parsePath(path)
	paras := make(map[string]string)
	router, ok := s.router[method]
	if !ok {
		return nil, nil
	}

	res := router.search(parts)
	if res != nil {
		p := parsePath(res.path)
		for idx, part := range p {
			if part[0] == ':' {
				paras[part[1:]] = parts[idx]
			}
			if part[0] == '*' && len(part) > 1 {
				paras[part[1:]] = strings.Join(parts[idx:], "/")
				break
			}
		}
		return res, paras
	}
	return nil, nil
}

func (s *Server) Get(path string, handler ReqHandler) *Server {
	return s.addRoute("GET", path, handler)
}

func (s *Server) Post(path string, handler ReqHandler) *Server {
	return s.addRoute("POST", path, handler)
}

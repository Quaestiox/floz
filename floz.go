package floz

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

type ReqHandler func(handler *Ctx)

type Floz struct {
	scope  *Scope
	server *Server
}

func New() *Floz {
	return &Floz{
		server: newServer(),
	}
}

func (floz *Floz) Server() *Server {
	return floz.server
}

func (floz *Floz) Handle(c *fasthttp.RequestCtx) {
	ctx := NewCtx(c)
	floz.server.handle(ctx)
}

func (floz *Floz) Run(port string) {
	err := fasthttp.ListenAndServe(port, floz.Handle)
	if err != nil {
		fmt.Println("[error]")
		os.Exit(1)
	}
}

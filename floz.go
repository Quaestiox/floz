package floz

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

type ReqHandler func(handler *Ctx)

type Floz struct {
	server *Server
	config *Config
}

func New() *Floz {
	return &Floz{
		server: newServer(),
		config: NewConfig(),
	}
}

func (floz *Floz) Server() *Server {
	return floz.server
}

func (floz *Floz) Handle(c *fasthttp.RequestCtx) {
	ctx := NewCtx(floz, c)
	floz.server.handle(ctx)
}

func (floz *Floz) Run(port string) {
	err := fasthttp.ListenAndServe(port, floz.Handle)
	if err != nil {
		fmt.Println("[error]")
		os.Exit(1)
	}
}

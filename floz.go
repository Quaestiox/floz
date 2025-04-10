package floz

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

type ReqHandler func(handler *fasthttp.RequestCtx)

type Floz struct {
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

func (floz *Floz) Run(port string) {
	err := fasthttp.ListenAndServe(port, floz.Server().Handle)
	if err != nil {
		fmt.Println("[error]")
		os.Exit(1)
	}
}

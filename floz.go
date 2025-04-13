package floz

import (
	"fmt"
	"github.com/Quaestiox/floz/middleware/rescue"
	"github.com/valyala/fasthttp"
	"os"
)

type Floz struct {
	server *Server
	config *Config
	mw     *MiddleWare
}

func New(middleware ...*MiddleWare) *Floz {
	var mw *MiddleWare
	if len(middleware) == 0 {
		mw = NewMW()
	} else {
		mw = middleware[0]
	}
	return &Floz{
		server: newServer(),
		config: NewConfig(),
		mw:     mw,
	}
}

func Default() *Floz {
	mw := NewMW(rescue.New())
	return &Floz{
		server: newServer(),
		config: NewConfig(),
		mw:     mw,
	}
}

func (floz *Floz) Server() *Server {
	return floz.server
}

func (floz *Floz) Handle(c *fasthttp.RequestCtx) {
	mws := make([]ReqHandler, 0)

	mws = append(mws, floz.mw.list...)
	if floz.server != nil {
		for _, scope := range floz.server.scopes {
			if string(c.URI().Path()) == scope.prefix {
				mws = append(mws, scope.mw.list...)
			}
		}
	}

	ctx := NewCtx(floz, c)
	ctx.handler = mws
	floz.server.handle(ctx)
}

func (floz *Floz) Run(port string) {
	err := fasthttp.ListenAndServe(port, floz.Handle)
	if err != nil {
		fmt.Println("[error]")
		os.Exit(1)
	}
}

func (floz *Floz) Wrap(middlewares ...ReqHandler) *Floz {
	floz.mw.addMW(middlewares...)
	return floz
}

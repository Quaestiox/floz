package floz

import "github.com/valyala/fasthttp"

type Ctx struct {
	fasthttp *fasthttp.RequestCtx
	paras    map[string]string
}

func NewCtx(ctx *fasthttp.RequestCtx) *Ctx {
	return &Ctx{ctx, make(map[string]string)}
}

func (c *Ctx) Status() int {
	return c.fasthttp.Response.StatusCode()
}

func (c *Ctx) Para(key string) string {
	return c.paras[key]
}

func (c *Ctx) Paras() map[string]string {
	return c.paras
}

func (c *Ctx) Method() string {
	return string(c.fasthttp.Method())
}

func (c *Ctx) Path() string {
	return string(c.fasthttp.Path())
}

func (c *Ctx) URI() string {
	return c.fasthttp.URI().String()
}

func (c *Ctx) Write(p []byte) (n int, err error) {
	return c.fasthttp.Write(p)
}

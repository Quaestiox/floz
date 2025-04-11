package floz

import "github.com/valyala/fasthttp"

type Ctx struct {
	floz       *Floz
	fasthttp   *fasthttp.RequestCtx
	paras      map[string]string
	handler    []ReqHandler
	handlerIdx int
}

func NewCtx(floz *Floz, ctx *fasthttp.RequestCtx) *Ctx {
	return &Ctx{floz, ctx, make(map[string]string), make([]ReqHandler, 0), -1}
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

func (c *Ctx) JSON(data any, code ...int) error {
	raw, err := c.floz.config.JSONEncoder(data)
	if err != nil {
		return err
	}
	c.fasthttp.Response.SetBodyRaw(raw)
	c.fasthttp.Response.Header.SetContentType(MIMEApplicationJSON)
	if len(code) == 0 {
		c.Status(fasthttp.StatusOK)
	} else {
		c.Status(code[0])
	}
	return nil
}

func (c *Ctx) String(str string, code ...int) error {
	c.fasthttp.Response.SetBodyString(str)
	if len(code) == 0 {
		c.Status(fasthttp.StatusOK)
	} else {
		c.Status(code[0])
	}
	return nil
}

func (c *Ctx) Status(code int) *Ctx {
	c.fasthttp.Response.SetStatusCode(code)
	return c
}

func (c *Ctx) Set(key, v string) {
	c.fasthttp.Response.Header.Set(key, v)
}

func (c *Ctx) Next() {
	c.handlerIdx++
	l := len(c.handler)
	for ; c.handlerIdx < l; c.handlerIdx++ {
		c.handler[c.handlerIdx](c)
	}
}

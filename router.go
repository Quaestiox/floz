package floz

import "github.com/valyala/fasthttp"

type Router interface {
	func(ctx *fasthttp.RequestCtx)
}

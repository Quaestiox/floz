# Floz

An efficient and easy-to-use web framework based on [fasthttp](https://github.com/valyala/fasthttp) for Go.

### QuickStart

```go
func main() {
    app := floz.New()
    app.Server().Get("/", handleRoot).
        Scope("/v1").Get("/hello", handle1).
        Scope("/v2").Get("/world", handle2)

    app.Run(":8788")
}

func handleRoot(ctx *floz.Ctx) {
    ctx.String("here is the root")
}

func handle1(ctx *floz.Ctx) {
    ctx.JSON(floz.H{
        "username": "x",
        "password": "123456",
    })
}

func handle2(ctx *floz.Ctx) {
    ctx.String("hello world!")
}
```
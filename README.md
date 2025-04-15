# Floz

An efficient and easy-to-use web framework based on [fasthttp](https://github.com/valyala/fasthttp) for Go.

### Install
```shell
go get github/Quaestiox@v0.1.0
```

### QuickStart

```go
func main() {
    app := floz.Default()
    app.Server().Get("/", handleRoot).
        Scope("/v1").Get("/hello", handle1)

    app.Run(":8788")
}

func handleRoot(ctx *floz.Ctx) {
    ctx.JSON(floz.H{
        "username": "x",
        "password": "123456",
    })
}

func handle1(ctx *floz.Ctx) {
    ctx.String("hello world!")
}
```

### Middleware

You can use Wrap() to add middleware for Floz or Scopes.<br>
Here are two ways to add middleware for Floz application.

```go
app := floz.New().Wrap(Log)
```
```go
mw  := floz.NewMW(Log)
app := floz.New(mw)
```

### Data

The Data is shared by all routes and scopes in the application.

example:
```go
type State struct {
    name  string
    count int
}

func main() {
    app := floz.Default().Data(State{name: "try", count: 0})
    app.Server().Get("/data", handleData)
    app.Run(":8788")
}

func handleData(ctx *floz.Ctx) {
    data := ctx.Data().(State)
    ctx.JSON(floz.H{
        "appName": data.name,
        "count":   data.count,
    })
}
```


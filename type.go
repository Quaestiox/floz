package floz

type ReqHandler func(handler *Ctx)

type JSONMarshal func(v any) ([]byte, error)
type H map[string]any

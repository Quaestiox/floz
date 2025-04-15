package floz

type ReqHandler func(*Ctx)

type JSONMarshal func(v any) ([]byte, error)
type H map[string]any

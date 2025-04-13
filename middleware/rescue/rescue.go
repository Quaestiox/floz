package rescue

import (
	"github.com/Quaestiox/floz"
)

func New() floz.ReqHandler {
	return func(c *floz.Ctx) {
		defer func() error {
			if err := recover(); err != nil {
				return err.(error)
			}
			return nil
		}()
	}
}

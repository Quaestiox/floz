package floz

import (
	"fmt"
)

func MWRecover() ReqHandler {
	return func(c *Ctx) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("err:", r.(error))
				return
			}
			return
		}()
		c.Next()

	}
}

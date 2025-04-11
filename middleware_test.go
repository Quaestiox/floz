package floz

import (
	"fmt"
	"testing"
)

func TestNewMW(t *testing.T) {
	mw := NewMW(test1, test2)
	ctx := &Ctx{}
	mw.list[0](ctx)
	mw.list[1](ctx)
}

func test1(c *Ctx) {
	fmt.Println("test1")
}

func test2(c *Ctx) {
	fmt.Println("test2")
}

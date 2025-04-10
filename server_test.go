package floz

import (
	"testing"
)

func TestGet(t *testing.T) {
	server := newServer()
	server.Get("/", nil).
		Get("/hello", nil).
		Get("/world/:name", nil).
		Get("/hello/*file", nil)

}

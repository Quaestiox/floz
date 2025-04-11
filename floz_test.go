package floz

import "testing"

func TestNew(t *testing.T) {
	app := New()
	if app == nil {
		t.Fatal("wrong")
	}
}

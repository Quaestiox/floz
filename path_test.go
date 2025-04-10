package floz

import (
	"reflect"
	"testing"
)

func TestParsePath(t *testing.T) {

	ok := reflect.DeepEqual(parsePath("/hello/world"), []string{"hello", "world"}) &&
		reflect.DeepEqual(parsePath("/"), []string{}) &&
		reflect.DeepEqual(parsePath("/hello/:name"), []string{"hello", ":name"}) &&
		reflect.DeepEqual(parsePath("/hello/*file"), []string{"hello", "*file"}) &&
		reflect.DeepEqual(parsePath("/hello/*file/path"), []string{"hello", "*file"})
	if !ok {
		t.Fatal("[Failed]: TestParsePath")
	}
}

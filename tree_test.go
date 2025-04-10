package floz

import "testing"

func TestMatch(t *testing.T) {
	ch1, ch2 := &node{part: "p1"}, &node{part: "p2"}
	node := &node{part: "a", children: []*node{ch1, ch2}}

	res := node.match("p1")
	if res != ch1 {
		t.Fatal()
	}
}

func TestMatchAll(t *testing.T) {
	ch1, ch2, ch3 := &node{part: "p1"}, &node{part: "p2"}, &node{part: "p2"}
	node := &node{part: "a", children: []*node{ch1, ch2, ch3}}

	res := node.matchAll("p2")
	if res[0] != ch2 && res[1] != ch3 {
		t.Fatal()
	}
}

func TestTree(t *testing.T) {
	tree := newTrie()
	tree.insert("/a/bb", []string{"a", "bb"}, nil)
	node := tree.search([]string{"a", "bb"})
	if node == nil {
		t.Fatal()
	}

	tree.insert("/a/:name/b", []string{"a", ":name", "b"}, nil)
	node2 := tree.search([]string{"a", "jack", "b"})
	if node2 == nil {
		t.Fatal()
	}

	tree.insert("/a/*file/p", []string{"a", "*file", "p"}, nil)
	node3 := tree.search([]string{"a", "ooo", "p"})
	if node3 == nil {
		t.Fatal()
	}
}

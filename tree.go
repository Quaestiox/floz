package floz

import "strings"

type node struct {
	path     string
	part     string
	children []*node
	handler  ReqHandler
	wildcard bool
}

func (n *node) match(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.wildcard {
			return child
		}
	}
	return nil
}

func (n *node) matchAll(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.wildcard {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchAll(part)

	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{&node{}}
}

func (t *trie) insert(path string, parts []string, handler ReqHandler) {
	current := t.root
	for height := 0; height < len(parts); height++ {
		part := parts[height]
		child := current.match(part)
		if child == nil {
			isWild := len(part) > 0 && (part[0] == ':' || part[0] == '*')
			child = &node{part: part, wildcard: isWild}
			current.children = append(current.children, child)
		}
		current = child
	}
	current.path = path
	current.handler = handler
}

func (t *trie) search(parts []string) *node {
	return t.root.search(parts, 0)
}

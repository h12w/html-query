// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package query

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	. "h12.me/html-query/expr"
)

// Node represents a HTML node.
// Wrap html.Node so that chainable interface is possible
// Use pointer of it because we want to test with nil.
type Node struct {
	n html.Node
}

func NewNode(n *html.Node) *Node {
	if n == nil {
		return nil
	}
	return &Node{*n}
}

func Parse(r io.Reader) (*Node, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return NewNode(n), nil
}

func (n *Node) InternalNode() *html.Node {
	return &n.n
}

func (n *Node) Attr(key string, pat ...string) *string {
	if n == nil {
		return nil
	}
	return GetAttrSubmatch(n.InternalNode(), key, GetPat(pat))
}

/*
func (n *Node) AttrSubmatch(key, pat string) *string {
	if n == nil {
		return nil
	}
	return GetAttrSubmatch(n.InternalNode(), key, pat)
}
*/

func (n *Node) getAttr(key string, pat ...string) *string {
	return n.Attr(key, pat...)
}

func (n *Node) Text(pat ...string) *string {
	if n == nil {
		return nil
	}
	return GetSubmatch(GetText(&n.n), GetPat(pat))
}

func (n *Node) AllText(pat ...string) *string {
	ss := []string{}
	for _, n := range n.Descendants(TextNode).All() {
		if text := n.Text(pat...); text != nil && *text != "" {
			ss = append(ss, *text)
		}
	}
	s := strings.Join(ss, " ")
	if s != "" {
		return &s
	}
	return nil
}

func (n *Node) PlainText() string {
	if n == nil {
		return ""
	}
	var buf bytes.Buffer
	for _, s := range n.Descendants(TextNode).Strings(GetText) {
		buf.WriteString(s)
	}
	return buf.String()
}

func (n *Node) Render() *string {
	if n == nil {
		return nil
	}
	var b bytes.Buffer
	err := html.Render(&b, &n.n)
	if err != nil {
		return nil
	}
	s := b.String()
	return &s
}

func (n *Node) RenderTagOnly() *string {
	if n == nil {
		return nil
	}
	var b bytes.Buffer

	if n.n.Type == html.ElementNode {
		err := renderOpeningTag(&b, &n.n)
		if err != nil {
			return nil
		}
	} else {
		err := renderSimpleNode(&b, &n.n)
		if err != nil {
			return nil
		}
	}
	s := b.String()
	return &s

}

func (n *Node) RenderChildren() *string {
	if n == nil {
		return nil
	}
	var b bytes.Buffer
	node := FirstChild(&n.n)
	for node != nil {
		err := html.Render(&b, node)
		if err != nil {
			return nil
		}
		node = node.NextSibling
	}
	s := b.String()
	return &s
}

type TextNodeNode struct {
	Node
	rx *regexp.Regexp
}

func NewTextNodeNode(n *Node, rx *regexp.Regexp) *TextNodeNode {
	if n == nil {
		return nil
	}
	return &TextNodeNode{*n, rx}
}

func (n *TextNodeNode) Submatch() *string {
	val := n.Text()
	if val == nil {
		return nil
	}
	m := n.rx.FindStringSubmatch(*val)
	if m == nil || len(m) < 2 {
		return nil
	}
	return &m[1]
}

type NodeIter struct {
	Iter
}

func (i NodeIter) find(c Checker, cs []Checker) NodeIter {
	return NodeIter{IterIter(i.Iter, also(c, cs)...)}
}

func (i NodeIter) Find(cs ...Checker) NodeIter {
	return NodeIter{IterIter(i.Iter, cs...)}
}

func (i NodeIter) For(visit func(n *Node)) {
	for n := i.Next(); n != nil; n = i.Next() {
		visit(n)
	}
}

func (i NodeIter) Next() *Node {
	next := i.Iter
	if next == nil {
		return nil
	}
	if node := next(); node != nil {
		return NewNode(node)
	}
	return nil
}

func (i NodeIter) All() (nodes []*Node) {
	next := i.Iter
	for node := next(); node != nil; node = next() {
		nodes = append(nodes, NewNode(node))
	}
	return
}

func (i NodeIter) Strings(f StringGetter, pat ...string) []string {
	if i.Iter == nil {
		return nil
	}
	return Strings(i.Iter, f, pat...)
}

func (i NodeIter) Integers(f StringGetter) []int {
	if i.Iter == nil {
		return nil
	}
	return Integers(i.Iter, f)
}

type NodeStack struct {
	*Stack
}

/*
func (s NodeStack) All() (nodes []*Node) {
	for _, node := range s.Stack.s {
		nodes = append(nodes, NewNode(node))
	}
	return
}
*/

// ---------------------------------
// If needed, autogenerate these routines

// node methods

func (i NodeIter) A(cs ...Checker) NodeIter {
	return i.find(A, cs)
}

func (i NodeIter) H2(cs ...Checker) NodeIter {
	return i.find(H2, cs)
}

func (i NodeIter) H3(cs ...Checker) NodeIter {
	return i.find(H3, cs)
}
func (i NodeIter) H4(cs ...Checker) NodeIter {
	return i.find(H4, cs)
}

func (i NodeIter) Div(cs ...Checker) NodeIter {
	return i.find(Div, cs)
}

func (i NodeIter) Td(cs ...Checker) NodeIter {
	return i.find(Td, cs)
}

// attr methods

func (i NodeIter) Href(pat ...string) []string {
	if i.Iter == nil {
		return nil
	}
	return Strings(i.Iter, GetHref, pat...)
}

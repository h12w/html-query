// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expr

import (
	"strconv"

	"golang.org/x/net/html"
)

type Iter func() *html.Node

// Pre-order depth first traversal in all descendants
func Descendants(n *html.Node, cs ...Checker) Iter {
	c := And(cs...)
	s := NewStack()
	node := FirstChild(n)
	return func() *html.Node {
		for node != nil || s.Len() > 0 {
			if node != nil {
				s.Push(node)
				r := node
				node = FirstChild(node)
				if c(r) != nil {
					return r
				}
			} else {
				node = s.Pop()
				node = NextSibling(node)
			}
		}
		return nil
	}
}

func IterIter(next Iter, cs ...Checker) Iter {
	find := Find(cs...)
	return func() *html.Node {
		node := next()
		for node != nil {
			if n := find(node); n != nil {
				return n
			}
		}
		return nil
	}
}

func Children(n *html.Node, cs ...Checker) Iter {
	c := And(cs...)
	node := FirstChild(n)
	return func() *html.Node {
		for node != nil {
			r := node
			node = NextSibling(node)
			if c(r) != nil {
				return r
			}
		}
		return nil
	}
}

func Strings(next Iter, f StringGetter, pat ...string) []string {
	ss := []string{}
	p := GetPat(pat)

	// TODO: I have met a bug here once that the program hangs at the next()
	// function call, but I cannot find the data to reproduce it. So Just wait
	// and see it happens again.
	for node := next(); node != nil; node = next() {
		if s := f(node); s != nil {
			ss = append(ss, *GetSubmatch(s, p))
		}
	}
	return ss
}

func Integers(next Iter, f StringGetter) []int {
	ss := []int{}
	for node := next(); node != nil; node = next() {
		s := f(node)
		if s != nil {
			if i, err := strconv.Atoi(*s); err == nil {
				ss = append(ss, i)
			}
		}
	}
	return ss
}

// FILO stack.
type Stack struct {
	s []*html.Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return len(s.s)
}

func (s *Stack) Push(n *html.Node) {
	s.s = append(s.s, n)
}

func (s *Stack) Pop() (n *html.Node) {
	if s.Len() == 0 {
		return nil
	}
	n, s.s = s.s[len(s.s)-1], s.s[:len(s.s)-1]
	return n
}

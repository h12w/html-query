// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expr

import (
	"container/list"
	"code.google.com/p/go.net/html"
)

// Broad first search in all descendants
func Find(cs ...Checker) Checker {
	c := And(cs...)
	return func(n *html.Node) *html.Node {
		q := NewQueue()
		q.PushNodes(Children(n))
		for q.Len() > 0 {
			t := q.Pop()
			if c(t) != nil {
				return t
			} else {
				q.PushNodes(Children(t))
			}
		}
		return nil
	}
}

// Find in direct children
func FindChild(cs ...Checker) Checker {
	c := And(cs...)
	return func(n *html.Node) *html.Node {
		for child := FirstChild(n); child != nil; child = NextSibling(child) {
			if c(child) != nil {
				return child
			}
		}
		return nil
	}
}

// Find in sibling nodes
func FindSibling(cs ...Checker) Checker {
	c := And(cs...)
	return func(n *html.Node) *html.Node {
		for sibling := NextSibling(n); sibling != nil; sibling = NextSibling(sibling) {
			if c(sibling) != nil {
				return sibling
			}
		}
		return nil
	}
}

// FIFO queue.
type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Len() int {
	return q.l.Len()
}

func (q *Queue) Push(n *html.Node) {
	q.l.PushBack(n)
}

func (q *Queue) PushNodes(next Iter) {
	for node := next(); node != nil; node = next() {
		q.Push(node)
	}
}

func (q *Queue) Pop() *html.Node {
	if q.l.Front() == nil {
		return nil
	}
	return q.l.Remove(q.l.Front()).(*html.Node)
}


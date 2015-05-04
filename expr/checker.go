// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expr // import "h12.me/html-query/expr"

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Checker func(*html.Node) *html.Node

func Not(c Checker) Checker {
	return func(n *html.Node) *html.Node {
		if c(n) == nil {
			return n
		}
		return nil
	}
}

func And(cs ...Checker) Checker {
	return func(n *html.Node) *html.Node {
		for _, c := range cs {
			if c(n) == nil {
				return nil
			}
		}
		return n
	}
}

func Pipe(cs ...Checker) Checker {
	return func(n *html.Node) *html.Node {
		for _, c := range cs {
			r := c(n)
			if r == nil {
				return nil
			} else {
				n = r
			}
		}
		return n
	}
}

func Or(cs ...Checker) Checker {
	return func(n *html.Node) *html.Node {
		for _, c := range cs {
			if c(n) != nil {
				return n
			}
		}
		return nil
	}
}

func FirstChild(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.FirstChild
}

func Parent(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.Parent
}

func NextSibling(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.NextSibling
}

func PrevSibling(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.PrevSibling
}

// Node Checkers
// =============

func TypeChecker(t html.NodeType) Checker {
	return func(n *html.Node) *html.Node {
		if n != nil && n.Type == t {
			return n
		}
		return nil
	}
}

var (
	ErrorNode    = TypeChecker(html.ErrorNode)
	TextNode     = TypeChecker(html.TextNode)
	DocumentNode = TypeChecker(html.DocumentNode)
	ElementNode  = TypeChecker(html.ElementNode)
	CommentNode  = TypeChecker(html.CommentNode)
	DoctypeNode  = TypeChecker(html.DoctypeNode)
)

func NonemptyTextNode(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if TextNode(n) != nil && strings.TrimSpace(n.Data) != "" {
		return n
	}
	return nil
}

func AtomChecker(a atom.Atom) Checker {
	return func(n *html.Node) *html.Node {
		if n.DataAtom == a {
			return n
		}
		return nil
	}
}

func ElementChecker(a atom.Atom) Checker {
	return And(ElementNode, AtomChecker(a))
}

// Attribute Checkers
// ==================

func AttributeCmpChecker(key string, cmp func(string) bool) Checker {
	return func(n *html.Node) *html.Node {
		attr := GetAttr(n, key)
		if attr != nil && cmp(*attr) {
			return n
		}
		return nil
	}
}

func Attr(key, pat string) Checker {
	rx := regexp.MustCompile(pat)
	return AttributeCmpChecker(key, func(val string) bool {
		return rx.MatchString(val)
	})
}

func AttrChecker(key string) func(string) Checker {
	return func(pat string) Checker {
		return Attr(key, pat)
	}
}

func HasAttr(key string) Checker {
	return func(n *html.Node) *html.Node {
		if GetAttr(n, key) != nil {
			return n
		}
		return nil
	}
}

func NoAttr(key string) Checker {
	return func(n *html.Node) *html.Node {
		if GetAttr(n, key) != nil {
			return nil
		}
		return n
	}
}

func fieldsToSet(val string, sep rune) map[string]bool {
	m := make(map[string]bool)
	fields := strings.FieldsFunc(val, func(r rune) bool { return r == sep })
	for _, field := range fields {
		m[field] = true
	}
	return m
}

func SeperatedAttrChecker(name string, sep rune) func(...string) Checker {
	return func(classes ...string) Checker {
		return AttributeCmpChecker(name, func(val string) bool {
			s := fieldsToSet(val, sep)
			for _, class := range classes {
				if !s[class] {
					return false
				}
			}
			return true
		})
	}
}

func Text_(rx *regexp.Regexp) Checker {
	return func(n *html.Node) *html.Node {
		if s := GetText(n); s != nil {
			if rx.MatchString(*s) {
				return n
			}
		}
		return nil
	}
}

func Text(pat string) Checker {
	return Text_(regexp.MustCompile(pat))
}

func CaptionText(pat string) Checker {
	return Find(Caption, Text(pat))
}

var (
	Ahref = And(A, HasAttr("href"))
)

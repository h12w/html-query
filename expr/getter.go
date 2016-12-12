// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expr

import (
	"regexp"

	"golang.org/x/net/html"
)

type StringGetter func(*html.Node) *string

func GetAttr(n *html.Node, key string) *string {
	if n == nil {
		return nil
	}
	for _, a := range n.Attr {
		if a.Key == key {
			return &a.Val
		}
	}
	return nil
}

func GetAttrSubmatch(n *html.Node, key, pat string) *string {
	return GetSubmatch(GetAttr(n, key), pat)
}

func GetSubmatch_(s *string, rx *regexp.Regexp) *string {
	if s == nil {
		return nil
	}
	m := rx.FindStringSubmatch(*s)
	if m == nil || len(m) < 2 {
		return nil
	}
	return &m[1]
}

func GetSubmatch(s *string, pat string) *string {
	if pat == "" {
		return s
	}
	return GetSubmatch_(s, regexp.MustCompile(pat))
}

func GetTextNodeText(n *html.Node) *string {
	if NonemptyTextNode(n) != nil {
		return &n.Data
	}
	return nil
}

func GetText(n *html.Node) *string {
	if s := GetTextNodeText(n); s != nil {
		return s
	}

	for c := FirstChild(n); c != nil; c = NextSibling(c) {
		if s := GetTextNodeText(c); s != nil {
			return s
		}
	}
	return nil
}

func GetHref(n *html.Node) *string {
	if n == nil {
		return nil
	}
	return GetAttr(n, "href")
}

func GetSrc(n *html.Node) *string {
	if n == nil {
		return nil
	}
	return GetAttr(n, "src")
}

func GetPat(pat []string) string {
	if len(pat) > 1 {
		panic("pat should be either ommited or only one string.")
	} else if len(pat) == 0 {
		return "" // empty string indicates that the whole string should be got.
	}
	return pat[0]
}

/*
func AttrValueGetter(key string) StringGetter {
	return func(n *html.Node) *string {
		return GetAttrValue(n, key)
	}
}
*/

// Copyright 2005, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package query

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
	"h12.io/html-query/expr"
)

func (n *Node) PlainText() *string {
	if n == nil {
		return nil
	}
	var w bytes.Buffer
	if err := renderPlain(&w, &n.n); err != nil {
		return nil
	}
	s := strings.TrimSpace(w.String())
	return &s
}

func renderPlain(w writer, n *html.Node) error {
	switch n.Type {
	case html.TextNode:
		w.WriteString(n.Data)
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := renderPlain(w, c); err != nil {
				return err
			}
		}
		return nil
	case html.ElementNode:
		return renderPlainElementNode(w, n)
	}
	return nil
}

func renderPlainElementNode(w writer, n *html.Node) error {
	if c := n.FirstChild; c != nil && c.Type == html.TextNode && strings.HasPrefix(c.Data, "\n") {
		switch n.Data {
		case "pre", "listing", "textarea":
			if err := w.WriteByte('\n'); err != nil {
				return err
			}
		}
	}

	switch n.Data {
	case "iframe", "noembed", "noframes", "noscript", "plaintext", "script", "style", "xmp":
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				if _, err := w.WriteString(c.Data); err != nil {
					return err
				}
			} else {
				if err := renderPlain(w, c); err != nil {
					return err
				}
			}
		}
		if n.Data == "plaintext" {
			return plaintextAbort
		}
		return nil
	case "a":
		if n.FirstChild != nil && isURL(n.FirstChild.Data) {
			renderPlainChild(w, n)
		} else if url := expr.GetAttr(n, "href"); url != nil && *url != "" {
			w.WriteString("[")
			renderPlainChild(w, n)
			w.WriteString("](")
			w.WriteString(*url)
			w.WriteString(")")
		} else {
			renderPlainChild(w, n)
		}
		return nil
	}

	renderPlainChild(w, n)
	// write break after children are written
	switch n.Data {
	case "p", "br", "div":
		writeBreak(w)
	}
	return nil
}

func isURL(s string) bool {
	s = strings.TrimSpace(s)
	return strings.Contains(s, "http://") ||
		strings.Contains(s, "@")
}

func renderPlainChild(w writer, n *html.Node) error {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err := renderPlain(w, c); err != nil {
			return err
		}
	}
	return nil
}

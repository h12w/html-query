// The code in this file is copied and modified from
// http://code.google.com/p/go.net.

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file under
// http://code.google.com/p/go.net.

package query

import (
	"bufio"
	"code.google.com/p/go.net/html"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type writer interface {
	io.Writer
	WriteByte(c byte) error // in Go 1.1, use io.ByteWriter
	WriteString(string) (int, error)
}

func (n *Node) PrettyPrint() {
	n.PrettyRender(os.Stdout, 4)
}

// PrettyRender renders prettily the parse tree n to the given writer,
// for easily viewing as plain text.
func (n *Node) PrettyRender(w io.Writer, indentSize int) error {
	if n == nil {
		return nil
	}
	if x, ok := w.(writer); ok {
		return render(x, &n.n, indentSize)
	}
	buf := bufio.NewWriter(w)
	if err := render(buf, &n.n, indentSize); err != nil {
		return err
	}
	return buf.Flush()
}

// plaintextAbort is returned from render1 when a <plaintext> element
// has been rendered. No more end tags should be rendered after that.
var plaintextAbort = errors.New("html: internal error (plaintext abort)")

func render(w writer, n *html.Node, size int) error {
	err := render1(w, n, -1, size)
	if err == plaintextAbort {
		err = nil
	}
	return err
}

func render1(w writer, n *html.Node, level, size int) error {
	if !isSpace(n) && n.Type != html.DocumentNode && n.Type != html.DoctypeNode {
		if err := writeBreak(w); err != nil {
			return err
		}
		if err := writeIndent(w, level, size); err != nil {
			return err
		}
	}

	if err := renderSimpleNode(w, n); err != nil {
		return err
	}

	switch n.Type {
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := render1(w, c, level+1, size); err != nil {
				return err
			}
		}
		return nil
	case html.ElementNode:
		return renderElementNode(w, n, level, size)
	}
	return nil
}

func renderSimpleNode(w writer, n *html.Node) error {
	// Render non-element nodes; these are the easy cases.
	switch n.Type {
	case html.ErrorNode:
		return errors.New("html: cannot render an html.ErrorNode node")
	case html.TextNode:
		return escape(w, n.Data)
	case html.CommentNode:
		if _, err := w.WriteString("<!--"); err != nil {
			return err
		}
		if _, err := w.WriteString(n.Data); err != nil {
			return err
		}
		if _, err := w.WriteString("-->"); err != nil {
			return err
		}
		return nil
	case html.DoctypeNode:
		if _, err := w.WriteString("<!DOCTYPE "); err != nil {
			return err
		}
		if _, err := w.WriteString(n.Data); err != nil {
			return err
		}
		if n.Attr != nil {
			var p, s string
			for _, a := range n.Attr {
				switch a.Key {
				case "public":
					p = a.Val
				case "system":
					s = a.Val
				}
			}
			if p != "" {
				if _, err := w.WriteString(" PUBLIC "); err != nil {
					return err
				}
				if err := writeQuoted(w, p); err != nil {
					return err
				}
				if s != "" {
					if err := w.WriteByte(' '); err != nil {
						return err
					}
					if err := writeQuoted(w, s); err != nil {
						return err
					}
				}
			} else if s != "" {
				if _, err := w.WriteString(" SYSTEM "); err != nil {
					return err
				}
				if err := writeQuoted(w, s); err != nil {
					return err
				}
			}
		}
		if err := w.WriteByte('>'); err != nil {
			return err
		}
		return nil
	case html.ElementNode, html.DocumentNode:
		// No-op.
	default:
		return errors.New("html: unknown node type")
	}
	return nil
}

func renderElementNode(w writer, n *html.Node, level, size int) error {
	if err := renderOpeningTag(w, n); err != nil {
		return err
	}

	// Add initial newline where there is danger of a newline beging ignored.
	if c := n.FirstChild; c != nil && c.Type == html.TextNode && strings.HasPrefix(c.Data, "\n") {
		switch n.Data {
		case "pre", "listing", "textarea":
			if err := w.WriteByte('\n'); err != nil {
				return err
			}
		}
	}

	// Render any child nodes.
	switch n.Data {
	case "iframe", "noembed", "noframes", "noscript", "plaintext", "script", "style", "xmp":
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				if _, err := w.WriteString(c.Data); err != nil {
					return err
				}
			} else {
				if err := render1(w, c, level+1, size); err != nil {
					return err
				}
			}
		}
		if n.Data == "plaintext" {
			// Don't render anything else. <plaintext> must be the
			// last element in the file, with no closing tag.
			return plaintextAbort
		}
	default:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := render1(w, c, level+1, size); err != nil {
				return err
			}
		}
	}

	// Render the </xxx> closing tag.
	if err := writeBreak(w); err != nil {
		return err
	}
	if err := writeIndent(w, level, size); err != nil {
		return err
	}
	if _, err := w.WriteString("</"); err != nil {
		return err
	}
	if _, err := w.WriteString(n.Data); err != nil {
		return err
	}
	if err := w.WriteByte('>'); err != nil {
		return err
	}
	return nil
}

// writeQuoted writes s to w surrounded by quotes. Normally it will use double
// quotes, but if s contains a double quote, it will use single quotes.
// It is used for writing the identifiers in a doctype declaration.
// In valid HTML, they can't contain both types of quotes.
func writeQuoted(w writer, s string) error {
	var q byte = '"'
	if strings.Contains(s, `"`) {
		q = '\''
	}
	if err := w.WriteByte(q); err != nil {
		return err
	}
	if _, err := w.WriteString(s); err != nil {
		return err
	}
	if err := w.WriteByte(q); err != nil {
		return err
	}
	return nil
}

// Section 12.1.2, "Elements", gives this list of void elements. Void elements
// are those that can't have any contents.
var voidElements = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}

const escapedChars = "&'<>\"\r"

func escape(w writer, s string) error {
	s = strings.TrimSpace(s)

	i := strings.IndexAny(s, escapedChars)
	for i != -1 {
		if _, err := w.WriteString(s[:i]); err != nil {
			return err
		}
		var esc string
		switch s[i] {
		case '&':
			esc = "&amp;"
		case '\'':
			// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
			esc = "&#39;"
		case '<':
			esc = "&lt;"
		case '>':
			esc = "&gt;"
		case '"':
			// "&#34;" is shorter than "&quot;".
			esc = "&#34;"
		case '\r':
			esc = "&#13;"
		default:
			panic("unrecognized escape character")
		}
		s = s[i+1:]
		if _, err := w.WriteString(esc); err != nil {
			return err
		}
		i = strings.IndexAny(s, escapedChars)
	}
	_, err := w.WriteString(s)
	return err
}

func writeIndent(w writer, level, size int) error {
	for i := 0; i < level*size; i++ {
		if _, err := w.WriteString(` `); err != nil {
			return err
		}
	}
	return nil
}

func writeBreak(w writer) error {
	_, err := w.Write([]byte{'\n'})
	return err
}

func isSpace(n *html.Node) bool {
	return n != nil && n.Type == html.TextNode && strings.TrimSpace(n.Data) == ""
}

func renderOpeningTag(w writer, n *html.Node) error {
	// Render the <xxx> opening tag.
	if err := w.WriteByte('<'); err != nil {
		return err
	}
	if _, err := w.WriteString(n.Data); err != nil {
		return err
	}
	for _, a := range n.Attr {
		if err := w.WriteByte(' '); err != nil {
			return err
		}
		if a.Namespace != "" {
			if _, err := w.WriteString(a.Namespace); err != nil {
				return err
			}
			if err := w.WriteByte(':'); err != nil {
				return err
			}
		}
		if _, err := w.WriteString(a.Key); err != nil {
			return err
		}
		if _, err := w.WriteString(`="`); err != nil {
			return err
		}
		if err := escape(w, a.Val); err != nil {
			return err
		}
		if err := w.WriteByte('"'); err != nil {
			return err
		}
	}
	if voidElements[n.Data] {
		if n.FirstChild != nil {
			return fmt.Errorf("html: void element <%s> has child nodes", n.Data)
		}
		_, err := w.WriteString("/>")
		if err != nil {
			return err
		}
		return nil
	}
	if err := w.WriteByte('>'); err != nil {
		return err
	}
	return nil
}

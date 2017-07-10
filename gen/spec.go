// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	"h12.me/html-query"
	. "h12.me/html-query/expr"
)

const (
	SpecUrl = `https://html.spec.whatwg.org/multipage/indices.html`

//	SPEC_URL = `http://www.w3.org/html/wg/drafts/html/master/single-page.html`
)

// skip the two because they does not appear in exp/html
var (
	ignoredElements = map[string]bool{
		"MathML":   true,
		"SVG":      true,
		"main":     true,
		"menuitem": true,
		"template": true,
		"picture":  true,
		"slot":     true,

		"autonomous custom elements": true,
	}
)

type Element struct {
	Name       string
	Attributes []*Attribute
}

type Attribute struct {
	Name     string
	Type     string
	IsGlobal bool
}

type AttributeTable struct {
	Set map[string]*Attribute
}

type ElementTable struct {
	Set  map[string]*Element
	Skip map[string]bool
}

type Spec struct {
	ElemTable *ElementTable
	AttrTable *AttributeTable
}

func DownloadSpec(file string) {
	resp, err := http.Get(SpecUrl)
	c(err)
	defer resp.Body.Close()
	f, err := os.Create(file)
	c(err)
	defer f.Close()
	io.Copy(f, resp.Body)
}

func parseSpec(file io.Reader) *Spec {
	root, err := query.Parse(file)
	c(err)
	attrTable := parseAttributeTable(root)
	elemTable := parseElementTable(root, attrTable)
	return &Spec{elemTable, attrTable}
}

func (t *ElementTable) Elements() []*Element {
	names := make([]string, len(t.Set))
	i := 0
	for k, _ := range t.Set {
		names[i] = k
		i++
	}
	sort.Sort(sort.StringSlice(names))
	elements := make([]*Element, len(names))
	for i := range elements {
		elements[i] = t.Set[names[i]]
	}
	return elements
}

func (t *AttributeTable) Attributes() []*Attribute {
	names := make([]string, len(t.Set))
	i := 0
	for k, _ := range t.Set {
		names[i] = k
		i++
	}
	sort.Sort(sort.StringSlice(names))
	attrs := make([]*Attribute, len(names))
	for i := range attrs {
		attrs[i] = t.Set[names[i]]
	}
	return attrs
}

func parseAttributeTable(root *query.Node) *AttributeTable {
	attrSet := make(map[string]*Attribute)
	attrTable := root.Table(CaptionText("List of attributes"))
	if attrTable == nil {
		panic("Cannot find List of attributes")
	}
	for _, tr := range attrTable.Tbody().Children(Tr).All() {
		name := *tr.Th().Code().Text()
		attr := &Attribute{Name: name}
		td := tr.Children(Td).All()
		if elemName := td[0].A().Text(); elemName != nil {
			if *elemName == "HTML elements" {
				attr.IsGlobal = true
			}
		}
		attr.Type = strings.Replace(*td[2].PlainText(), "\n", "", -1)
		// Attention: attribute may be duplicated, just choose the first one
		// but set isglobal if one of it is global
		if attrSet[name] == nil {
			attrSet[name] = attr
		} else if attr.IsGlobal {
			attrSet[name].IsGlobal = true
		}
	}
	return &AttributeTable{Set: attrSet}
}
func parseElementTable(root *query.Node, attrTable *AttributeTable) *ElementTable {
	elemSet := make(map[string]*Element)
	attrSet := attrTable.Set
	elementTable := root.Table(CaptionText("List of elements"))
	for _, tr := range elementTable.Tbody().Children(Tr).All() {
		td := tr.Children(Td).All()
		for _, elemLink := range tr.Th().Descendants(Ahref).All() {
			elem := &Element{Name: strings.TrimSpace(*elemLink.Text())}
			for _, attrLink := range td[4].Descendants(Ahref).All() {
				attrName := strings.TrimSpace(*attrLink.Text())
				if attr := attrSet[attrName]; attr != nil {
					elem.Attributes = append(elem.Attributes, attr)
				}
			}
			elemSet[elem.Name] = elem
		}
	}
	return &ElementTable{
		Set:  elemSet,
		Skip: ignoredElements,
	}
}

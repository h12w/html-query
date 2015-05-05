// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func (spec *Spec) GenerateExpr() {
	elemTable, attrTable := spec.ElemTable, spec.AttrTable
	file := "output/auto_expr.go"
	f, err := os.Create(file)
	c(err)
	fp(f, `package expr // import "h12.me/html-query/expr"`)
	fp(f, "import (")
	fp(f, `"golang.org/x/net/html/atom"`)
	fp(f, ")")

	fp(f, "var (")
	for _, elem := range elemTable.Elements() {
		if elemTable.Skip[elem.Name] {
			continue
		}
		nodeId := toid(elem.Name)
		fp(f, nodeId, "=", "ElementChecker(atom.", toid(elem.Name), ")")
	}
	fp(f, ")")
	fp(f, "")

	fp(f, "var (")
	for _, attr := range attrTable.Attributes() {
		attrId := toid(attr.Name)
		if elemTable.Set[attr.Name] != nil {
			attrId += "_"
		}
		if strings.Contains(attr.Type, "space-separated tokens") {
			fp(f, attrId, ` = SeperatedAttrChecker("`, attr.Name, `", ' ')`)
		} else {
			fp(f, attrId, ` = AttrChecker("`, attr.Name, `")`)
		}
	}
	fp(f, ")")

	f.Close()
	format(file)
}

func (spec *Spec) GenerateChain() {
	file := "output/auto_chain.go"
	f, err := os.Create(file)
	c(err)
	fp(f, "package query")
	fp(f, "import (")
	fp(f, `. "h12.me/html-query/expr"`)
	fp(f, ")")

	spec.generateChainSmall(f)

	//spec.generateChainBloaded(f)

	fp(f, "")
	f.Close()
	format(file)
}

func (spec *Spec) generateChainSmall(f io.Writer) {
	elemTable, attrTable := spec.ElemTable, spec.AttrTable
	for _, elem := range elemTable.Elements() {
		if elemTable.Skip[elem.Name] {
			continue
		}
		nodeId := toid(elem.Name)
		nodeFinderSmall(f, nodeId)
	}

	for _, attr := range attrTable.Attributes() {
		attrId := toid(attr.Name)
		if elemTable.Set[attr.Name] != nil {
			attrId += "_"
		}
		nodeAttribute(f, attr.Name, attrId, "Node")
	}
}

func (spec *Spec) generateChainBloaded(f io.Writer) {
	elemTable, attrTable := spec.ElemTable, spec.AttrTable
	for _, elem := range elemTable.Elements() {
		if elemTable.Skip[elem.Name] {
			continue
		}
		nodeId := toid(elem.Name)
		nodeType := nodeId + "Node"
		fp(f, "")
		fp(f, "// ", nodeId)
		fp(f, "")
		nodeDefinition(f, nodeType)
		nodeConstructor(f, nodeType)
		finderName := nodeId
		if a, ok := attrTable.Set[elem.Name]; ok && a.IsGlobal {
			finderName += "Node"
		}
		nodeFinder(f, finderName, nodeId, nodeType)
		for _, attr := range elem.Attributes {
			if !attr.IsGlobal {
				attrId := toid(attr.Name)
				nodeAttribute(f, attr.Name, attrId, nodeType)
			}
		}
		fp(f, "")
	}

	for _, attr := range attrTable.Attributes() {
		if attr.IsGlobal {
			attrId := toid(attr.Name)
			nodeAttribute(f, attr.Name, attrId, "Node")
		}
	}
}

func nodeDefinition(f io.Writer, name string) {
	fp(f, "type ", name, " struct {")
	fp(f, "Node")
	fp(f, "}")
	fp(f, "")
}

func nodeConstructor(f io.Writer, name string) {
	fp(f, "func New", name, "(n *Node) *", name, "{")
	fp(f, "if n == nil {")
	fp(f, "return nil")
	fp(f, "}")
	fp(f, "return &", name, "{*n}")
	fp(f, "}")
	fp(f, "")
}

func nodeFinder(f io.Writer, finderName, nodeId, nodeType string) {
	fp(f, "func (n *Node) ", finderName, "(cs ...Checker) *", nodeType, " {")
	fp(f, "return New", nodeType, "(n.find(", nodeId, ", cs))")
	fp(f, "}")
	fp(f, "")
}

func nodeFinderSmall(f io.Writer, nodeId string) {
	fp(f, "func (n *Node) ", nodeId, "(cs ...Checker) *Node {")
	fp(f, "return n.find(", nodeId, ", cs)")
	fp(f, "}")
	fp(f, "")
}

func nodeAttribute(f io.Writer, attrName, attrId, nodeType string) {
	fp(f, "func (n *", nodeType, ") ", attrId, "(pat ...string) *string {")
	fp(f, `return n.Attr("`, attrName, `", pat...)`)
	fp(f, "}")
	fp(f, "")

}

func toid(s string) string {
	return strings.Replace(strings.Title(s), "-", "", -1)
}

func format(file string) {
	cmd := exec.Command("go", "fmt", file)
	err := cmd.Start()
	c(err)
	err = cmd.Wait()
	c(err)
}

func fp(w io.Writer, v ...interface{}) {
	fmt.Fprint(w, v...)
	fmt.Fprintln(w)
}

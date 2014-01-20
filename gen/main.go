// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
)

func main() {
	htmlSpec := "html_spec.htm"
	f, err := os.Open(htmlSpec)
	if err != nil {
		DownloadSpec(htmlSpec)
		f, err = os.Open(htmlSpec)
		c(err)
	}
	defer f.Close()

	spec := parseSpec(f)

	os.Mkdir("output", 0755)
	spec.GenerateExpr()
	spec.GenerateChain()
}

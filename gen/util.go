// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

func c(err error) {
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
}

func p(v ...interface{}) {
	log.Println(v...)
}

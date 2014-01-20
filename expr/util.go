// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expr

import (
	"strconv"
	"strings"
	"time"
)

func ToInt(ps *string) *int {
	if ps == nil {
		return nil
	}
	i, err := strconv.Atoi(strings.TrimSpace(*ps))
	if err != nil {
		return nil
	}
	return &i
}

func ToFloat(ps *string) *float64 {
	if ps == nil {
		return nil
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(*ps), 64)
	if err != nil {
		return nil
	}
	return &f
}

func ToHex(ps *string) *int {
	if ps == nil {
		return nil
	}
	i64, err := strconv.ParseInt(strings.TrimSpace(*ps), 16, 64)
	if err != nil {
		return nil
	}
	i := int(i64)
	return &i
}

func ToTime(ps *string, layout string) *time.Time {
	if ps == nil {
		return nil
	}
	t, err := time.Parse(layout, *ps)
	if err != nil {
		return nil
	}
	return &t
}

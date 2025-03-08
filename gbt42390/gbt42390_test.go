// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

package gbt42390

import (
	"maps"
	"strings"
	"testing"

	"github.com/issue9/assert/v4"
)

func TestAll(t *testing.T) {
	a := assert.New(t, false)

	a.Length(maps.Collect(All(0)), len(codes))

	for k := range All(1) {
		a.True(strings.HasSuffix(k, "0000"))
	}

	for k := range All(2) {
		a.True(strings.HasSuffix(k, "00"))
	}

	a.PanicString(func() {
		All(3)
	}, "无效的 level 参数")
}

func TestIsValid(t *testing.T) {
	a := assert.New(t, false)
	a.True(IsValid("10000")).
		False(IsValid("90000"))
}

func TestGet(t *testing.T) {
	a := assert.New(t, false)
	a.Equal(Get("10000"), "快递封套").
		Equal(Get("10202"), "聚酯（PET）可循环封套").
		Empty(Get("90000"))
}

func TestParent(t *testing.T) {
	a := assert.New(t, false)
	a.Equal(Parent("10202"), "10200").
		Equal(Parent("10200"), "10000").
		Equal(Parent("10000"), "00000").
		Equal("00000", "00000")
}

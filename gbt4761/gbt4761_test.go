// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

package gbt4761

import (
	"maps"
	"strings"
	"testing"

	"github.com/issue9/assert/v4"
)

func TestAll(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(maps.Collect(All(false)), codes)

	for k := range All(true) {
		a.True(strings.HasSuffix(k, "0"))
	}
}

func TestGet(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(Get("1"), "配偶").
		Equal(Get("0"), "本人或户主").
		Equal(Get("9"), "其他").
		Equal(Get("11"), "夫").
		Equal(Get("112"), "夫之子").
		Equal(Get("1122"), "夫之长子").
		Zero(Get("11223"))
}

func TestIsValid(t *testing.T) {
	a := assert.New(t, false)

	a.True(IsValid("1")).
		True(IsValid("0")).
		True(IsValid("9")).
		True(IsValid("11")).
		True(IsValid("112")).
		True(IsValid("1122")).
		False(IsValid("11223"))
}

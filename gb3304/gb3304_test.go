// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

package gb3304

import (
	"slices"
	"testing"

	"github.com/issue9/assert/v4"
)

func TestAll(t *testing.T) {
	a := assert.New(t, false)
	a.Equal(len(slices.Collect(All())), 56)
}

func TestGet(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(Get("01").Roman, "Han").
		Equal(Get("Blang").Numeric, "34").
		Equal(Get("BN").Numeric, "47").
		Equal(Get("布依族").Numeric, "09").
		Nil(Get("不存在的民族"))
}

func TestIsValid(t *testing.T) {
	a := assert.New(t, false)
	a.True(IsValid("01")).
		True(IsValid("BN")).
		True(IsValid("Blang")).
		True(IsValid("布依族")).
		False(IsValid("不存在的民族"))
}

// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

package gbt4658

import (
	"testing"

	"github.com/issue9/assert/v4"
)

func TestIsValid(t *testing.T) {
	a := assert.New(t, false)

	a.True(IsValid("10")).
		True(IsValid("30")).
		False(IsValid("99")).
		False(IsValid("9"))
}

func TestName(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(Name("10"), "研究生教育").
		Equal(Name("17"), "研究生班毕业").
		Equal(Name("22"), "大学本科结业").
		Equal(Name("1"), "")
}

func TestAll(t *testing.T) {
	a := assert.New(t, false)
	a.Equal(All(), educations)
}

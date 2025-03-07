// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

//go:generate web locale -l=und -m -f=yaml ./
//go:generate web update-locale -src=./locales/und.yaml -dest=./locales/zh.yaml

// Package gb 国标代码
package gb

import "github.com/issue9/localeutil"

var errInvalidFromat = localeutil.Error("invalid format")

// ErrInvalidFormat 表示格式错误
func ErrInvalidFormat() error { return errInvalidFromat }

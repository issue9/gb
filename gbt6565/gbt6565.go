// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

// Package gbt6565 GB/T 6565-2015 职业分类代码
package gbt6565

import (
	"iter"
	"strings"
)

// All 遍历所有职业
//
// level 表示分类re级别，0 表示所有，1 表示大类，2 表示中类
func All(level int) iter.Seq2[string, string] {
	switch level {
	case 0:
		return func(yield func(string, string) bool) {
			for k, v := range classes {
				if !yield(k, v) {
					break
				}
			}
		}
	case 1:
		return func(yield func(string, string) bool) {
			for k, v := range classes {
				if !strings.HasSuffix(k, "0000") {
					continue
				}
				if !yield(k, v) {
					break
				}
			}
		}
	case 2:
		return func(yield func(string, string) bool) {
			for k, v := range classes {
				if !strings.HasSuffix(k, "00") {
					continue
				}
				if !yield(k, v) {
					break
				}
			}
		}
	default:
		panic("无效的 level 参数")
	}
}

// Parent 返回 code 的父级职业代码
//
// 如果 code 必须为一个有效的职业代码
func Parent(code string) string {
	switch {
	case code[3:] != "00":
		return code[:3] + "00"
	case code[1:] != "0000":
		return code[:1] + "0000"
	default:
		return "00000"
	}
}

// Get 根据 code 返回对应的职业名称
func Get(code string) string { return classes[code] }

// IsValid 判断 code 是否是一个有效的职业代码
func IsValid(code string) bool {
	_, found := classes[code]
	return found
}

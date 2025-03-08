// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

// Package gbt42390 GB/T 42390-2023 快递包装分类与代码
package gbt42390

import (
	"iter"
	"strings"
)

// All 遍历所有代码
//
// level 返回的级别，可以有以下几个值：
//   - 0 表示所有；
//   - 1 表示一级分类；
//   - 2 表示二级分类；
func All(level int) iter.Seq2[string, string] {
	switch level {
	case 0:
		return func(yield func(string, string) bool) {
			for k, v := range codes {
				if !yield(k, v) {
					break
				}
			}
		}
	case 1:
		return func(yield func(string, string) bool) {
			for k, v := range codes {
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
			for k, v := range codes {
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

// Parent 返回 code 的父级分类
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

// Get 获取指定代码表示的家庭关系
func Get(code string) string { return codes[code] }

// IsValid 验证数据的正确性
func IsValid(code string) bool {
	_, found := codes[code]
	return found
}

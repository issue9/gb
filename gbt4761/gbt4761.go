// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

// Package gbt4761 GB/T 4761-2008 家庭关系代码
//
// 代码有以下几种长度的格式：
//   - 1 表示基本的关系，值为 1-8；
//   - 2 表示所有的默认家庭关系；
//   - 3 表示组合关系，拆为两段，长度分别为 2 和 1，根据长度分别根据 1 和 2 规则查找；
//   - 4 表示给合关系，拆为两段，长度分别为 2 和 2，根据 2 规则查找；
package gbt4761

import "iter"

// All 遍历所有代码
//
// basic 只遍历基本的关系；
func All(basic bool) iter.Seq2[string, string] {
	if basic {
		return func(yield func(string, string) bool) {
			for k, v := range codes {
				if k[1] == '0' && k[0] != '9' { // 9 属于 8 的延续
					if !yield(k, v) {
						break
					}
				}
			}
		}
	}

	return func(yield func(string, string) bool) {
		for k, v := range codes {
			if !yield(k, v) {
				break
			}
		}
	}
}

// Get 获取指定代码表示的家庭关系
//
// code 的长度为 1，2，3，4 的都可以。具体可参考文档。
func Get(code string) string {
	switch len(code) {
	case 1:
		if code == "9" { // 9 属于 8 的延续
			code = "8"
		}
		return codes[code+"0"]
	case 2:
		return codes[code]
	case 3, 4:
		if name1 := Get(code[:2]); name1 != "" {
			if name2 := Get(code[2:]); name2 != "" {
				return name1 + "之" + name2
			}
		}
	}
	return ""
}

// IsValid 验证数据的正确性
//
// code 的长度为 1，2，3，4 的都可以。具体可参考文档。
func IsValid(code string) (ok bool) {
	switch len(code) {
	case 1:
		if code == "9" { // 9 属于 8 的延续
			code = "8"
		}
		_, ok = codes[code+"0"]
	case 2:
		_, ok = codes[code]
	case 3, 4:
		ok = IsValid(code[:2]) && IsValid(code[2:])
	}
	return ok
}

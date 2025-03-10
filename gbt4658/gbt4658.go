// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

// Package gbt4658 GB/T 4658-2006 学历代码
package gbt4658

import "maps"

var educations = map[string]string{
	"10": "研究生教育",
	"11": "博士研究生毕业",
	"12": "博士研究生结业",
	"13": "博士研究生肄业",
	"14": "硕士研究生毕业",
	"15": "硕士研究生结业",
	"16": "硕士研究生肄业",
	"17": "研究生班毕业",
	"18": "研究生班结业",
	"19": "研究生班肄业",

	"20": "大学本科教育",
	"21": "大学本科毕业",
	"22": "大学本科结业",
	"23": "大学本科肄业",

	"28": "大学普通班毕业",

	"30": "大学专科教育",
	"31": "大学专科毕业",
	"32": "大学专科结业",
	"33": "大学专科肄业",

	"40": "中等职业教育",
	"41": "中等专科毕业",
	"42": "中等专科结业",
	"43": "中等专科肄业",
	"44": "职业高中毕业",
	"45": "职业高中结业",
	"46": "职业高中肄业",
	"47": "技工学校毕业",
	"48": "技工学校结业",
	"49": "技工学校肄业",

	"60": "普通高级中学教育",
	"61": "普通高中毕业",
	"62": "普通高中结业",
	"63": "普通高中肄业",

	"70": "初级中学教育",
	"71": "初中毕业",
	"72": "初中肄业",

	"80": "小学教育",
	"81": "小学毕业",
	"82": "小学肄业",

	"90": "其他",
}

// All 返回所有学历代码
func All() map[string]string { return maps.Clone(educations) }

// IsValid 验证数据的正确性
func IsValid(id string) bool {
	_, found := educations[id]
	return found
}

// Name 获取 ID 对应的名称
func Name(id string) string {
	return educations[id]
}

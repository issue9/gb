// SPDX-FileCopyrightText: 2025 caixw
//
// SPDX-License-Identifier: MIT

// Package gb3304 GB3304-1991 民族代码
package gb3304

import (
	"iter"
	"slices"
)

var nations = []*Nation{
	{Roman: "Achang", Alpha: "AC", Name: "阿昌族", Numeric: "39"},
	{Roman: "Bai", Alpha: "BA", Name: "白族", Numeric: "14"},
	{Roman: "Blang", Alpha: "BL", Name: "布朗族", Numeric: "34"},
	{Roman: "Bonan", Alpha: "BN", Name: "保安族", Numeric: "47"},
	{Roman: "Buyei", Alpha: "BY", Name: "布依族", Numeric: "09"},
	{Roman: "Chosen", Alpha: "CS", Name: "朝鲜族", Numeric: "10"},
	{Roman: "Dai", Alpha: "DA", Name: "傣族", Numeric: "18"},
	{Roman: "Daur", Alpha: "DU", Name: "达斡尔族", Numeric: "31"},
	{Roman: "Deang", Alpha: "DE", Name: "德昂族", Numeric: "46"},
	{Roman: "Derung", Alpha: "DR", Name: "独龙族", Numeric: "51"},
	{Roman: "Dong", Alpha: "DO", Name: "侗族", Numeric: "12"},
	{Roman: "Dongxiang", Alpha: "DX", Name: "东乡族", Numeric: "26"},
	{Roman: "Ewenki", Alpha: "EW", Name: "鄂温克族", Numeric: "45"},
	{Roman: "Gaoshan", Alpha: "GS", Name: "高山族", Numeric: "23"},
	{Roman: "Gelao", Alpha: "GL", Name: "仡佬族", Numeric: "37"},
	{Roman: "Gin", Alpha: "GI", Name: "京族", Numeric: "49"},
	{Roman: "Han", Alpha: "HA", Name: "汉族", Numeric: "01"},
	{Roman: "Hani", Alpha: "HN", Name: "哈尼族", Numeric: "16"},
	{Roman: "Hezhen", Alpha: "HZ", Name: "赫哲族", Numeric: "53"},
	{Roman: "Hui", Alpha: "HU", Name: "回族", Numeric: "03"},
	{Roman: "Jingpo", Alpha: "JP", Name: "景颇族", Numeric: "28"},
	{Roman: "Jino", Alpha: "JN", Name: "基诺族", Numeric: "56"},
	{Roman: "Kazak", Alpha: "KZ", Name: "哈萨克族", Numeric: "17"},
	{Roman: "Kirgiz", Alpha: "KG", Name: "柯尔克孜族", Numeric: "29"},
	{Roman: "Lahu", Alpha: "LH", Name: "拉祜族", Numeric: "24"},
	{Roman: "Lhoba", Alpha: "LB", Name: "珞巴族", Numeric: "55"},
	{Roman: "Li", Alpha: "LI", Name: "黎族", Numeric: "19"},
	{Roman: "Lisu", Alpha: "LS", Name: "傈僳族", Numeric: "20"},
	{Roman: "Man", Alpha: "MA", Name: "满族", Numeric: "11"},
	{Roman: "Maonan", Alpha: "MN", Name: "毛南族", Numeric: "36"},
	{Roman: "Miao", Alpha: "MH", Name: "苗族", Numeric: "06"},
	{Roman: "Monba", Alpha: "MB", Name: "门巴族", Numeric: "54"},
	{Roman: "Mongol", Alpha: "MG", Name: "蒙古族", Numeric: "02"},
	{Roman: "Mulao", Alpha: "ML", Name: "仫佬族", Numeric: "32"},
	{Roman: "Naxi", Alpha: "NX", Name: "纳西族", Numeric: "27"},
	{Roman: "Nu", Alpha: "NU", Name: "怒族", Numeric: "42"},
	{Roman: "Oroqen", Alpha: "OR", Name: "鄂伦春族", Numeric: "52"},
	{Roman: "Pumi", Alpha: "PM", Name: "普米族", Numeric: "40"},
	{Roman: "Qiang", Alpha: "QI", Name: "羌族", Numeric: "33"},
	{Roman: "Russ", Alpha: "RS", Name: "俄罗斯族", Numeric: "44"},
	{Roman: "Salar", Alpha: "SL", Name: "撒拉族", Numeric: "35"},
	{Roman: "She", Alpha: "SH", Name: "畲族", Numeric: "22"},
	{Roman: "Sui", Alpha: "SU", Name: "水族", Numeric: "25"},
	{Roman: "Tajik", Alpha: "TA", Name: "塔吉克族", Numeric: "41"},
	{Roman: "Tatar", Alpha: "TT", Name: "塔塔尔族", Numeric: "50"},
	{Roman: "Tu", Alpha: "TU", Name: "土族", Numeric: "30"},
	{Roman: "Tujia", Alpha: "TJ", Name: "土家族", Numeric: "15"},
	{Roman: "Uygur", Alpha: "UG", Name: "维吾尔族", Numeric: "05"},
	{Roman: "Uzbek", Alpha: "UZ", Name: "乌孜别克族", Numeric: "43"},
	{Roman: "Va", Alpha: "VA", Name: "佤族", Numeric: "21"},
	{Roman: "Xibe", Alpha: "XB", Name: "锡伯族", Numeric: "38"},
	{Roman: "Yao", Alpha: "YA", Name: "瑶族", Numeric: "13"},
	{Roman: "Yi", Alpha: "YI", Name: "彝族", Numeric: "07"},
	{Roman: "Yugur", Alpha: "YG", Name: "裕固族", Numeric: "48"},
	{Roman: "Zang", Alpha: "ZA", Name: "藏族", Numeric: "04"},
	{Roman: "Zhuang", Alpha: "ZH", Name: "壮族", Numeric: "08"},
}

// Nation 民族信息
type Nation struct {
	Roman   string // 罗马字母
	Alpha   string // 字母表示
	Name    string // 民族名称
	Numeric string // 数字表示
}

func (n *Nation) from(nation *Nation) *Nation {
	n.Roman = nation.Roman
	n.Alpha = nation.Alpha
	n.Name = nation.Name
	n.Numeric = nation.Numeric
	return n
}

// All 遍历所有民族信息
func All() iter.Seq[*Nation] {
	return func(yield func(*Nation) bool) {
		n := &Nation{}
		for _, nation := range nations {
			if !yield(n.from(nation)) {
				return
			}
		}
	}
}

// Get 通过 Roman、Alpha、Name、Numeric 等任意字段获取民族信息
func Get(v string) *Nation {
	if GetByRoman(v) != nil {
		return GetByRoman(v)
	}

	if GetByAlpha(v) != nil {
		return GetByAlpha(v)
	}

	if GetByName(v) != nil {
		return GetByName(v)
	}

	return GetByNumeric(v)
}

// GetByRoman 通过罗马字母获取民族信息
func GetByRoman(v string) *Nation {
	return getValue(func(n *Nation) bool { return n.Roman == v })
}

// GetByAlpha 通过字母表示获取民族信息
func GetByAlpha(v string) *Nation {
	return getValue(func(n *Nation) bool { return n.Alpha == v })
}

// GetByName 通过民族名称获取民族信息
func GetByName(v string) *Nation {
	return getValue(func(n *Nation) bool { return n.Name == v })
}

// GetByNumeric 通过数字表示获取民族信息
func GetByNumeric(v string) *Nation {
	return getValue(func(n *Nation) bool { return n.Numeric == v })
}

func getValue(comp func(*Nation) bool) *Nation {
	if index := slices.IndexFunc(nations, comp); index >= 0 {
		return nations[index]
	}
	return nil
}

// IsValid 判断 v 是否是一个有效的民族信息字段
func IsValid(v string) bool {
	return IsValidRoman(v) || IsValidAlpha(v) || IsValidName(v) || IsValidNumeric(v)
}

func IsValidRoman(v string) bool {
	return isExists(func(n *Nation) bool { return n.Roman == v })
}

func IsValidAlpha(v string) bool {
	return isExists(func(n *Nation) bool { return n.Alpha == v })
}

func IsValidName(v string) bool {
	return isExists(func(n *Nation) bool { return n.Name == v })
}

func IsValidNumeric(v string) bool {
	return isExists(func(n *Nation) bool { return n.Numeric == v })
}

func isExists(comp func(*Nation) bool) bool {
	return slices.IndexFunc(nations, comp) >= 0
}

package word

import (
	"strings"
	"unicode"
)

// ToUpper all to upper
func ToUpper(s string) string {
	return strings.ToUpper(strings.Replace(s, " ", "_", -1))
}

// ToLower all to lower
func ToLower(s string) string {
	return strings.ToLower(strings.Replace(s, " ", "_", -1))
}

// UnderlineToUpperCamelCase underline to upper camelCase
func UnderlineToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)   // _ 转空格
	s = strings.Title(s)                   // 首字母大写
	return strings.Replace(s, " ", "", -1) // 空格转空
}

// UnderlineToLowerCamelCase underline to lower camelCase
func UnderlineToLowerCamelCase(s string) string {
	s = UnderlineToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderline camelCase to underline
func CamelCaseToUnderline(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) { // 第一个大写字母转_
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

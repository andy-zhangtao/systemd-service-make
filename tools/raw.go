package tools

import (
	"fmt"
	"strings"
)

// ParseRawArgs 将原生Docker Run解析为数组类型数据
func ParseRawArgs(raw string) []string {
	_rawArray := strings.Split(raw, " ")
	var _raw []string

	for i, r := range _rawArray {
		r = strings.TrimSpace(r)
		if strings.HasPrefix(r, "-") {
			_raw = append(_raw, fmt.Sprintf("%s %s", r, _rawArray[i+1]))
		} else {
			continue
		}
	}

	return _raw
}

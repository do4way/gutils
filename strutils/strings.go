package strutils

import "bytes"

func Join(strs []string, sep string) string {
	var buf bytes.Buffer
	lastIdx := len(strs) - 1
	for i, str := range strs {
		if str == "" {
			buf.WriteString("\"\"")
		} else {
			buf.WriteString(str)
		}
		if i != lastIdx {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}

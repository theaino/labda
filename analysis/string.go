package analysis

import (
	"strconv"
)

func LexString(body string) (string, int) {
	quote := body[0]
	escapeCount := 0
	for idx, char := range []byte(body) {
		if idx == 0 { continue }
		switch char {
		case '\\':
			escapeCount += 1
		case quote:
			if escapeCount % 2 == 0 {
				quoted := body[:idx+1]
				unquoted, err := strconv.Unquote(quoted)
				if err != nil {
					panic(err)
				}
				return unquoted, idx
			}
		default:
			escapeCount = 0
		}
	}
	panic("String not closed")
}

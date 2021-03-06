package leecode

import (
	"fmt"
	"testing"
)

func TestLongestSubStrWithoutRepeatChar(t *testing.T) {
	testData := []struct {
		Str    string
		LS     string
		Length int
	}{
		{
			Str:    "abcabcdd",
			LS:     "abcd",
			Length: 4,
		},
		{
			Str:    "",
			LS:     "",
			Length: 0,
		},
		{
			Str:    "pwwkeed",
			LS:     "wke",
			Length: 3,
		},
		{
			Str:    "dvdf",
			LS:     "vdf",
			Length: 3,
		},
	}
	for _, datum := range testData {
		length := lengthOfLongestSubstring(datum.Str)
		if length != datum.Length {
			t.Error("not right", length, datum.Str, datum.LS)
		}
	}
}

// Runtime: 520 ms, faster than 5.04% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 8.1 MB, less than 6.10% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(str string) int {
	substr := ""
	tmpSubstr := ""
	start := 0
	curIdx := 0
	curChars := make(map[string]interface{})
	// 1. 检查重复char
	// 2. 如果发现重复char，需考虑重复后起始位置
	for curIdx < len(str) {
		curChar := string(str[curIdx])
		if _, ok := curChars[curChar]; ok {
			curChars = make(map[string]interface{})
			for j := 0; j < len(tmpSubstr); j++ {
				if string(tmpSubstr[j]) == curChar {
					start += j
				}
			}
			start++
			curIdx = start
			curChars[string(str[curIdx])] = nil
			tmpSubstr = str[start : curIdx+1]
			curIdx++
		} else {
			curChars[curChar] = nil
			tmpSubstr = str[start : curIdx+1]
			curIdx++
		}

		if len(tmpSubstr) > len(substr) {
			substr = tmpSubstr
		}
		fmt.Println(curChars, string(curChar), start, curIdx, tmpSubstr, substr)

	}
	return len(substr)
}

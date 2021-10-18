package leecode

import (
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//执行用时：8 ms, 在所有 Go 提交中击败了70.39%的用户
//内存消耗：2.9 MB, 在所有 Go 提交中击败了55.02%的用户
func lengthOfLongestSubstring(str string) int {
	var chars = make(map[byte]int, 26)
	longestSubStringLength := 0
	start := 0
	byteList := []byte(str)
	for i := 0; i < len(byteList); i++ {
		curByte := byteList[i]
		lastCharIdx, ok := chars[curByte]
		if !ok {
			longestSubStringLength = max(longestSubStringLength, i-start+1)
			chars[curByte] = i
		} else {
			// 已经存在，需要重置start到重复字符串索引的后一位
			for j := start; j <= lastCharIdx; j++ {
				delete(chars, byteList[j])
			}
			start = lastCharIdx + 1
			chars[curByte] = i
		}
	}
	return longestSubStringLength
}

// Runtime: 520 ms, faster than 5.04% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 8.1 MB, less than 6.10% of Go online submissions for Longest Substring Without Repeating Characters.
//func lengthOfLongestSubstring(str string) int {
//	substr := ""
//	tmpSubstr := ""
//	start := 0
//	curIdx := 0
//	curChars := make(map[string]interface{})
//	// 1. 检查重复char
//	// 2. 如果发现重复char，需考虑重复后起始位置
//	for curIdx < len(str) {
//		curChar := string(str[curIdx])
//		if _, ok := curChars[curChar]; ok {
//			curChars = make(map[string]interface{})
//			for j := 0; j < len(tmpSubstr); j++ {
//				if string(tmpSubstr[j]) == curChar {
//					start += j
//				}
//			}
//			start++
//			curIdx = start
//			curChars[string(str[curIdx])] = nil
//			tmpSubstr = str[start : curIdx+1]
//			curIdx++
//		} else {
//			curChars[curChar] = nil
//			tmpSubstr = str[start : curIdx+1]
//			curIdx++
//		}
//
//		if len(tmpSubstr) > len(substr) {
//			substr = tmpSubstr
//		}
//		fmt.Println(curChars, string(curChar), start, curIdx, tmpSubstr, substr)
//
//	}
//	return len(substr)
//}

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	var tmp string
	for _, runee := range s {
		if string(runee) == " " {
			addWord(tmp, ret)
			tmp = ""
		} else {
			tmp += string(runee)
		}
	}
	addWord(tmp, ret)
	return ret
}

func addWord(tmp string, ret map[string]int) {
	if len(tmp) > 0 {
		word := tmp
		if cnt, ok := ret[word]; !ok {
			ret[word] = 1
		} else {
			ret[word] = cnt + 1
		}
	}
}

func WordCount2(s string) map[string]int {
	ret := make(map[string]int)
	splits := strings.Split(s, " ")
	for _, str := range splits {
		addWord(str, ret)
	}
	return ret
}

func main() {
	wc.Test(WordCount2)
}

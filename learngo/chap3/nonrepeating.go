package main

import "fmt"

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	/**
	先遍历所有的字符串，然后新建一个map[key] ->每个字符串 value是遍历的下标
	如果key不存在 lastI, ok := lastOccurred[ch] ok得到的是false
	i: 1-0+1 > 0 maxLength
	   2-0+1 > 2 maxLength
	*/

	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI > start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		} else {
			fmt.Println(string(ch))
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	ret := lengthOfNonRepeatingSubstr("hellobp")
	fmt.Println(ret)
}

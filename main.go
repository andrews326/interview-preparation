// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	str := "pwwkew"
	l := lengthOfSubStr(str)
	fmt.Println(l)
	
}

func lengthOfSubStr(str string) int {
	left, res := 0, 0

	charSet := make(map[byte]bool)
	r := 0
	for right := 0; right < len(str); right++ {
		r = right
		for charSet[str[right]] {
			delete(charSet, str[left])
			left++
		}

		charSet[str[right]] = true

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	for key := range charSet {
		fmt.Printf("%c =  \n", key)
	}

	fmt.Println("Longest Substring:", str[left:r+1])
	return res
}

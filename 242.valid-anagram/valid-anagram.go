package main

import "fmt"

func main() {
	s := "car"
	t := "rac"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	smap := make(map[int32]int)
	for _, sc := range s {
		if _, ok := smap[sc]; !ok {
			smap[sc] = 0
		}
		smap[sc]++
	}

	for _, tc := range t {
		if _, ok := smap[tc]; !ok {
			return false
		}
		smap[tc]--
	}

	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true
}

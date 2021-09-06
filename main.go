package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	groups := assembleAnagrams(strs)

	for _, group := range groups {
		fmt.Println(group)
	}
}

func assembleAnagrams(strs []string) (anagramGroups [][]string) {
	anagramMap := map[string][]string{}

	for _, str := range strs {
		insertToGroup(anagramMap, str)
	}

	for _, v := range anagramMap {
		anagramGroups = append(anagramGroups, v)
	}

	return anagramGroups
}

func insertToGroup(m map[string][]string, str string) {
	sortedStr := sortString(str)

	if group, ok := m[sortedStr]; ok {
		// put into the existing group
		group = append(group, str)
		m[sortedStr] = group
	} else {
		// create one if none exist
		m[sortedStr] = []string{str}
	}
}

func sortString(str string) string {
	slc := []rune{}

	for _, c := range str {
		slc = append(slc, c)
	}

	sort.SliceStable(slc, func(i, j int) bool { return slc[i] < slc[j] })
	return string(slc)
}

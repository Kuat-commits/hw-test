package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type keyValue struct {
	Key   string
	Value int
}

func Top10(str string) []string {
	var st []string
	var pr string
	cach := make(map[string]int)
	valpr := 0
	sortedStruct := make([]keyValue, 0)

	if str == "" {
		return st
	}
	s := strings.Fields(str)

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	for _, val := range s {
		if pr == val {
			cach[val]++
		} else {
			cach[val] = 1
		}
		pr = val
	}

	for _, val := range cach {
		if val > valpr {
			valpr = val
		}
	}

	for key, value := range cach {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		if sortedStruct[i].Value == sortedStruct[j].Value {
			return sortedStruct[i].Key < sortedStruct[j].Key
		}
		return sortedStruct[i].Value > sortedStruct[j].Value
	})

	if len(sortedStruct) < 10 {
		for i := 0; i < len(sortedStruct); i++ {
			st = append(st, sortedStruct[i].Key)
		}
		return st
	}

	for i := 0; i < 10; i++ {
		st = append(st, sortedStruct[i].Key)
	}
	return st
}

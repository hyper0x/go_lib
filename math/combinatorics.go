package go_lib

import (
	"fmt"
	"go_lib/collection"
	"sort"
)

func NextItem(current []int, n int) ([]int, bool) {
	r := len(current)
	next := make([]int, r)
	copy(next, current)
	ni := n - 1
	ri := r - 1
	carryPos := -1
	end := false
	for i := 0; i <= ri; i++ {
		if next[i] == ni {
			if i == ri {
				end = true
				break
			}
			carryPos = i + 1
			continue
		} else {
			if carryPos == i {
				for j := (i - 1); j >= 0; j-- {
					next[j] = 0
				}
			}
			next[i] = next[i] + 1
			break
		}
	}
	return next, end
}

func NextPermutation(current []int, n int, repeatable bool) ([]int, bool) {
	var next []int
	var end bool
	next, end = NextItem(current, n)
	if end {
		return nil, true
	}
	needNext := false
	if !repeatable {
		hasSame := false
		for i, v := range next {
			for j, w := range next {
				if v == w && i != j {
					hasSame = true
					break
				}
			}
			if hasSame {
				break
			}
		}
		needNext = hasSame
	}
	if !needNext {
		return next, end
	}
	return NextPermutation(next, n, repeatable)
}

func GetPermutationCount(n int, r int, repeatable bool) int {
	count := 0
	if repeatable {
		count = 1
		for i := 0; i < r; i++ {
			count *= n
		}
	} else {
		count = 1
		limit := n - r + 1
		for i := n; i >= limit; i-- {
			count *= i
		}
	}
	return count
}

func GetPermutations(origin []interface{}, r int, repeatable bool) [][]interface{} {
	n := len(origin)
	if n <= 0 || n <= 0 || r <= 0 || r > n {
		return nil
	}
	indexs := make([][]int, 0)
	current := make([]int, r)
	var end bool
	for {
		current, end = NextPermutation(current, n, repeatable)
		if end {
			break
		}
		indexs = append(indexs, current)
	}
	permutations := make([][]interface{}, len(indexs))
	for i, v := range indexs {
		item := make([]interface{}, r)
		for j, w := range v {
			item[j] = origin[w]
		}
		permutations[i] = item
	}
	return permutations
}

func GetCombinationCount(n int, r int, repeatable bool) int {
	count := 0
	if repeatable {
		dividend := 1
		limit := n + r - 1
		for i := 1; i <= limit; i++ {
			dividend *= i
		}
		divider := n - 1
		for i := 1; i <= r; i++ {
			divider *= i
		}
		count = dividend / divider
	} else {
		dividend := GetPermutationCount(n, r, repeatable)
		divider := 1
		for i := 1; i <= r; i++ {
			divider *= i
		}
		count = dividend / divider
	}
	return count
}

func GetCombinations(origin []interface{}, r int, repeatable bool) [][]interface{} {
	n := len(origin)
	if n <= 0 || n <= 0 || r <= 0 || r > n {
		return nil
	}
	indexSet := collection.SimpleSet{KeyGenerator: generateIntArrayKey, Comparator: compareIntArray}
	current := make([]int, r)
	var end bool
	for {
		current, end = NextPermutation(current, n, repeatable)
		if end {
			break
		}
		indexSet.Add(current)
	}
	indexSetLength := indexSet.Len()
	combinations := make([][]interface{}, indexSetLength)
	indexIterator := indexSet.Iterator()
	for i := 0; i < indexSetLength; i++ {
		v, has := indexIterator()
		if !has {
			break
		}
		ia := v.([]int)
		item := make([]interface{}, r)
		for j, w := range ia {
			item[j] = origin[w]
		}
		combinations[i] = item
	}
	return combinations
}

func generateIntArrayKey(x interface{}) string {
	xa := interface{}(x).([]int)
	xac := make([]int, len(xa))
	copy(xac, xa)
	sort.Ints(xac)
	return fmt.Sprintf("%v", xac)
}

func compareIntArray(i, j interface{}) int {
	ia := interface{}(i).([]int)
	ja := interface{}(j).([]int)
	sort.Ints(ia)
	sort.Ints(ja)
	il := len(ia)
	jl := len(ja)
	result := 0
	if il < jl {
		result = -1
	} else if il > jl {
		result = 1
	} else {
		for i, iv := range ia {
			jv := ja[i]
			if iv != jv {
				if iv < jv {
					result = -1
				} else if iv > jv {
					result = 1
				}
				break
			}
		}
	}
	return result
}

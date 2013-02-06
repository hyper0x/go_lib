package go_lib

import (
	"math/rand"
	"runtime/debug"
	"testing"
)

func toInt(a []int, n int) uint64 {
	r := len(a)
	var result uint64
	var temp uint64
	scale := uint64(n)
	for i := 1; i <= r; i++ {
		temp = uint64(a[i-1])
		for j := 1; j < i; j++ {
			temp *= scale
		}
		result += temp
	}
	return result
}

func TestNextItem(t *testing.T) {
	debugTag := false
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error: %s\n", err)
		}
	}()
	parameterMap := make(map[int]int, 3)
	parameterMap[5] = 3
	parameterMap[8] = 6
	parameterMap[12] = 5
	for n, r := range parameterMap {
		t.Log("Testing function 'NextItem'...")
		t.Logf("N: %d, R: %d", n, r)
		initial := make([]int, r)
		t.Logf("Initial: %v\n", initial)
		currentValue := uint64(0)
		previousValue := uint64(0)
		max := make([]int, r)
		for i := 0; i < len(max); i++ {
			max[i] = n - 1
		}
		t.Logf("Max: %v\n", max)
		maxValue := toInt(max, n)
		count := 0
		expectedCount := GetPermutationCount(n, r, true)
		t.Logf("Expected Count: %d\n", expectedCount)
		var end bool
		current := initial
		currentValue = toInt(current, n)
		count++
		if debugTag {
			t.Logf("%d: %v (%d)\n", count, current, currentValue)
		}
		previousValue = currentValue
		for {
			current, end = NextItem(current, n)
			if end {
				break
			}
			currentValue = toInt(current, n)
			if currentValue < 0 {
				t.Errorf("Error: %d < 0 (count=%d)\n", currentValue, count)
				t.FailNow()
			}
			if currentValue > maxValue {
				t.Errorf("Error: %d > %d (count=%d)\n", currentValue, maxValue, count)
				t.FailNow()
			}
			if previousValue >= 0 {
				if currentValue <= previousValue {
					t.Errorf("Error: %d <= %d (count=%d)\n", currentValue, previousValue, count)
					t.FailNow()
				}
			}
			count++
			if debugTag {
				t.Logf("%d: %v (%d)\n", count, current, currentValue)
			}
			previousValue = currentValue
		}
		if count != expectedCount {
			t.Errorf("Error: The count should be %d but %d.\n", expectedCount, count)
			t.FailNow()
		}
		t.Logf("Total: %d\n", count)
	}
}

func TestNextPermutation(t *testing.T) {
	debugTag := false
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error: %s\n", err)
		}
	}()
	parameterMap := make(map[int]int, 3)
	parameterMap[5] = 3
	parameterMap[8] = 6
	parameterMap[12] = 5
	repeatable := false
	for n, r := range parameterMap {
		t.Log("Testing function 'NextPermutation'...")
		t.Logf("N: %d, R: %d, Repeatable: %v", n, r, repeatable)
		initial := make([]int, r)
		t.Logf("Initial: %v\n", initial)
		currentValue := uint64(0)
		previousValue := uint64(0)
		max := make([]int, r)
		for i := 0; i < len(max); i++ {
			max[i] = n - 1
		}
		t.Logf("Max: %v\n", max)
		maxValue := toInt(max, n)
		count := 0
		expectedCount := GetPermutationCount(n, r, repeatable)
		t.Logf("Expected Count: %d\n", expectedCount)
		var end bool
		current := initial
		currentValue = toInt(current, n)
		if debugTag {
			t.Logf("%d: %v (%d)\n", count, current, currentValue)
		}
		previousValue = currentValue
		for {
			current, end = NextPermutation(current, n, repeatable)
			if end {
				break
			}
			currentValue = toInt(current, n)
			if currentValue < 0 {
				t.Errorf("Error: %d < 0 (count=%d)\n", currentValue, count)
				t.FailNow()
			}
			if currentValue > maxValue {
				t.Errorf("Error: %d > %d (count=%d)\n", currentValue, maxValue, count)
				t.FailNow()
			}
			hasSame := false
			for i, v := range current {
				for j, w := range current {
					if v == w && i != j {
						hasSame = true
						break
					}
				}
				if hasSame {
					t.Errorf("Error: %v has same numbers (count=%d)\n", current, count)
					t.FailNow()
				}
			}
			if previousValue >= 0 {
				if currentValue <= previousValue {
					t.Errorf("Error: %d <= %d (count=%d)\n", currentValue, previousValue, count)
					t.FailNow()
				}
			}
			count++
			if debugTag {
				t.Logf("%d: %v (%d)\n", count, current, currentValue)
			}
			previousValue = currentValue
		}
		if count != expectedCount {
			t.Errorf("Error: The count should be %d but %d.\n", expectedCount, count)
			t.FailNow()
		}
		t.Logf("Total: %d\n", count)
	}
}

func TestGetPermutations(t *testing.T) {
	debugTag := false
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error: %s\n", err)
		}
	}()
	parameterMap := make(map[int]int, 3)
	parameterMap[5] = 3
	parameterMap[8] = 6
	parameterMap[12] = 5
	repeatable := false
	for n, r := range parameterMap {
		t.Log("Testing function 'GetPermutations'...")
		t.Logf("N: %d, R: %d, Repeatable: %v", n, r, repeatable)
		origin := make([]interface{}, n)
		for i := 0; i < n; i++ {
			origin[i] = rand.Intn(30)
		}
		t.Logf("Origin: %v\n", origin)
		expectedCount := GetPermutationCount(n, r, repeatable)
		t.Logf("Expected Count: %d\n", expectedCount)
		result := GetPermutations(origin, r, repeatable)
		count := len(result)
		if count != expectedCount {
			t.Errorf("Error: The count should be %d but %d.\n", expectedCount, count)
			t.FailNow()
		}
		if debugTag {
			for i, v := range result {
				t.Logf("%d: %v\n", (i + 1), v)
			}
		}
		t.Logf("Total: %d\n", count)
	}
}

func TestGetCombinations(t *testing.T) {
	debugTag := false
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error: %s\n", err)
		}
	}()
	parameterMap := make(map[int]int, 3)
	parameterMap[5] = 3
	parameterMap[8] = 6
	parameterMap[12] = 5
	repeatable := false
	for n, r := range parameterMap {
		t.Log("Testing function 'GetCombinations'...")
		t.Logf("N: %d, R: %d, Repeatable: %v", n, r, repeatable)
		origin := make([]interface{}, n)
		for i := 0; i < n; i++ {
			origin[i] = rand.Intn(30)
		}
		t.Logf("Origin: %v\n", origin)
		expectedCount := GetCombinationCount(n, r, false)
		t.Logf("Expected Count: %d\n", expectedCount)
		result := GetCombinations(origin, r, false)
		count := len(result)
		if count != expectedCount {
			t.Errorf("Error: The count should be %d but %d.\n", expectedCount, count)
			t.FailNow()
		}
		if debugTag {
			for i, v := range result {
				t.Logf("%d: %v\n", (i + 1), v)
			}
		}
		t.Logf("Total: %d\n", count)
	}
}

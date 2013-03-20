package versionstring

import (
	"strconv"
	"strings"
)

func ParseVersionString(vs string) []int {
	tmp := strings.Split(vs, ".")
	arr := make([]int, 0, len(tmp))

	for _, val := range tmp {
		n, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		arr = append(arr, n)
	}

	return arr
}

// 0 = equal, 1 = greater, -1 = smaller
// for now: can only compare strings of equal length
func CompareStrings(lvs, rvs string) int {
	left := ParseVersionString(lvs)
	right := ParseVersionString(rvs)

	// compare the version strings as long as is possible
	min := min(len(left), len(right))
	for i := 0; i < min; i++ {
		if left[i] == right[i] {
			continue
		}

		if left[i] < right[i] {
			return 1
		}
		return -1
	}
	// both versions are equal as far as we know
	// - if they are of equal length => 0
	if len(left) == len(right) {
		return 0
	}
	// Worst case
	// - version string are of unequal length

	// the longest of the version strings has to
	// have just one value != 0 to win
	if len(left) < len(right) {
		for i := min; i < len(right); i++ {
			if right[i] != 0 {
				return 1
			}
		}
	} else {
		for i := min; i < len(left); i++ {
			if left[i] != 0 {
				return -1
			}
		}
	}

	// alright, they're equal.
	return 0
}

func max(n, m int) int {
	if n < m {
		return m
	}
	// n >= m
	return n
}

func min(n, m int) int {
	if n < m {
		return n
	}
	// m <= n
	return m
}

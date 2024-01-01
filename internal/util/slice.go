package util

import "slices"

func GetSetDiff[T ~[]E, E comparable](slice1 T, slice2 T) T {
	diff := make([]E, 0)
	for _, elem := range slice1 {
		if !slices.Contains(slice2, elem) {
			diff = append(diff, elem)
		}
	}

	for _, elem := range slice2 {
		if !slices.Contains(slice1, elem) {
			diff = append(diff, elem)
		}
	}

	return diff
}

func MakeUnique[T ~[]E, E comparable](s T) T {
	seen := make(map[E]bool, len(s))
	uniqueSlice := make([]E, 0)
	for _, elem := range s {
		if !seen[elem] {
			seen[elem] = true
			uniqueSlice = append(uniqueSlice, elem)
		}
	}

	return uniqueSlice
}

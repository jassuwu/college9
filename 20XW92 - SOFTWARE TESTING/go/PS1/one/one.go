package one

import (
	"slices"
)

func FindMin(nums []int) int {
	return slices.Min(nums)
}

func FindMax(nums []int) int {
	return slices.Max(nums)
}

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums) - 1
	res := make([][]int, 0)
	for i := 0; i < n-1; i += 1 {
		j := i + 1
		for j < n {
			if nums[j] == nums[j+1] {
				if nums[i] == nums[j] && nums[i] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[n]})
				}
				j++
				continue
			}
			if nums[i]+nums[j]+nums[n] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[n] > 0 {
				n--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[n]})
				j++
				n--
			}
		}
	}
	return res
}

func main() {
	nums := []int{0, 0, 0}
	res := threeSum(nums)
	fmt.Println(res)
}

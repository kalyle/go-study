package main

import (
	"fmt"
	"sort"
)

// var path = []int64{}
// var res = []int64{}

// func maximumEvenSplit(finalSum int64) []int64 {
// 	if finalSum%2 == 1 {
// 		return []int64{}
// 	}
// 	path = append(path, 2)
// 	backtracing(finalSum, 4)
// 	return res
// }

// // var path = make([]int64, 0)
// // var res = make([]int64, 0)

// func backtracing(finalSum int64, index int64) {

// 	if sumVal := sum(path); sumVal == finalSum {

// 		res = path[:len(path)-1]

// 		return
// 	}
// 	for ; index < finalSum; index += 2 {
// 		path = append(path, index)

// 		if temp := sum(path); temp >= finalSum {
// 			break
// 		} else {
// 			backtracing(finalSum, index+2)
// 			path = path[:len(path)-1]
// 		}
// 	}
// }
// func sum(slice []int64) int64 {
// 	var sum int64
// 	for _, val := range slice {
// 		sum += val
// 	}
// 	return sum
// }
// func maximumEvenSplit(finalSum int64) []int64 {
// 	if finalSum%2 == 1 {
// 		return []int64{}
// 	}

// 	backtracing(finalSum, 2)
// 	return res
// }

// var path = make([]int64, 0)

// var res = make([]int64, 0)

// func backtracing(finalSum int64, index int64) {
// 	fmt.Println(&path, index)

// 	if sumVal := sum(path); sumVal == finalSum {
// 		if len(path) > len(res) {
// 			res = path
// 		}

// 		return
// 	}
// 	for ; index <= finalSum; index += 2 {

//			if temp := sum(path); temp > finalSum {
//				break
//			} else {
//				path = append(path, index)
//				backtracing(finalSum, index+2)
//				path = path[:len(path)-1]
//			}
//		}
//	}
//
//	func sum(slice []int64) int64 {
//		var sum int64
//		for _, val := range slice {
//			sum += val
//		}
//		return sum
//	}
func maximumEvenSplit(finalSum int64) []int64 {
	// 贪心
	if finalSum%2 == 1 {
		return []int64{}
	}
	q := make([]int64, 0)
	var i int64 = 2 // syntax error: var declaration not allowed in for initializer
	for ; finalSum >= 0; i += 2 {
		q = append(q, i)
		finalSum -= i
		if finalSum == 0 {
			return q
		} else if i >= finalSum {
			finalSum += q[0]
			q = q[1:]
		}
	}
	return []int64{}
}

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	preSum := make([]int, len(nums))
	sufSum := make([]int, len(nums))
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			preSum[i] = 0
		} else if i == 1 {
			preSum[i] = nums[i-1]
		} else {
			preSum[i] = preSum[i-1] * nums[i-1]
		}
	}
	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			sufSum[j] = 0
		} else if j == len(nums)-2 {
			sufSum[j] = nums[j+1]
		} else {
			sufSum[j] = sufSum[j+1] * nums[j+1]
		}
	}
	for k := 0; k < len(nums); k++ {
		if k == 0 {
			ans[k] = sufSum[k]
		} else if k == len(nums)-1 {
			ans[k] = preSum[k]
		} else {
			ans[k] = preSum[k] * sufSum[k]
		}
	}
	return ans
}

// func firstMissingPositive(nums []int) int {
// 	sort.Ints(nums)
// 	set := make([]int, 0)
// 	for index, _ := range nums {
// 		if index < len(nums)-1 && nums[index] == nums[index+1] {
// 			continue
// 		}
// 		set = append(set, nums[index])
// 	}
// 	postIndex := 0
// 	ans := 1
// 	for index, _ := range set {

// 		if (index == 0 && set[index] > 0) || set[index] > 0 && set[index-1] <= 0 {
// 			postIndex = index
// 		}
// 		if index < len(set)-1 && set[index] > 0 && set[index+1]-set[index] != 1 {
// 			ans = set[index] + 1
// 			break
// 		}

//		}
//		if set[postIndex] != 1 {
//			ans = 1
//			return ans
//		} else if ans == 1 {
//			return set[len(set)-1] + 1
//		} else {
//			return ans
//		}
//	}
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	set := make([]int, nums[len(nums)-1]+1)
	for _, val := range nums {

		set[val] = 1

	}
	for index, val := range set {
		if val != 1 && index != 0 {
			return index
		}
	}
	return len(set)
}
func main() {
	// var nums []int = [1,2,3,4]
	// array := [...]int{-1, 1, 0, -3, 3}
	// var nums []int = array[:]
	// ans := productExceptSelf(nums)
	// fmt.Println("ans =", ans)
	array := [...]int{1, 2, 0}
	var nums []int = array[:]
	fmt.Println("ans = ", firstMissingPositive(nums))
}

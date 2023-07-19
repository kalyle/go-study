package main

import "fmt"

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

// 		if temp := sum(path); temp > finalSum {
// 			break
// 		} else {
// 			path = append(path, index)
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

func main() {
	var finalSum int64 = 28
	val := maximumEvenSplit(finalSum)
	fmt.Println(&val)
}

package _case

import "fmt"

// 常见泛型类型 --- 函数泛型
func getSumOfNum(a, b int) int {
	var c int = a + b
	return c
}

func getSumOfStr(a, b string) string {
	var c string = a + b
	return c
}

func getSum[T int | string | float32 | float64](a, b T) T {
	var c T = a + b
	return c
}

func FuncCase() {
	a1, b1 := 1, 2
	c1 := getSumOfNum(a1, b1)
	fmt.Println("getSumOfNum c1 =", c1)

	var a2, b2 string = "abc", "efg"
	c2 := getSumOfStr(a2, b2)
	fmt.Println("getSumOfStr c2 =", c2)

	a3, b3 := 1.2, 2.3
	c3 := getSum[float64](a3, b3)
	fmt.Println("generic getSum c3 =", c3)

}

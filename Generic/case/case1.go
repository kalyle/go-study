package _case

import "fmt"

func SimpleCase() {
	a1, b1 := 1, 2
	a2, b2 := "abc", "def"

	c1 := getSumOfNum(a1, b1)
	fmt.Println("getSumOfNum c1 =", c1)
	c2 := getSumOfStr(a2, b2)
	fmt.Println("getSumOfStr c2 =", c2)

	// 由编译器自动推断输入类型
	c3 := getSum(a1, b1)
	fmt.Println("getSumOfNum c3 =", c3)
	// 显示指定输入类型（使用）
	c4 := getSum[string](a2, b2)
	fmt.Println("getSumOfNum c4 =", c4)
	m := getMax[int](a1, b1)
	fmt.Println("getSumOfNum max =", m)

}

func getSumOfNum(a, b int) int {
	var c int = a + b
	return c
}

func getSumOfStr(a, b string) string {
	var c string = a + b
	return c
}

func getSum[T int | string](a, b T) T {
	c := a + b
	return c
}

func getMax[T NumT](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

//
type NumT interface {
	// interface接受类型集，方法集（常见）
	// interface接受类型集 用于支持泛型，如下所示
	int8 | int16 | int32 | int64 | float32 | float64 | ~int
	int16 | int32 | int64 | float32 | float64 | ~int
	// ~ 指支持类型的 *衍生类型*（使用type自定义类型）
	// | 指取交集
	// 多行取交集
}

// eg:
// 衍生类型,MyInt 为int 衍生类型，两者为不同类型
type MyInt int

// 别名 ，MyStr为 string别名，两者为同一类型
type MyStr = string

// go 内置两个泛型类型 any \ comparable
// comparable 只支持 能使用==  ！=两个操作的类型，eg:int,float,string

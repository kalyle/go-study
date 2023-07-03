package _case

import "fmt"

// 类型泛型
type user struct {
	ID   int64
	Name string
	Age  int
}
type address struct {
	ID       int
	Province string
	City     string
}

func Map2List[k comparable, v any](mp map[k]v) []v {
	list := make([]v, len(mp))
	var i int
	for _, val := range mp {
		list[i] = val
		i++
	}
	return list
}

// 定义类型泛型
type Map[k comparable, v any] map[k]v //泛型map定义，声明两个泛型k,v
type List[T any] []T                  //泛型切片定义

func TSimpleCase() {
	userMp := make(map[int64]user, 0)
	// fmt.Printf("userMp初始化地址 = %p\n", &userMp)
	userMp[0] = user{ID: 1, Name: "Tom", Age: 18}
	userMp[1] = user{ID: 2, Name: "Li", Age: 19}
	// fmt.Printf("userMp赋值后地址 = %p\n", &userMp)

	userList := Map2List[int64, user](userMp)
	fmt.Println("Map2List userList = ", userList)

	addressMp := make(map[int64]address, 0)
	addressMp[0] = address{ID: 1, Province: "四川", City: "成都"}
	addressMp[1] = address{ID: 2, Province: "浙江", City: "杭州"}
	addressList := Map2List[int64, address](addressMp)
	fmt.Println("Map2List addressList = ", addressList)

}
func TypeCase() {
	// 使用泛型
	fmt.Println("使用泛型")
	userMp := make(Map[int, user], 0)
	userMp[0] = user{ID: 1, Name: "Tom", Age: 18}
	userMp[1] = user{ID: 2, Name: "Li", Age: 19}

	var userList List[user] = Map2List[int, user](userMp)
	fmt.Println("Map2List userList = ", userList)

	addressMp := make(Map[int, address], 0)
	addressMp[0] = address{ID: 1, Province: "四川", City: "成都"}
	addressMp[1] = address{ID: 2, Province: "浙江", City: "杭州"}

	var addressList List[address] = Map2List[int, address](addressMp)
	fmt.Println("Map2List addressList = ", addressList)
}

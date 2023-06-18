package check_if_exist

import "fmt"

func CheckIfExist(arr []int) bool {
	m := make(map[int]int)
	exist := false
	for _, v := range arr {
		fmt.Printf("%v\n", m)
		if v%2 == 0 {
			_, exist = m[v/2]
			if exist {
				fmt.Printf("v=%d\n", v)
				break
			}
		}
		_, exist = m[v*2]
		if exist {
			fmt.Printf("v=%d\n", v)
			break
		}
		m[v] = v
	}
	fmt.Printf("result=%v\n", exist)
	return exist
}

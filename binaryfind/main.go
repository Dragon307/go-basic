// 二分查找
package main

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int)  {
	// 判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("不存在")
		return
	}
	// 先找到 中间的下标
	middle := (leftIndex+rightIndex)/2
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle -1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("Find it index is %v", middle)
	}
}

func main() {

	arr := [6]int{1,8, 10, 89, 1000, 1234}

	BinaryFind(&arr, 0, len(arr) - 1, 8)

}

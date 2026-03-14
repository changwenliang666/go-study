package main

import (
	"fmt"
	"math"
	"slices"
)

// 向切片头部插入数据
func appendToHead(newData []int, originData []int) []int {
	return append(newData, originData...)
}

// 向切片中间下标i 插入数据
func appendToMiddle(newData []int, originData []int, index int) []int {
	return append(originData[:index+1], append(newData, originData[index+1:]...)...)
}

// 删除切片头部n个元素
func deleteToHead(originData []int, n int) []int {
	return originData[n:]
}

// 删除切片尾部n个元素
func deleteToEnd(originData []int, n int) []int {
	return originData[:len(originData)-n]
}

// 从切片中间删除n个元素
func deleteToMiddle(originData []int, index int, n int) []int {
	spliceIndex := math.Min(float64(index+n), float64(len(originData)-1))
	spliceIndex2 := int(spliceIndex)

	return append(originData[:index], originData[spliceIndex2:]...)
}

func main() {
	var arr = [10]int{1, 2, 3, 4, 5}
	var slice = arr[:]                // 用这种方法会将数组类型转换成切片类型，但是两者指向同一块内存地址
	var slice2 = slices.Clone(arr[:]) // 用这种方式将数组转换位切片类型，不会共享内存地址
	slice[0] = 10
	slice2[0] = 100

	newData := []int{-2, -1}
	newData2 := []int{-3, -4}
	slice2 = appendToHead(newData, slice2)       // 向切片头部插入数据
	slice2 = appendToMiddle(newData2, slice2, 2) // 向切片中间插入数据
	slice2 = deleteToHead(slice2, 1)             // 从头部删除元素
	slice2 = deleteToEnd(slice2, 3)              // 从尾部删除元素
	slice2 = deleteToMiddle(slice2, 1, 4)
	fmt.Println(slice2)
}

/**
切片与数组的区别：
1.外观上数组需要指定长度，长度一旦制定便不可以变化，切片的长度可以动态变化
*/

package main

import (
	"fmt"
)

func init(){
	fmt.Println("首先执行了init函数")

	//make
}

func main(){
	fmt.Println("在执行了main函数")

	//测试数组的入参复制问题
	//aguArr()

	//数组切片
	arrSlice()

	//数组遍历
	arrFor()

	//动态添加数组
	arrAppend()

	//数组复制
	arrCopy()
}

func arrCopy(){
	arr := []int{1,2,3,4,5,6}
	arrB := []int{23,57,87,9}

	fmt.Println(arr,arrB)

	copy(arr,arrB)

	fmt.Println(arr)

	copy(arrB,arr)

	fmt.Println(arrB)
}



func arrAppend(){
	arr := []int{22,43,64,75,33}

	arrB := append(arr,48,56,78,32)

	fmt.Println(arrB)
	
}




func arrFor(){
	arr := []int{1,2,3}

	for key,val := range arr{
		fmt.Println(key,val)
	}
}



//数组切片
func arrSlice(){
	var arr [10]int = [10]int{1,2,3,4,5,6,7,8,9,0}

	arrB := arr[:5]
	arrC := arr[3:6]
	fmt.Println(arr)
	fmt.Println("切片后的数组B为:",arrB)
	fmt.Println("切片后的数组C为:",arrC)
}





//数组作为参数传入时会复制一份
func aguArr(){
	arr := [7]int{2,3,4,5,6,7,8}
	arrFun(arr)
	fmt.Println("传入的参数",arr)

}
//如果不写多少个数组则是使用了引用类型
func arrFun(getagu [7]int){
	getagu[0] = 222
	fmt.Println(getagu)
}



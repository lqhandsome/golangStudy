package main

import (
	"fmt"
)

func main()  {
	arr := [...]int{1,8,10,89,1000,2000}

	//res := minFind(&arr,0,5,1000)
	res := midFindForWhile(arr,20001)
	fmt.Println(res)
}

func minFind(arr *[6]int,leftIndex int, rightIndex int,find int)int{
	if rightIndex <  leftIndex{
		fmt.Println("没找到")
		return -1
	}
	var res int
	mid := (leftIndex + rightIndex) /2
	if (*arr)[mid] == find {
		fmt.Println("找到了index=",mid)
		return mid

	}
  	if (*arr)[mid] >  find{
	res = minFind(arr,leftIndex,mid -1,find)
	} else  {
		res = minFind(arr,mid + 1,rightIndex,find)
	}
	return res
}

func midFindForWhile(arr [6]int,find int)int{

	l := 0
	r := len(arr)
	for  l < r{
		mid  := ( l + r) /2
		if arr[mid] == find {
			fmt.Println("找到了")
			return mid
		}
		if arr[mid] < find  {
			l = mid +1
		}	else {
			r = mid - 1
		}
	}
	fmt.Println("没找到")
	return  -1
}
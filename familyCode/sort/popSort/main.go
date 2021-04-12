package main

import (
	"fmt"
)

func main()  {
	arr := [...]int{12,32,11213,412,5,31,23,2}
	fmt.Printf("%p",&arr[1])
	//for i := 0; i < count;i++{
	//	for j := i;j < count;j++ {
	//		if arr[j] > arr[i] {
	//			tmp := arr[j]
	//			arr[j] = arr[i]
	//			arr[i] = tmp
	//		}
	//	}
	//}
	popSort(&arr)
	fmt.Println(arr)
}

func popSort(arr *[8]int){
	count := len(*arr)
	fmt.Println(arr) //&[12 32 11213 412 5 31 23 2]

	fmt.Println((*arr)) //[12 32 11213 412 5 31 23 2]
	fmt.Printf("%v %T %p %v\n",arr,*(arr),&arr,arr)
	fmt.Printf("%p,%p %p %p\n",&arr,&(*arr),&(*arr)[1],&(arr)[1]) //0xc000148020,0xc000134080 0xc000134088 0xc000134088
	for i := 0;i < count;i++{
		for j := 0;j < count -1 -i;j++{
			if (arr)[j] > (arr)[j+1] {
				tmp :=(arr)[j]
				(arr)[j] = (arr)[j+1]
				(arr)[j+1] = tmp
			}
		}
	}
}

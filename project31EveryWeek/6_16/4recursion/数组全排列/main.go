package main

import "fmt"

func main() {
	arr := [5]int{1, 2,2}
	printArray(arr, 3, 3)
	fmt.Println(xibao(5))
}

//数组 0全排列
//k代表要处理的数组个数
func printArray(arr [5]int, len int, k int) {
	if k == 1 {
		for i := 0 ;i < len;i++{
			fmt.Print(arr[i])
		}
		fmt.Println()
	}
	for j :=0 ; j < k;j++ {
		arr[j],arr[k-1] = arr[k-1],arr[j]
		printArray(arr,len,k-1)
		arr[j],arr[k-1] = arr[k-1],arr[j]
	}
}

//https://time.geekbang.org/column/article/91541
//一个细胞生命周期是三小时，一小时分裂一次，n小时后容器内还剩多少细胞
func xibao(n int) (val int) {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return  2
	}
	if n == 2 {
		return 4
	}
	if n == 3 {
		return 7
	}
	return 2 * xibao(n -1) - xibao(n-4)
}

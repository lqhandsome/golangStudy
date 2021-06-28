package main

import "fmt"

func main() {

	//fmt.Println(lengthOfLongestSubstring("abcdfabcdefgt"))
	//fmt.Println(maxLengthString("abcacb"))
	//matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	//matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	//matrix := [][]int{{-1,3}}
	//fmt.Println(searchMatrix(matrix,3))
	num1 := []int{1,2,3,0,0,0}
	num2 := []int{2,5,6}
	fmt.Println(merge(num1,3,num2,3))
}
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return  0
	}
	maxLen := 1
	for i := 0;i < len(s) ;i++{
		j := i +1
		for ;j < len(s);j++ {
			b,_ := isOk(s,i,j)
			if !b {
				break
			}
		}
		fmt.Println(111,i,j)
		if maxLen <  j - i  {
			maxLen = j - i
			fmt.Println(17)
		}
	}
	return maxLen
}
func isOk(s string,start int,end int) (bool,int) {
	for i := start;i < end;i++ {
		if s[i] == s[end] {
			fmt.Printf("%c-%c\n",s[i],s[end])
			return false,i
		}
	}
	return true,0
}

func maxLengthString(s string) int {
	if len(s) == 0  {
		return 0
	}
	ma := make(map[uint8]int)
	max := 0
	//定义最左边
	left := 0
	for i := 0;i < len(s);i++ {
		//如果存在了
		if index,ok := ma[s[i]] ;ok{
			//判断存在的是否大于最左边界
			if index >= left {
				//如果存在，left就变成他下一个
				left = index +1
			}
		}
		//如果max大于i-最左边界+1
		if max < i - left  {
			fmt.Println(ma)
			max = i - left
		}
		//讲这个值的位置改变
		ma[s[i]] = i
	}
	return max
}

func singleNumber(nums []int) int {

	var num int
	for _,val := range nums {
		num ^= val
	}
	return num
}

func searchMatrix(matrix [][]int, target int) bool {
	clo,row := len(matrix) -1,0

	for clo >= 0 && row < len(matrix[0]) {
		if matrix[clo][row] == target  {
			return true
		}
		if matrix[clo][row] > target {
			clo--
		}
		if clo >= 0 && matrix[clo][row] < target   {
			row++
		}
	}
	return false
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var newArr = make([]int,0)

	i ,j:= 0,0
	for m  != i && n != j {
		if nums1[i] > nums2[j] {
			newArr = append(newArr,nums2[j])

			j++
		} else {
			newArr = append(newArr,nums1[i])
			i++
		}
	}
	for m  != i {
		newArr = append(newArr,nums1[i])
		i++

	}
	for n != j {
		newArr = append(newArr,nums2[j])
		j++
	}
	//nums1 = newArr
	for index,val := range newArr{
		nums1[index] = val
	}
	return nums1
}
package main

import "fmt"

var sum int64

func sumWater(arr []int,l,r int)  {
	//计算lr之间的 和小的值的差值。
	if r-l==1{
		return
	}
	if arr[l]>=arr[r]{
		for a:=l+1;a<r;a++{
			sum+=int64(arr[r]-arr[a])
		}
	}else {
		for a:=l+1;a<r;a++{
			sum+=int64(arr[l]-arr[a])
		}
	}
}
//定义两个指针，左指针指向起点，右指针也是如此
//右指针右移，如果下一个区间比左指针小，则继续右移，直到找到大于等于作指针大小得数
//如果 比左指针大，则左指针直接右移。然后计算sums
func maxWater( arr []int ) int64 {
	// write code here
	length:=len(arr)
	if length<=2{return 0}
	//sum:=0
	l:=0
	r:=0
	for k,v:=range arr{
		if v>=arr[l]||k==length-1{
			r=k
			sumWater(arr,l,r)
			l=k
		}else {
			r=k
		}
	}
	//l=length-1
	//r=length-1
	//for i := len(arr)-1; i >= 0; i-- {
	//	if arr[i]>=arr[l]{
	//		r=i
	//		sumWater(arr,r,l)
	//		l=i
	//	}else {
	//		r=i
	//	}
	//}

	return sum
}


func main() {
	//zbc:=[]int{1000000000,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1000000000}
	zbc:=[]int{4,5,1,3,2}
	fmt.Println(maxWater(zbc))
}

package main

import "fmt"

func insertionSortIncreasing(arr []int) []int{
	for j:=1; j<len(arr); j++{
		var key int = arr[j]
		i := j-1
		for i>=0 && arr[i]>key {
			arr[i+1] = arr[i]
			i = i-1
		}
		arr[i+1] = key
	}
	return arr
}
func insertionSortDecreasing(arr []int) []int{
	for j:=1; j<len(arr); j++{
		var key int = arr[j]
		i := j-1
		for i>=0 && arr[i]<key {
			arr[i+1] = arr[i]
			i = i-1
		}
		arr[i+1] = key
	}	
	return arr
}

func main(){
	var arr = []int{31, 41, 59, 26, 41, 58}
	fmt.Println(insertionSortIncreasing(arr))
	fmt.Println(insertionSortDecreasing(arr))
}
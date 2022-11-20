package main

import "fmt"

type MaxHeap struct {
	slice []int
}

func (heap *MaxHeap) Insert(newChild int) {
	heap.slice = append(heap.slice, newChild)
	heap.HeapifyUp(len(heap.slice) - 1)
}

func (heap *MaxHeap) HeapifyUp(index int) {
	for heap.slice[getParent(index)]<heap.slice[index] {
		heap.swap(getParent(index), index)
		index = getParent(index)
	}
}

func getParent(index int) int {
	return (index - 1) / 2
}

func getLeftChild(index int) int{
	return (index * 2) + 1
}

func getRightChild(index int) int{
	return (index *2) + 2
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.slice[index1], h.slice[index2] = h.slice[index2], h.slice[index1]
}


func main() {
	maxHeap := &MaxHeap{}
	createHeap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range createHeap {
		maxHeap.Insert(value)
		fmt.Println(maxHeap.slice)
	}
}
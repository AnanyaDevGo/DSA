package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(data int) {
	h.array = append(h.array, data)
	h.BuildHeap()
}
func (h *MaxHeap) BuildHeap() {
	n := len(h.array)
	for i := n/2 - 1; i >= 0; i-- {
		h.maxHeapify(i)
	}
}
func (h *MaxHeap) maxHeapify(Index int) {
	n := len(h.array) - 1
	var k int
	for {
		l := leftChild(Index)
		r := rightChild(Index)
		if l > n {
			break
		}
		if r <= n && h.array[r] > h.array[l] {
			k = r
		} else {
			k = l
		}
		if h.array[k] > h.array[Index] {
			swap(h.array, k, Index)
			Index = k
		} else {
			break
		}
	}
}

func (h *MaxHeap) Delete() int {
	if len(h.array) == 0 {
		return -1
	}
	maxValue := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.maxHeapify(0)
	return maxValue
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func display(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}
func main() {
	var arr = []int{11, 2, 7, 4, 12, 5, 9, 10}
	h := &MaxHeap{}
	for _, v := range arr {
		h.insert(v)
	}
	fmt.Println("Array:", arr)
	h.BuildHeap()
	fmt.Print("Heap:")
	display(h.array)
	h.Delete()
	fmt.Println("After delete ")
	display(h.array)
}

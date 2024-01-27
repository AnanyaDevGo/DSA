package main

import "fmt"

type Heap struct {
	array []int
}

func leftChild(i int) int {
	return (2 * i) + 1
}
func rightChild(i int) int {
	return (2 * i) + 2
}

func (h *Heap) Insert(i int) {
	h.array = append(h.array, i)
	h.Build()
}
func (h *Heap) Build() {
	n := len(h.array) //visit parent node
	for i := n/2 - 1; i >= 0; i-- {
		h.minHeapify(i)
	}
}
func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
func (h *Heap) minHeapify(Index int) {
	n := len(h.array) - 1
	var k int
	for {
		l := leftChild(Index)
		r := rightChild(Index)
		if l > n {
			break
		}
		if r <= n && h.array[r] < h.array[l] {
			k = r
		} else {
			k = l
		}
		if h.array[k] < h.array[Index] {
			swap(h.array, k, Index)
			Index = k
		} else {
			break
		}
	}
}

func (h *Heap) Remove() int {
	if len(h.array) == 0 {
		return -1
	}
	minVal := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.minHeapify(0)
	return minVal
}
func Display(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}
func main() {
	arr := []int{2, 3, 4, 6, 9, 8, 7}
	heap := &Heap{}
	for _, v := range arr {
		heap.Insert(v)
	}
	fmt.Println("Array is :-", arr)
	heap.Build()
	fmt.Println("Heap is :-")
	Display(heap.array)
	heap.Remove()
	fmt.Println("After Delete :-")
	Display(heap.array)

}

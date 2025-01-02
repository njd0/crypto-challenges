package structures

import "fmt"

type MinHeap struct {
	Data []int
	Size int
}

func NewSortedMinHeap (size int) *MinHeap {
	return &MinHeap{
		Data: make([]int, 0, size),
		Size: size,
	}
}

func (h MinHeap) Len() int {
	return len(h.Data)
}

func (h *MinHeap) Pop() int {
	old := h.Data
	length := h.Len()
	x := old[length - 1]
	h.Data = old[0 : length - 1]
	return x
}

func (h *MinHeap) Add(x int) {
	fmt.Println("x: ", x)
	length := h.Len()

	// empty array, add
	if length == 0 {
		h.Data = append(h.Data, x)
		return
	}

	i := length

	// array is full and value is too large
	if i == h.Size && x >= h.Data[i - 1] {
		return 
	}

	fmt.Println("length: ", length)

	// find sorted index for new element
	for {
		if x >= h.Data[i - 1] { break }
		i--
	}

	// shift array and insert into ith spot
	h.Data = append(h.Data[:i+1], h.Data[i:]...)
	h.Data[i] = x

	fmt.Println(h.Data)
} 
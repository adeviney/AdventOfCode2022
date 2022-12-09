package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"container/heap"
)

type MaxThreeItemsHeap []int

func (h MaxThreeItemsHeap) Len() int           { return len(h) }
func (h MaxThreeItemsHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxThreeItemsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *MaxThreeItemsHeap) Push(x any) {
	*h = append(*h, x.(int))

	// If the length of the heap is greater than 3,
	// remove the item with the
	if h.Len() > 3 {
		heap.Pop(h)
	}
}

func (h *MaxThreeItemsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Sum returns the sum of all the ints in the heap
func (h *MaxThreeItemsHeap) Sum() int {
	var sum int
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return sum
}

func (h *MaxThreeItemsHeap) Peek() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("error reading file: %s", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	TopThreeCalories := &MaxThreeItemsHeap{}
	heap.Init(TopThreeCalories)

	currentElfCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentElfCalories > TopThreeCalories.Peek() {
				TopThreeCalories.Push(currentElfCalories)
			}
			currentElfCalories = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentElfCalories += int(num)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(TopThreeCalories.Sum())
}

package main

import (
	"fmt"
	"sync"
)

// Find intersection of two slices

func findIntersection(slice1, slice2 []int) chan int {

	existed := map[int]struct{}{}
	//intersection := []int{}
	intersection := make(chan int, len(slice2))

	for _, j := range slice1 {
		//fmt.Println(i, j)
		existed[j] = struct{}{}
	}

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < len(slice2); i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, ok := existed[slice2[i]]; ok {
				//intersection = append(intersection, item)
				intersection <- slice2[i]
				mu.Lock()
				delete(existed, slice2[i])
				mu.Unlock()
			}
		}()
	}

	func() {
		wg.Wait()
		close(intersection)
	}()

	return intersection
}

func main() {
	sl1 := []int{1, 7, 5, 9, 3}
	sl2 := []int{3, 4, 5, 8, 7}

	for el := range findIntersection(sl1, sl2) {
		fmt.Println(el)
	}
}

package main

import "fmt"

// Find intersection of two slices

func findIntersection(slice1, slice2 []int) []int {

	existed := map[int]struct{}{}

	for _, j := range slice1 {
		//fmt.Println(i, j)
		existed[j] = struct{}{}
	}

	intersection := []int{}

	for _, item := range slice2 {
		if _, ok := existed[item]; ok {
			intersection = append(intersection, item)
			delete(existed, item)
		}
	}

	return intersection
}

func main() {
	sl1 := []int{1, 7, 5, 9, 3}
	sl2 := []int{3, 4, 5, 8, 7}

	fmt.Println(findIntersection(sl1, sl2))
}

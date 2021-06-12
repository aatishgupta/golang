package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	binarySearch(10, a)

}
func binarySearch(needle int, haystack []int) {

	start := 0
	end := len(haystack)
	counter := 1
	for end > start {
		middle := (start + end) / 2
		fmt.Println(counter, middle, haystack[middle])

		if haystack[middle] == needle {
			fmt.Println(haystack[middle], "Found at position of", middle)
			break
		} else if needle > haystack[middle] {
			start = middle
		} else {
			end = middle
		}
		counter = counter + 1
	}

}

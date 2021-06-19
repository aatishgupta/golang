package main

import "fmt"

func main() {
	a := []int{2, 4, 5, 3, 1, 32, 4, 5, 7, 5}
	for i := 0; i < len(a); i = i + 2 {
		a[i+1], a[i] = a[i], a[i+1]
	}
	fmt.Println(a)
}

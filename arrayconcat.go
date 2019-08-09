package main

import (
	"fmt"
)

func main() {
	dizi1 := []int{1, 2, 3}
	dizi2 := []int{4, 5, 6}

	result := append(dizi1, dizi2...)
	fmt.Println(result)
}

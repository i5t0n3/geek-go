package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}
	s = Add(s, 0, 5)
	fmt.Println(s, len(s), cap(s)) // [5 1 2 4 7] 5 8

	s = Add(s, 1, 9)
	fmt.Println(s, len(s), cap(s)) // [5 9 1 2 4 7] 6 8

	s = Add(s, 6, 13)
	fmt.Println(s, len(s), cap(s)) // [5 9 1 2 4 7 13] 7 8

	s = Delete(s, 2)
	fmt.Println(s, len(s), cap(s)) // [5 9 2 4 7 13] 6 8

	s = Delete(s, 0)
	fmt.Println(s, len(s), cap(s)) // [9 2 4 7 13] 5 8

	s = Delete(s, 4)
	fmt.Println(s, len(s), cap(s)) // [9 2 4 7] 4 8
}

func Add(s []int, index int, value int) []int {
	back := append([]int{}, s[index:]...)
	return append(append(s[:index], value), back...)
}

func Delete(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

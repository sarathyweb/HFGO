package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d")
	fmt.Println(slice, len(slice))

	s1 := []string{"s1", "s2"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
}

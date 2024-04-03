package main

import (
	"fmt"
)

func main() {
	test1()
	test2()
	test3()
	test4()
}

func test1() {
	fmt.Println("test1...")
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func test2() {
	fmt.Println("test2...")
	s := make([]int, 0) // s:= []int{}
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func test3() {
	fmt.Println("test3...")
	s := []int{}

	for i := 0; i < 10; i++ {
		func(s []int) {
			for i := 0; i < 10; i++ {
				s = append(s, i)
			}
		}(s)
	}

	fmt.Println(s)
}

func test4() {
	fmt.Println("test4...")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{-2, -3, -4, -5}

	s3 := append(s1[:1], s2...)

	fmt.Println(s1)
	fmt.Println(s3)
}

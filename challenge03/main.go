package main

import "log"

func main() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"

	log.Printf("s: %v\n", s)
}

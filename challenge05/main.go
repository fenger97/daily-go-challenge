package main

import "log"

type Interger int

// A
func (a Interger) Add(b Interger) Interger {
	return a + b
}

// B
// func (a Interger) Add(b *Interger) Interger {
// 	return a + *b
// }

// C
// func (a *Interger) Add(b Interger) Interger {
// 	return *a + b
// }

// D
// func (a *Interger) Add(b *Interger) Interger {
// 	return *a + *b
// }

func main() {
	var a Interger = 1
	var b Interger = 2
	var i interface{} = &a

	sum := i.(*Interger).Add(b)
	log.Print(sum)
}

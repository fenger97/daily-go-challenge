package main

import "log"

func main() {
	type myInt int
	var j myInt
	// test A
	{
		// var i int = 1
		// j = i
		// log.Print(j)
	}

	// test B
	{
		// var i int = 1
		// j = (myInt)i
		// log.Print(j)
	}

	// test C
	{
		var i int = 1
		j = myInt(i)
		log.Print(j)
	}

	// test D
	{
		// var i int = 1
		// j = i.(myInt)
		// log.Print(j)
	}
}

package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

// this function intSeq return another function, which we define anonymously in the body of intSeq.
// the returned function closees over the variable i to form a clouse.

//we call intSeq, assigning the result (a function) to nextInt. this function value captures its own-
//+ i value which will be updated each time we call nextInt.
//see the eeffect of the closuree by calling nextInt a few times.
//to confirm that the state is unique to that particular function, create and test a new one.

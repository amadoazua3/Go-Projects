package main

import "fmt"

func main() {

	// variable defintions in Go can be declared like this

	var i, j, k int
	var c, ch byte
	var f, salary float32

	fmt.Printf("i, j, and k are of type %T %T and %T\n", i, j, k)
	fmt.Printf("c and ch are of type %T and %T\n", c, ch)
	fmt.Printf("f and salary are of type %T and %T\n", f, salary)

	var student1 string = "Liam" //type is string
	var student2 = "Noel"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1 + " waived at " + student2)
	fmt.Printf("x is of type %T\n", x)

}

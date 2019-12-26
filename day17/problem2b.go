package main

import (
	"fmt"

	"./vm"
)

func main() {
	program := "A,C,A,B,C,B,C,A,B,C\n"
	A := "L,10,L,6,R,10\n"
	B := "L,10,R,8,R,8,L,10\n"
	C := "R,6,R,8,R,8,L,6,R,8\n"

	m := vm.CreateVMFromFile("input")
	m.Ram[0] = 2
	go m.Run()

	done := make(chan int)
	go func() {
		var last int
		for value := range m.Outputs {
			last = value
		}
		done <- last
		close(done)
	}()

	// Movement routines
	for _, ch := range program {
		m.Inputs <- int(ch)
	}
	for _, ch := range A {
		m.Inputs <- int(ch)
	}
	for _, ch := range B {
		m.Inputs <- int(ch)
	}
	for _, ch := range C {
		m.Inputs <- int(ch)
	}

	// Video feed
	m.Inputs <- int('n')
	m.Inputs <- int('\n')

	// Outputs
	fmt.Println(<-done)

}

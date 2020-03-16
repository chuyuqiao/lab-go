package main

import (
	"fmt"
)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %d\n", val)
		}()
	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func main() {
	x := 133
	f1 := foo1(&x)
	f2 := foo2(x)

	fmt.Println("Q1.1 start")
	for i := 0; i < 3; i++ {
		f1()
		f2()
	}

	fmt.Println("Q1.2 start")
	x = 233
	for i := 0; i < 3; i++ {
		f1()
		f2()
	}

	fmt.Println("Q1.3 start")
	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()
	foo2(x)()

	fmt.Println("Q2 start")
	foo3()
	foo4()
	foo5()

	fmt.Println("Q4 start")
	f7s := foo7(11)
	for _, fs := range f7s {
		fs()
	}
}

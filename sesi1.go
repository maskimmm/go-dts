package main

import "fmt"

func main() {
	var i = 21
	var j bool = true
	var k float64 = 123.456
	ja := 'Я'
	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	fmt.Println(j)
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%U\n", ja)
	fmt.Printf("%f\n", k)
	fmt.Printf("%E\n", k)

}

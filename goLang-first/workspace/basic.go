package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"

	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello word")
	variableZeroValue()
	variableInitaValue()
	variableTypeDeduction()
}

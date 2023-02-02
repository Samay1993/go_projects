package main

import "fmt"

func main() {
	fmt.Println("ABC\n \tDEF")
	fmt.Println(`
	This is multiline . \n
	String
	`)
	fmt.Println("😆")
	fmt.Println("\u2272")
	// fmt.Println('Linux Academy') --> This is rune, and can only have one character, this line will give error "illegal rune literal"
	fmt.Println('S')

	//Variable defining
	var a = "xyz"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)
}

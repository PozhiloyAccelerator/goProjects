package main

import "fmt"

func test() (a, b, c, d string) {
	a = "a"
	b = "b"
	c = "c"
	d = "d"
	return a, b, c, d
}

func main() {
	fmt.Print(test)

}

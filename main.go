package main

import "michael/trash/vec"

func main() {
	a := vec.NewVec2(1, 3)
	a.Print()
	a = a.Multiply(3)
	a.Print()
	println(a.Dot(vec.NewVec2(-1, 4)))
}

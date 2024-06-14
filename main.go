package main

import "michael/trash/vec"

func main() {
	a := vec.NewVec2(1, 3)
	a = a.Multiply(3)
	println(a.Dot(vec.NewVec2(-1, 4)))
}

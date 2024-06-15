package main

import "michael/trash/pkg/vec"

func main() {
	a := vec.NewVec2(1, 3)
	println(a.Dot(vec.NewVec2(-1, 4)))
}

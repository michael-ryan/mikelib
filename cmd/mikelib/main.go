package main

import (
	"fmt"
	"michael/trash/pkg/vec"
)

func main() {
	a := vec.NewVec3(1, 2, 3)
	b := vec.NewVec3(4, 5, 6)

	fmt.Println(a.Cross(b))
}

package main

import (
	"fmt"

	"github.com/niketa1493/solid-spork/pkg/mathutils"
)

func main() {
	fmt.Println("Add:", mathutils.Add(2, 3))
	fmt.Println("Subtract:", mathutils.Subtract(5, 2))
	fmt.Println("Multiply:", mathutils.Multiply(3, 4))

	result, err := mathutils.Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", result)
	}
}

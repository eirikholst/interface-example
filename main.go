package main

import (
	"fmt"
	"interface-example/pkg/geometry"
)

func main() {
	v2 := geometry.Vector2D{X: 3, Y: 4}
	printLength(v2)

	v3 := geometry.Vector3D{X: 3, Y: 4, Z: 5}
	printLength(v3)
}

func printLength(vector geometry.Vector) {
	fmt.Println(vector.Length())
}

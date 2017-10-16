package main

import (
	"fmt"
	"interface-example/pkg/geometry"
)

func main() {
	v2 := geometry.Vector2D{X: 3, Y: 4}
	printComponents(v2)
	printLength(v2)

	v3 := geometry.Vector3D{X: 3, Y: 4, Z: 5}
	printComponents(v3)
	printLength(v3)
}

func printComponents(vector geometry.Vector) {
	fmt.Println(vector.GetComponents())
}

func printLength(vector geometry.Vector) {
	fmt.Println(vector.Length())
}

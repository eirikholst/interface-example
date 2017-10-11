package main

import (
	"fmt"
	"interface-example/pkg/geometry"
)

func main() {
	v2 := geometry.Vector2D{X: 3, Y: 4}
	printComponents(v2)
	printAbs(v2)

	v3 := geometry.Vector3D{X: 3, Y: 4, Z: 5}
	printComponents(v3)
	printAbs(v3)
}

func printComponents(vector geometry.Vector) {
	fmt.Println(vector.GetComponents())
}

func printAbs(vector geometry.Vector) {
	fmt.Println(vector.Abs())
}

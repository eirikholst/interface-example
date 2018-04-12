package geometry

import "math"

type Vector interface {
	Length() float64
}

type Vector2D struct {
	X, Y float64
}

type Vector3D struct {
	X, Y, Z float64
}

func (v Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func (v Vector3D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}
